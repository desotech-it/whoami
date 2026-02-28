# Whoami

> **DesoLabs by Desotech** — Kubernetes demo and stress-testing application.

Whoami is a Go web app designed for Kubernetes training labs. It displays system information (hostname, network interfaces, HTTP request details) and provides CPU/memory stress-test endpoints, Prometheus metrics, and Kubernetes health/readiness probes.

## Screenshot

![Preview](static/images/preview.png)

## Quick Start

### Run locally

```sh
make
./bin/whoami -p 8080
```

### Run in Kubernetes

```sh
kubectl run whoami --image=r.deso.tech/whoami/whoami:0.5.0 --port=80
kubectl expose pod whoami --type=NodePort --port=80
```

### Docker

```sh
docker build -t whoami .
docker run -p 8080:80 whoami
```

## Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Landing page: hostname, network interfaces, HTTP request, client info |
| `/cpustress` | GET/POST | CPU usage overview. POST to stress all cores for a given duration |
| `/memstress` | GET/POST | Memory usage overview. POST to allocate memory for a given duration |
| `/metrics` | GET | Prometheus metrics (Go runtime) |
| `/healthz` | GET | Kubernetes health probe (HTTP 200 or 503) |
| `/readiness` | GET | Kubernetes readiness probe (HTTP 200 or 503) |
| `/cpuusage` | GET | JSON API: current CPU usage per core |
| `/memusage` | GET | JSON API: current memory usage |
| `/goldie` | GET | Goldie character image + system info |
| `/zee` | GET | Zee character image + system info |
| `/captainkube` | GET | Captain Kube character image + system info |
| `/phippy` | GET | Phippy character image + system info |

All pages have dual output: **HTML** for browsers, **ASCII tables** for `curl`.

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `LOG_INTERVAL` | `2s` | How often CPU/memory stats are logged to console. Set to `0` to disable |
| `READINESS_DELAY` | `15s` | Delay before the app reports ready (simulates slow startup) |
| `HEALTH_DELAY` | `0s` | Delay before the app reports healthy |
| `NAME_APPLICATION` | _(none)_ | Override landing page with a character: `goldie`, `zee`, `captainkube`, `phippy` |

> Values are parsed by [time.ParseDuration](https://pkg.go.dev/time#ParseDuration).

## Architecture

- **Go 1.26** with static binary (`CGO_ENABLED=0`)
- Multi-arch: **linux/amd64**, **linux/arm64**, **windows/amd64**
- Multi-stage Docker build → minimal Alpine runtime image
- MVC-like structure: `app/` (model), `server/` (controller), `view/` (templates)

## Project Structure

```
.
├── main.go                     # Entry point
├── app/                        # Configuration, system info, health probes, logging
├── server/                     # HTTP handlers and routing
│   └── util/                   # Request parsing, JSON responses, stress generators
├── view/                       # View interface + HTML/text rendering
│   └── util/                   # Terminal output (ASCII tables, image-to-text)
├── template/                   # Go HTML templates (base layout + pages)
├── static/                     # CSS, JS, images (Bootstrap, K8s characters)
├── Dockerfile                  # Linux multi-stage build
├── windows.dockerfile          # Windows Server Core build
├── Makefile                    # Build, cross-compile, Docker targets
├── unix.mk                    # Unix-specific Make variables
├── windows.mk                 # Windows-specific Make variables
└── .github/workflows/
    └── release.yaml            # CI: build images + GitHub release on tag push
```

## Build

```sh
# Default build (current OS/arch)
make

# Cross-compile all platforms
make xcompress

# Docker (linux multi-arch)
make docker-linux

# Docker (windows)
make docker-windows

# Clean
make clean
```

## CI/CD

The release pipeline triggers on `v*` tag pushes and:

1. Builds and pushes Linux images (amd64 + arm64)
2. Builds and pushes Windows Server Core image
3. Creates multi-arch manifest lists
4. Publishes a GitHub Release with cross-compiled archives

## License

MIT - See [LICENSE](LICENSE) for details.
