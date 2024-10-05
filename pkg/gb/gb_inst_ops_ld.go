package gb

func addLDBHandlers() {
	// LD B, B
	h.add(0x40, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.CPU.BC().Hi())
		return 4
	})

	// LD B, C
	h.add(0x41, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.CPU.BC().Lo())
		return 4
	})

	// LD B, D
	h.add(0x42, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.CPU.DE().Hi())
		return 4
	})

	// LD B, E
	h.add(0x43, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.CPU.DE().Lo())
		return 4
	})

	// LD B, H
	h.add(0x44, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.CPU.HL().Hi())
		return 4
	})

	// LD B, L
	h.add(0x45, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.CPU.HL().Lo())
		return 4
	})

	// LD B, [HL]
	h.add(0x46, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.MMU.Read(g.CPU.HL().Val()))
		return 8
	})

	// LD B, A
	h.add(0x47, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.CPU.AF().Hi())
		return 4
	})
}

func addLDCHandlers() {
	// LD C, B
	h.add(0x48, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.CPU.BC().Hi())
		return 4
	})

	// LD C, C
	h.add(0x49, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.CPU.BC().Lo())
		return 4
	})

	// LD C, D
	h.add(0x4A, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.CPU.DE().Hi())
		return 4
	})

	// LD C, E
	h.add(0x4B, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.CPU.DE().Lo())
		return 4
	})

	// LD C, H
	h.add(0x4C, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.CPU.HL().Hi())
		return 4
	})

	// LD C, L
	h.add(0x4D, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.CPU.HL().Lo())
		return 4
	})

	// LD C, [HL]
	h.add(0x4E, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.MMU.Read(g.CPU.HL().Val()))
		return 8
	})

	// LD C, A
	h.add(0x4F, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.CPU.AF().Hi())
		return 4
	})
}

func addLDDHandlers() {
	// LD D, B
	h.add(0x50, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.CPU.BC().Hi())
		return 4
	})

	// LD D, C
	h.add(0x51, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.CPU.BC().Lo())
		return 4
	})

	// LD D, D
	h.add(0x52, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.CPU.DE().Hi())
		return 4
	})

	// LD D, E
	h.add(0x53, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.CPU.DE().Lo())
		return 4
	})

	// LD D, H
	h.add(0x54, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.CPU.HL().Hi())
		return 4
	})

	// LD D, L
	h.add(0x55, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.CPU.HL().Lo())
		return 4
	})

	// LD D, [HL]
	h.add(0x56, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.MMU.Read(g.CPU.HL().Val()))
		return 8
	})

	// LD D, A
	h.add(0x57, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.CPU.AF().Hi())
		return 4
	})
}

func addLDEHandlers() {
	// LD E, B
	h.add(0x58, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.CPU.BC().Hi())
		return 4
	})

	// LD E, C
	h.add(0x59, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.CPU.BC().Lo())
		return 4
	})

	// LD E, D
	h.add(0x5A, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.CPU.DE().Hi())
		return 4
	})

	// LD E, E
	h.add(0x5B, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.CPU.DE().Lo())
		return 4
	})

	// LD E, H
	h.add(0x5C, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.CPU.HL().Hi())
		return 4
	})

	// LD E, L
	h.add(0x5D, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.CPU.HL().Lo())
		return 4
	})

	// LD E, [HL]
	h.add(0x5E, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.MMU.Read(g.CPU.HL().Val()))
		return 8
	})

	// LD E, A
	h.add(0x5F, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.CPU.AF().Hi())
		return 4
	})
}

func addLDHHandlers() {
	// LD H, B
	h.add(0x60, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.CPU.BC().Hi())
		return 4
	})

	// LD H, C
	h.add(0x61, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.CPU.BC().Lo())
		return 4
	})

	// LD H, D
	h.add(0x62, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.CPU.DE().Hi())
		return 4
	})

	// LD H, E
	h.add(0x63, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.CPU.DE().Lo())
		return 4
	})

	// LD H, H
	h.add(0x64, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.CPU.HL().Hi())
		return 4
	})

	// LD H, L
	h.add(0x65, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.CPU.HL().Lo())
		return 4
	})

	// LD H, [HL]
	h.add(0x66, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.MMU.Read(g.CPU.HL().Val()))
		return 8
	})

	// LD H, A
	h.add(0x67, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.CPU.AF().Hi())
		return 4
	})
}

func addLDLHandlers() {
	// LD L, B
	h.add(0x68, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.CPU.BC().Hi())
		return 4
	})

	// LD L, C
	h.add(0x69, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.CPU.BC().Lo())
		return 4
	})

	// LD L, D
	h.add(0x6A, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.CPU.DE().Hi())
		return 4
	})

	// LD L, E
	h.add(0x6B, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.CPU.DE().Lo())
		return 4
	})

	// LD L, H
	h.add(0x6C, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.CPU.HL().Hi())
		return 4
	})

	// LD L, L
	h.add(0x6D, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.CPU.HL().Lo())
		return 4
	})

	// LD L, [HL]
	h.add(0x6E, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.MMU.Read(g.CPU.HL().Val()))
		return 8
	})

	// LD L, A
	h.add(0x6F, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.CPU.AF().Hi())
		return 4
	})
}

func addLDHLHandlers() {
	// LD [HL], B
	h.add(0x70, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.HL().Val(), g.CPU.BC().Hi())
		return 8
	})

	// LD [HL], C
	h.add(0x71, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.HL().Val(), g.CPU.BC().Lo())
		return 8
	})

	// LD [HL], D
	h.add(0x72, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.HL().Val(), g.CPU.DE().Hi())
		return 8
	})

	// LD [HL], E
	h.add(0x73, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.HL().Val(), g.CPU.DE().Lo())
		return 8
	})

	// LD [HL], H
	h.add(0x74, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.HL().Val(), g.CPU.HL().Hi())
		return 8
	})

	// LD [HL], L
	h.add(0x75, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.HL().Val(), g.CPU.HL().Lo())
		return 8
	})

	// LD [HL], A
	h.add(0x77, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.HL().Val(), g.CPU.AF().Hi())
		return 8
	})
}

func addLDAHandlers() {
	// LD A, B
	h.add(0x78, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.CPU.BC().Hi())
		return 4
	})

	// LD A, C
	h.add(0x79, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.CPU.BC().Lo())
		return 4
	})

	// LD A, D
	h.add(0x7A, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.CPU.DE().Hi())
		return 4
	})

	// LD A, E
	h.add(0x7B, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.CPU.DE().Lo())
		return 4
	})

	// LD A, H
	h.add(0x7C, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.CPU.HL().Hi())
		return 4
	})

	// LD A, L
	h.add(0x7D, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.CPU.HL().Lo())
		return 4
	})

	// LD A, [HL]
	h.add(0x7E, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.MMU.Read(g.CPU.HL().Val()))
		return 8
	})

	// LD A, A
	h.add(0x7F, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.CPU.AF().Hi())
		return 4
	})
}

func addLD2Handlers() {
	// LD [BC], A
	h.add(0x02, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.BC().Val(), g.CPU.AF().Hi())
		return 8
	})

	// LD A, [BC]
	h.add(0x0A, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.MMU.Read(g.CPU.BC().Val()))
		return 8
	})

	// LD [DE], A
	h.add(0x12, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.DE().Val(), g.CPU.AF().Hi())
		return 8
	})

	// LD A, [DE]
	h.add(0x1A, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.MMU.Read(g.CPU.DE().Val()))
		return 8
	})

	// LD [HL+], A
	h.add(0x22, func(g *GameBoy) (ticks byte) {
		hl := g.CPU.HL().Val()
		g.MMU.Write(hl, g.CPU.AF().Hi())
		g.CPU.HL().Set(hl + 1)
		return 8
	})

	// LD A, [HL+]
	h.add(0x2A, func(g *GameBoy) (ticks byte) {
		hl := g.CPU.HL().Val()
		g.CPU.AF().SetHi(g.MMU.Read(hl))
		g.CPU.HL().Set(hl + 1)
		return 8
	})

	// LD [HL-], A
	h.add(0x32, func(g *GameBoy) (ticks byte) {
		hl := g.CPU.HL().Val()
		g.MMU.Write(hl, g.CPU.AF().Hi())
		g.CPU.HL().Set(hl - 1)
		return 8
	})

	// LD A, [HL-]
	h.add(0x3A, func(g *GameBoy) (ticks byte) {
		hl := g.CPU.HL().Val()
		g.CPU.AF().SetHi(g.MMU.Read(hl))
		g.CPU.HL().Set(hl - 1)
		return 8
	})
}

func addLDn8Handlers() {
	// LD B, n8
	h.add(0x06, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetHi(g.Fetch8())
		return 8
	})

	// LD C, n8
	h.add(0x0E, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().SetLo(g.Fetch8())
		return 8
	})

	// LD D, n8
	h.add(0x16, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetHi(g.Fetch8())
		return 8
	})

	// LD E, n8
	h.add(0x1E, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().SetLo(g.Fetch8())
		return 8
	})

	// LD H, n8
	h.add(0x26, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetHi(g.Fetch8())
		return 8
	})

	// LD L, n8
	h.add(0x2E, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().SetLo(g.Fetch8())
		return 8
	})

	// LD [HL], n8
	h.add(0x36, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.CPU.HL().Val(), g.Fetch8())
		return 12
	})

	// LD A, n8
	h.add(0x3E, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.Fetch8())
		return 8
	})
}

func addLDHa8Handlers() {
	// LDH [a8], A
	h.add(0xE0, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(0xFF00|uint16(g.Fetch8()), g.CPU.AF().Hi())
		return 12
	})

	// LDH A, [a8]
	h.add(0xF0, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.MMU.Read(0xFF00 | uint16(g.Fetch8())))
		return 12
	})
}

func addLDn16Handlers() {
	// LD BC, n16
	h.add(0x01, func(g *GameBoy) (ticks byte) {
		g.CPU.BC().Set(g.Fetch16())
		return 12
	})

	// LD DE, n16
	h.add(0x11, func(g *GameBoy) (ticks byte) {
		g.CPU.DE().Set(g.Fetch16())
		return 12
	})

	// LD HL, n16
	h.add(0x21, func(g *GameBoy) (ticks byte) {
		g.CPU.HL().Set(g.Fetch16())
		return 12
	})

	// LD SP, n16
	h.add(0x31, func(g *GameBoy) (ticks byte) {
		g.CPU.SP().Set(g.Fetch16())
		return 12
	})
}

func addLDMiscHandlers() {
	// LD [n16], SP
	h.add(0x08, func(g *GameBoy) (ticks byte) {
		address := g.Fetch16()

		g.MMU.Write(address, g.CPU.SP().Lo())
		g.MMU.Write(address+1, g.CPU.SP().Hi())

		return 20
	})

	// LD HL, SP + s8
	h.add(0xF8, func(g *GameBoy) (ticks byte) {
		return instAdd16Signed(g.CPU, g.CPU.HL(), g.CPU.SP(), int8(g.Fetch8()))
	})

	// LD SP, HL
	h.add(0xF9, func(g *GameBoy) (ticks byte) {
		g.CPU.SP().Set(g.CPU.HL().Val())
		return 8
	})

	// LD [C], A
	h.add(0xE2, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(0xFF00+uint16(g.CPU.BC().Lo()), g.CPU.AF().Hi())
		return 8
	})

	// LD A, [C]
	h.add(0xF2, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.MMU.Read(0xFF00 + uint16(g.CPU.BC().Lo())))
		return 8
	})

	// LD [a16], A
	h.add(0xEA, func(g *GameBoy) (ticks byte) {
		g.MMU.Write(g.Fetch16(), g.CPU.AF().Hi())
		return 16
	})

	// LD A, [a16]
	h.add(0xFA, func(g *GameBoy) (ticks byte) {
		g.CPU.AF().SetHi(g.MMU.Read(g.Fetch16()))
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
