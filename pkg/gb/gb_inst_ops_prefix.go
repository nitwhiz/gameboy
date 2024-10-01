package gb

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/cpu"
)

func initPHandlers() {
	getters := [8]func(g *GameBoy) (result byte, ticks byte){
		func(g *GameBoy) (result byte, ticks byte) {
			return g.CPU.BC.Hi(), 4
		},
		func(g *GameBoy) (result byte, ticks byte) {
			return g.CPU.BC.Lo(), 4
		},
		func(g *GameBoy) (result byte, ticks byte) {
			return g.CPU.DE.Hi(), 4
		},
		func(g *GameBoy) (result byte, ticks byte) {
			return g.CPU.DE.Lo(), 4
		},
		func(g *GameBoy) (result byte, ticks byte) {
			return g.CPU.HL.Hi(), 4
		},
		func(g *GameBoy) (result byte, ticks byte) {
			return g.CPU.HL.Lo(), 4
		},
		func(g *GameBoy) (result byte, ticks byte) {
			return g.MMU.Read(g.CPU.HL.Val()), 8
		},
		func(g *GameBoy) (result byte, ticks byte) {
			return g.CPU.AF.Hi(), 4
		},
	}

	setters := [8]func(g *GameBoy, val byte) (ticks byte){
		func(g *GameBoy, val byte) (ticks byte) {
			g.CPU.BC.SetHi(val)
			return 4
		},
		func(g *GameBoy, val byte) (ticks byte) {
			g.CPU.BC.SetLo(val)
			return 4
		},
		func(g *GameBoy, val byte) (ticks byte) {
			g.CPU.DE.SetHi(val)
			return 4
		},
		func(g *GameBoy, val byte) (ticks byte) {
			g.CPU.DE.SetLo(val)
			return 4
		},
		func(g *GameBoy, val byte) (ticks byte) {
			g.CPU.HL.SetHi(val)
			return 4
		},
		func(g *GameBoy, val byte) (ticks byte) {
			g.CPU.HL.SetLo(val)
			return 4
		},
		func(g *GameBoy, val byte) (ticks byte) {
			g.MMU.Write(g.CPU.HL.Val(), val)
			return 8
		},
		func(g *GameBoy, val byte) (ticks byte) {
			g.CPU.AF.SetHi(val)
			return 4
		},
	}

	for x := byte(0); x < 8; x++ {
		i := x

		// RLC
		p.add(0x00+i, func(g *GameBoy) (ticks byte) {
			val, getTicks := getters[i](g)

			carry := val >> 7
			rot := (val<<1)&0xFF | carry

			setTicks := setters[i](g, rot)

			g.CPU.SetFlag(cpu.Z, rot == 0)
			g.CPU.SetFlag(cpu.N, false)
			g.CPU.SetFlag(cpu.H, false)
			g.CPU.SetFlag(cpu.C, carry == 1)

			return getTicks + setTicks
		})

		// RRC
		p.add(0x08+i, func(g *GameBoy) (ticks byte) {
			val, getTicks := getters[i](g)

			carry := val & 1
			rot := (val >> 1) | (carry << 7)

			setTicks := setters[i](g, rot)

			g.CPU.SetFlag(cpu.Z, rot == 0)
			g.CPU.SetFlag(cpu.N, false)
			g.CPU.SetFlag(cpu.H, false)
			g.CPU.SetFlag(cpu.C, carry == 1)

			return getTicks + setTicks
		})

		// RL
		p.add(0x10+i, func(g *GameBoy) (ticks byte) {
			val, getTicks := getters[i](g)

			carry := val >> 7
			oldCarry := byte(0)

			if g.CPU.Flag(cpu.C) {
				oldCarry = 1
			}

			rot := (val<<1)&0xFF | oldCarry

			setTicks := setters[i](g, rot)

			g.CPU.SetFlag(cpu.Z, rot == 0)
			g.CPU.SetFlag(cpu.N, false)
			g.CPU.SetFlag(cpu.H, false)
			g.CPU.SetFlag(cpu.C, carry == 1)

			return getTicks + setTicks
		})

		// RR
		p.add(0x18+i, func(g *GameBoy) (ticks byte) {
			val, getTicks := getters[i](g)

			carry := val & 1
			oldCarry := byte(0)

			if g.CPU.Flag(cpu.C) {
				oldCarry = 1
			}

			rot := (val >> 1) | (oldCarry << 7)

			setTicks := setters[i](g, rot)

			g.CPU.SetFlag(cpu.Z, rot == 0)
			g.CPU.SetFlag(cpu.N, false)
			g.CPU.SetFlag(cpu.H, false)
			g.CPU.SetFlag(cpu.C, carry == 1)

			return getTicks + setTicks
		})

		// SLA
		p.add(0x20+i, func(g *GameBoy) (ticks byte) {
			val, getTicks := getters[i](g)

			carry := val >> 7
			rot := (val << 1) & 0xFF

			setTicks := setters[i](g, rot)

			g.CPU.SetFlag(cpu.Z, rot == 0)
			g.CPU.SetFlag(cpu.N, false)
			g.CPU.SetFlag(cpu.H, false)
			g.CPU.SetFlag(cpu.C, carry == 1)

			return getTicks + setTicks
		})

		// SRA
		p.add(0x28+i, func(g *GameBoy) (ticks byte) {
			val, getTicks := getters[i](g)

			rot := (val >> 1) | (val & 0x80)

			setTicks := setters[i](g, rot)

			g.CPU.SetFlag(cpu.Z, rot == 0)
			g.CPU.SetFlag(cpu.N, false)
			g.CPU.SetFlag(cpu.H, false)
			g.CPU.SetFlag(cpu.C, val&1 == 1)

			return getTicks + setTicks
		})

		// SWAP
		p.add(0x30+i, func(g *GameBoy) (ticks byte) {
			val, getTicks := getters[i](g)

			swapped := (val<<4)&0xF0 | (val>>4)&0x0F

			setTicks := setters[i](g, swapped)

			g.CPU.SetFlag(cpu.Z, swapped == 0)
			g.CPU.SetFlag(cpu.N, false)
			g.CPU.SetFlag(cpu.H, false)
			g.CPU.SetFlag(cpu.C, false)

			return getTicks + setTicks
		})

		// SRL
		p.add(0x38+i, func(g *GameBoy) (ticks byte) {
			val, getTicks := getters[i](g)

			carry := val & 1
			rot := val >> 1

			setTicks := setters[i](g, rot)

			g.CPU.SetFlag(cpu.Z, rot == 0)
			g.CPU.SetFlag(cpu.N, false)
			g.CPU.SetFlag(cpu.H, false)
			g.CPU.SetFlag(cpu.C, carry == 1)

			return getTicks + setTicks
		})

		for y := byte(0); y < 8; y++ {
			j := y

			// BIT
			p.add(0x40+0x08*j+i, func(g *GameBoy) (ticks byte) {
				val, ticks := getters[i](g)

				g.CPU.SetFlag(cpu.Z, (val>>j)&1 == 0)
				g.CPU.SetFlag(cpu.N, false)
				g.CPU.SetFlag(cpu.H, true)

				return ticks + 4
			})

			// RES
			p.add(0x80+0x08*j+i, func(g *GameBoy) (ticks byte) {
				val, getTicks := getters[i](g)
				setTicks := setters[i](g, bits.Reset(val, j))

				return getTicks + setTicks + 4
			})

			// SET
			p.add(0xC0+0x08*j+i, func(g *GameBoy) (ticks byte) {
				val, getTicks := getters[i](g)
				setTicks := setters[i](g, bits.Set(val, j))

				return getTicks + setTicks + 4
			})
		}
	}
}
