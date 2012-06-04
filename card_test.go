package stripe

import (
	"testing"
)

type card struct {
	Number string
	Type   string
	Valid  bool
}

var cards = []*card{
	&card{"4242424242424242", Visa, true},            // should pass
	&card{"4213729238347292", Visa, false},           // should fail
	&card{"79927398713", UnknownCard, true},          // should pass
	&card{"79927398710", UnknownCard, false},         // should fail
	&card{"601134239348202", Discover, false},        // should fail
	&card{"344347386473833", AmericanExpress, false}, // should fail
	&card{"374347386473833", AmericanExpress, false}, // should fail
	&card{"361134239348202", DinersClub, false},      // should fail
	&card{"300134239348202", DinersClub, false},      // should fail
	&card{"521134239348202", MasterCard, false},      // should fail
	&card{"380134239348202", JCB, false},             // should fail
	&card{"180034239348202", JCB, false},             // should fail
}

func TestLuhn(t *testing.T) {
	for _, card := range cards {
		valid, _ := IsLuhnValid(card.Number)
		cardType := GetCardType(card.Number)

		if valid != card.Valid {
			t.Errorf("card validation [%v]; want [%v]", valid, card.Valid)
		}
		if cardType != card.Type {
			t.Errorf("card type [%s]; want [%s]", cardType, card.Type)
		}
	}
}
