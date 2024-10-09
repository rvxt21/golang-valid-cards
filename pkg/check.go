package pkg

import (
	"errors"
	"regexp"
	"time"
)

// Patterns is a map that associates common credit card issuers with their corresponding regular expressions (regex).
// Each entry in the map contains a key (the card issuer's name) and a value (a regex pattern).
var Patterns = map[string]string{
	"Amex":          `^3[47][0-9]{13}$`,
	"BCGlobal":      `^(6541|6556)[0-9]{12}$`,
	"Carte Blanche": `^389[0-9]{11}$`,
	"Diners Club":   `^3(?:0[0-5]|[68][0-9])[0-9]{11}$`,
	"Discover":      `^(65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9]{0,9}|9[01][0-9]|92[0-5])[0-9]{10})$`,
	"Insta Payment": `^63[7-9][0-9]{13}$`,
	"JCB":           `^(?:2131|1800|35\d{3})\d{11}$`,
	"KoreanLocal":   `^9[0-9]{15}$`,
	"Laser":         `^(6304|6706|6709|6771)[0-9]{12,15}$`,
	"Maestro":       `^(5018|5020|5038|5893|6304|6759|6761|6762|6763)[0-9]{8,15}$`,
	"Mastercard":    `^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$`,
	"Solo":          `^(6334|6767)[0-9]{12,15}$`,
	"Switch":        `^(4903|4905|4911|4936|6333|6759)[0-9]{12,15}$`,
	"Union Pay":     `^(62[0-9]{14,17})$`,
	"Visa":          `^4[0-9]{12}(?:[0-9]{3})?$`,
	"Visa Master":   `^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$`,
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
