package router

import (
	mux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func RouterConfig() *mux.Router {
	log.Info("Staring Server")
	router := mux.NewRouter()
	return router
}
