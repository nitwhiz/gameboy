package gb

import (
	"github.com/nitwhiz/gameboy/pkg/cpu"
)

func addJRHandlers() {
	// JR NZ, e8
	h.add(0x20, func(g *GameBoy) (ticks byte) {
		return instJrCond(g, cpu.Z, false)
	})

	// JR NC, e8
	h.add(0x30, func(g *GameBoy) (ticks byte) {
		return instJrCond(g, cpu.C, false)
	})

	// JR Z, e8
	h.add(0x28, func(g *GameBoy) (ticks byte) {
		return instJrCond(g, cpu.Z, true)
	})

	// JR C, e8
	h.add(0x38, func(g *GameBoy) (ticks byte) {
		return instJrCond(g, cpu.C, true)
	})

	// JR e8
	h.add(0x18, func(g *GameBoy) (ticks byte) {
		return instJr(g, g.Fetch8()) + 4
	})
}

func addJPHandlers() {
	// JP NZ, a16
	h.add(0xC2, func(g *GameBoy) (ticks byte) {
		return instJpCond(g, cpu.Z, false)
	})

	// JP NC, a16
	h.add(0xD2, func(g *GameBoy) (ticks byte) {
		return instJpCond(g, cpu.C, false)
	})

	// JP Z, a16
	h.add(0xCA, func(g *GameBoy) (ticks byte) {
		return instJpCond(g, cpu.Z, true)
	})

	// JP C, a16
	h.add(0xDA, func(g *GameBoy) (ticks byte) {
		return instJpCond(g, cpu.C, true)
	})

	// JP a18
	h.add(0xC3, func(g *GameBoy) (ticks byte) {
		return instJp(g, g.Fetch16()) + 12
	})

	// JP HL
	h.add(0xE9, func(g *GameBoy) (ticks byte) {
		return instJp(g, g.CPU.HL.Val())
	})
}
