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

	mmu types.MMU

	ime     bool
	halting bool
}

func New(mmu types.MMU) *CPU {

	c := CPU{
		af: NewAFRegister(0x01B0),
		bc: NewRegister(0x0013),
		de: NewRegister(0x00D8),
		hl: NewRegister(0x014D),
		sp: NewRegister(0xFFFE),
		pc: NewRegister(0x0100),

		mmu: mmu,

		ime:     false,
		halting: false,
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
	return c.halting
}

func (c *CPU) SetHalt(halt bool) {
	c.halting = halt
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

func (c *CPU) Fetch8() byte {
	pc := c.pc.Val()
	c.pc.Set(pc + 1)

	return c.mmu.Read(pc)
}

func (c *CPU) Fetch16() uint16 {
	pc := c.pc.Val()

	v1 := c.mmu.Read(pc)
	v2 := c.mmu.Read(pc + 1)

	c.pc.Set(pc + 2)

	return uint16(v1) | (uint16(v2) << 8)
}
