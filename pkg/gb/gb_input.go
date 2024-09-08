package gb

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/input"
)

func (g *GameBoy) PressButton(button input.Button) {
	g.Input.Press(button)
}

func (g *GameBoy) ReleaseButton(button input.Button) {
	g.Input.Release(button)
	g.MMU.RequestInterrupt(addr.InterruptJoypad)
}
