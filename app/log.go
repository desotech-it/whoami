package app

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func LogRequest(r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"remote_address": r.RemoteAddr,
		"resource": r.RequestURI,
	}).Infof("Handling %s request for %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
}
