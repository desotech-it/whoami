FROM golang:1.17-alpine
RUN ["apk", "add", "--no-cache", "gcc", "musl-dev"]
RUN ["go", "install", "github.com/go-delve/delve/cmd/dlv@latest"]
RUN rm -rf -- "$GOPATH/pkg" "$HOME/.cache/"*
WORKDIR /app/src
EXPOSE 5000
