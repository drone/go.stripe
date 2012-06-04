package stripe

import (
	"net/url"
	"strconv"
)

// Plan Intervals
const (
	IntervalMonth = "month"
	IntervalYear  = "year"
)

// see https://stripe.com/docs/api#plan_object
type Plan struct {
	Id              string `json:"id"`
	Name            string `json:"name"`              // Display name of the plan
	Amount          int64  `json:"amount"`            // The amount in cents to be charged on the interval specified
	Interval        string `json:"interval"`          // One of month or year. The frequency with which a subscription should be billed.
	Currency        string `json:"currency"`          // Currency in which subscription will be charged
	TrialPeriodDays Int    `json:"trial_period_days"` // Number of trial period days granted when subscribing a customer to this plan. Null if the plan has no trial period.
	Livemode        bool   `json:"livemode"`
}

type PlanClient struct{}

type CreatePlanReq struct {
	Id              string // Unique string of your choice that will be used to identify this plan when subscribing a customer.
	Amount          int64  // A positive integer in cents (or 0 for a free plan) representing how much to charge (on a recurring basis)
	Currency        string // 3-letter ISO code for currency. Currently, only 'usd' is supported.
	Interval        string // Specifies billing frequency. Either month or year.
	Name            string // Name of the plan, to be displayed on invoices and in the web interface.
	TrialPeriodDays int    // (Optional) Specifies a trial period in (an integer number of) days. If you include a trial period, the customer won't be billed for the first time until the trial period ends. If the customer cancels before the trial period is over, she'll never be billed at all.
}

// see https://stripe.com/docs/api?lang=java#create_plan
func (self *PlanClient) Create(req *CreatePlanReq) (*Plan, error) {
	plan := Plan{}
	values := url.Values{
		"id":       {req.Id},
		"name":     {req.Name},
		"amount":   {strconv.FormatInt(req.Amount, 10)},
		"interval": {req.Interval},
		"currency": {req.Currency},
	}

	// trial_period_days is optional, add if specified
	if req.TrialPeriodDays != 0 {
		values.Add("trial_period_days", strconv.Itoa(req.TrialPeriodDays))
	}

	err := query("POST", "/v1/plans", values, &plan)
	return &plan, err
}

// Retrieves the plan with the given ID.
//
// see https://stripe.com/docs/api?lang=java#retrieve_plan
func (self *PlanClient) Retrieve(id string) (*Plan, error) {
	plan := Plan{}
	path := "/v1/plans/" + url.QueryEscape(id)
	err := query("GET", path, nil, &plan)
	return &plan, err
}

// Updates the name of a plan. Other plan details (price, interval, etc.) are,
// by design, not editable.
//
// see https://stripe.com/docs/api?lang=java#update_plan
func (self *PlanClient) Update(id string, newName string) (*Plan, error) {
	values := url.Values{"name": {newName}}
	plan := Plan{}
	path := "/v1/plans/" + url.QueryEscape(id)
	err := query("POST", path, values, &plan)
	return &plan, err
}

// see https://stripe.com/docs/api?lang=java#delete_plan
func (self *PlanClient) Delete(id string) (*Plan, error) {
	plan := Plan{}
	path := "/v1/plans/" + url.QueryEscape(id)
	err := query("DELETE", path, nil, &plan)
	return &plan, err
}

// Returns a list of your Plans.
//
// see https://stripe.com/docs/api?lang=java#list_Plans
func (self *PlanClient) List() ([]*Plan, error) {
	return self.ListN(10, 0)
}

// Returns a list of your Plans with the specified count and at the specified
// offset.
//
// see https://stripe.com/docs/api?lang=java#list_Plans
func (self *PlanClient) ListN(count int, offset int) ([]*Plan, error) {
	// define a wrapper function for the Plan List, so that we can
	// cleanly parse the JSON
	type listPlanResp struct{ Data []*Plan }

	resp := listPlanResp{}
	err := query("GET", "/v1/plans", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
