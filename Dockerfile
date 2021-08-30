FROM golang:1.17-alpine as builder

WORKDIR /go/whoami

COPY . .

ENV CGO_ENABLED=0

RUN apk update --no-cache && \
    apk upgrade --no-cache && \
    apk add --no-cache --update ca-certificates && \
    rm -rf /var/cache/apk* && \
    go build -tags netgo -ldflags '-extldflags "-static"'

# Create a minimal container to run a Golang static binary
FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/whoami/template/ /whoami/template/
COPY --from=builder /go/whoami/static/   /whoami/static/
COPY --from=builder /go/whoami/whoami    /whoami/

WORKDIR /whoami

RUN touch readiness

ENTRYPOINT ["./whoami", "-p", "80"]

EXPOSE 80
