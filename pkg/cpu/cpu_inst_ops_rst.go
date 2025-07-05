package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func AddRSTHandlers() {
	// RST $00
	H.add(0xC7, func(g types.GameBoy) {
		instCall(g, 0x0000)
	})

	// RST $08
	H.add(0xCF, func(g types.GameBoy) {
		instCall(g, 0x0008)
	})

	// RST $10
	H.add(0xD7, func(g types.GameBoy) {
		instCall(g, 0x0010)
	})

	// RST $18
	H.add(0xDF, func(g types.GameBoy) {
		instCall(g, 0x0018)
	})

	// RST $20
	H.add(0xE7, func(g types.GameBoy) {
		instCall(g, 0x0020)
	})

	// RST $28
	H.add(0xEF, func(g types.GameBoy) {
		instCall(g, 0x0028)
	})

	// RST $30
	H.add(0xF7, func(g types.GameBoy) {
		instCall(g, 0x0030)
	})

	// RST $38
	H.add(0xFF, func(g types.GameBoy) {
		instCall(g, 0x0038)
	})
}
