FROM golang:1.19.2 AS builder

MAINTAINER runningriven@gmail.com

EXPOSE 8081
WORKDIR /usr/src/go-zero-demo

COPY . .

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -v -o account ./cmd/account \
    && cp account ./cmd/account/account

FROM ubuntu
WORKDIR /usr/local/bin/account

RUN mkdir ./etc

COPY --from=builder /usr/src/go-zero-demo/cmd/account/account ./account
COPY --from=builder /usr/src/go-zero-demo/cmd/account/etc/account-api.yaml ./etc/account-api.yaml

CMD ["./account"]