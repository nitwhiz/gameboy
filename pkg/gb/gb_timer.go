package gb

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
)

func (g *GameBoy) TickTimers(ticks int) {
	increaseTima := g.Timer.Update(ticks, g.MMU.Read(addr.TAC))

	if increaseTima > 0 {
		tima := int(g.MMU.Read(addr.TIMA)) + increaseTima

		if tima > 0xFF {
			g.MMU.Write(addr.TIMA, g.MMU.Read(addr.TMA))
			g.IM.Request(interrupt.Timer)
		} else {
			g.MMU.Write(addr.TIMA, byte(tima))
		}
	}
}
