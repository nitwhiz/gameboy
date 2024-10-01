package gb

func addADDHandlers() {
	// ADD A, B
	h.add(0x80, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.CPU.BC.Hi(), false)
	})

	// ADD A, C
	h.add(0x81, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.CPU.BC.Lo(), false)
	})

	// ADD A, D
	h.add(0x82, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.CPU.DE.Hi(), false)
	})

	// ADD A, E
	h.add(0x83, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.CPU.DE.Lo(), false)
	})

	// ADD A, H
	h.add(0x84, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.CPU.HL.Hi(), false)
	})

	// ADD A, L
	h.add(0x85, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.CPU.HL.Lo(), false)
	})

	// ADD A, [HL]
	h.add(0x86, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.MMU.Read(g.CPU.HL.Val()), false) + 4
	})

	// ADD A, A
	h.add(0x87, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.CPU.AF.Hi(), false)
	})

	// ADD HL, BC
	h.add(0x09, func(g *GameBoy) (ticks byte) {
		return instAdd16HL(g.CPU, g.CPU.BC.Val())
	})

	// ADD HL, DE
	h.add(0x19, func(g *GameBoy) (ticks byte) {
		return instAdd16HL(g.CPU, g.CPU.DE.Val())
	})

	// ADD HL, HL
	h.add(0x29, func(g *GameBoy) (ticks byte) {
		return instAdd16HL(g.CPU, g.CPU.HL.Val())
	})

	// ADD HL, SP
	h.add(0x39, func(g *GameBoy) (ticks byte) {
		return instAdd16HL(g.CPU, g.CPU.SP.Val())
	})

	// ADD A, n8
	h.add(0xC6, func(g *GameBoy) (ticks byte) {
		return instAdd(g.CPU, g.Fetch8(), false) + 4
	})

	// ADD SP, e8
	h.add(0xE8, func(g *GameBoy) (ticks byte) {
		return instAdd16Signed(g.CPU, g.CPU.SP, g.CPU.SP, int8(g.Fetch8())) + 4
	})
}
