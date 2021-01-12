package json

import (
	"encoding/json"
	"testing"
)

func TestBool_UnmarshalJSON(t *testing.T) {
	type obj struct {
		Key Bool `json:"key"`
	}

	tests := []struct {
		name    string
		arg     []byte
		expect  bool
		wantErr bool
	}{
		{name: "empty", arg: []byte("{\"key\": \"\"}"), wantErr: true},
		{name: "empty", arg: []byte("{\"key\": \" \"}"), wantErr: true},
		{name: "incorrect", arg: []byte("{\"key\": \"lorem\"}"), wantErr: true},
		{name: "true", expect: true, arg: []byte("{\"key\": true}"), wantErr: false},
		{name: "true", expect: true, arg: []byte("{\"key\": \"true\"}"), wantErr: false},
		{name: "true", expect: true, arg: []byte("{\"key\": \"on\"}"), wantErr: false},
		{name: "true", expect: true, arg: []byte("{\"key\": \"1\"}"), wantErr: false},
		{name: "true", expect: true, arg: []byte("{\"key\": \"y\"}"), wantErr: false},
		{name: "true", expect: true, arg: []byte("{\"key\": \"yes\"}"), wantErr: false},
		{name: "false", expect: false, arg: []byte("{\"key\": false}"), wantErr: false},
		{name: "false", expect: false, arg: []byte("{\"key\": \"false\"}"), wantErr: false},
		{name: "false", expect: false, arg: []byte("{\"key\": \"off\"}"), wantErr: false},
		{name: "false", expect: false, arg: []byte("{\"key\": \"0\"}"), wantErr: false},
		{name: "false", expect: false, arg: []byte("{\"key\": \"n\"}"), wantErr: false},
		{name: "false", expect: false, arg: []byte("{\"key\": \"no\"}"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := new(obj)
			if er := json.Unmarshal(tt.arg, o); (er != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", er, tt.wantErr)
			}

			if tt.expect != o.Key.Is() {
				t.Errorf("UnmarshalJSON() expected %v, got %v", tt.expect, o.Key.Is())
			}
		})
	}
}
