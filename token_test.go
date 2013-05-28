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

// Sample Tokens to use when creating tokens
var (

	// Charge with only the required fields
	token1 = TokenParams{
		Card: &CardParams{
			Name:     "George Costanza",
			Number:   "4242424242424242",
			ExpYear:  time.Now().Year() + 1,
			ExpMonth: 5,
		},
	}
)

// TestCreateToken will test that we can successfully Create a Card Token,
// parse the JSON reponse from Stripe, and that all values are populated as
// expected.
func TestCreateToken(t *testing.T) {

	// Create the token
	resp, err := Tokens.Create(&token1)

	if err != nil {
		t.Errorf("Expected Token Created, got Error %s", err.Error())
	}
	if resp.Amount != 0 {
		t.Errorf("Expected Token Amount 0, got %v", resp.Amount)
	}
	if resp.Used == true {
		t.Errorf("Expected Token Used false, got true")
		return
	}
	if resp.Card == nil {
		t.Errorf("Expected Token Card not nil")
		return
	}
	if string(resp.Card.Name) != token1.Card.Name {
		t.Errorf("Expected Token Card Name %s, got %s", token1.Card.Name, resp.Card.Name)
	}
	if resp.Card.ExpMonth != token1.Card.ExpMonth {
		t.Errorf("Expected Token Card ExpMonth %d, got %d", token1.Card.ExpMonth, resp.Card.ExpMonth)
	}
	if resp.Card.ExpYear != token1.Card.ExpYear {
		t.Errorf("Expected Token Card ExpYear %d, got %d", token1.Card.ExpYear, resp.Card.ExpYear)
	}
	if resp.Card.Last4 != "4242" {
		t.Errorf("Expected Token Card Last4 4242, got %d", resp.Card.Last4)
	}
}

// TestCreateToken will test that we can successfully Retrieve a Card Token.
func TestRetrieveToken(t *testing.T) {
	// Create the token
	resp, err := Tokens.Create(&token1)
	if err != nil {
		t.Errorf("Expected Successful Token, got Error %s", err.Error())
		return
	}

	// Retrieve the Token from the database
	_, err = Tokens.Retrieve(resp.Id)
	if err != nil {
		t.Errorf("Expected to retrieve Token by Id, got Error %s", err.Error())
		return
	}
}
