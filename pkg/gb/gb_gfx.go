package gb

func (g *GameBoy) UpdateGFX(ticks int) {
	g.GFX.Update(ticks)
}
