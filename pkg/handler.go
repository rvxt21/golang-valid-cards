package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PostAndValidateCards handles the HTTP POST request for validating a credit card.
// It reads the request body to extract the card information, validates it using the isValidCardNumber function,
// and sends a response indicating whether the card is valid or not.
// If the request body is malformed, or the card is invalid, it responds with an appropriate error.
func PostAndValidateCards(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var creditCardReqBody creditCardRequestBody

	err := json.NewDecoder(r.Body).Decode(&creditCardReqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "001", "Invalid request body")
		return
	}

	ok, statusCode, code, message := isValidCardNumber(creditCardReqBody)
	if !ok {
		respondWithError(w, statusCode, code, message)
		return
	}
	respondWithSuccess(w, true)
}

// respondWithError sends an error response in JSON format with a specified status code, error code, and message.
// It sets the "Valid" field to false and populates the "Error" field with the provided error details.
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

// respondWithSuccess sends a success response in JSON format with a valid status.
// The "Valid" field indicates if the card passed validation, and the "Error" field is nil.
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
