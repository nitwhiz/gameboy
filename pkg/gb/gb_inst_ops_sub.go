package gb

func addSUBHandlers() {
	// SUB A, B
	h.add(0x90, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.BC.Hi(), false)
	})

	// SUB A, C
	h.add(0x91, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.BC.Lo(), false)
	})

	// SUB A, D
	h.add(0x92, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.DE.Hi(), false)
	})

	// SUB A, E
	h.add(0x93, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.DE.Lo(), false)
	})

	// SUB A, H
	h.add(0x94, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.HL.Hi(), false)
	})

	// SUB A, L
	h.add(0x95, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.HL.Lo(), false)
	})

	// SUB A, [HL]
	h.add(0x96, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.MMU.Read(g.CPU.HL.Val()), false) + 4
	})

	// SUB A, A
	h.add(0x97, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.CPU.AF.Hi(), false)
	})

	// SUB A, n8
	h.add(0xD6, func(g *GameBoy) (ticks byte) {
		return instSub(g.CPU, g.Fetch8(), false) + 4
	})
}
