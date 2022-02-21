package router

import (
	controller "firstpayment/controller"

	"github.com/gorilla/mux"
)

func ServerConfig() *mux.Router {
	router := mux.NewRouter()
	controller.CardDetailsList = append(controller.CardDetailsList, controller.CardDetails{ID: "1012201230124012", Name: "Anuradha", CardType: "VISA", Cvv: "444"})
	controller.CardDetailsList = append(controller.CardDetailsList, controller.CardDetails{ID: "1021202130214021", Name: "Ashok", CardType: "MasterCard", Cvv: "555"})
	controller.CardDetailsList = append(controller.CardDetailsList, controller.CardDetails{ID: "9999888877776666", Name: "Raj", CardType: "VISA", Cvv: "333"})
	controller.CardDetailsList = append(controller.CardDetailsList, controller.CardDetails{ID: "8888777766665555", Name: "Kumar", CardType: "AMEX", Cvv: "222"})

	router.HandleFunc("/carddetails", controller.GetAllCardDetails).Methods("GET")
	router.HandleFunc("/carddetails/{id}", controller.GetCardDetails).Methods("GET")
	router.HandleFunc("/carddetails", controller.PostCardDetails).Methods("POST")
	router.HandleFunc("/carddetails/{id}", controller.DeleteCardDetails).Methods("DELETE")
	router.HandleFunc("/health", controller.HealthCheck).Methods("GET")
	router.HandleFunc("/verifycard", controller.VerifyCardDetails).Methods("POST")

	return router

}
