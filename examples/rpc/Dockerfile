FROM alpine:3.7

RUN apk --no-cache --update add go musl-dev

COPY ./ /root/go/src/github.com/percolate/shisa

WORKDIR /root/go/src/github.com/percolate/shisa/examples/rpc

RUN go build -ldflags '-linkmode external -extldflags "-static" -s' -o /root/go/bin/rpc github.com/percolate/shisa/examples/rpc

FROM alpine:3.7

COPY --from=0 /root/go/bin/rpc rpc

EXPOSE 9001

CMD ["./rpc", "-addr=127.0.0.1:9001"]
