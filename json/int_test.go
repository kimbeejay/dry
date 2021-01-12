package json

import (
	"encoding/json"
	"testing"
)

func TestInt_UnmarshalJSON(t *testing.T) {
	type obj struct {
		Key Int `json:"key"`
	}

	tests := []struct {
		name    string
		arg     []byte
		expect  int64
		wantErr bool
	}{
		{name: "empty", arg: nil, expect: 0, wantErr: true},
		{name: "empty", arg: []byte("{\"key\": \"\"}"), expect: 0, wantErr: false},
		{name: "empty", arg: []byte("{\"key\": \" \"}"), expect: 0, wantErr: false},
		{name: "value", arg: []byte("{\"key\": \"1\"}"), expect: 1, wantErr: false},
		{name: "value", arg: []byte("{\"key\": 1}"), expect: 1, wantErr: false},
		{name: "value", arg: []byte("{\"key\": 1.0}"), expect: 1, wantErr: false},
		{name: "value", arg: []byte("{\"key\": 100.1}"), expect: 100, wantErr: false},
		{name: "value", arg: []byte("{\"key\": \"100.1\"}"), expect: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := new(obj)

			if er := json.Unmarshal(tt.arg, o); (er != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", er, tt.wantErr)
			}

			if tt.expect != o.Key.Int64() {
				t.Errorf("UnmarshalJSON() expected %v, got %v", tt.expect, o.Key.Int64())
			}
		})
	}
}
