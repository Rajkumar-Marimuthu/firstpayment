package router

import (
	c "firstpayment/controller"

	mux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func RouterConfig() *mux.Router {
	log.Info("Staring Server")
	router := mux.NewRouter()
	router.HandleFunc("/health", c.HealthCheck).Methods("GET")
	return router
}
