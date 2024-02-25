package gb

func (g *GameBoy) ServiceInterrupts() (ticks int) {
	return g.IM.Service()
}
