package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PostAndValidateCards(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var creditCardReqBody creditCardRequestBody

	err := json.NewDecoder(r.Body).Decode(&creditCardReqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "001", "Invalid request body")
		return
	}

	if creditCardReqBody.CardNumber == "" {
		respondWithError(w, http.StatusBadRequest, "002", "Card number is required")
		return
	}

	err = checkCardForPattern(creditCardReqBody.CardNumber)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "003", err.Error())
		return
	}

	_, err = luhnAlgorithmCheck(creditCardReqBody.CardNumber)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "004", err.Error())
		return
	}

	_, err = checkIsDateValid(creditCardReqBody.ExpirationMonth, creditCardReqBody.ExpirationYear)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "005", err.Error())
		return
	}

	respondWithSuccess(w, true)
}

func respondWithError(w http.ResponseWriter, statusCode int, code, message string) {
	response := Response{
		Valid: false,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("Error encoding response: %v", err)
		return
	}

}

func respondWithSuccess(w http.ResponseWriter, isValid bool) {
	response := Response{
		Valid: isValid,
		Error: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
		fmt.Printf("Error encoding success response: %v\n", encodeErr)
		return
	}
}
