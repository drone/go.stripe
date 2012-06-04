package stripe

import (
	"strconv"
)

// Int is a special type of integer that can unmarshall a JSON value of
// "null", which cannot be parsed by the Go JSON parser as of Go v1.
//
// see http://code.google.com/p/go/issues/detail?id=2540
type Int int

func (self *Int) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

	*self = Int(i)
	return nil
}

// Int64 is a special type of int64 that can unmarshall a JSON value of
// "null", which cannot be parsed by the Go JSON parser as of Go v1.
//
// see http://code.google.com/p/go/issues/detail?id=2540
type Int64 int64

func (self *Int64) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	*self = Int64(i)
	return nil
}

// Bool is a special type of bool that can unmarshall a JSON value of
// "null", which cannot be parsed by the Go JSON parser as of Go v1.
//
// see http://code.google.com/p/go/issues/detail?id=2540
type Bool bool

func (self *Bool) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}

	b, err := strconv.ParseBool(str)
	if err != nil {
		return err
	}

	*self = Bool(b)
	return nil
}

// String is a special type of string that can unmarshall a JSON value of
// "null", which cannot be parsed by the Go JSON parser as of Go v1.
//
// see http://code.google.com/p/go/issues/detail?id=2540
type String string

func (self *String) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}

	str, err := strconv.Unquote(str)
	if err != nil {
		return err
	}

	*self = String(str)
	return nil
}
