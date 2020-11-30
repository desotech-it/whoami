package server

import (
	"desotech/whoami/app"
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
	v := view.ImageView{
		Title:    "Goldie",
		Filename: "goldie.png",
		AltText:  "goldie",
	}
	v.Write(w)
}

func zeeHandler(w http.ResponseWriter, r *http.Request) {
	v := view.ImageView{
		Title:    "Zee",
		Filename: "zee.png",
		AltText:  "zee",
	}
	v.Write(w)
}

func captainkubeHandler(w http.ResponseWriter, r *http.Request) {
	v := view.ImageView{
		Title:    "Captain Kube",
		Filename: "captain-kube.png",
		AltText:  "captain kube",
	}
	v.Write(w)
}

func phippyHandler(w http.ResponseWriter, r *http.Request) {
	v := view.ImageView{
		Title:    "Phippy",
		Filename: "phippy.png",
		AltText:  "phippy",
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
	if fileinfo, err := image.Stat(); err == nil {
		extension := filepath.Ext(fileinfo.Name())
		mimeType := mime.TypeByExtension(extension)
		header := w.Header()
		header.Set("Content-Type", mimeType)
		header.Set("Content-Length", strconv.FormatInt(fileinfo.Size(), 10))
	}
	io.Copy(w, image)
}

func Start(config *app.Config) {
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

	address := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
