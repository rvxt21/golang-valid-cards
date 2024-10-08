package pkg

import (
	"testing"
)

func TestLuhnAlgorithmCheck(t *testing.T) {
	tests := []struct {
		cardNumber string
		expected   bool
	}{
		{"5512789002271854", true},
		{"5446137756964913", true},
		{"4916618047059690", true},
		{"6011438022129973", true},
		{"4556233516534075", false},
		{"1234567890123456", false},
		{"4485086012345670", false},
		{"invalid_card_number", false},
	}

	for _, test := range tests {
		t.Run(test.cardNumber, func(t *testing.T) {
			result, _ := luhnAlgorithmCheck(test.cardNumber)
			if result != test.expected {
				t.Errorf("expected %v, got %v for card number %s", test.expected, result, test.cardNumber)
			}
		})
	}
}
