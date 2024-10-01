package gb

func addPOPHandlers() {
	// POP BC
	h.add(0xC1, func(g *GameBoy) (ticks byte) {
		g.CPU.BC.Set(g.Stack.Pop())
		return 12
	})

	// POP DE
	h.add(0xD1, func(g *GameBoy) (ticks byte) {
		g.CPU.DE.Set(g.Stack.Pop())
		return 12
	})

	// POP HL
	h.add(0xE1, func(g *GameBoy) (ticks byte) {
		g.CPU.HL.Set(g.Stack.Pop())
		return 12
	})

	// POP AF
	h.add(0xF1, func(g *GameBoy) (ticks byte) {
		g.CPU.AF.Set(g.Stack.Pop())
		return 12
	})
}

func addPUSHHandlers() {
	// PUSH BC
	h.add(0xC5, func(g *GameBoy) (ticks byte) {
		g.Stack.Push(g.CPU.BC.Val())
		return 16
	})

	// PUSH DE
	h.add(0xD5, func(g *GameBoy) (ticks byte) {
		g.Stack.Push(g.CPU.DE.Val())
		return 16
	})

	// PUSH HL
	h.add(0xE5, func(g *GameBoy) (ticks byte) {
		g.Stack.Push(g.CPU.HL.Val())
		return 16
	})

	// PUSH AF
	h.add(0xF5, func(g *GameBoy) (ticks byte) {
		g.Stack.Push(g.CPU.AF.Val())
		return 16
	})
}
