package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostAndValidateCards(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    CreditCard
		expectedStatus int
		expectedValid  bool
		expectedError  *Error
	}{
		{
			name: "Valid card",
			requestBody: CreditCard{
				CardNumber:      "5512789002271854",
				ExpirationMonth: "12",
				ExpirationYear:  "2025",
			},
			expectedStatus: http.StatusOK,
			expectedValid:  true,
			expectedError:  nil,
		},
		{
			name: "Invalid card numer",
			requestBody: CreditCard{
				CardNumber:      "4111111111111112",
				ExpirationMonth: "12",
				ExpirationYear:  "2025",
			},
			expectedStatus: http.StatusBadRequest,
			expectedValid:  false,
			expectedError:  &Error{"004", "the card number is incorrect"},
		},
		{
			name: "Expired card",
			requestBody: CreditCard{
				CardNumber:      "341073406242763",
				ExpirationMonth: "10",
				ExpirationYear:  "2022",
			},
			expectedStatus: http.StatusBadRequest,
			expectedValid:  false,
			expectedError:  &Error{"005", "credit card has expired"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			PostAndValidateCards(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			var response Response
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("could not parse response: %v", err)
			}

			if response.Valid != tt.expectedValid {
				t.Errorf("handler returned unexpected body: got %v want %v", response.Valid, tt.expectedValid)
			}

			if tt.expectedError != nil {
				if response.Error == nil || response.Error.Code != tt.expectedError.Code || response.Error.Message != tt.expectedError.Message {
					t.Errorf("handler returned unexpected error: got %v want %v", response.Error, tt.expectedError)
				}
			} else {
				if response.Error != nil {
					t.Errorf("handler returned unexpected error: got %v want nil", response.Error)
				}
			}
		})
	}
}
