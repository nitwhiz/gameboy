package gb

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addBitInstructions() {
	// RLCA
	h.add(0x07, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF().Hi()
		result := (a << 1) | (a >> 7)

		g.CPU.AF().SetHi(result)

		g.CPU.SetFlag(types.FlagZ, false)
		g.CPU.SetFlag(types.FlagN, false)
		g.CPU.SetFlag(types.FlagH, false)
		g.CPU.SetFlag(types.FlagC, a > 0x7F)

		return 4
	})

	// RLA
	h.add(0x17, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF().Hi()

		c := byte(0)

		if g.CPU.Flag(types.FlagC) {
			c = 0x01
		}

		r := (a << 1) + c

		g.CPU.AF().SetHi(r)

		g.CPU.SetFlag(types.FlagZ, false)
		g.CPU.SetFlag(types.FlagN, false)
		g.CPU.SetFlag(types.FlagH, false)
		g.CPU.SetFlag(types.FlagC, a > 0x7F)

		return 4
	})

	// RRCA
	h.add(0x0F, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF().Hi()

		a = (a >> 1) | (a&1)<<7

		g.CPU.AF().SetHi(a)

		g.CPU.SetFlag(types.FlagZ, false)
		g.CPU.SetFlag(types.FlagN, false)
		g.CPU.SetFlag(types.FlagH, false)
		g.CPU.SetFlag(types.FlagC, a > 0x7F)

		return 4
	})

	// RRA
	h.add(0x1F, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF().Hi()

		c := byte(0)

		if g.CPU.Flag(types.FlagC) {
			c = 0x80
		}

		r := a>>1 | c

		g.CPU.AF().SetHi(r)

		g.CPU.SetFlag(types.FlagZ, false)
		g.CPU.SetFlag(types.FlagN, false)
		g.CPU.SetFlag(types.FlagH, false)
		g.CPU.SetFlag(types.FlagC, (1&a) == 1)

		return
	})

	// PREFIX
	h.add(0xCB, func(g *GameBoy) (ticks byte) {
		return p.executeNextOpcode(g) + 4
	})

	// DAA
	h.add(0x27, func(g *GameBoy) (ticks byte) {
		a := g.CPU.AF().Hi()

		if !g.CPU.Flag(types.FlagN) {
			if g.CPU.Flag(types.FlagC) || a > 0x99 {
				a += 0x60
				g.CPU.SetFlag(types.FlagC, true)
			}

			if g.CPU.Flag(types.FlagH) || a&0x0F > 0x9 {
				a += 0x06
				g.CPU.SetFlag(types.FlagH, false)
			}
		} else if g.CPU.Flag(types.FlagC) && g.CPU.Flag(types.FlagH) {
			a += 0x9A
			g.CPU.SetFlag(types.FlagH, false)
		} else if g.CPU.Flag(types.FlagC) {
			a += 0xA0
		} else if g.CPU.Flag(types.FlagH) {
			a += 0xFA
			g.CPU.SetFlag(types.FlagH, false)
		}

		g.CPU.SetFlag(types.FlagZ, a == 0)
		g.CPU.AF().SetHi(a)

		return 4
	})

	// CCF
	h.add(0x3F, func(g *GameBoy) (ticks byte) {
		g.CPU.SetFlag(types.FlagC, !g.CPU.Flag(types.FlagC))

		g.CPU.SetFlag(types.FlagN, false)
		g.CPU.SetFlag(types.FlagH, false)

		return 4
	})

	// CPL
	h.add(0x2F, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(^g.CPU.AF().Hi())

		g.CPU.SetFlag(types.FlagN, true)
		g.CPU.SetFlag(types.FlagH, true)

		return 4
	})

	// SCF
	h.add(0x37, func(g *GameBoy) (ticks byte) {
		g.CPU.SetFlag(types.FlagC, true)
		g.CPU.SetFlag(types.FlagN, false)
		g.CPU.SetFlag(types.FlagH, false)

		return 4
	})
}
