package input

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/types"
)

const (
	ButtonA      = types.ButtonType(0)
	ButtonB      = types.ButtonType(1)
	ButtonSelect = types.ButtonType(2)
	ButtonStart  = types.ButtonType(4)

	ButtonRight = types.ButtonType(5)
	ButtonLeft  = types.ButtonType(6)
	ButtonUp    = types.ButtonType(7)
	ButtonDown  = types.ButtonType(8)
)

type State struct {
	abss byte
	dpad byte
}

func NewState() *State {
	return &State{
		abss: 0xFF,
		dpad: 0xFF,
	}
}

func (s *State) Press(button types.ButtonType) {
	if button > ButtonStart {
		s.abss = bits.Reset(s.abss, byte(button))
	} else {
		s.dpad = bits.Reset(s.dpad, byte(button-ButtonRight))
	}
}

func (s *State) Release(button types.ButtonType) {
	if button > ButtonStart {
		s.abss = bits.Set(s.abss, byte(button))
	} else {
		s.dpad = bits.Set(s.dpad, byte(button-ButtonRight))
	}
}

func (s *State) Value(sel types.SelectType) byte {
	if sel == types.InputSelectButtons {
		return s.abss
	}

	return s.dpad
}
