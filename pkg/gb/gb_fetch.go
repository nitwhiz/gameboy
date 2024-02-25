package gb

func (g *GameBoy) Fetch8() byte {
	pc := g.CPU.PC.Val()
	g.CPU.PC.Set(pc + 1)

	return g.MMU.Read(pc)
}

func (g *GameBoy) Fetch16() uint16 {
	return uint16(g.Fetch8()) | (uint16(g.Fetch8()) << 8)
}
