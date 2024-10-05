package gb

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/types"
)

func (g *GameBoy) PressButton(button types.ButtonType) {
	g.Input.Press(button)
}

func (g *GameBoy) ReleaseButton(button types.ButtonType) {
	g.Input.Release(button)
	g.MMU.RequestInterrupt(addr.InterruptJoypad)
}
