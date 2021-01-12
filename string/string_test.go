package string

import "testing"

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{name: "non-empty", arg: "lorem ipsum", want: false},
		{name: "non-empty", arg: " a", want: false},
		{name: "empty", arg: " ", want: true},
		{name: "empty", arg: "   ", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.arg); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidUrl(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{name: "empty", arg: "", want: false},
		{name: "invalid", arg: "192.168.0.1", want: false},
		{name: "invalid", arg: "localhost", want: false},
		{name: "valid", arg: "http://localhost", want: true},
		{name: "valid", arg: "https://localhost/way/to/die", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUrl(tt.arg); got != tt.want {
				t.Errorf("IsValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
