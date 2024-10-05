package cpu

type Register uint16

func NewRegister(v uint16) *Register {
	r := Register(v)
	return &r
}

func (r *Register) Set(v uint16) {
	*r = Register(v)
}

func (r *Register) Val() uint16 {
	return uint16(*r)
}

func (r *Register) SetLo(v byte) {
	r.Set((uint16(*r) & 0xFF00) | uint16(v))
}

func (r *Register) SetHi(v byte) {
	r.Set((uint16(*r) & 0x00FF) | (uint16(v) << 8))
}

func (r *Register) Lo() byte {
	return byte(*r)
}

func (r *Register) Hi() byte {
	return byte(*r >> 8)
}
