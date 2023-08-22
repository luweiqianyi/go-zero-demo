FROM golang:1.19.2 AS builder

MAINTAINER runningriven@gmail.com

EXPOSE 8081
WORKDIR /usr/src/go-zero-demo

COPY . .

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -v -o userinfo-api ./cmd/userinfo \
    && cp userinfo-api ./cmd/userinfo/userinfo-api

FROM ubuntu
WORKDIR /usr/local/bin/userinfo-api

RUN mkdir ./etc

COPY --from=builder /usr/src/go-zero-demo/cmd/userinfo/userinfo-api ./userinfo-api
COPY --from=builder /usr/src/go-zero-demo/cmd/userinfo/etc/userinfo-api.yaml ./etc/userinfo-api.yaml

CMD ["./userinfo-api"]