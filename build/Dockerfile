FROM golang:1.14-alpine3.11 as build

LABEL maintainer="https://github.com/takutakukatokatojapan"

WORKDIR /go/app

COPY . .

ENV GO111MODULE=off

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl && \
  go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

ENV GO111MODULE on

RUN set -eux && \
  go build -o go_redis ./main.go

FROM alpine:3.11

WORKDIR /app

COPY --from=build /go/app/go_redis .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/go_redis

CMD ["./go_redis"]