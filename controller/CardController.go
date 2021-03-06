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

// give array of cards to add
// verify card details
// Enter unique card number
func PostCardDetails(res http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	var addCard []CardDetails
	err := json.NewDecoder(req.Body).Decode(&addCard)
	//addCard.ID = params["id"]
	for _, card := range addCard {
		CardDetailsList = append(CardDetailsList, card)
	}
	jsonBytes, err := json.Marshal(CardDetailsList)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func VerifyCardDetails(res http.ResponseWriter, req *http.Request) {
	var enteredCard CardDetails
	err := json.NewDecoder(req.Body).Decode(&enteredCard)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	for _, card := range CardDetailsList {
		if card.ID == enteredCard.ID && card.Name == enteredCard.Name && card.Cvv == enteredCard.Cvv && card.CardType == enteredCard.CardType {
			res.WriteHeader(http.StatusOK)
			res.Write([]byte("Card verified successfully!!!"))
			return
		}
	}
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("Card details doesn't match with any record!!!"))
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
