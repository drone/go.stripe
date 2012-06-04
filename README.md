# go.stripe

a simple Credit Card processing library for Go using the Stripe API

```sh
go get https://github.com/bradrydzewski/go.stripe
```

Stripe API Documentation:<br/>
https://stripe.com/docs/api

Go Package Documentation:<br/>
http://gopkgdoc.appspot.com/pkg/github.com/bradrydzewski/go.stripe

## Examples

### Create Customer

```go
// set your API key
stripe.SetKey("vtUQeOtUnYr7PGCLQ96Ul4zqpDUO4sOE")

// define the Customer
params := stripe.CustomerParams{
	Email:  "george.costanza@mail.com",
	Desc:   "short, bald",
	Coupon: c1.Id,
	Plan:   p1.Id,
	Card:   &stripe.CardParams {
		Name     : "George Costanza",
		Number   : "4242424242424242",
		ExpYear  : 2012,
		ExpMonth : 5,
		CVC      : "26726",
	},
}

// Create the Customer via the Stripe REST API
customer, err := stripe.Customers.Create(&params)
```

### Create Charge

```go
// set your API key from environment (an alternative to hard-coding)
stripe.SetKeyEnv()

// setup the charge for $4.00 (expressed as 400 cents)
params := stripe.ChargeParams{
	Desc:     "Calzone",
	Amount:   400,
	Currency: "usd",
	Card:     &stripe.CardParams {
		Name     : "George Costanza",
		Number   : "4242424242424242",
		ExpYear  : 2012,
		ExpMonth : 5,
		CVC      : "26726",
	},
}

// Charge the Card via the Stripe REST API
charge, err := stripe.Charges.Create(&params)
```

## Unit Tests

The unit tests include sample code for nearly every API call. They are your best resource until I have the opportunity to write more thorough documentation.

In order to run your tests, you must have a Stripe account and a **Test** Secret Key. The Test Secret Key must be set in environment variable `STRIPE_API_KEY`:

```sh
export STRIPE_API_KEY="vtUQeOtUnYr7PGCLQ96Ul4zqpDUO4sOE"
go test -v
```

The unit tests attempt to cleanup after themselves whenever possible. You can manually clear all test data from the Stripe console by navigating to: Your Account » Account Settings » Test Data. Then click the "Remove All Test Data" button.
