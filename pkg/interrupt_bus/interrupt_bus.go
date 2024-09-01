package interrupt_bus

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
)

type Type byte

const (
	VBlank = Type(0)
	LCD    = Type(1)
	Timer  = Type(2)
	Serial = Type(3)
	Joypad = Type(4)
)

// todo: i don't like this solution
type Bus struct {
	IF byte
}

func (b *Bus) Request(t Type) {
	b.IF = bits.Set(b.IF, byte(t))
}
