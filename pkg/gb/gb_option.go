package gb

import (
	"errors"
	"github.com/nitwhiz/gameboy/pkg/cartridge"
)

var ErrMissingMMU = errors.New("game boy is missing a MMU, cannot load rom data")

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

		g.MMU.Cartridge = cart

		return nil
	}
}

func WithSerialReceiver(receiver func(byte)) GameBoyOption {
	return func(g *GameBoy) error {
		if g.MMU == nil {
			return ErrMissingMMU
		}

		g.MMU.SerialReceiver = receiver

		return nil
	}
}

func WithExecuteNextOpcodeFunc(fun ExecuteNextOpcodeFunc) GameBoyOption {
	return func(g *GameBoy) error {
		g.ExecuteNextOpcodeFunc = fun

		return nil
	}
}
