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
	Email:  "test2@test.com",
	Desc:   "a 2nd test customer",
	Coupon: c1.Id,
	Plan:   p1.Id,
	Card:   &stripe.CardParams {
		Name     : "John Smith",
		Number   : "4242424242424242",
		ExpYear  : time.Now().Year()+1,
		ExpMonth : 1,
	},
}

// Invoke the Customer Create function
customer, err := stripe.Customers.Create(&params)
```

