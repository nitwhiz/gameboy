package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/types"
)

func initPHandlers() {
	getters := [8]func(g types.GameBoy) byte{
		func(g types.GameBoy) byte {
			return g.CPU().BC().Hi()
		},
		func(g types.GameBoy) byte {
			return g.CPU().BC().Lo()
		},
		func(g types.GameBoy) byte {
			return g.CPU().DE().Hi()
		},
		func(g types.GameBoy) byte {
			return g.CPU().DE().Lo()
		},
		func(g types.GameBoy) byte {
			return g.CPU().HL().Hi()
		},
		func(g types.GameBoy) byte {
			return g.CPU().HL().Lo()
		},
		func(g types.GameBoy) byte {
			return g.Read8(g.CPU().HL().Val())
		},
		func(g types.GameBoy) byte {
			return g.CPU().AF().Hi()
		},
	}

	setters := [8]func(g types.GameBoy, val byte){
		func(g types.GameBoy, val byte) {
			g.CPU().BC().SetHi(val)
		},
		func(g types.GameBoy, val byte) {
			g.CPU().BC().SetLo(val)
		},
		func(g types.GameBoy, val byte) {
			g.CPU().DE().SetHi(val)
		},
		func(g types.GameBoy, val byte) {
			g.CPU().DE().SetLo(val)
		},
		func(g types.GameBoy, val byte) {
			g.CPU().HL().SetHi(val)
		},
		func(g types.GameBoy, val byte) {
			g.CPU().HL().SetLo(val)
		},
		func(g types.GameBoy, val byte) {
			g.Write(g.CPU().HL().Val(), val)
		},
		func(g types.GameBoy, val byte) {
			g.CPU().AF().SetHi(val)
		},
	}

	for x := byte(0); x < 8; x++ {
		i := x

		// RLC
		p.add(0x00+i, func(g types.GameBoy) {
			val := getters[i](g)

			carry := val >> 7
			rot := (val<<1)&0xFF | carry

			setters[i](g, rot)

			g.CPU().SetFlag(types.FlagZ, rot == 0)
			g.CPU().SetFlag(types.FlagN, false)
			g.CPU().SetFlag(types.FlagH, false)
			g.CPU().SetFlag(types.FlagC, carry == 1)
		})

		// RRC
		p.add(0x08+i, func(g types.GameBoy) {
			val := getters[i](g)

			carry := val & 1
			rot := (val >> 1) | (carry << 7)

			setters[i](g, rot)

			g.CPU().SetFlag(types.FlagZ, rot == 0)
			g.CPU().SetFlag(types.FlagN, false)
			g.CPU().SetFlag(types.FlagH, false)
			g.CPU().SetFlag(types.FlagC, carry == 1)
		})

		// RL
		p.add(0x10+i, func(g types.GameBoy) {
			val := getters[i](g)

			carry := val >> 7
			oldCarry := byte(0)

			if g.CPU().Flag(types.FlagC) {
				oldCarry = 1
			}

			rot := (val<<1)&0xFF | oldCarry

			setters[i](g, rot)

			g.CPU().SetFlag(types.FlagZ, rot == 0)
			g.CPU().SetFlag(types.FlagN, false)
			g.CPU().SetFlag(types.FlagH, false)
			g.CPU().SetFlag(types.FlagC, carry == 1)
		})

		// RR
		p.add(0x18+i, func(g types.GameBoy) {
			val := getters[i](g)

			carry := val & 1
			oldCarry := byte(0)

			if g.CPU().Flag(types.FlagC) {
				oldCarry = 1
			}

			rot := (val >> 1) | (oldCarry << 7)

			setters[i](g, rot)

			g.CPU().SetFlag(types.FlagZ, rot == 0)
			g.CPU().SetFlag(types.FlagN, false)
			g.CPU().SetFlag(types.FlagH, false)
			g.CPU().SetFlag(types.FlagC, carry == 1)
		})

		// SLA
		p.add(0x20+i, func(g types.GameBoy) {
			val := getters[i](g)

			carry := val >> 7
			rot := (val << 1) & 0xFF

			setters[i](g, rot)

			g.CPU().SetFlag(types.FlagZ, rot == 0)
			g.CPU().SetFlag(types.FlagN, false)
			g.CPU().SetFlag(types.FlagH, false)
			g.CPU().SetFlag(types.FlagC, carry == 1)
		})

		// SRA
		p.add(0x28+i, func(g types.GameBoy) {
			val := getters[i](g)

			rot := (val >> 1) | (val & 0x80)

			setters[i](g, rot)

			g.CPU().SetFlag(types.FlagZ, rot == 0)
			g.CPU().SetFlag(types.FlagN, false)
			g.CPU().SetFlag(types.FlagH, false)
			g.CPU().SetFlag(types.FlagC, val&1 == 1)
		})

		// SWAP
		p.add(0x30+i, func(g types.GameBoy) {
			val := getters[i](g)

			swapped := (val<<4)&0xF0 | (val>>4)&0x0F

			setters[i](g, swapped)

			g.CPU().SetFlag(types.FlagZ, swapped == 0)
			g.CPU().SetFlag(types.FlagN, false)
			g.CPU().SetFlag(types.FlagH, false)
			g.CPU().SetFlag(types.FlagC, false)
		})

		// SRL
		p.add(0x38+i, func(g types.GameBoy) {
			val := getters[i](g)

			carry := val & 1
			rot := val >> 1

			setters[i](g, rot)

			g.CPU().SetFlag(types.FlagZ, rot == 0)
			g.CPU().SetFlag(types.FlagN, false)
			g.CPU().SetFlag(types.FlagH, false)
			g.CPU().SetFlag(types.FlagC, carry == 1)
		})

		for y := byte(0); y < 8; y++ {
			j := y

			// BIT
			p.add(0x40+0x08*j+i, func(g types.GameBoy) {
				val := getters[i](g)

				g.CPU().SetFlag(types.FlagZ, (val>>j)&1 == 0)
				g.CPU().SetFlag(types.FlagN, false)
				g.CPU().SetFlag(types.FlagH, true)
			})

			// RES
			p.add(0x80+0x08*j+i, func(g types.GameBoy) {
				val := getters[i](g)
				setters[i](g, bits.Reset(val, j))
			})

			// SET
			p.add(0xC0+0x08*j+i, func(g types.GameBoy) {
				val := getters[i](g)
				setters[i](g, bits.Set(val, j))
			})
		}
	}
}
