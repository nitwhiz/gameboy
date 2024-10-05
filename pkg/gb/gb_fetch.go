package gb

func (g *GameBoy) Fetch8() byte {
	pc := g.CPU.PC().Val()
	g.CPU.PC().Set(pc + 1)

	return g.MMU.Read(pc)
}

func (g *GameBoy) Fetch16() uint16 {
	pc := g.CPU.PC().Val()

	v1 := g.MMU.Read(pc)
	v2 := g.MMU.Read(pc + 1)

	g.CPU.PC().Set(pc + 2)

	return uint16(v1) | (uint16(v2) << 8)
}
