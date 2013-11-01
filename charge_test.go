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

// Sample Charges to use when creating, deleting, updating Charge data.
var (

	// Charge with only the required fields
	charge1 = ChargeParams{
		Desc:     "Calzone",
		Amount:   400,
		Currency: USD,
		Card: &CardParams{
			Name:     "George Costanza",
			Number:   "4242424242424242",
			ExpYear:  time.Now().Year() + 1,
			ExpMonth: 5,
		},
	}
)

// TestCreateCharge will test that we can successfully Charge a credit card,
// parse the JSON reponse from Stripe, and that all values are populated as
// expected.
func TestCreateCharge(t *testing.T) {

	// Create the charge
	resp, err := Charges.Create(&charge1)

	if err != nil {
		t.Errorf("Expected Successful Charge, got Error %s", err.Error())
	}
	if string(resp.Desc) != charge1.Desc {
		t.Errorf("Expected Charge Desc %s, got %s", charge1.Desc, resp.Desc)
	}
	if resp.Amount != charge1.Amount {
		t.Errorf("Expected Charge Amount %v, got %v", charge1.Amount, resp.Amount)
	}
	if resp.Card == nil {
		t.Errorf("Expected Charge Response to include the Charged Credit Card")
		return
	}
	if resp.Paid != true {
		t.Errorf("Expected Charge was paid, got %v", resp.Paid)
	}
}

// TestCreateChargeToken attempts to charge using a Card Token.
func TestCreateChargeToken(t *testing.T) {

	// Create a Token for the credit card
	token, err := Tokens.Create(&token1)
	if err != nil {
		t.Errorf("Expected Token Creation, got Error %s", err.Error())
	}

	// Create a Charge that uses a Token
	charge := ChargeParams{
		Desc:     "Calzone",
		Amount:   400,
		Currency: USD,
		Token:    token.Id,
	}

	// Create the charge
	_, err = Charges.Create(&charge)
	if err != nil {
		t.Errorf("Expected Successful Charge, got Error %s", err.Error())
	}
}

// TestCreateChargeCustomer attempts to charge a pre-defined customer, meaning
// we don't specify the credit card or token when Creating the charge.
func TestCreateChargeCustomer(t *testing.T) {

	// Create a Customer and defer deletion
	// This customer should have a credit card setup
	cust, _ := Customers.Create(&cust4)
	defer Customers.Delete(cust.Id)
	if cust.Cards.Count == 0 {
		t.Errorf("Cannot test charging a customer with no pre-defined Card")
		return
	}

	// Create a Charge that uses a Token
	charge := ChargeParams{
		Desc:     "Calzone",
		Amount:   400,
		Currency: USD,
		Customer: cust.Id,
	}

	// Create the charge
	_, err := Charges.Create(&charge)
	if err != nil {
		t.Errorf("Expected Successful Charge, got Error %s", err.Error())
	}
}

func TestRetrieveCharge(t *testing.T) {
	// Create the charge
	resp, err := Charges.Create(&charge1)
	if err != nil {
		t.Errorf("Expected Successful Charge, got Error %s", err.Error())
		return
	}

	// Retrieve the charge from the database
	_, err = Charges.Retrieve(resp.Id)
	if err != nil {
		t.Errorf("Expected to retrieve Charge by Id, got Error %s", err.Error())
		return
	}
}

func TestRefundCharge(t *testing.T) {
	// Create the charge
	resp, err := Charges.Create(&charge1)
	if err != nil {
		t.Errorf("Expected Successful Charge, got Error %s", err.Error())
		return
	}

	// Refund the full amount
	charge, err := Charges.Refund(resp.Id)
	if err != nil {
		t.Errorf("Expected Refund, got Error %s", err.Error())
		return
	}

	if charge.Refunded == false {
		t.Errorf("Expected Refund, however Refund flag was set to false")
	}
	if int64(charge.AmountRefunded) != charge1.Amount {
		t.Errorf("Expected AmountRefunded %v, but got %v", charge1.Amount, int64(charge.AmountRefunded))
		return
	}
}
