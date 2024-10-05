package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/types"
)

type CPU struct {
	af types.Register
	bc types.Register
	de types.Register
	hl types.Register
	sp types.Register
	pc types.Register

	ime  bool
	halt bool
}

func New() *CPU {
	c := CPU{
		af: NewAFRegister(0x01B0),
		bc: NewRegister(0x0013),
		de: NewRegister(0x00D8),
		hl: NewRegister(0x014D),
		sp: NewRegister(0xFFFE),
		pc: NewRegister(0x0100),

		ime:  false,
		halt: false,
	}

	return &c
}

func (c *CPU) AF() types.Register {
	return c.af
}

func (c *CPU) BC() types.Register {
	return c.bc
}

func (c *CPU) DE() types.Register {
	return c.de
}

func (c *CPU) HL() types.Register {
	return c.hl
}

func (c *CPU) SP() types.Register {
	return c.sp
}

func (c *CPU) PC() types.Register {
	return c.pc
}

func (c *CPU) IME() bool {
	return c.ime
}

func (c *CPU) SetIME(ime bool) {
	c.ime = ime
}

func (c *CPU) Halt() bool {
	return c.halt
}

func (c *CPU) SetHalt(halt bool) {
	c.halt = halt
}

func (c *CPU) SetFlag(flag types.Flag, v bool) {
	if v {
		c.AF().SetLo(bits.Set(c.AF().Lo(), byte(flag)))
	} else {
		c.AF().SetLo(bits.Reset(c.AF().Lo(), byte(flag)))
	}
}

func (c *CPU) Flag(flag types.Flag) bool {
	return bits.Test(c.AF().Lo(), byte(flag))
}

func (c *CPU) IncPC() uint16 {
	v := c.pc.Val()
	c.pc.Set(v + 1)

	return v
}
