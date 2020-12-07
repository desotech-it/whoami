package server

import (
	"desotech/whoami/app"
	"desotech/whoami/server/util"
	"desotech/whoami/view"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// Prometheus
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.IndexView{
		Title:   "WhoAmI",
		Info:    app.Info,
		Request: util.GetRequestAsString(r),
	}
	if util.IsFromCurl(r) {
		v.WriteAsText(w)
	} else {
		v.Write(w)
	}
}

func cpustressHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	stats := app.CPUInfo()
	v := view.CPUStressView{
		Title: "CPU Load",
		Stats: stats,
	}
	v.Write(w)

	duration, err := time.ParseDuration(r.URL.Query().Get("d"))
	if err == nil {
		go util.GenerateCPULoadFor(duration)
	}
}

func memstressHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	stats := app.MemInfo()
	v := view.MemStressView{
		Title: "Memory Usage",
		Stats: stats,
	}
	v.Write(w)

	duration, err := time.ParseDuration(r.URL.Query().Get("d"))
	if err == nil {
		go util.GenerateHighMemoryUsageFor(duration)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	health := app.GetHealth()
	json, err := health.GetJsonResponse()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if health == app.STOPPING || health == app.DOWN {
		http.Error(w, string(json), http.StatusServiceUnavailable)
		return
	}

	if health == app.ERRORED {
		http.Error(w, string(json), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	liveness := app.GetLiveness()
	json, err := liveness.GetJsonResponse()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if liveness == app.DOWN {
		http.Error(w, string(json), http.StatusServiceUnavailable)
		return
	}

	if liveness == app.ERRORED {
		http.Error(w, string(json), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	promhttp.Handler().ServeHTTP(w, r)
}

func goldieHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.ImageView{
		Title:    "Goldie",
		Filename: "goldie.png",
		AltText:  "goldie",
		Info:     app.Info,
		Request:  util.GetRequestAsString(r),
	}
	if util.IsFromCurl(r) {
		v.WriteAsPlainText(w)
	} else {
		v.Write(w)
	}
}

func zeeHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.ImageView{
		Title:    "Zee",
		Filename: "zee.png",
		AltText:  "zee",
		Info:     app.Info,
		Request:  util.GetRequestAsString(r),
	}
	if util.IsFromCurl(r) {
		v.WriteAsPlainText(w)
	} else {
		v.Write(w)
	}
}

func captainkubeHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.ImageView{
		Title:    "Captain Kube",
		Filename: "captain-kube.png",
		AltText:  "captain kube",
		Info:     app.Info,
		Request:  util.GetRequestAsString(r),
	}
	if util.IsFromCurl(r) {
		v.WriteAsPlainText(w)
	} else {
		v.Write(w)
	}
}

func phippyHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.ImageView{
		Title:    "Phippy",
		Filename: "phippy.png",
		AltText:  "phippy",
		Info:     app.Info,
		Request:  util.GetRequestAsString(r),
	}
	if util.IsFromCurl(r) {
		v.WriteAsPlainText(w)
	} else {
		v.Write(w)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	if err := util.ServeLocalResource(w, r.URL.Path[1:]); err != nil {
		if _, ok := err.(*os.PathError); ok {
			http.Error(w, "404 - Not found!", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error!", http.StatusInternalServerError)
		}
	}
}

func cpuusageHandler(w http.ResponseWriter, _ *http.Request) {
	cpusUsage := app.CPULoad()
	bytes, _ := json.Marshal(cpusUsage)
	w.Write(bytes)
}

func memusageHandler(w http.ResponseWriter, _ *http.Request) {
	memUsage := app.MemInfo()
	bytes, _ := json.Marshal(memUsage)
	w.Write(bytes)
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

	http.HandleFunc("/", targetHandler)
}

func (s *Server) Start() {
	assignRootHandler()
	http.HandleFunc("/cpustress", cpustressHandler)
	http.HandleFunc("/memstress", memstressHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/liveness", livenessHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/goldie", goldieHandler)
	http.HandleFunc("/zee", zeeHandler)
	http.HandleFunc("/captainkube", captainkubeHandler)
	http.HandleFunc("/phippy", phippyHandler)

	http.HandleFunc("/static/", staticHandler)

	http.HandleFunc("/cpuusage", cpuusageHandler)
	http.HandleFunc("/memusage", memusageHandler)

	address := fmt.Sprintf(":%d", s.Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
