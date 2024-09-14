package memory

type IO interface {
	Val() byte
	Set(v byte)
}

type MappedIO struct {
	UnusedBits byte
	Value      byte
}

func NewMappedIO(address uint16) *MappedIO {
	return &MappedIO{
		UnusedBits: GetUnusedBits(address),
		Value:      0xFF,
	}
}

func (io *MappedIO) Val() byte {
	return io.Value | io.UnusedBits
}

func (io *MappedIO) Set(v byte) {
	io.Value = v
}

type UnmappedIO struct {
	UnusedBits byte
	Value      byte
}

func NewUnmappedIO() *UnmappedIO {
	return &UnmappedIO{}
}

func (*UnmappedIO) Val() byte {
	return 0xFF
}

func (*UnmappedIO) Set(byte) {
}

func NewIO(address uint16) IO {
	if IsUnmapped(address) {
		return NewUnmappedIO()
	}

	return NewMappedIO(address)
}
