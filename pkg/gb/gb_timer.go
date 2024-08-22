package gb

import (
	"errors"
	"fmt"
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/quarz"
)

func (g *GameBoy) TickTimers(ticks int) {
	divWrapCount := g.DIVTimerTicks.Increase(float64(ticks) * quarz.TimerTicksPerCPUTick)

	if divWrapCount > 0 {
		g.MMU.SetDIV(byte(int(g.MMU.Read(addr.DIV)) + divWrapCount))
	}

	tac := g.MMU.Read(addr.TAC)

	if bits.IsTACEnabled(tac) {
		clockSpeed := bits.GetTACClockSelect(tac)
		ticksPerCPUTick, ok := quarz.TACClockTicksPerCPUTick[clockSpeed]

		if !ok {
			panic(errors.New(fmt.Sprintf("missing clock speed: %d", clockSpeed)))
		}

		timaWrapCount := g.TIMATimerTicks.Increase(float64(ticks) * ticksPerCPUTick)

		if timaWrapCount > 0 {
			tima := int(g.MMU.Read(addr.TIMA)) + timaWrapCount

			if tima > 0xFF {
				g.MMU.Write(addr.TIMA, g.MMU.Read(addr.TMA))
				g.IM.Request(interrupt.Timer)
			} else {
				g.MMU.Write(addr.TIMA, byte(tima))
			}
		}
	}
}
