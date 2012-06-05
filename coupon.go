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

// see https://stripe.com/docs/api#coupon_object
type Coupon struct {
	// Coupon's Unique Identifier in the system.
	Id string `json:"id"`

	// Describes how long a customer who applies this coupon will get the
	// discount. Possible values are: forever, once, and repeating. 
	Duration string `json:"duration"`

	// Percent that will be taken off the subtotal of any invoices for this
	// customer for the duration of the coupon. For example, a coupon with
	// percent_off of 50 will make a $100 invoice $50 instead. 
	PercentOff int `json:"percent_off"`

	// If duration is repeating, the number of months the coupon applies. Null
	// if coupon duration is forever or once. 
	DurationInMonths Int `json:"duration_in_months,omitempty"`

	// Maximum number of times this coupon can be redeemed by a customer before
	// it is no longer valid. 
	MaxRedemptions Int `json:"max_redemptions,omitempty"`

	// Date after which the coupon can no longer be redeemed
	RedeemBy Int64 `json:"redeem_by,omitempty"`

	// Number of times this coupon has been applied to a customer.
	TimesRedeemed int  `json:"times_redeemed,omitempty"`
	Livemode      bool `json:"livemode"`
}

type CouponClient struct{}

// see https://stripe.com/docs/api?lang=java#create_coupon
type CouponParams struct {
	// Unique string of your choice that will be used to identify this coupon
	// when applying it a customer. 
	Id string

	// A positive integer between 1 and 100 that represents the discount the
	// coupon will apply.
	PercentOff int

	// Specifies how long the discount will be in effect. Can be forever, once,
	// or repeating.
	Duration string

	// If duration is repeating, a positive integer that specifies the number of
	// months the discount will be in effect.
	DurationInMonths int

	// A positive integer specifying the number of times the coupon can be
	// redeemed before it's no longer valid. For example, you might have a 50%
	// off coupon that the first 20 readers of your blog can use.
	MaxRedemptions int

	// UTC timestamp specifying the last time at which the coupon can be
	// redeemed. After the redeem_by date, the coupon can no longer be applied
	// to new customers.
	RedeemBy int64
}

// see https://stripe.com/docs/api?lang=java#create_coupon
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
// see https://stripe.com/docs/api?lang=java#retrieve_coupon
func (self *CouponClient) Retrieve(id string) (*Coupon, error) {
	coupon := Coupon{}
	path := "/v1/coupons/" + url.QueryEscape(id)
	err := query("GET", path, nil, &coupon)
	return &coupon, err
}

// Deletes the coupon with the given ID.
//
// see https://stripe.com/docs/api?lang=java#delete_coupon
func (self *CouponClient) Delete(id string) (*Coupon, error) {
	coupon := Coupon{}
	path := "/v1/coupons/" + url.QueryEscape(id)
	err := query("DELETE", path, nil, &coupon)
	return &coupon, err
}

// Returns a list of your coupons.
//
// see https://stripe.com/docs/api?lang=java#list_coupons
func (self *CouponClient) List() ([]*Coupon, error) {
	return self.ListN(10, 0)
}

// Returns a list of your coupons with the specified count and at the specified
// offset.
//
// see https://stripe.com/docs/api?lang=java#list_coupons
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
