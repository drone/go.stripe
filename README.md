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

// Invoke the Customer Create function
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

// Invoke the Charge Create function
charge, err := stripe.Charges.Create(&params)
```
