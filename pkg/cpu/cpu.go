package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
)

type Flag byte

const (
	// Z - Zero Flag
	Z = Flag(7)
	// N - Subtraction Flag
	N = Flag(6)
	// H - Half Carry Flag
	H = Flag(5)
	// C - Carry Flag
	C = Flag(4)
)

type CPU struct {
	// AF - Accumulator and Flags
	AF *Register

	// BC - Register
	BC *Register
	// DE - Register
	DE *Register
	// HL - Register
	HL *Register

	// SP - Stack Pointer
	SP *Register
	// PC - Program Counter
	PC *Register

	// IME - IM master enable flag
	IME bool

	// Halt - This cpu is halting
	Halt bool
}

func New() *CPU {
	c := CPU{
		AF: &Register{},

		BC: &Register{},
		DE: &Register{},

		HL: &Register{},

		SP: &Register{},
		PC: &Register{},

		IME: false,
	}

	c.AF.value = 0x01B0

	c.BC.value = 0x0013
	c.DE.value = 0x00D8

	c.HL.value = 0x014D

	c.SP.value = 0xFFFE
	c.PC.value = 0x0100

	c.AF.mask = 0xFFF0

	c.BC.mask = 0xFFFF
	c.DE.mask = 0xFFFF

	c.HL.mask = 0xFFFF

	c.SP.mask = 0xFFFF
	c.PC.mask = 0xFFFF

	return &c
}

func (c *CPU) SetFlag(flag Flag, v bool) {
	if v {
		c.AF.SetLo(bits.Set(c.AF.Lo(), byte(flag)))
	} else {
		c.AF.SetLo(bits.Reset(c.AF.Lo(), byte(flag)))
	}
}

func (c *CPU) Flag(flag Flag) bool {
	return bits.Test(c.AF.Lo(), byte(flag))
}
