FROM golang:1-alpine as builder

RUN apk --no-cache --no-progress add ca-certificates \
        && update-ca-certificates \
        && rm -rf /var/cache/apk/*

WORKDIR /go/whoami

COPY . .

RUN go build -tags netgo -ldflags '-extldflags "-static"'

# Create a minimal container to run a Golang static binary
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/whoami/template/ /whoami/template/
COPY --from=builder /go/whoami/static/   /whoami/static/
COPY --from=builder /go/whoami/whoami    /whoami/

WORKDIR /whoami

ENTRYPOINT ["/whoami/whoami", "-p", "80"]

EXPOSE 80
