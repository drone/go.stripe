package stripe

type Card struct {
    Id                string `json:"id"`
    Name              string `json:"name"`                // Cardholder name
    Type              string `json:"type"`                // Card brand. Can be Visa, American Express, MasterCard, Discover, JCB, Diners Club, or Unknown
    ExpMonth          int    `json:"exp_month"`
    ExpYear           int    `json:"exp_year"`
    Last4             int    `json:"last4"`
    Fingerprint       string `json:"fingerprint"`         // Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number
    Country           string `json:"country"`             // Two-letter ISO code representing the country of the card (as accurately as we can determine it). You could use this attribute to get a sense of the international breakdown of cards you've collected.
    Address1          string `json:"address_line1"`
	Address2          string `json:"address_line2"`
    AddressCountry    string `json:"address_country"`     // Billing address country, if provided when creating card
	AddressState      string `json:"address_state"`
    AddressZip        string `json:"address_zip"`
    AddressLine1Check string `json:"address_line1_check"` // If address_line1 was provided, results of the check: pass, fail, or unchecked
	AddressZipCheck   string `json:"address_zip_check"`   // If address_zip was provided, results of the check: pass, fail, or unchecked 
    CVCCheck          string `json:"cvc_check"`           // If a CVC was provided, results of the check: pass, fail, or unchecked
}

// TODO handle A common source of error is an invalid or expired card, or a valid card with insufficient available balance.
func (self *Card) IsExpired() bool {
	return false
}

// TODO check out the Luhn Algorithm to verify credit card numbers
// see http://en.wikipedia.org/wiki/Luhn_algorithm
func (self *Card) IsLuhnValid() bool {
	return false
}
