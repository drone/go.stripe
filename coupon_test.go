package stripe

import (
	"testing"
)

func init() {
	// In order to execute Unit Test, you must set your Stripe API Key as
	// environment variable, STRIPE_API_KEY=xxxx
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

// Sample Coupons to use when creating, deleting, updating Coupon data.
var (
	// Coupon with only the required fields
	c1 = CouponParams{
		Id:         "test coupon 1",
		PercentOff: 5,
		Duration:   DurationOnce,
	}

	// Coupon with all required + optional fields.
	c2 = CouponParams{
		Id:               "test coupon 2",
		PercentOff:       10,
		Duration:         DurationRepeating,
		MaxRedemptions:   100,
		DurationInMonths: 6,
	}
)

// TestCreateCoupon will test that we can successfully Create a coupon, parse
// the JSON reponse from Stripe, and that all values are populated as expected.
//
// Second, we will test that error handling works correctly by attempting to
// create a duplicate Coupon, which should thrown an exception.
func TestCreateCoupon(t *testing.T) {

	// Create the coupon, and defer its deletion
	coupon, err := Coupons.Create(&c1)
	defer Coupons.Delete(c1.Id)

	if coupon.Id != c1.Id {
		t.Errorf("Expected Coupon Id %s, got %s", c1.Id, coupon.Id)
	}
	if coupon.Duration != c1.Duration {
		t.Errorf("Expected Coupon Duration %v, got %v",
			c1.Duration, coupon.Duration)
	}
	if coupon.MaxRedemptions != Int(c1.MaxRedemptions) {
		t.Errorf("Expected Coupon MaxRedemptions %v, got %v",
			c1.MaxRedemptions, coupon.MaxRedemptions)
	}
	if coupon.PercentOff != c1.PercentOff {
		t.Errorf("Expected Coupon PercentOff %v, got %v",
			c1.PercentOff, coupon.PercentOff)
	}

	// Now try to re-create the existing coupon, which should throw an exception
	coupon, err = Coupons.Create(&c1)
	if err == nil {
		t.Error("Expected non-null Error when creating a duplicate coupon.")
	} else if err.Error() != "Coupon already exists." {
		t.Errorf("Expected %s, got %s", "Coupon already exists.", err.Error())
	}
}

// TestRetrieveCoupon will test that we can successfully Retrieve a Coupon,
// parse the JSON response, and that all values are populated as expected.
//
// Second, we will test that error handling works correctly by attempting to
// retrieve a coupon that does not exist. This should yield a Not Found error.
func TestRetrieveCoupon(t *testing.T) {
	// create a request that we can retrieve, defer deletion in case test fails
	Coupons.Create(&c2)
	defer Coupons.Delete(c2.Id)

	// now let's retrieve the recently added coupon
	coupon, err := Coupons.Retrieve(c2.Id)
	if err != nil {
		t.Errorf("Expected Coupon %s, got Error %s", c2.Id, err.Error())
	}
	if coupon.Id != c2.Id {
		t.Errorf("Expected Coupon Id %s, got %s", c2.Id, coupon.Id)
	}
	if coupon.PercentOff != c2.PercentOff {
		t.Errorf("Expected Coupon PercentOff %v, got %v",
			c2.PercentOff, coupon.PercentOff)
	}
	if coupon.Duration != c2.Duration {
		t.Errorf("Expected Coupon Duration %v, got %v",
			c2.Duration, coupon.Duration)
	}

	// now let's try to retrieve a coupon that doesn't exist, and make sure
	// we can handle the error
	_, err = Coupons.Retrieve("free for life")
	if err == nil {
		t.Error("Expected non-null Error when coupon not found.")
	}
}

// TestDeleteCoupon will test that we can successfully remove a Coupon, parse
// the JSON reponse, and that the deletion flag is captured as a boolean value.
func TestDeleteCoupon(t *testing.T) {
	// create a request that we can delete
	Coupons.Create(&c1)

	// let's try to delete the coupon
	ok, err := Coupons.Delete(c1.Id)
	if err != nil {
		t.Errorf("Expected Coupon deletion, got Error %s", err.Error())
	}
	if !ok {
		t.Errorf("Expected Coupon deletion true, got false")
	}
}

// TestListCoupon will test that we can successfully retrieve a list of Coupons,
// parse the JSON reponse, and that the length of the coupon array matches our
// expectations.
func TestListCoupon(t *testing.T) {

	// create 2 dummy coupons that we can retrieve
	Coupons.Create(&c1)
	Coupons.Create(&c2)
	defer Coupons.Delete(c1.Id)
	defer Coupons.Delete(c2.Id)

	// get the list from Stripe
	coupons, err := Coupons.List()
	if err != nil {
		t.Errorf("Expected Coupon List, got Error %s", err.Error())
	}

	// since we added 2 dummy coupons, we expect the array to be a size of 2
	if len(coupons) != 2 {
		t.Errorf("Expected 2 Coupons, got %s", len(coupons))
	}
}
