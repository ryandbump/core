// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// at https://github.com/julienschmidt/httprouter/blob/master/LICENSE

package gateway

import (
	"net/url"
	"strings"
	"unicode"

	"github.com/percolate/shisa/service"

	"github.com/ansel1/merry"
)

type methodTree struct {
	method string
	root   *node
}

type treeSet []methodTree

func newTreeSet() treeSet {
	return make(treeSet, 0, 9)
}

func (s *treeSet) addEndpoint(endpoint *service.Endpoint) error {
	root := s.get(endpoint.Method)
	if root == nil {
		root = new(node)
		*s = append(*s, methodTree{method: endpoint.Method, root: root})
	}

	return root.addRoute(endpoint.Route, endpoint)

}

func (trees treeSet) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countParams(path string) uint8 {
	var n uint
	for i := 0; i < len(path); i++ {
		if path[i] != ':' && path[i] != '*' {
			continue
		}
		n++
	}
	if n >= 255 {
		return 255
	}
	return uint8(n)
}

type nodeType uint8

const (
	static nodeType = iota // default
	root
	param
	catchAll
)

type node struct {
	path      string
	wildChild bool
	nType     nodeType
	maxParams uint8
	indices   string
	children  []*node
	endpoint  *service.Endpoint
	priority  uint32
}

// addRoute adds a node with the given handle to the path. Not concurrency-safe!
func (n *node) addRoute(path string, endpoint *service.Endpoint) error {
	fullPath := path
	n.priority++
	numParams := countParams(path)

	// non-empty tree
	if len(n.path) > 0 || len(n.children) > 0 {
	walk:
		for {
			// Update maxParams of the current node
			if numParams > n.maxParams {
				n.maxParams = numParams
			}

			// Find the longest common prefix.
			// This also implies that the common prefix contains no ':' or '*'
			// since the existing key can't contain those chars.
			i := 0
			max := min(len(path), len(n.path))
			for i < max && path[i] == n.path[i] {
				i++
			}

			// Split edge
			if i < len(n.path) {
				child := node{
					path:      n.path[i:],
					wildChild: n.wildChild,
					indices:   n.indices,
					children:  n.children,
					endpoint:  n.endpoint,
					priority:  n.priority - 1,
				}

				// Update maxParams (max of all children)
				for i := range child.children {
					if child.children[i].maxParams > child.maxParams {
						child.maxParams = child.children[i].maxParams
					}
				}

				n.children = []*node{&child}
				// []byte for proper unicode char conversion, see #65
				n.indices = string([]byte{n.path[i]})
				n.path = path[:i]
				n.endpoint = nil
				n.wildChild = false
			}

			// Make new node a child of this node
			if i < len(path) {
				path = path[i:]

				if n.wildChild {
					n = n.children[0]
					n.priority++

					// Update maxParams of the child node
					if numParams > n.maxParams {
						n.maxParams = numParams
					}
					numParams--

					// Check if the wildcard matches
					if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
						// check for longer wildcard, e.g. :name and :names
						if len(n.path) >= len(path) || path[len(n.path)] == '/' {
							continue walk
						}
					}

					return merry.New("path expression conflict").WithUserMessagef(
						"path segment %q conflicts with existing wildcard %q in path %q",
						path, n.path, fullPath)
				}

				c := path[0]

				// slash after param
				if n.nType == param && c == '/' && len(n.children) == 1 {
					n = n.children[0]
					n.priority++
					continue walk
				}

				// Check if a child with the next path byte exists
				for i := 0; i < len(n.indices); i++ {
					if c == n.indices[i] {
						i = n.incrementChildPrio(i)
						n = n.children[i]
						continue walk
					}
				}

				// Otherwise insert it
				if c != ':' && c != '*' {
					// []byte for proper unicode char conversion, see #65
					n.indices += string([]byte{c})
					child := &node{
						maxParams: numParams,
					}
					n.children = append(n.children, child)
					n.incrementChildPrio(len(n.indices) - 1)
					n = child
				}
				return n.insertChild(numParams, path, fullPath, endpoint)

			} else if i == len(path) { // Make node a (in-path) leaf
				if n.endpoint != nil {
					return merry.New("path handler conflict").WithUserMessagef("endpoint already registered for path %q", fullPath)
				}
				n.endpoint = endpoint
			}
			return nil
		}
	} else { // Empty tree
		err := n.insertChild(numParams, path, fullPath, endpoint)
		n.nType = root
		return err
	}
}

// increments priority of the given child and reorders if necessary.
func (n *node) incrementChildPrio(pos int) int {
	n.children[pos].priority++
	prio := n.children[pos].priority

	// adjust position (move to front)
	newPos := pos
	for newPos > 0 && n.children[newPos-1].priority < prio {
		// swap node positions
		n.children[newPos-1], n.children[newPos] = n.children[newPos], n.children[newPos-1]

		newPos--
	}

	// build new index char string
	if newPos != pos {
		n.indices = n.indices[:newPos] + // unchanged prefix, might be empty
			n.indices[pos:pos+1] + // the index char we move
			n.indices[newPos:pos] + n.indices[pos+1:] // rest without char at 'pos'
	}

	return newPos
}

func (n *node) insertChild(numParams uint8, path string, fullPath string, endpoint *service.Endpoint) error {
	var offset int // already handled bytes of the path

	// find prefix until first wildcard (beginning with ':'' or '*'')
	for i, max := 0, len(path); numParams > 0; i++ {
		c := path[i]
		if c != ':' && c != '*' {
			continue
		}

		// find wildcard end (either '/' or path end)
		end := i + 1
		for end < max && path[end] != '/' {
			switch path[end] {
			// the wildcard name must not contain ':' and '*'
			case ':', '*':
				return merry.New("path wildcard conflict").WithUserMessagef(
					"only one wildcard per path segment is allowed, have: %q in path %q", path[i:], fullPath)
			default:
				end++
			}
		}

		// check if this Node existing children which would be
		// unreachable if we insert the wildcard here
		if len(n.children) > 0 {
			return merry.New("path wildcard conflict").WithUserMessagef(
				"wildcard route %q conflicts with existing children in path %q", path[i:end], fullPath)
		}

		// check if the wildcard has a name
		if end-i < 2 {
			return merry.New("invalid path wildcard").WithUserMessagef(
				"wildcards must be named with a non-empty name in path %q", fullPath)
		}

		if c == ':' { // param
			// split path at the beginning of the wildcard
			if i > 0 {
				n.path = path[offset:i]
				offset = i
			}

			child := &node{
				nType:     param,
				maxParams: numParams,
			}
			n.children = []*node{child}
			n.wildChild = true
			n = child
			n.priority++
			numParams--

			// if the path doesn't end with the wildcard, then there
			// will be another non-wildcard subpath starting with '/'
			if end < max {
				n.path = path[offset:end]
				offset = end

				child := &node{
					maxParams: numParams,
					priority:  1,
				}
				n.children = []*node{child}
				n = child
			}

		} else { // catchAll
			if end != max || numParams > 1 {
				return merry.New("invalid path catch-all").WithUserMessagef(
					"catch-all routes are only allowed at the end of the path in path %q", fullPath)
			}

			if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
				return merry.New("catch-all conflict").WithUserMessagef(
					"catch-all conflicts with existing handle for the path segment root in path %q", fullPath)
			}

			// currently fixed width 1 for '/'
			i--
			if path[i] != '/' {
				return merry.New("invalid path catch-all").WithUserMessagef("no / before catch-all in path %q", fullPath)
			}

			n.path = path[offset:i]

			// first node: catchAll node with empty path
			child := &node{
				wildChild: true,
				nType:     catchAll,
				maxParams: 1,
			}
			n.children = []*node{child}
			n.indices = string(path[i])
			n = child
			n.priority++

			// second node: node holding the variable
			child = &node{
				path:      path[i:],
				nType:     catchAll,
				maxParams: 1,
				endpoint:  endpoint,
				priority:  1,
			}
			n.children = []*node{child}

			return nil
		}
	}

	// insert remaining path part and handle to the leaf
	n.path = path[offset:]
	n.endpoint = endpoint

	return nil
}

// getValue returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.
// If no handle can be found, a TSR (trailing slash redirect) recommendation is
// made if a handle exists with an extra (without the) trailing slash for the
// given path.
func (n *node) getValue(path string, unescape bool) (endpoint *service.Endpoint, p Params, tsr bool, err error) {
walk: // Outer loop for walking the tree
	for {
		if len(path) > len(n.path) {
			if path[:len(n.path)] == n.path {
				path = path[len(n.path):]
				// If this node does not have a wildcard (param or catchAll)
				// child,  we can just look up the next child node and continue
				// to walk down the tree
				if !n.wildChild {
					c := path[0]
					for i := 0; i < len(n.indices); i++ {
						if c == n.indices[i] {
							n = n.children[i]
							continue walk
						}
					}

					// Nothing found.
					// We can recommend to redirect to the same URL without a
					// trailing slash if a leaf exists for that path.
					tsr = (path == "/" && n.endpoint != nil)
					return
				}

				// handle wildcard child
				n = n.children[0]
				switch n.nType {
				case param:
					// find param end (either '/' or path end)
					end := 0
					for end < len(path) && path[end] != '/' {
						end++
					}

					// save param value
					if cap(p) < int(n.maxParams) {
						p = make(Params, 0, n.maxParams)
					}
					i := len(p)
					p = p[:i+1] // expand slice within preallocated capacity
					p[i].Key = n.path[1:]
					val := path[:end]
					if unescape {
						var err error
						if p[i].Value, err = url.QueryUnescape(val); err != nil {
							p[i].Value = val // fallback, in case of error
						}
					} else {
						p[i].Value = val
					}

					// we need to go deeper!
					if end < len(path) {
						if len(n.children) > 0 {
							path = path[end:]
							n = n.children[0]
							continue walk
						}

						// ... but we can't
						tsr = (len(path) == end+1)
						return
					}

					if endpoint = n.endpoint; endpoint != nil {
						return
					}
					if len(n.children) == 1 {
						// No handle found. Check if a handle for this path + a
						// trailing slash exists for TSR recommendation
						n = n.children[0]
						tsr = (n.path == "/" && n.endpoint != nil)
					}

					return

				case catchAll:
					// save param value
					if cap(p) < int(n.maxParams) {
						p = make(Params, 0, n.maxParams)
					}
					i := len(p)
					p = p[:i+1] // expand slice within preallocated capacity
					p[i].Key = n.path[2:]
					if unescape {
						var err error
						if p[i].Value, err = url.QueryUnescape(path); err != nil {
							p[i].Value = path // fallback, in case of error
						}
					} else {
						p[i].Value = path
					}

					endpoint = n.endpoint
					return

				default:
					err = merry.New("internal error").WithUserMessage("internal error: invalid node type")
					return
				}
			}
		} else if path == n.path {
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
			if endpoint = n.endpoint; endpoint != nil {
				return
			}

			if path == "/" && n.wildChild && n.nType != root {
				tsr = true
				return
			}

			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
			for i := 0; i < len(n.indices); i++ {
				if n.indices[i] == '/' {
					n = n.children[i]
					tsr = (len(n.path) == 1 && n.endpoint != nil) ||
						(n.nType == catchAll && n.children[0].endpoint != nil)
					return
				}
			}

			return
		}

		// Nothing found. We can recommend to redirect to the same URL with an
		// extra trailing slash if a leaf exists for that path
		tsr = (path == "/") ||
			(len(n.path) == len(path)+1 && n.path[len(path)] == '/' &&
				path == n.path[:len(n.path)-1] && n.endpoint != nil)
		return
	}
}

// findCaseInsensitivePath makes a case-insensitive lookup of the given path and tries to find a handler.
// It can optionally also fix trailing slashes.
// It returns the case-corrected path and a bool indicating whether the lookup
// was successful.
func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) (ciPath []byte, found bool, err error) {
	ciPath = make([]byte, 0, len(path)+1) // preallocate enough memory

	// Outer loop for walking the tree
	for len(path) >= len(n.path) && strings.ToLower(path[:len(n.path)]) == strings.ToLower(n.path) {
		path = path[len(n.path):]
		ciPath = append(ciPath, n.path...)

		if len(path) > 0 {
			// If this node does not have a wildcard (param or catchAll) child,
			// we can just look up the next child node and continue to walk down
			// the tree
			if !n.wildChild {
				r := unicode.ToLower(rune(path[0]))
				for i, index := range n.indices {
					// must use recursive approach since both index and
					// ToLower(index) could exist. We must check both.
					if r == unicode.ToLower(index) {
						out, found, err := n.children[i].findCaseInsensitivePath(path, fixTrailingSlash)
						if err != nil {
							return nil, false, err
						}
						if found {
							return append(ciPath, out...), true, nil
						}
					}
				}

				// Nothing found. We can recommend to redirect to the same URL
				// without a trailing slash if a leaf exists for that path
				found = (fixTrailingSlash && path == "/" && n.endpoint != nil)
				return
			}

			n = n.children[0]
			switch n.nType {
			case param:
				// find param end (either '/' or path end)
				k := 0
				for k < len(path) && path[k] != '/' {
					k++
				}

				// add param value to case insensitive path
				ciPath = append(ciPath, path[:k]...)

				// we need to go deeper!
				if k < len(path) {
					if len(n.children) > 0 {
						path = path[k:]
						n = n.children[0]
						continue
					}

					// ... but we can't
					if fixTrailingSlash && len(path) == k+1 {
						return ciPath, true, nil
					}
					return
				}

				if n.endpoint != nil {
					return ciPath, true, nil
				} else if fixTrailingSlash && len(n.children) == 1 {
					// No handle found. Check if a handle for this path + a
					// trailing slash exists
					n = n.children[0]
					if n.path == "/" && n.endpoint != nil {
						return append(ciPath, '/'), true, nil
					}
				}
				return

			case catchAll:
				return append(ciPath, path...), true, nil

			default:
				err = merry.New("internal error").WithUserMessage("internal error: invalid node type")
				return
			}
		} else {
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
			if n.endpoint != nil {
				return ciPath, true, nil
			}

			// No handle found.
			// Try to fix the path by adding a trailing slash
			if fixTrailingSlash {
				for i := 0; i < len(n.indices); i++ {
					if n.indices[i] == '/' {
						n = n.children[i]
						if (len(n.path) == 1 && n.endpoint != nil) ||
							(n.nType == catchAll && n.children[0].endpoint != nil) {
							return append(ciPath, '/'), true, nil
						}
						return
					}
				}
			}
			return
		}
	}

	// Nothing found.
	// Try to fix the path by adding / removing a trailing slash
	if fixTrailingSlash {
		if path == "/" {
			return ciPath, true, nil
		}
		if len(path)+1 == len(n.path) && n.path[len(path)] == '/' &&
			strings.ToLower(path) == strings.ToLower(n.path[:len(path)]) &&
			n.endpoint != nil {
			return append(ciPath, n.path...), true, nil
		}
	}
	return
}