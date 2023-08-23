FROM golang:1.19.2 AS builder

MAINTAINER runningriven@gmail.com

WORKDIR /usr/src/go-zero-demo

COPY . .

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -v -o account-rpc ./cmd/account-rpc \
    && cp account-rpc ./cmd/account-rpc/account-rpc

FROM ubuntu:22.04
WORKDIR /usr/local/bin/account-rpc

RUN mkdir ./etc

COPY --from=builder /usr/src/go-zero-demo/cmd/account-rpc/account-rpc ./account-rpc
COPY --from=builder /usr/src/go-zero-demo/cmd/account-rpc/etc/account.yaml ./etc/account.yaml
COPY --from=builder /usr/src/go-zero-demo/cmd/account-rpc/etc/redis.yaml ./etc/redis.yaml

CMD ["./account-rpc"]