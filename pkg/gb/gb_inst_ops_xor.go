package gb

func addXORHandlers() {
	// XOR A, B
	h.add(0xA8, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.CPU.BC.Hi())
	})

	// XOR A, C
	h.add(0xA9, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.CPU.BC.Lo())
	})

	// XOR A, D
	h.add(0xAA, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.CPU.DE.Hi())
	})

	// XOR A, E
	h.add(0xAB, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.CPU.DE.Lo())
	})

	// XOR A, H
	h.add(0xAC, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.CPU.HL.Hi())
	})

	// XOR A, L
	h.add(0xAD, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.CPU.HL.Lo())
	})

	// XOR A, [HL]
	h.add(0xAE, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.MMU.Read(g.CPU.HL.Val())) + 4
	})

	// XOR A, A
	h.add(0xAF, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.CPU.AF.Hi())
	})

	// XOR A, n8
	h.add(0xEE, func(g *GameBoy) (ticks byte) {
		return instXor(g.CPU, g.Fetch8()) + 4
	})
}
