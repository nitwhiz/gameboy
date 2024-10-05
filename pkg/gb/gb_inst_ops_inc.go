package gb

func addINCHandlers() {
	// INC BC
	h.add(0x03, func(g *GameBoy) (ticks byte) {
		return instIncReg(g.CPU.BC())
	})

	// INC DE
	h.add(0x13, func(g *GameBoy) (ticks byte) {
		return instIncReg(g.CPU.DE())
	})

	// INC HL
	h.add(0x23, func(g *GameBoy) (ticks byte) {
		return instIncReg(g.CPU.HL())
	})

	// INC SP
	h.add(0x33, func(g *GameBoy) (ticks byte) {
		return instIncReg(g.CPU.SP())
	})

	// INC B
	h.add(0x04, func(g *GameBoy) (ticks byte) {
		return instIncRegHi(g.CPU, g.CPU.BC())
	})

	// INC C
	h.add(0x0C, func(g *GameBoy) (ticks byte) {
		return instIncRegLo(g.CPU, g.CPU.BC())
	})

	// INC D
	h.add(0x14, func(g *GameBoy) (ticks byte) {
		return instIncRegHi(g.CPU, g.CPU.DE())
	})

	// INC E
	h.add(0x1C, func(g *GameBoy) (ticks byte) {
		return instIncRegLo(g.CPU, g.CPU.DE())
	})

	// INC H
	h.add(0x24, func(g *GameBoy) (ticks byte) {
		return instIncRegHi(g.CPU, g.CPU.HL())
	})

	// INC L
	h.add(0x2C, func(g *GameBoy) (ticks byte) {
		return instIncRegLo(g.CPU, g.CPU.HL())
	})

	// INC [HL]
	h.add(0x34, func(g *GameBoy) (ticks byte) {
		t, r := instInc8(g.CPU, g.MMU.Read(g.CPU.HL().Val()))

		g.MMU.Write(g.CPU.HL().Val(), r)

		return t + 8
	})

	// INC A
	h.add(0x3C, func(g *GameBoy) (ticks byte) {
		return instIncRegHi(g.CPU, g.CPU.AF())
	})
}
