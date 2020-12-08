# Whoami

Whoami is an app developed to demo the capabilities of Kubernetes. It provides endpoints to stress-test and get information on the running system.

## Available Endpoints

- `/`: the app's landing page only shows the hostname, the list of network interfaces and their IP address and the HTTP request that generated the response. See [Environment](#Environment) for how to slightly tweak this page.
- `/cpustress`: shows an overview of the CPU usage. You can also send all the CPU threads to maximum load for a chosen amount of time.
- `/memstress`: shows an overview of the memory usage. You can also pick a time interval during which the application will continuously perform bite-sized allocations until completely running out of memory (if you picked a high enough interval, otherwise the allocations will stop but the memory won't be reclaimed by the OS until the next run of the garbage collector).
- `/metrics`: useful metrics on the Go runtime, provided by [Prometheus](https://prometheus.io/).
- `/readiness` and `/healthz`: produce a JSON report on the current status of the container. Orchestrators such as Kubernetes can use the HTTP status code returned by these endpoints to manage containers autonomously.
- `/zee`, `/captainkube`, `/phippy`, `/goldie`: some cute pictures followed by the system info (as in `/`).

## Environment
- `LOG_INTERVAL`: controls how often the app logs CPU and memory usage to the console*
- `READINESS_DELAY`: controls how long it'll take for the app to report itself as ready to handle new incoming requests (15 seconds by default)*
- `HEALTH_DELAY`: controls how long it'll take for the app to report a healthy status*
- `NAME_APPLICATION`: you can optionally show an image in the landing page as per by `/goldie`, `/zee`, `/captainkube`, `/phippy`.
Possible values: `goldie`, `zee`, `captainkube`, `phippy`.

> (*) The value of those variables is fed to [time.ParseDuration](https://pkg.go.dev/time#ParseDuration) so any value accepted by that function is also accepted by those environment variables.

###### License
This code is published under the terms of the MIT license. As long as you credit us, you're free to use this code for any purpose you like!
