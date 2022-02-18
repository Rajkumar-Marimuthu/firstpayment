package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CardDetails struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	CardType string `json:"cardType"`
	Cvv      string `json:cvv`
}

var CardDetailsList []CardDetails

func GetCardDetails(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, Card := range CardDetailsList {
		if Card.ID == params["id"] {
			jsonBytes, err := json.Marshal(Card)
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
			}
			res.WriteHeader(http.StatusOK)
			res.Write(jsonBytes)
		}
	}
}

func GetAllCardDetails(res http.ResponseWriter, req *http.Request) {
	jsonBytes, err := json.Marshal(CardDetailsList)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func PostCardDetails(res http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	var addCard CardDetails
	err := json.NewDecoder(req.Body).Decode(&addCard)
	//addCard.ID = params["id"]
	CardDetailsList = append(CardDetailsList, addCard)
	jsonBytes, err := json.Marshal(CardDetailsList)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func DeleteCardDetails(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, Card := range CardDetailsList {
		if Card.ID == params["id"] {
			CardDetailsList = append(CardDetailsList[:index], CardDetailsList[index+1:]...)
			break
		}
	}
	jsonBytes, err := json.Marshal(CardDetailsList)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}
