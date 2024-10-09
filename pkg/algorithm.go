package pkg

import (
	"errors"
	"strings"
)

// // luhnAlgorithmCheck checks if the provided card number is valid based on the Luhn algorithm.
// The function iterates over the card number from right to left, doubling every second digit.
// If the doubled digit is greater than 9, 9 is subtracted from it. The resulting sum is then
// calculated, and if the sum modulo 10 is equal to 0, the card number is valid.
// Parameters:
// - cardNumber: A string representing the credit card number.
// Returns:
// - bool: true if the card number is valid according to the Luhn algorithm, false otherwise.
// - error: An error if the card number is invalid
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
