package pkg

import (
	"errors"
	"strings"
)

func luhnAlgorithmCheck(cardNumber string) (bool, error) {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
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

	if sum%10 != 0 {
		return false, errors.New("the card number is incorrect")
	}
	return true, nil
}
