package stripe

import (
	"net/url"
	"strconv"
)

// Subscription Statuses
const (
	SubscriptionTrialing = "trialing"
	SubscriptionActive   = "active"
	SubscriptionPastDue  = "past_due"
	SubscriptionCanceled = "canceled"
	SubscriptionUnpaid   = "unpaid"
)

// see https://stripe.com/docs/api#subscription_object
type Subscription struct {
	Customer string `json:"customer"`
	Status   string `json:"status"`
	Plan     *Plan  `json:"plan"`

	// Date the subscription started
	Start int64 `json:"start"`

	// If the subscription has ended (either because it was canceled or because
	// the customer was switched to a subscription to a new plan), the date the
	// subscription ended
	EndedAt Int64 `json:"ended_at"`

	// Start of the current period that the subscription has been invoiced for 
	CurrentPeriodStart Int64 `json:"current_period_start"`

	// End of the current period that the subscription has been invoiced for.
	// At the end of this period, a new invoice will be created.
	CurrentPeriodEnd Int64 `json:"current_period_end"`

	// If the subscription has a trial, the beginning of that trial
	TrialStart Int64 `json:"trial_start"`

	// If the subscription has a trial, the end of that trial.
	TrialEnd Int64 `json:"trial_end"`

	// If the subscription has been canceled, the date of that cancellation. If
	// the subscription was canceled with cancel_at_period_end, canceled_at will
	// still reflect the date of the initial cancellation request, not the end
	// of the subscription period when the subscription is automatically moved
	// to a canceled state.
	CanceledAt Int64 `json:"canceled_at"`

	// If the subscription has been canceled with the at_period_end flag set to
	// true, cancel_at_period_end on the subscription will be true. You can use
	// this attribute to determine whether a subscription that has a status of
	// active is scheduled to be canceled at the end of the current period.
	CancelAtPeriodEnd bool `json:"cancel_at_period_end"`
}

type SubscriptionClient struct{}

type SubscriptionParams struct {
	// The identifier of the plan to subscribe the customer to.
	Plan string

	// The code of the coupon to apply to the customer if you would like to
	// apply it at the same time as creating the subscription.
	Coupon string

	// Flag telling us whether to prorate switching plans during a billing cycle
	Prorate  bool

	// UTC integer timestamp representing the end of the trial period the
	// customer will get before being charged for the first time. If set,
	// trial_end will override the default trial period of the plan the customer
	// is being subscribed to.
	TrialEnd int64

	// A new card to attach to the customer. The card can either be a token,
	// like the ones returned by our Stripe.js, or a Map containing a user's
	// credit card details
	//Card     *Card
}

// Subscribes a customer to a plan, meaning the customer will be billed monthly
// starting from signup. If the customer already has an active subscription,
// we'll update it to the new plan and optionally prorate the price we charge
// next month to make up for any price changes.
// TODO unable to include the card parameter at this time
//
// see https://stripe.com/docs/api#update_subscription
func (self *SubscriptionClient) Update(customerId string, params *SubscriptionParams) (*Subscription, error) {
	values := url.Values{"plan": {params.Plan}}
	if len(params.Coupon) != 0 {
		values.Add("coupon", params.Coupon)
	}
	if params.Prorate {
		values.Add("prorate", "true")
	}
	if params.TrialEnd != 0 {
		values.Add("trial_end", strconv.FormatInt(params.TrialEnd, 10))
	}

	s := Subscription{}
	path := "/v1/customers/" + url.QueryEscape(customerId) + "/subscription"
	err := query("POST", path, values, &s)
	return &s, err
}

// Cancels the subscription if it exists. If you set the atPeriodEnd parameter
// to true, the subscription will remain active until the end of the period, at
// which point it will be cancelled and not renewed.
// TODO enable at_period_end parameter
//
// see https://stripe.com/docs/api#cancel_subscription
func (self *SubscriptionClient) Cancel(customerId string) (*Subscription, error) {
	s := Subscription{}
	path := "/v1/customers/" + url.QueryEscape(customerId) + "/subscription"
	err := query("DELETE", path, nil, &s)
	return &s, err
}
