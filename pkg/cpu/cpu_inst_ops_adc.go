package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addADCHandlers() {
	// ADC A, B
	H.add(0x88, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().BC().Hi(), true)
	})

	// ADC A, C
	H.add(0x89, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().BC().Lo(), true)
	})

	// ADC A, D
	H.add(0x8A, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().DE().Hi(), true)
	})

	// ADC A, E
	H.add(0x8B, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().DE().Lo(), true)
	})

	// ADC A, H
	H.add(0x8C, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().HL().Hi(), true)
	})

	// ADC A, L
	H.add(0x8D, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().HL().Lo(), true)
	})

	// ADC A, [HL]
	H.add(0x8E, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.MMU().Read(g.CPU().HL().Val()), true) + 4
	})

	// ADC A, A
	H.add(0x8F, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().AF().Hi(), true)
	})

	// ADC A, n8
	H.add(0xCE, func(g types.GameBoy) (ticks byte) {
		return instAdd(g.CPU(), g.CPU().Fetch8(), true) + 4
	})
}
