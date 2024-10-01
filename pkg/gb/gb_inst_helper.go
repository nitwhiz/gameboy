package gb

import (
	"github.com/nitwhiz/gameboy/pkg/cpu"
)

func instAdd(c *cpu.CPU, val byte, carry bool) (ticks byte) {
	a := c.AF.Hi()
	r := int16(a) + int16(val)
	cv := int16(a&0x0F) + int16(val&0xF)

	if carry && c.Flag(cpu.C) {
		r += 1
		cv += 1
	}

	result := byte(r)

	c.AF.SetHi(result)

	c.SetFlag(cpu.Z, result == 0)
	c.SetFlag(cpu.N, false)
	c.SetFlag(cpu.H, cv > 0x0F)
	c.SetFlag(cpu.C, r > 0xFF)

	return 4
}

func instAdd16Signed(c *cpu.CPU, dst *cpu.Register, src *cpu.Register, s8 int8) (ticks byte) {
	v := src.Val()
	r := uint16(int32(v) + int32(s8))

	dst.Set(r)

	o := v ^ uint16(s8) ^ r

	c.SetFlag(cpu.Z, false)
	c.SetFlag(cpu.N, false)
	c.SetFlag(cpu.H, (o&0x0010) == 0x0010)
	c.SetFlag(cpu.C, (o&0x0100) == 0x0100)

	return 12
}

func instIncReg(reg *cpu.Register) (ticks byte) {
	reg.Set(reg.Val() + 1)
	return 8
}

func instInc8(c *cpu.CPU, val byte) (ticks byte, result byte) {
	r := val + 1

	c.SetFlag(cpu.Z, r == 0)
	c.SetFlag(cpu.N, false)
	c.SetFlag(cpu.H, (val&0x0F)+1 > 0x0F)

	return 4, r
}

func instIncRegHi(c *cpu.CPU, reg *cpu.Register) (ticks byte) {
	t, r := instInc8(c, reg.Hi())

	reg.SetHi(r)

	return t
}

func instIncRegLo(c *cpu.CPU, reg *cpu.Register) (ticks byte) {
	t, r := instInc8(c, reg.Lo())

	reg.SetLo(r)

	return t
}

func instAdd16HL(c *cpu.CPU, val uint16) (ticks byte) {
	hl := c.HL.Val()
	r := int32(hl) + int32(val)

	result := uint16(r)

	c.HL.Set(result)

	c.SetFlag(cpu.N, false)
	c.SetFlag(cpu.H, int32(hl&0x0FFF) > (r&0x0FFF))
	c.SetFlag(cpu.C, r > 0xFFFF)

	return 8
}

func instAnd(c *cpu.CPU, val byte) (ticks byte) {
	result := c.AF.Hi() & val

	c.AF.SetHi(result)

	c.SetFlag(cpu.Z, result == 0)
	c.SetFlag(cpu.N, false)
	c.SetFlag(cpu.H, true)
	c.SetFlag(cpu.C, false)

	return 4
}

func instCp(c *cpu.CPU, val byte) (ticks byte) {
	a := c.AF.Hi()
	result := a - val

	c.SetFlag(cpu.Z, result == 0)
	c.SetFlag(cpu.N, true)
	c.SetFlag(cpu.H, (val&0x0F) > (a&0x0F))
	c.SetFlag(cpu.C, val > a)

	return 4
}

func instOr(c *cpu.CPU, val byte) (ticks byte) {
	result := c.AF.Hi() | val

	c.AF.SetHi(result)

	c.SetFlag(cpu.Z, result == 0)
	c.SetFlag(cpu.N, false)
	c.SetFlag(cpu.H, false)
	c.SetFlag(cpu.C, false)

	return 4
}

func instSub(c *cpu.CPU, val byte, carry bool) (ticks byte) {
	a := c.AF.Hi()
	r := int16(a) - int16(val)
	cv := int16(a&0x0F) - int16(val&0xF)

	if carry && c.Flag(cpu.C) {
		r -= 1
		cv -= 1
	}

	result := byte(r)

	c.AF.SetHi(result)

	c.SetFlag(cpu.Z, result == 0)
	c.SetFlag(cpu.N, true)
	c.SetFlag(cpu.H, cv < 0)
	c.SetFlag(cpu.C, r < 0)

	return 4
}

func instDecReg(reg *cpu.Register) (ticks byte) {
	reg.Set(reg.Val() - 1)
	return 8
}

func instDec8(c *cpu.CPU, val byte) (ticks byte, result byte) {
	r := val - 1

	c.SetFlag(cpu.Z, r == 0)
	c.SetFlag(cpu.N, true)
	c.SetFlag(cpu.H, val&0x0F == 0)

	return 4, r
}

func instDecRegHi(c *cpu.CPU, reg *cpu.Register) (ticks byte) {
	t, r := instDec8(c, reg.Hi())

	reg.SetHi(r)

	return t
}

func instDecRegLo(c *cpu.CPU, reg *cpu.Register) (ticks byte) {
	t, r := instDec8(c, reg.Lo())

	reg.SetLo(r)

	return t
}

func instXor(c *cpu.CPU, val byte) (ticks byte) {
	result := c.AF.Hi() ^ val

	c.AF.SetHi(result)

	c.SetFlag(cpu.Z, result == 0)
	c.SetFlag(cpu.N, false)
	c.SetFlag(cpu.H, false)
	c.SetFlag(cpu.C, false)

	return 4
}

func instJr(g *GameBoy, rel byte) (ticks byte) {
	g.CPU.PC.Set(uint16(int32(g.CPU.PC.Val()) + int32(int8(rel))))
	return 8
}

func instJrCond(g *GameBoy, flag cpu.Flag, cond bool) (ticks byte) {
	rel := g.Fetch8()

	if g.CPU.Flag(flag) == cond {
		return instJr(g, rel) + 4
	}

	return 8
}

func instJp(g *GameBoy, addr uint16) (ticks byte) {
	g.CPU.PC.Set(addr)
	return 4
}

func instJpCond(g *GameBoy, flag cpu.Flag, cond bool) (ticks byte) {
	rel := g.Fetch16()

	if g.CPU.Flag(flag) == cond {
		return instJp(g, rel) + 12
	}

	return 12
}

func instCall(g *GameBoy, addr uint16) (ticks byte) {
	g.Stack.Push(g.CPU.PC.Val())
	g.CPU.PC.Set(addr)

	return 16
}

func instCallCond(g *GameBoy, flag cpu.Flag, cond bool) (ticks byte) {
	addr := g.Fetch16()

	if g.CPU.Flag(flag) == cond {
		return instCall(g, addr) + 8
	}

	return 12
}

func instRet(g *GameBoy) (ticks byte) {
	g.CPU.PC.Set(g.Stack.Pop())

	return 8
}

func instRetCond(g *GameBoy, flag cpu.Flag, cond bool) (ticks byte) {
	if g.CPU.Flag(flag) == cond {
		return instRet(g) + 12
	}

	return 8
}
