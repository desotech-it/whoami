FROM golang:1.17-alpine as builder

WORKDIR /whoami

COPY . .

ENV CGO_ENABLED=0

RUN apk update --no-cache && \
    apk upgrade --no-cache && \
    apk add --no-cache --update ca-certificates make git && \
    rm -rf /var/cache/apk* && \
    make build

# Create a minimal container to run a Golang static binary
FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /whoami/template/   /whoami/template/
COPY --from=builder /whoami/static/     /whoami/static/
COPY --from=builder /whoami/bin/whoami* /whoami/whoami

WORKDIR /whoami

RUN touch readiness

ENTRYPOINT ["./whoami", "-p", "80"]

EXPOSE 80
