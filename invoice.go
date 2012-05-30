package stripe


// see https://stripe.com/docs/api#invoice_object
type Invoice struct { }

// see https://stripe.com/docs/api#invoiceitem_object
type InvoiceItem struct {
	Id       string "id"
	Amount   int64  "amount"
	Currency string "currency"
	Customer string "customer"
	Date     int64  "date"
	Desc     string "description"
	Invoice  string "invoice"
	Livemode bool   "livemode"
}
