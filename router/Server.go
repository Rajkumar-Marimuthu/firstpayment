package router

import (
	controller "firstpayment/controller"

	"github.com/gorilla/mux"
)

func ServerConfig() *mux.Router {
	router := mux.NewRouter()
	controller.CardDetailsList = append(controller.CardDetailsList, controller.CardDetails{ID: "101", Name: "Anuradha", CardType: "VISA", Cvv: "444"})
	controller.CardDetailsList = append(controller.CardDetailsList, controller.CardDetails{ID: "102", Name: "Ashok", CardType: "MasterCard", Cvv: "555"})

	router.HandleFunc("/carddetails", controller.GetAllCardDetails).Methods("GET")
	router.HandleFunc("/carddetails/{id}", controller.GetCardDetails).Methods("GET")
	router.HandleFunc("/carddetails", controller.PostCardDetails).Methods("POST")
	router.HandleFunc("/carddetails/{id}", controller.DeleteCardDetails).Methods("DELETE")
	router.HandleFunc("/health", controller.HealthCheck).Methods("GET")

	return router

}
