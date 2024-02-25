package gb

import (
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
)

func (g *GameBoy) PressButton(button input.Button) {
	g.Input.Press(button)
	g.IM.Request(interrupt.Joypad)
}

func (g *GameBoy) ReleaseButton(button input.Button) {
	g.Input.Release(button)
}
