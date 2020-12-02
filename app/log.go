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
	logrus.Infof("%s requested %s", r.RemoteAddr, r.RequestURI)
}
