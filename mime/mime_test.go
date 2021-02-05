package mime

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    *Type
		wantErr bool
	}{
		{name: "empty", arg: "", want: nil, wantErr: true},
		{name: "simple", arg: "image/heic", want: &Type{
			registry:   Image,
			name:       "heic",
		}, wantErr: false},
		{name: "complicated", arg: "application/vnd.software602.filler.form-xml-zip", want: &Type{
			registry: Application,
			name: "vnd.software602.filler.form-xml-zip",
		}, wantErr: false},
		{name: "params", arg: "text/plain;charset=us-ascii", want: &Type{
			registry:   Text,
			name:       "plain",
			parameters: map[string]string{"charset": "us-ascii"},
		}, wantErr: false},
		{name: "incorrect", arg: "text/a", want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
