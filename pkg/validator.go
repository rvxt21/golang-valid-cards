package pkg

import "net/http"

// isValidCardNumber validates the provided credit card information, including the card number and expiration date.
// It checks the card number format, runs the Luhn algorithm to ensure the number is valid,
// and verifies that the expiration date is valid.
// Parameters:
// - creditCardReqBody: A struct containing the card number, expiration month, and expiration year.
// Returns:
// - A boolean indicating whether the card number and expiration date are valid.
// - An HTTP status code that represents the type of error if validation fails.
// - A string error code that represents the specific validation error.
// - A descriptive error message for the validation error.
func isValidCardNumber(creditCardReqBody creditCardRequestBody) (bool, int, string, string) {
	if creditCardReqBody.CardNumber == "" {
		return false, http.StatusBadRequest, "002", "Card number is required"
	}

	err := checkCardForPattern(creditCardReqBody.CardNumber)
	if err != nil {
		return false, http.StatusBadRequest, "003", err.Error()
	}

	_, err = luhnAlgorithmCheck(creditCardReqBody.CardNumber)
	if err != nil {
		return false, http.StatusBadRequest, "004", err.Error()
	}

	_, err = checkIsDateValid(creditCardReqBody.ExpirationMonth, creditCardReqBody.ExpirationYear)
	if err != nil {
		return false, http.StatusBadRequest, "005", err.Error()
	}

	return true, 0, "", ""
}
