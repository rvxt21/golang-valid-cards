package pkg

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CreditCard defines the structure of the request body
// that contains the credit card details for validation.
type CreditCard struct {
	CardNumber, ExpirationMonth, ExpirationYear string
}

// Response defines the structure of the response for card validation.
type Response struct {
	Valid bool   `json:"valid"`
	Error *Error `json:"error,omitempty"`
}

// Error defines the structure of the error message in the response.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	ErrInvalidYear          = errors.New("invalid year")
	ErrInvalidMonth         = errors.New("invalid month")
	ErrCreditCardHasExpired = errors.New("credit card has expired")
	ErrIncorrectCardNumber  = errors.New("the card number is incorrect")
)

func (c *CreditCard) ValidateExpiration() error {
	var year, month int
	var err error
	timeNow := time.Now()
	if len(c.ExpirationYear) < 3 {
		year, err = strconv.Atoi(strconv.Itoa(timeNow.UTC().Year())[:2] + c.ExpirationYear)
		if err != nil {
			return ErrInvalidYear
		}
	} else {
		year, err = strconv.Atoi(c.ExpirationYear)
		if err != nil {
			return ErrInvalidYear
		}
	}

	month, err = strconv.Atoi(c.ExpirationMonth)
	if err != nil {
		return ErrInvalidMonth
	}

	switch {
	case month < 1 || month > 12:
		return ErrInvalidMonth
	case year < timeNow.Year():
		return ErrCreditCardHasExpired
	case year == timeNow.Year() && month < int(timeNow.Month()):
		return ErrCreditCardHasExpired
	}
	return nil
}

func (c *CreditCard) isValid() (bool, int, string, string) {
	if c.CardNumber == "" {
		return false, http.StatusBadRequest, "002", "Card number is required"
	}

	err := c.checkCardForPattern()
	if err != nil {
		return false, http.StatusBadRequest, "003", err.Error()
	}

	ok := c.ValidateNumber()
	if !ok {
		return false, http.StatusBadRequest, "004", ErrIncorrectCardNumber.Error()
	}

	err = c.ValidateExpiration()
	if err != nil {
		return false, http.StatusBadRequest, "005", err.Error()
	}

	return true, 0, "", ""
}

func (c *CreditCard) ValidateNumber() bool {
	cardNumber := strings.ReplaceAll(c.CardNumber, " ", "")
	sum := 0
	length := len(cardNumber)
	isSecondDigit := false

	for i := length - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit

		isSecondDigit = !isSecondDigit
	}

	return sum%10 == 0
}

func (c *CreditCard) checkCardForPattern() error {
	cardNumber := strings.ReplaceAll(c.CardNumber, " ", "")
	for _, pattern := range Patterns {
		matched, err := regexp.MatchString(pattern, cardNumber)
		if err != nil {
			return errors.New("error checking regexp")
		}
		if matched {
			return nil
		}
	}
	return errors.New("unknown card format")
}
