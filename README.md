# go.stripe
[![](https://drone.io/drone/go.stripe/status.png)](https://drone.io/drone/go.stripe/latest)

a simple Credit Card processing library for Go using the Stripe API

```sh
go get https://github.com/drone/go.stripe
```

## Examples

In order to use the `go.stripe` API you will need to create an account with
stripe.com, and obtain an Secret Key. You must set this key by invoking the
following function:

```go
stripe.SetKey("vtUQeOtUnYr7PGCLQ96Ul4zqpDUO4sOE")
```

Or you can specify your Secret Key in environment variable `STRIPE_API_KEY`, and
then invoke the following function:

```go
stripe.SetKeyEnv()
```

### Create Customer

```go
params := stripe.CustomerParams{
	Email:  "george.costanza@mail.com",
	Desc:   "short, bald",
	Card:   &stripe.CardParams {
		Name     : "George Costanza",
		Number   : "4242424242424242",
		ExpYear  : 2012,
		ExpMonth : 5,
		CVC      : "26726",
	},
}

customer, err := stripe.Customers.Create(&params)
```

### Charge Card

```go
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

charge, err := stripe.Charges.Create(&params)
```

Note: the amount charged is $4.00, but is specified in cents (400 cents == $4)

## Documentation

* [Customers](https://github.com/drone/go.stripe/wiki/Customers)
* [Charges](https://github.com/drone/go.stripe/wiki/Charges)
* [Coupons](https://github.com/drone/go.stripe/wiki/Coupons)
* [Invoices](https://github.com/drone/go.stripe/wiki/Invoices)
* [Invoice Items](https://github.com/drone/go.stripe/wiki/Invoice-Items)
* [Plans](https://github.com/drone/go.stripe/wiki/Plans)
* [Subscriptions](https://github.com/drone/go.stripe/wiki/Subscriptions)
* [Tokens](https://github.com/drone/go.stripe/wiki/Tokens)

You can also have a look at the [Godocs](http://gopkgdoc.appspot.com/pkg/github.com/drone/go.stripe).

## Unit Tests

In order to run the unit tests, you must have a Stripe account and a **Test** Secret
Key. The Test Secret Key must be set in environment variable `STRIPE_API_KEY`:

```sh
export STRIPE_API_KEY="vtUQeOtUnYr7PGCLQ96Ul4zqpDUO4sOE"
go test -v
```

The unit tests attempt to cleanup after themselves whenever possible. You can
manually clear all test data from the Stripe console by navigating to: Your 
Account » Account Settings » Test Data. Then click the "Remove All Test Data" button.
