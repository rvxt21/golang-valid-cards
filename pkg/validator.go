package pkg

import "net/http"

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
