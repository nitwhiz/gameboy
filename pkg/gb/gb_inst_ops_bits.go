package gb

import (
	"github.com/nitwhiz/gameboy/pkg/cpu"
)

func addBitInstructions() {
	// RLCA
	h.add(0x07, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF.Hi()
		result := a<<1 | (a >> 7)

		g.CPU.AF.SetHi(result)

		g.CPU.SetFlag(cpu.Z, false)
		g.CPU.SetFlag(cpu.N, false)
		g.CPU.SetFlag(cpu.H, false)
		g.CPU.SetFlag(cpu.C, a > 0x7F)

		return 4
	})

	// RLA
	h.add(0x17, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF.Hi()

		c := byte(0)

		if g.CPU.Flag(cpu.C) {
			c = 0x01
		}

		r := a<<1 + c

		g.CPU.AF.SetHi(r)

		g.CPU.SetFlag(cpu.Z, false)
		g.CPU.SetFlag(cpu.N, false)
		g.CPU.SetFlag(cpu.H, false)
		g.CPU.SetFlag(cpu.C, a > 0x7F)

		return 4
	})

	// RRCA
	h.add(0x0F, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF.Hi()

		a = a>>1 | (a&1)<<7

		g.CPU.AF.SetHi(a)

		g.CPU.SetFlag(cpu.Z, false)
		g.CPU.SetFlag(cpu.N, false)
		g.CPU.SetFlag(cpu.H, false)
		g.CPU.SetFlag(cpu.C, a > 0x7F)

		return 4
	})

	// RRA
	h.add(0x1F, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF.Hi()

		c := byte(0)

		if g.CPU.Flag(cpu.C) {
			c = 0x80
		}

		r := a>>1 | c

		g.CPU.AF.SetHi(r)

		g.CPU.SetFlag(cpu.Z, false)
		g.CPU.SetFlag(cpu.N, false)
		g.CPU.SetFlag(cpu.H, false)
		g.CPU.SetFlag(cpu.C, (1&a) == 1)

		return
	})

	// PREFIX
	h.add(0xCB, func(g *GameBoy) (ticks byte) {
		return p.executeNextOpcode(g) + 4
	})

	// DAA
	h.add(0x27, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF.Hi()

		if !g.CPU.Flag(cpu.N) {
			if g.CPU.Flag(cpu.C) || a > 0x99 {
				a += 0x60
				g.CPU.SetFlag(cpu.C, true)
			}

			if g.CPU.Flag(cpu.H) || a&0x0F > 0x9 {
				a += 0x06
				g.CPU.SetFlag(cpu.H, false)
			}
		} else if g.CPU.Flag(cpu.C) && g.CPU.Flag(cpu.H) {
			a += 0x9A
			g.CPU.SetFlag(cpu.H, false)
		} else if g.CPU.Flag(cpu.C) {
			a += 0xA0
		} else if g.CPU.Flag(cpu.H) {
			a += 0xFA
			g.CPU.SetFlag(cpu.H, false)
		}

		g.CPU.SetFlag(cpu.Z, a == 0)
		g.CPU.AF.SetHi(a)

		return 4
	})

	// CCF
	h.add(0x3F, func(g *GameBoy) (ticks byte) {
		g.CPU.SetFlag(cpu.C, !g.CPU.Flag(cpu.C))

		g.CPU.SetFlag(cpu.N, false)
		g.CPU.SetFlag(cpu.H, false)

		return 4
	})

	// CPL
	h.add(0x2F, func(g *GameBoy) (ticks byte) {
		g.CPU.AF.SetHi(^g.CPU.AF.Hi())

		g.CPU.SetFlag(cpu.N, true)
		g.CPU.SetFlag(cpu.H, true)

		return 4
	})

	// SCF
	h.add(0x37, func(g *GameBoy) (ticks byte) {
		g.CPU.SetFlag(cpu.C, true)
		g.CPU.SetFlag(cpu.N, false)
		g.CPU.SetFlag(cpu.H, false)

		return 4
	})
}
