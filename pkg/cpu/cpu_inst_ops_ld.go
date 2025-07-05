package cpu

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

func addLDBHandlers() {
	// LD B, B
	H.add(0x40, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.CPU().BC().Hi())
	})

	// LD B, C
	H.add(0x41, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.CPU().BC().Lo())
	})

	// LD B, D
	H.add(0x42, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.CPU().DE().Hi())
	})

	// LD B, E
	H.add(0x43, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.CPU().DE().Lo())
	})

	// LD B, H
	H.add(0x44, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.CPU().HL().Hi())
	})

	// LD B, L
	H.add(0x45, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.CPU().HL().Lo())
	})

	// LD B, [HL]
	H.add(0x46, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.Read8(g.CPU().HL().Val()))
	})

	// LD B, A
	H.add(0x47, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.CPU().AF().Hi())
	})
}

func addLDCHandlers() {
	// LD C, B
	H.add(0x48, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.CPU().BC().Hi())
	})

	// LD C, C
	H.add(0x49, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.CPU().BC().Lo())
	})

	// LD C, D
	H.add(0x4A, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.CPU().DE().Hi())
	})

	// LD C, E
	H.add(0x4B, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.CPU().DE().Lo())
	})

	// LD C, H
	H.add(0x4C, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.CPU().HL().Hi())
	})

	// LD C, L
	H.add(0x4D, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.CPU().HL().Lo())
	})

	// LD C, [HL]
	H.add(0x4E, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.Read8(g.CPU().HL().Val()))
	})

	// LD C, A
	H.add(0x4F, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.CPU().AF().Hi())
	})
}

func addLDDHandlers() {
	// LD D, B
	H.add(0x50, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.CPU().BC().Hi())
	})

	// LD D, C
	H.add(0x51, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.CPU().BC().Lo())
	})

	// LD D, D
	H.add(0x52, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.CPU().DE().Hi())
	})

	// LD D, E
	H.add(0x53, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.CPU().DE().Lo())
	})

	// LD D, H
	H.add(0x54, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.CPU().HL().Hi())
	})

	// LD D, L
	H.add(0x55, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.CPU().HL().Lo())
	})

	// LD D, [HL]
	H.add(0x56, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.Read8(g.CPU().HL().Val()))
	})

	// LD D, A
	H.add(0x57, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.CPU().AF().Hi())
	})
}

func addLDEHandlers() {
	// LD E, B
	H.add(0x58, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.CPU().BC().Hi())
	})

	// LD E, C
	H.add(0x59, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.CPU().BC().Lo())
	})

	// LD E, D
	H.add(0x5A, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.CPU().DE().Hi())
	})

	// LD E, E
	H.add(0x5B, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.CPU().DE().Lo())
	})

	// LD E, H
	H.add(0x5C, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.CPU().HL().Hi())
	})

	// LD E, L
	H.add(0x5D, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.CPU().HL().Lo())
	})

	// LD E, [HL]
	H.add(0x5E, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.Read8(g.CPU().HL().Val()))
	})

	// LD E, A
	H.add(0x5F, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.CPU().AF().Hi())
	})
}

func addLDHHandlers() {
	// LD H, B
	H.add(0x60, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.CPU().BC().Hi())
	})

	// LD H, C
	H.add(0x61, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.CPU().BC().Lo())
	})

	// LD H, D
	H.add(0x62, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.CPU().DE().Hi())
	})

	// LD H, E
	H.add(0x63, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.CPU().DE().Lo())
	})

	// LD H, H
	H.add(0x64, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.CPU().HL().Hi())
	})

	// LD H, L
	H.add(0x65, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.CPU().HL().Lo())
	})

	// LD H, [HL]
	H.add(0x66, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.Read8(g.CPU().HL().Val()))
	})

	// LD H, A
	H.add(0x67, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.CPU().AF().Hi())
	})
}

func addLDLHandlers() {
	// LD L, B
	H.add(0x68, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.CPU().BC().Hi())
	})

	// LD L, C
	H.add(0x69, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.CPU().BC().Lo())
	})

	// LD L, D
	H.add(0x6A, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.CPU().DE().Hi())
	})

	// LD L, E
	H.add(0x6B, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.CPU().DE().Lo())
	})

	// LD L, H
	H.add(0x6C, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.CPU().HL().Hi())
	})

	// LD L, L
	H.add(0x6D, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.CPU().HL().Lo())
	})

	// LD L, [HL]
	H.add(0x6E, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.Read8(g.CPU().HL().Val()))
	})

	// LD L, A
	H.add(0x6F, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.CPU().AF().Hi())
	})
}

func addLDHLHandlers() {
	// LD [HL], B
	H.add(0x70, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), g.CPU().BC().Hi())
	})

	// LD [HL], C
	H.add(0x71, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), g.CPU().BC().Lo())
	})

	// LD [HL], D
	H.add(0x72, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), g.CPU().DE().Hi())
	})

	// LD [HL], E
	H.add(0x73, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), g.CPU().DE().Lo())
	})

	// LD [HL], H
	H.add(0x74, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), g.CPU().HL().Hi())
	})

	// LD [HL], L
	H.add(0x75, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), g.CPU().HL().Lo())
	})

	// LD [HL], A
	H.add(0x77, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), g.CPU().AF().Hi())
	})
}

func addLDAHandlers() {
	// LD A, B
	H.add(0x78, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.CPU().BC().Hi())
	})

	// LD A, C
	H.add(0x79, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.CPU().BC().Lo())
	})

	// LD A, D
	H.add(0x7A, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.CPU().DE().Hi())
	})

	// LD A, E
	H.add(0x7B, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.CPU().DE().Lo())
	})

	// LD A, H
	H.add(0x7C, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.CPU().HL().Hi())
	})

	// LD A, L
	H.add(0x7D, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.CPU().HL().Lo())
	})

	// LD A, [HL]
	H.add(0x7E, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.Read8(g.CPU().HL().Val()))
	})

	// LD A, A
	H.add(0x7F, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.CPU().AF().Hi())
	})
}

func addLD2Handlers() {
	// LD [BC], A
	H.add(0x02, func(g types.GameBoy) {
		g.Write(g.CPU().BC().Val(), g.CPU().AF().Hi())
	})

	// LD A, [BC]
	H.add(0x0A, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.Read8(g.CPU().BC().Val()))
	})

	// LD [DE], A
	H.add(0x12, func(g types.GameBoy) {
		g.Write(g.CPU().DE().Val(), g.CPU().AF().Hi())
	})

	// LD A, [DE]
	H.add(0x1A, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.Read8(g.CPU().DE().Val()))
	})

	// LD [HL+], A
	H.add(0x22, func(g types.GameBoy) {
		hl := g.CPU().HL().Val()
		g.Write(hl, g.CPU().AF().Hi())
		g.CPU().HL().Set(hl + 1)
	})

	// LD A, [HL+]
	H.add(0x2A, func(g types.GameBoy) {
		hl := g.CPU().HL().Val()
		g.CPU().AF().SetHi(g.Read8(hl))
		g.CPU().HL().Set(hl + 1)
	})

	// LD [HL-], A
	H.add(0x32, func(g types.GameBoy) {
		hl := g.CPU().HL().Val()
		g.Write(hl, g.CPU().AF().Hi())
		g.CPU().HL().Set(hl - 1)
	})

	// LD A, [HL-]
	H.add(0x3A, func(g types.GameBoy) {
		hl := g.CPU().HL().Val()
		g.CPU().AF().SetHi(g.Read8(hl))
		g.CPU().HL().Set(hl - 1)
	})
}

func addLDn8Handlers() {
	// LD B, n8
	H.add(0x06, func(g types.GameBoy) {
		g.CPU().BC().SetHi(g.Fetch8())
	})

	// LD C, n8
	H.add(0x0E, func(g types.GameBoy) {
		g.CPU().BC().SetLo(g.Fetch8())
	})

	// LD D, n8
	H.add(0x16, func(g types.GameBoy) {
		g.CPU().DE().SetHi(g.Fetch8())
	})

	// LD E, n8
	H.add(0x1E, func(g types.GameBoy) {
		g.CPU().DE().SetLo(g.Fetch8())
	})

	// LD H, n8
	H.add(0x26, func(g types.GameBoy) {
		g.CPU().HL().SetHi(g.Fetch8())
	})

	// LD L, n8
	H.add(0x2E, func(g types.GameBoy) {
		g.CPU().HL().SetLo(g.Fetch8())
	})

	// LD [HL], n8
	H.add(0x36, func(g types.GameBoy) {
		g.Write(g.CPU().HL().Val(), g.Fetch8())
	})

	// LD A, n8
	H.add(0x3E, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.Fetch8())
	})
}

func addLDHa8Handlers() {
	// LDH [a8], A
	H.add(0xE0, func(g types.GameBoy) {
		g.Write(0xFF00|uint16(g.Fetch8()), g.CPU().AF().Hi())
	})

	// LDH A, [a8]
	H.add(0xF0, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.Read8(0xFF00 | uint16(g.Fetch8())))
	})
}

func addLDn16Handlers() {
	// LD BC, n16
	H.add(0x01, func(g types.GameBoy) {
		g.CPU().BC().Set(g.Fetch16())
	})

	// LD DE, n16
	H.add(0x11, func(g types.GameBoy) {
		g.CPU().DE().Set(g.Fetch16())
	})

	// LD HL, n16
	H.add(0x21, func(g types.GameBoy) {
		g.CPU().HL().Set(g.Fetch16())
	})

	// LD SP, n16
	H.add(0x31, func(g types.GameBoy) {
		g.CPU().SP().Set(g.Fetch16())
	})
}

func addLDMiscHandlers() {
	// LD [n16], SP
	H.add(0x08, func(g types.GameBoy) {
		address := g.Fetch16()

		g.Write(address, g.CPU().SP().Lo())
		g.Write(address+1, g.CPU().SP().Hi())
	})

	// LD HL, SP + s8
	H.add(0xF8, func(g types.GameBoy) {
		instAdd16Signed(g, g.CPU().HL(), g.CPU().SP(), int8(g.Fetch8()))
	})

	// LD SP, HL
	H.add(0xF9, func(g types.GameBoy) {
		g.CPU().SP().Set(g.CPU().HL().Val())
		// todo: oam bug
	})

	// LD [C], A
	H.add(0xE2, func(g types.GameBoy) {
		g.Write(0xFF00|uint16(g.CPU().BC().Lo()), g.CPU().AF().Hi())
	})

	// LD A, [C]
	H.add(0xF2, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.Read8(0xFF00 | uint16(g.CPU().BC().Lo())))
	})

	// LD [a16], A
	H.add(0xEA, func(g types.GameBoy) {
		g.Write(g.Fetch16(), g.CPU().AF().Hi())
	})

	// LD A, [a16]
	H.add(0xFA, func(g types.GameBoy) {
		g.CPU().AF().SetHi(g.Read8(g.Fetch16()))
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
