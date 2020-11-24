package server

import (
	"desotech/whoami/view"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func unimplementedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, this endpoint is not implemented yet.")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
	fmt.Fprintln(w, "Root handler!")
}

func cpustressHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func memstressHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func goldieHandler(w http.ResponseWriter, r *http.Request) {
	view.WriteImageView(w, "goldie.png")
}

func zeeHandler(w http.ResponseWriter, r *http.Request) {
	view.WriteImageView(w, "zee.png")
}

func captainkubeHandler(w http.ResponseWriter, r *http.Request) {
	view.WriteImageView(w, "captain-kube.png")
}

func phippyHandler(w http.ResponseWriter, r *http.Request) {
	view.WriteImageView(w, "phippy.png")
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	basename := filepath.Base(r.URL.Path)
	image, err := os.Open("static/images/" + basename)
	if err != nil {
		panic(err) // TODO: handle error more gracefully
	}
	defer image.Close()
	if fileinfo, err := image.Stat(); err == nil {
		extension := filepath.Ext(fileinfo.Name())
		mimeType := mime.TypeByExtension(extension)
		header := w.Header()
		header.Set("Content-Type", mimeType)
		header.Set("Content-Length", strconv.FormatInt(fileinfo.Size(), 10))
	}
	io.Copy(w, image)
}

func Start() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/images/", imageHandler)
	http.HandleFunc("/cpustress", cpustressHandler)
	http.HandleFunc("/memstress", memstressHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/liveness", livenessHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/goldie", goldieHandler)
	http.HandleFunc("/zee", zeeHandler)
	http.HandleFunc("/captainkube", captainkubeHandler)
	http.HandleFunc("/phippy", phippyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
