package controller

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
