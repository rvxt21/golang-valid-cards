package pkg

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// Patterns is a map that associates common credit card issuers with their corresponding regular expressions (regex).
// Each entry in the map contains a key (the card issuer's name) and a value (a regex pattern).
var Patterns = map[string]string{
	"Amex":        `^3[47][0-9]{13}$`,
	"Diners Club": `^3(?:0[0-5]|[68][0-9])[0-9]{11}$`,
	"Discover":    `^(6011[0-9]{12}|64[4-9][0-9]{12,15}|65[0-9]{14,17})$$`,
	"JCB":         `^35(2[8-9]|[3-8][0-9])[0-9]{11,14}$`,
	"Maestro":     `^(5[0-9]{2}|6013|62|63|67)[0-9]{10,15}$`,
	"Mastercard":  `^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12,15})|5[1-5][0-9]{14})$`,
	"Visa":        `^4[0-9]{12}(?:[0-9]{3})?$`,
	"Visa Master": `^(4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$`,
}

// checkIsDateValid validates the provided credit card information about expiration month and expiration year.
// It verifies that the expiration date is valid.
// Parameters:
// - month, year: A int representing expiration month and expiration year.
// Returns:
// - A boolean indicating whether the card number and expiration date are valid.
// - A descriptive error message for the validation error.

func checkIsDateValid(month, year int) (bool, error) {
	currentYear, currentMonth, _ := time.Now().Date()

	if year < currentYear || (year == currentYear && month < int(currentMonth)) {
		return false, errors.New("the card is expired")
	} else if month < 1 || month > 12 {
		return false, errors.New("the month must be between 1 and 12")
	}

	return true, nil

}

// checkCardForPattern validates the provided card number for matching patterns.
// Parameters:
// - cardNumber: A string representing the credit card number.
// Returns:
// - A descriptive error message for the validation error.
func checkCardForPattern(cardNumber string) error {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	for _, pattern := range Patterns {
		matched, err := regexp.MatchString(pattern, cardNumber)
		if err != nil {
			return errors.New("error checking regexp")
		}
		if matched {
			return nil
		}
	}

	return errors.New("unkown card format")
}
