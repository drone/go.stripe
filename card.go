package stripe

import (
	"net/url"
	"strconv"
	"strings"
)

// Credit Card Types accepted by the Stripe API.
const (
	AmericanExpress = "American Express"
	DinersClub      = "Diners Club"
	Discover        = "Discover"
	JCB             = "JCB"
	MasterCard      = "MasterCard"
	Visa            = "Visa"
	UnknownCard     = "Unknown"
)

// Card represents details about a Credit Card entered into Stripe.
type Card struct {
	Id                string `json:"id"`
	Name              String `json:"name,omitempty"`
	Type              string `json:"type"`
	ExpMonth          int    `json:"exp_month"`
	ExpYear           int    `json:"exp_year"`
	Last4             string `json:"last4"`
	Fingerprint       string `json:"fingerprint"`
	Country           String `json:"country,omitempty"`
	Address1          String `json:"address_line1,omitempty"`
	Address2          String `json:"address_line2,omitempty"`
	AddressCountry    String `json:"address_country,omitempty"`
	AddressState      String `json:"address_state,omitempty"`
	AddressZip        String `json:"address_zip,omitempty"`
	AddressCity       String `json:"address_city"`
	AddressLine1Check String `json:"address_line1_check,omitempty"`
	AddressZipCheck   String `json:"address_zip_check,omitempty"`
	CVCCheck          String `json:"cvc_check,omitempty"`
}

// CardParams encapsulates options for Creating or Updating Credit Cards.
type CardParams struct {
	// (Optional) Cardholder's full name.
	Name string

	// The card number, as a string without any separators.
	Number string

	// Two digit number representing the card's expiration month.
	ExpMonth int

	// Four digit number representing the card's expiration year.
	ExpYear int

	// Card security code
	CVC string

	// (Optional) Billing address line 1
	Address1 string

	// (Optional) Billing address line 2
	Address2 string

	// (Optional) Billing address country
	AddressCountry string

	// (Optional) Billing address state
	AddressState string

	// (Optional) Billing address zip code
	AddressZip string
}

// CardClient encapsulates operations for creating, updating, deleting and
// querying cards using the Stripe REST API.
type CardClient struct{}

func (self *CardClient) Create(c *CardParams, customerId string) (*Card, error) {
	card := Card{}
	values := url.Values{}
	appendCardParamsToValues(c, &values)

	err := query("POST", "/v1/customers/"+customerId+"/cards", values, &card)
	return &card, err
}

func (self *CardClient) Delete(cardId string, customerId string) (*DeleteResp, error) {
	delResponse := DeleteResp{}
	values := url.Values{}

	err := query("DELETE", "/v1/customers/"+customerId+"/cards/"+cardId, values, &delResponse)
	return &delResponse, err
}

// IsLuhnValid uses the Luhn Algorithm (also known as the Mod 10 algorithm) to
// verify a credit cards checksum, which helps flag accidental data entry
// errors.
//
// see http://en.wikipedia.org/wiki/Luhn_algorithm
func IsLuhnValid(card string) (bool, error) {

	var sum = 0
	var digits = strings.Split(card, "")

	// iterate through the digits in reverse order
	for i, even := len(digits)-1, false; i >= 0; i, even = i-1, !even {

		// convert the digit to an integer
		digit, err := strconv.Atoi(digits[i])
		if err != nil {
			return false, err
		}

		// we multiply every other digit by 2, adding the product to the sum.
		// note: if the product is double digits (i.e. 14) we add the two digits
		//       to the sum (14 -> 1+4 = 5). A simple shortcut is to subtract 9
		//       from a double digit product (14 -> 14 - 9 = 5).
		switch {
		case even && digit > 4:
			sum += (digit * 2) - 9
		case even:
			sum += digit * 2
		case !even:
			sum += digit
		}
	}

	// if the sum is divisible by 10, it passes the check
	return sum%10 == 0, nil
}

// GetCardType is a simple algorithm to determine the Card Type (ie Visa,
// Discover) based on the Credit Card Number. If the Number is not recognized, a
// value of "Unknown" will be returned.
func GetCardType(card string) string {

	switch card[0:1] {
	case "4":
		return Visa
	case "2", "1":
		switch card[0:4] {
		case "2131", "1800":
			return JCB
		}
	case "6":
		switch card[0:4] {
		case "6011":
			return Discover
		}
	case "5":
		switch card[0:2] {
		case "51", "52", "53", "54", "55":
			return MasterCard
		}
	case "3":
		switch card[0:2] {
		case "34", "37":
			return AmericanExpress
		case "36":
			return DinersClub
		case "30":
			switch card[0:3] {
			case "300", "301", "302", "303", "304", "305":
				return DinersClub
			}
		default:
			return JCB
		}
	}

	return UnknownCard
}
