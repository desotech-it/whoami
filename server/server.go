package server

import (
	"fmt"
	"log"
	"net/http"
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
	unimplementedHandler(w, r)
}

func zeeHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func captainkubeHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func phippyHandler(w http.ResponseWriter, r *http.Request) {
	unimplementedHandler(w, r)
}

func Start() {
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
