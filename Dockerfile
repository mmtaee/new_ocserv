FROM golang:1.24.2 AS builder
LABEL authors="masoud"

ENV GIN_MODE=release
ENV CGO_ENABLED=1
ENV GOOS=linux

WORKDIR /app

COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod download

COPY . .

RUN go build -o ocserv_api main.go

FROM ubuntu:24.04

ENV SECRET_KEY_FILE_NAME=/tmp/init_secret
ENV LOG_FILE_PATH=/var/log/ocserv/ocserv.log

# dnsutils use for dig command in entrypoint
RUN apt update && apt install -y --no-install-recommends ocserv ca-certificates procps gnutls-bin build-essential iptables openssl less dnsutils jq &&\
    apt-get clean &&\
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN mkdir /app /db

COPY script/entrypoint.sh /entrypoint.sh

COPY script/server.sh /server.sh

COPY --from=builder /app/ocserv_api /ocserv_api

RUN chmod +x /entrypoint.sh /server.sh /ocserv_api

EXPOSE 443/tcp 443/udp

VOLUME ["/etc/ocserv", "/app", "/var/log/ocserv", "/db"]

ENTRYPOINT ["/entrypoint.sh"]

CMD ["/server.sh"]
