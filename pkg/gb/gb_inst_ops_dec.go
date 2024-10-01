package gb

func addDECHandlers() {
	// DEC BC
	h.add(0x0B, func(g *GameBoy) (ticks byte) {
		return instDecReg(g.CPU.BC)
	})

	// DEC DE
	h.add(0x1B, func(g *GameBoy) (ticks byte) {
		return instDecReg(g.CPU.DE)
	})

	// DEC HL
	h.add(0x2B, func(g *GameBoy) (ticks byte) {
		return instDecReg(g.CPU.HL)
	})

	// DEC SP
	h.add(0x3B, func(g *GameBoy) (ticks byte) {
		return instDecReg(g.CPU.SP)
	})

	// DEC B
	h.add(0x05, func(g *GameBoy) (ticks byte) {
		return instDecRegHi(g.CPU, g.CPU.BC)
	})

	// DEC C
	h.add(0x0D, func(g *GameBoy) (ticks byte) {
		return instDecRegLo(g.CPU, g.CPU.BC)
	})

	// DEC D
	h.add(0x15, func(g *GameBoy) (ticks byte) {
		return instDecRegHi(g.CPU, g.CPU.DE)
	})

	// DEC E
	h.add(0x1D, func(g *GameBoy) (ticks byte) {
		return instDecRegLo(g.CPU, g.CPU.DE)
	})

	// DEC H
	h.add(0x25, func(g *GameBoy) (ticks byte) {
		return instDecRegHi(g.CPU, g.CPU.HL)
	})

	// DEC L
	h.add(0x2D, func(g *GameBoy) (ticks byte) {
		return instDecRegLo(g.CPU, g.CPU.HL)
	})

	// DEC [HL]
	h.add(0x35, func(g *GameBoy) (ticks byte) {
		t, r := instDec8(g.CPU, g.MMU.Read(g.CPU.HL.Val()))

		g.MMU.Write(g.CPU.HL.Val(), r)

		return t + 8
	})

	// DEC A
	h.add(0x3D, func(g *GameBoy) (ticks byte) {
		return instDecRegHi(g.CPU, g.CPU.AF)
	})
}
