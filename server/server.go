package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/desotech-it/whoami/app"
	"github.com/desotech-it/whoami/server/util"
	"github.com/desotech-it/whoami/view"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// withCounter wraps a handler to increment the global request counter.
func withCounter(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.IncrementRequestCounter()
		handler(w, r)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewIndexView("WhoAmI", app.Info, util.GetRequestAsString(r), util.MakeClientInfo(r))
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func cpustressHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	switch r.Method {
	case "GET":
		if app.IsOSDarwin() {
			v := view.NewUnsupportedView(
				"Unsupported Operation",
				"Retrieving the CPU info on MacOS is not currently implemented.",
				"https://github.com/shirou/gopsutil/issues/1000",
			)
			v.Write(w)
			return
		}
		stats, err := app.CPUInfo()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		v := view.NewCPUStressView("CPU Load", stats)
		v.Write(w)
	case "POST":
		if app.IsOSDarwin() {
			http.Error(w, "CPU stress is not supported on macOS", http.StatusNotImplemented)
			return
		}
		r.ParseForm()
		form := r.Form
		magnitude := form.Get("magnitude")
		unit := form.Get("unit")
		duration, err := time.ParseDuration(magnitude + unit)
		if err == nil {
			go util.GenerateCPULoadFor(duration)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func memstressHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	switch r.Method {
	case "GET":
		stats := app.MemInfo()
		v := view.NewMemStressView("Memory Usage", stats)
		v.Write(w)
	case "POST":
		r.ParseForm()
		form := r.Form
		magnitude := form.Get("magnitude")
		unit := form.Get("unit")
		duration, err := time.ParseDuration(magnitude + unit)
		if err == nil {
			go util.GenerateHighMemoryUsageFor(duration)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	health := app.GetHealth()
	var status string
	if health == 200 {
		status = "UP"
	} else {
		status = "ERRORED"
	}
	v := view.HealthView{
		Status: status,
	}
	w.WriteHeader(health)
	v.Write(w)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	readiness := app.GetReadiness()
	var status string
	if readiness == 200 {
		status = "UP"
	} else {
		status = "STARTING"
	}
	v := view.ReadinessView{
		Status: status,
	}
	w.WriteHeader(readiness)
	v.Write(w)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	promhttp.Handler().ServeHTTP(w, r)
}

func goldieHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewImageView("Goldie", "goldie.png", "goldie", app.Info, util.GetRequestAsString(r), util.MakeClientInfo(r))
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func zeeHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewImageView("Zee", "zee.png", "zee", app.Info, util.GetRequestAsString(r), util.MakeClientInfo(r))
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func captainkubeHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewImageView("Captain Kube", "captain-kube.png", "captain kube", app.Info, util.GetRequestAsString(r), util.MakeClientInfo(r))
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func phippyHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewImageView("Phippy", "phippy.png", "phippy", app.Info, util.GetRequestAsString(r), util.MakeClientInfo(r))
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func cpuusageHandler(w http.ResponseWriter, _ *http.Request) {
	cpusUsage := app.CPULoad()
	util.WriteJSONResponse(w, cpusUsage)
}

func memusageHandler(w http.ResponseWriter, _ *http.Request) {
	memUsage := app.MemInfo()
	util.WriteJSONResponse(w, memUsage)
}

// --- System Info handlers ---

func envHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewEnvView("Environment", app.GetEnvironment())
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func k8sHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewK8sView("K8s Metadata", app.GetK8sMetadata())
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func resourcesHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewResourcesView("Resource Limits", app.GetResourceLimits())
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewVolumesView("Volumes", app.GetMountPoints())
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func osinfoHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewOSInfoView("OS & Runtime", app.GetOSInfo())
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func dnsHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewDNSView("DNS Config", app.GetDNSConfig())
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

// --- HTTP Tools handlers ---

func headersHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewHeadersView("Echo Headers", r.Header)
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	codeStr := r.URL.Query().Get("code")
	code, err := strconv.Atoi(codeStr)
	if err != nil || code < 100 || code > 599 {
		code = 200
	}
	text := http.StatusText(code)
	if text == "" {
		text = "Unknown"
	}
	w.WriteHeader(code)
	v := view.NewStatusCodeView("Status Code", code, text)
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func delayHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	durationStr := r.URL.Query().Get("duration")
	delay, err := app.ParseDuration(durationStr)
	if err != nil || delay < 0 {
		delay = 0
	}
	if delay > 60*time.Second {
		delay = 60 * time.Second
	}
	time.Sleep(delay)
	v := view.NewDelayView("Response Delay", delay)
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	count := app.GetRequestCount()
	v := view.NewCounterView("Request Counter", count)
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

// --- Networking handlers ---

func downstreamHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	targetURL := r.URL.Query().Get("url")
	var result *app.DownstreamResult
	if targetURL != "" {
		res := app.CallDownstream(targetURL)
		result = &res
	}
	v := view.NewDownstreamView("Downstream Call", result)
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

// --- Docs handler ---

func docsHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.NewDocsView("Documentation")
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

type Server struct {
	Port uint64
}

func assignRootHandler() {
	targetImage := os.Getenv("NAME_APPLICATION")
	var targetHandler func(http.ResponseWriter, *http.Request)

	switch targetImage {
	case "goldie":
		targetHandler = goldieHandler
	case "zee":
		targetHandler = zeeHandler
	case "captainkube":
		targetHandler = captainkubeHandler
	case "phippy":
		targetHandler = phippyHandler
	default:
		targetHandler = rootHandler
	}

	http.HandleFunc("/", withCounter(targetHandler))
}

func (s *Server) Start() {
	// Startup delay
	if delay := app.GetStartupDelay(); delay > 0 {
		log.Printf("Startup delay: waiting %v before accepting connections...", delay)
		time.Sleep(delay)
	}

	assignRootHandler()

	// Existing endpoints
	http.HandleFunc("/cpustress", withCounter(cpustressHandler))
	http.HandleFunc("/memstress", withCounter(memstressHandler))
	http.HandleFunc("/readiness", withCounter(readinessHandler))
	http.HandleFunc("/healthz", withCounter(healthHandler))
	http.HandleFunc("/metrics", withCounter(metricsHandler))
	http.HandleFunc("/goldie", withCounter(goldieHandler))
	http.HandleFunc("/zee", withCounter(zeeHandler))
	http.HandleFunc("/captainkube", withCounter(captainkubeHandler))
	http.HandleFunc("/phippy", withCounter(phippyHandler))
	http.HandleFunc("/cpuusage", withCounter(cpuusageHandler))
	http.HandleFunc("/memusage", withCounter(memusageHandler))

	// System Info
	http.HandleFunc("/env", withCounter(envHandler))
	http.HandleFunc("/k8s", withCounter(k8sHandler))
	http.HandleFunc("/resources", withCounter(resourcesHandler))
	http.HandleFunc("/volumes", withCounter(volumesHandler))
	http.HandleFunc("/osinfo", withCounter(osinfoHandler))
	http.HandleFunc("/dns", withCounter(dnsHandler))

	// HTTP Tools
	http.HandleFunc("/headers", withCounter(headersHandler))
	http.HandleFunc("/status", withCounter(statusHandler))
	http.HandleFunc("/delay", withCounter(delayHandler))
	http.HandleFunc("/counter", withCounter(counterHandler))

	// Networking
	http.HandleFunc("/downstream", withCounter(downstreamHandler))
	http.HandleFunc("/ws", wsHandler) // WebSocket handler manages its own counter

	// Docs
	http.HandleFunc("/docs", withCounter(docsHandler))

	// Static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Graceful shutdown
	address := fmt.Sprintf(":%d", s.Port)
	srv := &http.Server{Addr: address}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		log.Printf("Listening on %s", address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownDelay := app.GetShutdownDelay()
	if shutdownDelay > 0 {
		log.Printf("Waiting %v before shutting down...", shutdownDelay)
		time.Sleep(shutdownDelay)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
