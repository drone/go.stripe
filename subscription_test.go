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

// Sample Subscriptions to use for testing
var (

	// Subscriptions with only the required fields
	sub1 = SubscriptionParams{
		Plan: "plan1",
	}

	// Subscriptions with all fields, plus new Credit Card
	sub2 = SubscriptionParams{
		Plan:     "plan1",
		Coupon:   "test coupon 1",
		Prorate:  true,
		TrialEnd: time.Now().Unix() + 1000000,
		Quantity: 5,
		Card: &CardParams{
			Name:     "George Costanza",
			Number:   "4242424242424242",
			ExpYear:  time.Now().Year() + 1,
			ExpMonth: 6,
		},
	}
)

func TestUpdateSubscription(t *testing.T) {
	// Create the customer, and defer its deletion
	cust, _ := Customers.Create(&cust1)
	defer Customers.Delete(cust.Id)

	// Create the plan, and defer its deletion
	Plans.Create(&p1)
	defer Customers.Delete(p1.Id)

	// Subscribe the Customer to the Plan
	resp, err := Subscriptions.Update(cust.Id, &sub1)
	if err != nil {
		t.Errorf("Expected Subscription, got error %s", err.Error())
	}
	if resp.Customer != cust.Id {
		t.Errorf("Expected Customer %s, got %s", cust.Id, resp.Customer)
	}
	if resp.Status != SubscriptionActive {
		t.Errorf("Expected Active Subscription, got %s", resp.Status)
	}
}

func TestUpdateSubscriptionCard(t *testing.T) {

	// Create the customer, and defer its deletion
	cust, _ := Customers.Create(&cust1)
	defer Customers.Delete(cust.Id)
	if cust.Cards.Count != 0 {
		t.Errorf("Expected Customer to be created with a nil card")
		return
	}

	// Create the plan, and defer its deletion
	Plans.Create(&p1)
	defer Customers.Delete(p1.Id)

	// Create the coupon, and defer its deletion
	Coupons.Create(&c1)
	defer Coupons.Delete(c1.Id)

	// Subscribe a Customer to a new plan, using a new Credit Card
	resp, err := Subscriptions.Update(cust.Id, &sub2)
	if err != nil {
		t.Errorf("Expected Subscription, got error %s", err.Error())
	}
	if resp.Quantity != sub2.Quantity {
		t.Errorf("Expected Quantity %d, got %d", sub2.Quantity, resp.Quantity)
	}

	// Check to see if the customer's card was added
	cust, _ = Customers.Retrieve(cust.Id)
	if cust.Cards.Count == 0 {
		t.Errorf("Expected Subscription to assign a new active customer card")
	}
}

func TestUpdateSubscriptionToken(t *testing.T) {
	// Create the customer, and defer its deletion
	cust, _ := Customers.Create(&cust1)
	defer Customers.Delete(cust.Id)
	if cust.Cards.Count != 0 {
		t.Errorf("Expected Customer to be created with a nil card")
		return
	}

	// Create the plan, and defer its deletion
	Plans.Create(&p1)
	defer Customers.Delete(p1.Id)

	// Create a Token for the credit card
	token, _ := Tokens.Create(&token1)

	// Subscribe the Customer to the Plan, using the Token
	params := SubscriptionParams{Plan: "plan1", Token: token.Id}
	_, err := Subscriptions.Update(cust.Id, &params)
	if err != nil {
		t.Errorf("Expected Subscription with Token, got error %s", err.Error())
	}

	// Check to see if the customer's card was added
	cust, _ = Customers.Retrieve(cust.Id)
	if cust.Cards.Count == 0 {
		t.Errorf("Expected Subscription to assign a new active customer card")
	}
}

func TestCancelSubscription(t *testing.T) {
	// Create the customer, and defer its deletion
	cust, _ := Customers.Create(&cust1)
	defer Customers.Delete(cust.Id)

	// Create the plan, and defer its deletion
	Plans.Create(&p1)
	defer Customers.Delete(p1.Id)

	// Subscribe the Customer to the Plan
	_, err := Subscriptions.Update(cust.Id, &sub1)
	if err != nil {
		t.Errorf("Expected Subscription, got error %s", err.Error())
	}

	// Now cancel the subscription
	subs, err := Subscriptions.Cancel(cust.Id)
	if err != nil {
		t.Errorf("Expected Subscription Cancellation, got error %s", err.Error())
	}

	if subs.Status != SubscriptionCanceled {
		t.Errorf("Expected Subscription Status %s, got %s", SubscriptionCanceled, subs.Status)
	}
}

func TestCancelSubscriptionAtPeriodEnd(t *testing.T) {
	// Create the customer, and defer its deletion
	cust, _ := Customers.Create(&cust1)
	defer Customers.Delete(cust.Id)

	// Create the plan, and defer its deletion
	Plans.Create(&p1)
	defer Customers.Delete(p1.Id)

	// Subscribe the Customer to the Plan
	_, err := Subscriptions.Update(cust.Id, &sub1)
	if err != nil {
		t.Errorf("Expected Subscription, got error %s", err.Error())
	}

	// Now cancel the subscription
	subs, err := Subscriptions.CancelAtPeriodEnd(cust.Id)
	if err != nil {
		t.Errorf("Expected Subscription Cancellation, got error %s", err.Error())
	}

	if subs.Status != SubscriptionActive {
		t.Errorf("Expected Subscription Status %s, got %s", SubscriptionCanceled, subs.Status)
	}

	if subs.CancelAtPeriodEnd != true {
		t.Errorf("Expected CancelAtPeriodEnd to be %s, got %s", true, subs.CancelAtPeriodEnd)
	}
}
