package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func instAdd(g types.GameBoy, val byte, carry bool) {
	c := g.CPU()

	a := c.AF().Hi()
	r := int16(a) + int16(val)
	cv := int16(a&0x0F) + int16(val&0xF)

	if carry && c.Flag(types.FlagC) {
		r += 1
		cv += 1
	}

	result := byte(r)

	c.AF().SetHi(result)

	c.SetFlag(types.FlagZ, result == 0)
	c.SetFlag(types.FlagN, false)
	c.SetFlag(types.FlagH, cv > 0x0F)
	c.SetFlag(types.FlagC, r > 0xFF)
}

func instAdd16Signed(g types.GameBoy, dst types.Register, src types.Register, s8 int8) {
	c := g.CPU()

	v := src.Val()

	g.Cycle()

	r := uint16(int32(v) + int32(s8))

	dst.Set(r)

	o := v ^ uint16(s8) ^ r

	c.SetFlag(types.FlagZ, false)
	c.SetFlag(types.FlagN, false)
	c.SetFlag(types.FlagH, (o&0x0010) == 0x0010)
	c.SetFlag(types.FlagC, (o&0x0100) == 0x0100)
}

func instAdd16Signed2(g types.GameBoy, dst types.Register, src types.Register, s8 int8) {
	c := g.CPU()

	v := src.Val()

	// this is the same as `instAdd16Signed`, but for some reason `ADD SP, e8` has 1 more cycles
	g.Cycle()
	g.Cycle()

	r := uint16(int32(v) + int32(s8))

	dst.Set(r)

	o := v ^ uint16(s8) ^ r

	c.SetFlag(types.FlagZ, false)
	c.SetFlag(types.FlagN, false)
	c.SetFlag(types.FlagH, (o&0x0010) == 0x0010)
	c.SetFlag(types.FlagC, (o&0x0100) == 0x0100)
}

func instIncReg(reg types.Register) {
	reg.Set(reg.Val() + 1)
}

func instInc8(c types.CPU, val byte) byte {
	r := val + 1

	c.SetFlag(types.FlagZ, r == 0)
	c.SetFlag(types.FlagN, false)
	c.SetFlag(types.FlagH, (val&0x0F)+1 > 0x0F)

	return r
}

func instIncRegHi(c types.CPU, reg types.Register) {
	r := instInc8(c, reg.Hi())

	reg.SetHi(r)
}

func instIncRegLo(c types.CPU, reg types.Register) {
	r := instInc8(c, reg.Lo())

	reg.SetLo(r)
}

func instAdd16HL(g types.GameBoy, val uint16) {
	c := g.CPU()
	hl := c.HL().Val()

	g.Cycle()

	r := int32(hl) + int32(val)

	result := uint16(r)

	c.HL().Set(result)

	c.SetFlag(types.FlagN, false)
	c.SetFlag(types.FlagH, int32(hl&0x0FFF) > (r&0x0FFF))
	c.SetFlag(types.FlagC, r > 0xFFFF)
}

func instAnd(c types.CPU, val byte) {
	result := c.AF().Hi() & val

	c.AF().SetHi(result)

	c.SetFlag(types.FlagZ, result == 0)
	c.SetFlag(types.FlagN, false)
	c.SetFlag(types.FlagH, true)
	c.SetFlag(types.FlagC, false)
}

func instCp(c types.CPU, val byte) {
	a := c.AF().Hi()
	result := a - val

	c.SetFlag(types.FlagZ, result == 0)
	c.SetFlag(types.FlagN, true)
	c.SetFlag(types.FlagH, (val&0x0F) > (a&0x0F))
	c.SetFlag(types.FlagC, val > a)
}

func instOr(c types.CPU, val byte) {
	result := c.AF().Hi() | val

	c.AF().SetHi(result)

	c.SetFlag(types.FlagZ, result == 0)
	c.SetFlag(types.FlagN, false)
	c.SetFlag(types.FlagH, false)
	c.SetFlag(types.FlagC, false)
}

func instSub(c types.CPU, val byte, carry bool) {
	a := c.AF().Hi()
	r := int16(a) - int16(val)
	cv := int16(a&0x0F) - int16(val&0xF)

	if carry && c.Flag(types.FlagC) {
		r -= 1
		cv -= 1
	}

	result := byte(r)

	c.AF().SetHi(result)

	c.SetFlag(types.FlagZ, result == 0)
	c.SetFlag(types.FlagN, true)
	c.SetFlag(types.FlagH, cv < 0)
	c.SetFlag(types.FlagC, r < 0)
}

func instDecReg(reg types.Register) {
	reg.Set(reg.Val() - 1)
}

func instDec8(c types.CPU, val byte) byte {
	r := val - 1

	c.SetFlag(types.FlagZ, r == 0)
	c.SetFlag(types.FlagN, true)
	c.SetFlag(types.FlagH, val&0x0F == 0)

	return r
}

func instDecRegHi(c types.CPU, reg types.Register) {
	r := instDec8(c, reg.Hi())

	reg.SetHi(r)
}

func instDecRegLo(c types.CPU, reg types.Register) {
	r := instDec8(c, reg.Lo())

	reg.SetLo(r)
}

func instXor(c types.CPU, val byte) {
	result := c.AF().Hi() ^ val

	c.AF().SetHi(result)

	c.SetFlag(types.FlagZ, result == 0)
	c.SetFlag(types.FlagN, false)
	c.SetFlag(types.FlagH, false)
	c.SetFlag(types.FlagC, false)
}

func instJr(g types.GameBoy, rel byte) {
	g.CPU().PC().Set(uint16(int32(g.CPU().PC().Val()) + int32(int8(rel))))
}

func instJrCond(g types.GameBoy, flag types.Flag, cond bool) {
	rel := g.Fetch8()

	if g.CPU().Flag(flag) == cond {
		instJr(g, rel)
	}
}

func instJp(g types.GameBoy, addr uint16) {
	g.Cycle()
	g.CPU().PC().Set(addr)
}

func instJpCond(g types.GameBoy, flag types.Flag, cond bool) {
	addr := g.Fetch16()

	if g.CPU().Flag(flag) == cond {
		instJp(g, addr)
	}
}

func instCall(g types.GameBoy, addr uint16) {
	g.PushStack(g.CPU().PC().Val())
	g.CPU().PC().Set(addr)
}

func instCallCond(g types.GameBoy, flag types.Flag, cond bool) {
	addr := g.Fetch16()

	if g.CPU().Flag(flag) == cond {
		// todo: oam bug
		instCall(g, addr)
	}
}

func instRet(g types.GameBoy) {
	g.CPU().PC().Set(g.PopStack())
	g.Cycle()
}

func instRetCond(g types.GameBoy, flag types.Flag, cond bool) {
	g.Cycle()

	if g.CPU().Flag(flag) == cond {
		instRet(g)
	}
}

func instPopReg(g types.GameBoy, reg types.Register) {
	reg.Set(g.PopStack())
}

func instPushReg(g types.GameBoy, reg types.Register) {
	// todo: oam bug
	g.PushStack(reg.Val())
}
