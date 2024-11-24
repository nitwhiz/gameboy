package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addADCHandlers() {
	// ADC A, B
	h.add(0x88, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().BC().Hi(), true)
	})

	// ADC A, C
	h.add(0x89, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().BC().Lo(), true)
	})

	// ADC A, D
	h.add(0x8A, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().DE().Hi(), true)
	})

	// ADC A, E
	h.add(0x8B, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().DE().Lo(), true)
	})

	// ADC A, H
	h.add(0x8C, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().HL().Hi(), true)
	})

	// ADC A, L
	h.add(0x8D, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().HL().Lo(), true)
	})

	// ADC A, [HL]
	h.add(0x8E, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.MMU().Read(g.CPU().HL().Val()), true) + 4
	})

	// ADC A, A
	h.add(0x8F, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().AF().Hi(), true)
	})

	// ADC A, n8
	h.add(0xCE, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().Fetch8(), true) + 4
	})
}
