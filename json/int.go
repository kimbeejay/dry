package json

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Int int64

func (r *Int) UnmarshalJSON(data []byte) error {
	var v interface{}

	err := json.Unmarshal(data, &v)
	if err != nil {
		return errors.New("[Int] [UnmarshalJSON]: " + err.Error())
	}

	if _, ok := v.(int64); ok {
		*r = Int(v.(int64))
		return nil
	}

	if _, ok := v.(float64); ok {
		*r = Int(v.(float64))
		return nil
	}

	if _, ok := v.(string); ok {
		i, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			*r = 0
		} else {
			*r = Int(i)
		}

		return nil
	}

	return errors.New("[Int] [UnmarshalJSON]: Unknown Int value [" + string(data) + "]")
}

func (r Int) Int64() int64 {
	return int64(r)
}
