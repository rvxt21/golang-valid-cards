package pkg

// creditCardRequestBody defines the structure of the request body
// that contains the credit card details for validation.
type creditCardRequestBody struct {
	CardNumber      string `json:"card_number"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
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
