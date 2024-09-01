package gb

import (
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/interrupt_bus"
)

func (g *GameBoy) PressButton(button input.Button) {
	g.Input.Press(button)
}

func (g *GameBoy) ReleaseButton(button input.Button) {
	g.Input.Release(button)
	g.IMBus.Request(interrupt_bus.Joypad)
}
