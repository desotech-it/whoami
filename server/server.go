package server

import (
	"desotech/whoami/app"
	"desotech/whoami/server/util"
	"desotech/whoami/view"
	"fmt"
	"log"
	"os"
	"net/http"
	"path/filepath"
	"time"

	// Prometheus
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func unimplementedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, this endpoint is not implemented yet.")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	v := view.IndexView{
		Title:   "WhoAmI",
		Info:    app.Info,
		Request: util.GetRequestAsString(r),
	}
	v.Write(w)
}

func cpustressHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	duration, err := time.ParseDuration(r.URL.Query().Get("d"))
	if err == nil {
		go util.GenerateCPULoadFor(duration)
		http.Redirect(w, r, "/metrics", http.StatusPermanentRedirect)
		return
	}
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func memstressHandler(w http.ResponseWriter, r *http.Request) {
	duration, err := time.ParseDuration(r.URL.Query().Get("d"))
	if err == nil {
		go util.GenerateHighMemoryUsageFor(duration)
		http.Redirect(w, r, "/metrics", http.StatusPermanentRedirect)
		return
	}
	http.Error(w, err.Error(), http.StatusBadRequest)
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
	v.Write(w)
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
	v.Write(w)
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
	v.Write(w)
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
	v.Write(w)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	basename := filepath.Base(r.URL.Path)
	path := filepath.Join("static", "images", basename)
	if err := util.ServeLocalResource(w, path); err != nil {
		if _, ok := err.(*os.PathError); ok {
			http.Error(w, "404 - Not found!", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error!", http.StatusInternalServerError)
		}
	}
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)
	basename := filepath.Base(r.URL.Path)
	path := filepath.Join("static", "css", basename)
	if err := util.ServeLocalResource(w, path); err != nil {
		if _, ok := err.(*os.PathError); ok {
			http.Error(w, "404 - Not found!", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error!", http.StatusInternalServerError)
		}
	}
}

type Server struct {
	Port uint64
}

func (s *Server) Start() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/cpustress", cpustressHandler)
	http.HandleFunc("/memstress", memstressHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/liveness", livenessHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/goldie", goldieHandler)
	http.HandleFunc("/zee", zeeHandler)
	http.HandleFunc("/captainkube", captainkubeHandler)
	http.HandleFunc("/phippy", phippyHandler)

	http.HandleFunc("/images/", imageHandler)
	http.HandleFunc("/css/", cssHandler)

	address := fmt.Sprintf(":%d", s.Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
