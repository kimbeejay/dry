package flag

import "testing"

func Test_flags_HasFlag(t *testing.T) {
	tests := []struct {
		name string
		f    flags
		arg  int
		want bool
	}{
		{name: "empty", f: flags(0), arg: 1, want: false},
		{name: "simple", f: flags(1 << 0), arg: 1, want: true},
		{name: "simple", f: flags(5), arg: 1, want: true},
		{name: "simple", f: flags(5), arg: 4, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.HasFlag(tt.arg); got != tt.want {
				t.Errorf("HasFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flags_RemoveFlag(t *testing.T) {
	tests := []struct {
		name string
		f    flags
		arg  int
		want int
	}{
		{name: "RemoveFlag", f: flags(0), arg: 1, want: 0},
		{name: "RemoveFlag", f: flags(1), arg: 1, want: 0},
		{name: "RemoveFlag", f: flags(5), arg: 1, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.f.RemoveFlag(tt.arg); tt.f.Int() != tt.want {
				t.Errorf("RemoveFlag() = %v, want %v", tt.f.Int(), tt.want)
			}
		})
	}
}

func Test_flags_SetFlag(t *testing.T) {
	tests := []struct {
		name string
		f    flags
		arg  int
		want int
	}{
		{name: "SetFlag", f: flags(1), arg: 0, want: 1},
		{name: "SetFlag", f: flags(1), arg: 1 << 1, want: 3},
		{name: "SetFlag", f: flags(5), arg: 4, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.f.SetFlag(tt.arg); tt.f.Int() != tt.want {
				t.Errorf("SetFlag() = %v, want %v", tt.f.Int(), tt.want)
			}
		})
	}
}
