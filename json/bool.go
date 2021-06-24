package json

import (
	"encoding/json"
	"fmt"
	"strings"
)

var bools = map[string]bool{
	"true": true,
	"on":   true,
	"1":    true,
	"y":    true,
	"yes":  true,

	"false": false,
	"off":   false,
	"0":     false,
	"n":     false,
	"no":    false,
}

type Bool bool

func (r *Bool) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return fmt.Errorf("could not unmarshal `%s`: %v", data, err)
	}

	if s, ok := v.(bool); ok {
		*r = Bool(s)
		return nil
	}

	if s, ok := v.(string); ok {
		if v, has := bools[strings.ToLower(s)]; has {
			*r = Bool(v)
			return nil
		}
	}

	if i, ok := v.(float64); ok {
		if i == 1 {
			*r = true
			return nil
		} else if i == 0 {
			*r = false
			return nil
		}
	}

	return fmt.Errorf("could not unmarshal `%s` as a Bool value", data)
}

func (r Bool) Is() bool {
	return bool(r)
}
