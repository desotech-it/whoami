package server

import (
	"bytes"
	"desotech/whoami/app"
	"desotech/whoami/view"
	"desotech/whoami/server/util"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	// Prometheus
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func unimplementedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, this endpoint is not implemented yet.")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	v := view.IndexView{
		Title:   "WhoAmI",
		Info:    app.Info,
		Request: getRequestAsString(r),
	}
	v.Write(w)
}

func cpustressHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func memstressHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
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

func goldieHandler(w http.ResponseWriter, r *http.Request) {
	v := view.ImageView{
		Title:    "Goldie",
		Filename: "goldie.png",
		AltText:  "goldie",
		Info:     app.Info,
		Request:  getRequestAsString(r),
	}
	v.Write(w)
}

func zeeHandler(w http.ResponseWriter, r *http.Request) {
	v := view.ImageView{
		Title:    "Zee",
		Filename: "zee.png",
		AltText:  "zee",
		Info:     app.Info,
		Request:  getRequestAsString(r),
	}
	v.Write(w)
}

func captainkubeHandler(w http.ResponseWriter, r *http.Request) {
	v := view.ImageView{
		Title:    "Captain Kube",
		Filename: "captain-kube.png",
		AltText:  "captain kube",
		Info:     app.Info,
		Request:  getRequestAsString(r),
	}
	v.Write(w)
}

func phippyHandler(w http.ResponseWriter, r *http.Request) {
	v := view.ImageView{
		Title:    "Phippy",
		Filename: "phippy.png",
		AltText:  "phippy",
		Info:     app.Info,
		Request:  getRequestAsString(r),
	}
	v.Write(w)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	basename := filepath.Base(r.URL.Path)
	image, err := os.Open("static/images/" + basename)
	if err != nil {
		panic(err) // TODO: handle error more gracefully
	}
	defer image.Close()
	util.AddContentInfoToResponseHeades(w, image)
	io.Copy(w, image)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	basename := filepath.Base(r.URL.Path)
	stylesheet, err := os.Open("static/css/" + basename)
	if err != nil {
		panic(err)
	}
	defer stylesheet.Close()
	util.AddContentInfoToResponseHeades(w, stylesheet)
	io.Copy(w, stylesheet)
}

func getRequestAsString(r *http.Request) string {
	buf := new(bytes.Buffer)
	if err := r.Write(buf); err != nil {
		return err.Error()
	}
	return string(buf.Bytes())
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
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/goldie", goldieHandler)
	http.HandleFunc("/zee", zeeHandler)
	http.HandleFunc("/captainkube", captainkubeHandler)
	http.HandleFunc("/phippy", phippyHandler)

	http.HandleFunc("/images/", imageHandler)
	http.HandleFunc("/css/", cssHandler)

	address := fmt.Sprintf(":%d", s.Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
