package gb

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/types"
)

func (g *GameBoy) PressButton(button types.ButtonType) {
	g.input.Press(button)
}

func (g *GameBoy) ReleaseButton(button types.ButtonType) {
	g.input.Release(button)
	g.interruptController.Request(addr.InterruptJoypad)
}
