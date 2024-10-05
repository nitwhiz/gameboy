package gb

import (
	"errors"
	"github.com/nitwhiz/gameboy/pkg/cartridge"
)

var ErrMissingMMU = errors.New("game boy is missing a mmu, cannot load rom data")

type GameBoyOption func(g *GameBoy) error

func WithRom(romData []byte) GameBoyOption {
	return func(g *GameBoy) error {
		if g.MMU == nil {
			return ErrMissingMMU
		}

		cart, err := cartridge.New(romData)

		if err != nil {
			return err
		}

		g.MMU.SetCartridge(cart)

		return nil
	}
}

func WithSerialReceiver(receiver func(byte)) GameBoyOption {
	return func(g *GameBoy) error {
		if g.MMU == nil {
			return ErrMissingMMU
		}

		g.MMU.SetSerialReceiver(receiver)

		return nil
	}
}
