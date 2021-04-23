FROM golang:1 as builder

WORKDIR /go/whoami

COPY . .

ENV CGO_ENABLED=0

RUN apt-get update && \
    apt-get upgrade -y && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    go build -tags netgo -ldflags '-extldflags "-static"'

# Create a minimal container to run a Golang static binary
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/whoami/template/ /whoami/template/
COPY --from=builder /go/whoami/static/   /whoami/static/
COPY --from=builder /go/whoami/whoami    /whoami/

WORKDIR /whoami

ENTRYPOINT ["./whoami", "-p", "80"]

EXPOSE 80
