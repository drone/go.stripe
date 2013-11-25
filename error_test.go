package stripe

import (
	"testing"
	"time"
)

func init() {
	// In order to execute Unit Test, you must set your Stripe API Key as
	// environment variable, STRIPE_API_KEY=xxxx
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

var (
	// Cards from https://stripe.com/docs/testing
	// These cards will be successfully charged.
	goodCards = []string{
		"4242424242424242",
		"4012888888881881",
		"5555555555554444",
		"5105105105105100",
		"378282246310005",
		"371449635398431",
		"6011111111111117",
		"6011000990139424",
		"30569309025904",
		"38520000023237",
		"3530111333300000",
		"3566002020360505",
		"4000000000000010",
		"4000000000000028",
		"4000000000000036",
		"4000000000000044",
		"4000000000000101",
	}
	// "These cards will produce specific responses that are useful for testing different scenarios"
	badCardsAndErrorCodes = map[string]string{
		"4000000000000341": ErrCodeCardDeclined,
		"4000000000000002": ErrCodeCardDeclined,
		"4000000000000127": ErrCodeIncorrectCVC,
		"4000000000000069": ErrCodeExpiredCard,
		"4000000000000119": ErrCodeProcessingError,
	}
	// Charge with only the required fields
	charge = ChargeParams{
		Desc:     "Turkish Delight",
		Amount:   300,
		Currency: USD,
		Card: &CardParams{
			Name: "Orhan Pamuk",
			//Number:   "", // This gets changed per-test
			ExpYear:  time.Now().Year() + 1,
			ExpMonth: 5,
		},
	}
)

// TestGoodCards ensures we can charge all of Stripe's "good" test cards.
func TestGoodCards(t *testing.T) {
	for _, cardNumber := range goodCards {
		charge.Card.Number = cardNumber
		if _, err := Charges.Create(&charge); err != nil {
			t.Errorf("Expected Successful Charge, got Error %s", err.Error())
		}
	}
}

// TestBadCards ensures we can't charge any of Stripe's "bad" test cards,
// and that the resulting error types and codes are correctly mapped.
func TestBadCards(t *testing.T) {
	for cardNumber, errCode := range badCardsAndErrorCodes {
		charge.Card.Number = cardNumber
		_, err := Charges.Create(&charge)
		stripeErr := err.(*Error)
		if stripeErr.Detail.Type != ErrTypeCard {
			t.Errorf("Expected Error Type %s, got %s", ErrTypeCard, stripeErr.Detail.Type)
		}
		if stripeErr.Detail.Code != errCode {
			t.Errorf("Expected Error Code %s, got %s", errCode, stripeErr.Detail.Code)
		}
	}
}
