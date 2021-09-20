FROM golang:1.17-windowsservercore-1809 as builder

SHELL ["powershell","-NoProfile", "-Command"]

WORKDIR C:\\whoami

COPY . .

ENV CGO_ENABLED=0

RUN Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh') ; \
    scoop install make git ; \
    make build

FROM mcr.microsoft.com/windows/servercore:1809

WORKDIR C:\\whoami

COPY --from=builder C:\\whoami\\template template\\
COPY --from=builder C:\\whoami\\static   static\\
COPY --from=builder C:\\whoami\\bin      .\\

# COPY in Windows dockerfiles doesn't work like in Linux so we need to resort to this hack
# to normalize the executable name
RUN ren whoami*.exe whoami.exe

# Create an empty file for readiness tests
RUN copy NUL readiness

ENTRYPOINT ["whoami.exe", "-p", "80"]

EXPOSE 80
