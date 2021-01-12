package json

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
		s = strings.ToLower(s)

		if s == "true" ||
			s == "on" ||
			s == "1" ||
			s == "y" ||
			s == "yes" {
			*r = true
			return nil
		} else if s == "false" ||
			s == "off" ||
			s == "0" ||
			s == "n" ||
			s == "no" {
			*r = false
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
