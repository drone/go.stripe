/*

An example of how to integrate Stripe into your Go application using either
Stripe.js (https://stripe.com/docs/stripe.js) or
Stripe Checkout (https://stripe.com/docs/checkout).

These tools will prevent credit card data from hitting your application, making
it easier to remain PCI compliant and minimising security risks as a result.

See the testing section in Stripe's documentation for a list of test card numbers,
error codes and other details: https://stripe.com/docs/testing

*/

package main

import (
	"fmt"
	"github.com/drone/go.stripe"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Defines a template variable for your Stripe publishable key
type TemplateVars struct {
	PublishableKey template.HTML
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}

func stripeJSHandler(w http.ResponseWriter, r *http.Request) {
	pubKey := TemplateVars{PublishableKey: template.HTML(os.Getenv("STRIPE_PUB_KEY"))}
	t := template.Must(template.ParseFiles("stripejs_form.html"))
	t.Execute(w, pubKey)
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	pubKey := TemplateVars{PublishableKey: template.HTML(os.Getenv("STRIPE_PUB_KEY"))}
	t := template.Must(template.ParseFiles("checkout_form.html"))
	t.Execute(w, pubKey)
}

// stripeToken represents a valid card token returned by the Stripe API.
// We use this to create a charge against the card instead of directly handling
// the credit card details in our application. Note that you could potentially
// collect the expiry date to allow you to remind users to update their card
// details as it nears expiry.
func paymentHandler(w http.ResponseWriter, r *http.Request) {

	// Use stripe.SetKeyEnv() to read the STRIPE_API_KEY environmental variable or alternatively
	// use stripe.SetKey() to set it directly (just don't publish it to GitHub!)
	err := stripe.SetKeyEnv()

	if err != nil {
		log.Fatal(err)
	}

	params := stripe.ChargeParams{
		Desc: "Pastrami on Rye",
		// Amount as an integer: 2000 = $20.00
		Amount:   2000,
		Currency: "AUD",
		Token:    r.PostFormValue("stripeToken"),
	}

	_, err = stripe.Charges.Create(&params)

	if err == nil {
		fmt.Fprintf(w, "Successful test payment!")
	} else {
		fmt.Fprintf(w, "Unsuccessful test payment: "+err.Error())
	}

}

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/stripejs", stripeJSHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/payment/new", paymentHandler)
	http.ListenAndServe(":9000", nil)
}
