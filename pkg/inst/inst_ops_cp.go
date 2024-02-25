package inst

import (
	"github.com/nitwhiz/gameboy/pkg/gb"
)

func addCPHandlers() {
	// CP A, B
	h.add(0xB8, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.CPU.BC.Hi())
	})

	// CP A, C
	h.add(0xB9, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.CPU.BC.Lo())
	})

	// CP A, D
	h.add(0xBA, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.CPU.DE.Hi())
	})

	// CP A, E
	h.add(0xBB, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.CPU.DE.Lo())
	})

	// CP A, H
	h.add(0xBC, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.CPU.HL.Hi())
	})

	// CP A, L
	h.add(0xBD, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.CPU.HL.Lo())
	})

	// CP A, [HL]
	h.add(0xBE, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.MMU.Read(g.CPU.HL.Val())) + 4
	})

	// CP A, A
	h.add(0xBF, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.CPU.AF.Hi())
	})

	// CP A, n8
	h.add(0xFE, func(g *gb.GameBoy) (ticks byte) {
		return instCp(g.CPU, g.Fetch8()) + 4
	})
}
