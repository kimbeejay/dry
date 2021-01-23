package flag

type flags int

func (f *flags) HasFlag(flag int) bool {
	return f.Int()&flag == flag
}

func (f *flags) SetFlag(flag int) {
	*f = flags(f.Int() | flag)
}

func (f *flags) RemoveFlag(flag int) {
	*f = flags(f.Int() & ^flag)
}

func (f *flags) Int() int {
	return int(*f)
}
