FROM golang:1-nanoserver as builder

WORKDIR C:\\whoami

COPY . .

ENV CGO_ENABLED=0

RUN go build

FROM mcr.microsoft.com/windows/servercore:1809

COPY --from=builder C:\\whoami\\template\\ C:\\whoami\\template\\
COPY --from=builder C:\\whoami\\static\\ C:\\whoami\\static\\
COPY --from=builder C:\\whoami\\whoami.exe C:\\whoami\\

WORKDIR C:\\whoami

# Create an empty file for readiness tests
RUN copy NUL readiness

ENTRYPOINT ["whoami.exe", "-p", "80"]

EXPOSE 80
