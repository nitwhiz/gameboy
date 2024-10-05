package cpu

type AFRegister Register

func NewAFRegister(v uint16) *AFRegister {
	r := AFRegister(v)
	return &r
}

func (r *AFRegister) Set(v uint16) {
	*r = AFRegister(v & 0xFFF0)
}

func (r *AFRegister) Val() uint16 {
	return uint16(*r)
}

func (r *AFRegister) SetLo(v byte) {
	r.Set((uint16(*r) & 0xFF00) | uint16(v&0xF0))
}

func (r *AFRegister) SetHi(v byte) {
	r.Set((uint16(*r) & 0x00FF) | (uint16(v) << 8))
}

func (r *AFRegister) Lo() byte {
	return byte(*r)
}

func (r *AFRegister) Hi() byte {
	return byte(*r >> 8)
}
