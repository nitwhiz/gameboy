package cpu

import "github.com/nitwhiz/gameboy/pkg/types"

func addLDBHandlers() {
	// LD B, B
	H.add(0x40, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.CPU().BC().Hi())
		return 4
	})

	// LD B, C
	H.add(0x41, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.CPU().BC().Lo())
		return 4
	})

	// LD B, D
	H.add(0x42, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.CPU().DE().Hi())
		return 4
	})

	// LD B, E
	H.add(0x43, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.CPU().DE().Lo())
		return 4
	})

	// LD B, H
	H.add(0x44, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.CPU().HL().Hi())
		return 4
	})

	// LD B, L
	H.add(0x45, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.CPU().HL().Lo())
		return 4
	})

	// LD B, [HL]
	H.add(0x46, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.MMU().Read(g.CPU().HL().Val()))
		return 8
	})

	// LD B, A
	H.add(0x47, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.CPU().AF().Hi())
		return 4
	})
}

func addLDCHandlers() {
	// LD C, B
	H.add(0x48, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.CPU().BC().Hi())
		return 4
	})

	// LD C, C
	H.add(0x49, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.CPU().BC().Lo())
		return 4
	})

	// LD C, D
	H.add(0x4A, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.CPU().DE().Hi())
		return 4
	})

	// LD C, E
	H.add(0x4B, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.CPU().DE().Lo())
		return 4
	})

	// LD C, H
	H.add(0x4C, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.CPU().HL().Hi())
		return 4
	})

	// LD C, L
	H.add(0x4D, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.CPU().HL().Lo())
		return 4
	})

	// LD C, [HL]
	H.add(0x4E, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.MMU().Read(g.CPU().HL().Val()))
		return 8
	})

	// LD C, A
	H.add(0x4F, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.CPU().AF().Hi())
		return 4
	})
}

func addLDDHandlers() {
	// LD D, B
	H.add(0x50, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.CPU().BC().Hi())
		return 4
	})

	// LD D, C
	H.add(0x51, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.CPU().BC().Lo())
		return 4
	})

	// LD D, D
	H.add(0x52, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.CPU().DE().Hi())
		return 4
	})

	// LD D, E
	H.add(0x53, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.CPU().DE().Lo())
		return 4
	})

	// LD D, H
	H.add(0x54, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.CPU().HL().Hi())
		return 4
	})

	// LD D, L
	H.add(0x55, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.CPU().HL().Lo())
		return 4
	})

	// LD D, [HL]
	H.add(0x56, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.MMU().Read(g.CPU().HL().Val()))
		return 8
	})

	// LD D, A
	H.add(0x57, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.CPU().AF().Hi())
		return 4
	})
}

func addLDEHandlers() {
	// LD E, B
	H.add(0x58, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.CPU().BC().Hi())
		return 4
	})

	// LD E, C
	H.add(0x59, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.CPU().BC().Lo())
		return 4
	})

	// LD E, D
	H.add(0x5A, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.CPU().DE().Hi())
		return 4
	})

	// LD E, E
	H.add(0x5B, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.CPU().DE().Lo())
		return 4
	})

	// LD E, H
	H.add(0x5C, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.CPU().HL().Hi())
		return 4
	})

	// LD E, L
	H.add(0x5D, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.CPU().HL().Lo())
		return 4
	})

	// LD E, [HL]
	H.add(0x5E, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.MMU().Read(g.CPU().HL().Val()))
		return 8
	})

	// LD E, A
	H.add(0x5F, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.CPU().AF().Hi())
		return 4
	})
}

func addLDHHandlers() {
	// LD H, B
	H.add(0x60, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.CPU().BC().Hi())
		return 4
	})

	// LD H, C
	H.add(0x61, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.CPU().BC().Lo())
		return 4
	})

	// LD H, D
	H.add(0x62, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.CPU().DE().Hi())
		return 4
	})

	// LD H, E
	H.add(0x63, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.CPU().DE().Lo())
		return 4
	})

	// LD H, H
	H.add(0x64, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.CPU().HL().Hi())
		return 4
	})

	// LD H, L
	H.add(0x65, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.CPU().HL().Lo())
		return 4
	})

	// LD H, [HL]
	H.add(0x66, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.MMU().Read(g.CPU().HL().Val()))
		return 8
	})

	// LD H, A
	H.add(0x67, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.CPU().AF().Hi())
		return 4
	})
}

func addLDLHandlers() {
	// LD L, B
	H.add(0x68, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.CPU().BC().Hi())
		return 4
	})

	// LD L, C
	H.add(0x69, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.CPU().BC().Lo())
		return 4
	})

	// LD L, D
	H.add(0x6A, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.CPU().DE().Hi())
		return 4
	})

	// LD L, E
	H.add(0x6B, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.CPU().DE().Lo())
		return 4
	})

	// LD L, H
	H.add(0x6C, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.CPU().HL().Hi())
		return 4
	})

	// LD L, L
	H.add(0x6D, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.CPU().HL().Lo())
		return 4
	})

	// LD L, [HL]
	H.add(0x6E, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.MMU().Read(g.CPU().HL().Val()))
		return 8
	})

	// LD L, A
	H.add(0x6F, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.CPU().AF().Hi())
		return 4
	})
}

func addLDHLHandlers() {
	// LD [HL], B
	H.add(0x70, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().HL().Val(), g.CPU().BC().Hi())
		return 8
	})

	// LD [HL], C
	H.add(0x71, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().HL().Val(), g.CPU().BC().Lo())
		return 8
	})

	// LD [HL], D
	H.add(0x72, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().HL().Val(), g.CPU().DE().Hi())
		return 8
	})

	// LD [HL], E
	H.add(0x73, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().HL().Val(), g.CPU().DE().Lo())
		return 8
	})

	// LD [HL], H
	H.add(0x74, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().HL().Val(), g.CPU().HL().Hi())
		return 8
	})

	// LD [HL], L
	H.add(0x75, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().HL().Val(), g.CPU().HL().Lo())
		return 8
	})

	// LD [HL], A
	H.add(0x77, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().HL().Val(), g.CPU().AF().Hi())
		return 8
	})
}

func addLDAHandlers() {
	// LD A, B
	H.add(0x78, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.CPU().BC().Hi())
		return 4
	})

	// LD A, C
	H.add(0x79, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.CPU().BC().Lo())
		return 4
	})

	// LD A, D
	H.add(0x7A, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.CPU().DE().Hi())
		return 4
	})

	// LD A, E
	H.add(0x7B, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.CPU().DE().Lo())
		return 4
	})

	// LD A, H
	H.add(0x7C, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.CPU().HL().Hi())
		return 4
	})

	// LD A, L
	H.add(0x7D, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.CPU().HL().Lo())
		return 4
	})

	// LD A, [HL]
	H.add(0x7E, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.MMU().Read(g.CPU().HL().Val()))
		return 8
	})

	// LD A, A
	H.add(0x7F, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.CPU().AF().Hi())
		return 4
	})
}

func addLD2Handlers() {
	// LD [BC], A
	H.add(0x02, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().BC().Val(), g.CPU().AF().Hi())
		return 8
	})

	// LD A, [BC]
	H.add(0x0A, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.MMU().Read(g.CPU().BC().Val()))
		return 8
	})

	// LD [DE], A
	H.add(0x12, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().DE().Val(), g.CPU().AF().Hi())
		return 8
	})

	// LD A, [DE]
	H.add(0x1A, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.MMU().Read(g.CPU().DE().Val()))
		return 8
	})

	// LD [HL+], A
	H.add(0x22, func(g types.GameBoy) (ticks byte) {
		hl := g.CPU().HL().Val()
		g.MMU().Write(hl, g.CPU().AF().Hi())
		g.CPU().HL().Set(hl + 1)
		return 8
	})

	// LD A, [HL+]
	H.add(0x2A, func(g types.GameBoy) (ticks byte) {
		hl := g.CPU().HL().Val()
		g.CPU().AF().SetHi(g.MMU().Read(hl))
		g.CPU().HL().Set(hl + 1)
		return 8
	})

	// LD [HL-], A
	H.add(0x32, func(g types.GameBoy) (ticks byte) {
		hl := g.CPU().HL().Val()
		g.MMU().Write(hl, g.CPU().AF().Hi())
		g.CPU().HL().Set(hl - 1)
		return 8
	})

	// LD A, [HL-]
	H.add(0x3A, func(g types.GameBoy) (ticks byte) {
		hl := g.CPU().HL().Val()
		g.CPU().AF().SetHi(g.MMU().Read(hl))
		g.CPU().HL().Set(hl - 1)
		return 8
	})
}

func addLDn8Handlers() {
	// LD B, n8
	H.add(0x06, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetHi(g.CPU().Fetch8())
		return 8
	})

	// LD C, n8
	H.add(0x0E, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().SetLo(g.CPU().Fetch8())
		return 8
	})

	// LD D, n8
	H.add(0x16, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetHi(g.CPU().Fetch8())
		return 8
	})

	// LD E, n8
	H.add(0x1E, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().SetLo(g.CPU().Fetch8())
		return 8
	})

	// LD H, n8
	H.add(0x26, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetHi(g.CPU().Fetch8())
		return 8
	})

	// LD L, n8
	H.add(0x2E, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().SetLo(g.CPU().Fetch8())
		return 8
	})

	// LD [HL], n8
	H.add(0x36, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().HL().Val(), g.CPU().Fetch8())
		return 12
	})

	// LD A, n8
	H.add(0x3E, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.CPU().Fetch8())
		return 8
	})
}

func addLDHa8Handlers() {
	// LDH [a8], A
	H.add(0xE0, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(0xFF00|uint16(g.CPU().Fetch8()), g.CPU().AF().Hi())
		return 12
	})

	// LDH A, [a8]
	H.add(0xF0, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.MMU().Read(0xFF00 | uint16(g.CPU().Fetch8())))
		return 12
	})
}

func addLDn16Handlers() {
	// LD BC, n16
	H.add(0x01, func(g types.GameBoy) (ticks byte) {
		g.CPU().BC().Set(g.CPU().Fetch16())
		return 12
	})

	// LD DE, n16
	H.add(0x11, func(g types.GameBoy) (ticks byte) {
		g.CPU().DE().Set(g.CPU().Fetch16())
		return 12
	})

	// LD HL, n16
	H.add(0x21, func(g types.GameBoy) (ticks byte) {
		g.CPU().HL().Set(g.CPU().Fetch16())
		return 12
	})

	// LD SP, n16
	H.add(0x31, func(g types.GameBoy) (ticks byte) {
		g.CPU().SP().Set(g.CPU().Fetch16())
		return 12
	})
}

func addLDMiscHandlers() {
	// LD [n16], SP
	H.add(0x08, func(g types.GameBoy) (ticks byte) {
		address := g.CPU().Fetch16()

		g.MMU().Write(address, g.CPU().SP().Lo())
		g.MMU().Write(address+1, g.CPU().SP().Hi())

		return 20
	})

	// LD HL, SP + s8
	H.add(0xF8, func(g types.GameBoy) (ticks byte) {
		return instAdd16Signed(g.CPU(), g.CPU().HL(), g.CPU().SP(), int8(g.CPU().Fetch8()))
	})

	// LD SP, HL
	H.add(0xF9, func(g types.GameBoy) (ticks byte) {
		g.CPU().SP().Set(g.CPU().HL().Val())
		return 8
	})

	// LD [C], A
	H.add(0xE2, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(0xFF00+uint16(g.CPU().BC().Lo()), g.CPU().AF().Hi())
		return 8
	})

	// LD A, [C]
	H.add(0xF2, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.MMU().Read(0xFF00 + uint16(g.CPU().BC().Lo())))
		return 8
	})

	// LD [a16], A
	H.add(0xEA, func(g types.GameBoy) (ticks byte) {
		g.MMU().Write(g.CPU().Fetch16(), g.CPU().AF().Hi())
		return 16
	})

	// LD A, [a16]
	H.add(0xFA, func(g types.GameBoy) (ticks byte) {
		g.CPU().AF().SetHi(g.MMU().Read(g.CPU().Fetch16()))
		return 16
	})
}

func addLDHandlers() {
	addLDBHandlers()
	addLDCHandlers()
	addLDDHandlers()
	addLDEHandlers()
	addLDHHandlers()
	addLDLHandlers()
	addLDHLHandlers()
	addLDAHandlers()

	addLD2Handlers()

	addLDn8Handlers()
	addLDHa8Handlers()

	addLDn16Handlers()

	addLDMiscHandlers()
}
