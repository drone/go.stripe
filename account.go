package stripe

type Account struct {
	Id                  string   `json:"id"`
	Email               string   `json:"email,omitempty"`
	StatementDescriptor string   `json:"statement_descriptor"`
	DetailsSubmitted    bool     `json:"details_submitted"`
	ChargeEnabled       bool     `json:"charge_enabled"`
	TransferEnabled     bool     `json:"transfer_enabled"`
	CurrenciesSupported []string `json:"currencies_supported"`
}

type AccountClient struct {
	Client
}

func (self *AccountClient) Retrieve() (*Account, error) {
	account := Account{}
	path := "/v1/account"
	err := self.query("GET", path, nil, &account)
	return &account, err
}
