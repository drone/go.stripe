package stripe

import (
	"net/url"
	"strconv"
)

// Coupon Durations
const (
	DurationForever   = "forever"
	DurationOnce      = "once"
	DurationRepeating = "repeating"
)

// Coupon represents percent-off discount you might want to apply to a customer.
//
// see https://stripe.com/docs/api#coupon_object
type Coupon struct {
	Id               string `json:"id"`
	Duration         string `json:"duration"`
	PercentOff       int    `json:"percent_off"`
	DurationInMonths Int    `json:"duration_in_months,omitempty"`
	MaxRedemptions   Int    `json:"max_redemptions,omitempty"`
	RedeemBy         Int64  `json:"redeem_by,omitempty"`
	TimesRedeemed    int    `json:"times_redeemed,omitempty"`
	Livemode         bool   `json:"livemode"`
}

// CouponClient encapsulates operations for creating, updating, deleting and
// querying coupons using the Stripe REST API.
type CouponClient struct{}

// CouponParams encapsulates options for creating a new Coupon.
type CouponParams struct {
	// (Optional) Unique string of your choice that will be used to identify
	// this coupon when applying it a customer.
	Id string

	// A positive integer between 1 and 100 that represents the discount the
	// coupon will apply.
	PercentOff int

	// Specifies how long the discount will be in effect. Can be forever, once,
	// or repeating.
	Duration string

	// (Optional) If duration is repeating, a positive integer that specifies
	// the number of months the discount will be in effect.
	DurationInMonths int

	// (Optional) A positive integer specifying the number of times the coupon
	// can be redeemed before it's no longer valid. For example, you might have
	// a 50% off coupon that the first 20 readers of your blog can use.
	MaxRedemptions int

	// (Optional) UTC timestamp specifying the last time at which the coupon can
	// be redeemed. After the redeem_by date, the coupon can no longer be
	// applied to new customers.
	RedeemBy int64
}

// Creates a new Coupon.
//
// see https://stripe.com/docs/api#create_coupon
func (self *CouponClient) Create(params *CouponParams) (*Coupon, error) {
	coupon := Coupon{}
	values := url.Values{
		"duration":    {params.Duration},
		"percent_off": {strconv.Itoa(params.PercentOff)},
	}

	// coupon id is optional, add if specified
	if len(params.Id) != 0 {
		values.Add("id", params.Id)
	}

	// duration in months is optional, add if specified
	if params.DurationInMonths != 0 {
		values.Add("duration_in_months", strconv.Itoa(params.DurationInMonths))
	}

	// max_redemptions is optional, add if specified
	if params.MaxRedemptions != 0 {
		values.Add("max_redemptions", strconv.Itoa(params.MaxRedemptions))
	}

	// redeem_by is optional, add if specified
	if params.RedeemBy != 0 {
		values.Add("redeem_by", strconv.FormatInt(params.RedeemBy, 10))
	}
	err := query("POST", "/v1/coupons", values, &coupon)
	return &coupon, err
}

// Retrieves the coupon with the given ID.
//
// see https://stripe.com/docs/api#retrieve_coupon
func (self *CouponClient) Retrieve(id string) (*Coupon, error) {
	coupon := Coupon{}
	path := "/v1/coupons/" + url.QueryEscape(id)
	err := query("GET", path, nil, &coupon)
	return &coupon, err
}

// Deletes the coupon with the given ID.
//
// see https://stripe.com/docs/api#delete_coupon
func (self *CouponClient) Delete(id string) (bool, error) {
	resp := DeleteResp{}
	path := "/v1/coupons/" + url.QueryEscape(id)
	if err := query("DELETE", path, nil, &resp); err != nil {
		return false, err
	}
	return resp.Deleted, nil
}

// Returns a list of your coupons.
//
// see https://stripe.com/docs/api#list_coupons
func (self *CouponClient) List() ([]*Coupon, error) {
	return self.ListN(10, 0)
}

// Returns a list of your coupons at the specified range.
//
// see https://stripe.com/docs/api#list_coupons
func (self *CouponClient) ListN(count int, offset int) ([]*Coupon, error) {
	// define a wrapper function for the Coupon List, so that we can
	// cleanly parse the JSON
	type listCouponResp struct{ Data []*Coupon }
	resp := listCouponResp{}

	// add the count and offset to the list of url values
	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := query("GET", "/v1/coupons", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
