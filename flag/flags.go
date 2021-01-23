package flag

type Flags int

func (f *Flags) HasFlag(flag int) bool {
	return f.Int()&flag == flag
}

func (f *Flags) SetFlag(flag int) {
	*f = Flags(f.Int() | flag)
}

func (f *Flags) RemoveFlag(flag int) {
	*f = Flags(f.Int() & ^flag)
}

func (f *Flags) Int() int {
	return int(*f)
}
