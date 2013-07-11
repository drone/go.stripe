package stripe

import (
	"fmt"
	"testing"
)

func init() {
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

func TestRetrieveAccount(t *testing.T) {
	account, err := Accounts.Retrieve()

	if err != nil {
		t.Errorf("Expected account, got Error %s", err.Error())
	}

	fmt.Println(account)
}
