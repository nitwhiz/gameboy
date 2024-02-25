package inst

import (
	"github.com/nitwhiz/gameboy/pkg/gb"
)

func AddRSTHandlers() {
	// RST $00
	h.add(0xC7, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, 0x0000)
	})

	// RST $08
	h.add(0xCF, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, 0x0008)
	})

	// RST $10
	h.add(0xD7, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, 0x0010)
	})

	// RST $18
	h.add(0xDF, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, 0x0018)
	})

	// RST $20
	h.add(0xE7, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, 0x0020)
	})

	// RST $28
	h.add(0xEF, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, 0x0028)
	})

	// RST $30
	h.add(0xF7, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, 0x0030)
	})

	// RST $38
	h.add(0xFF, func(g *gb.GameBoy) (ticks byte) {
		return instCall(g, 0x0038)
	})
}
