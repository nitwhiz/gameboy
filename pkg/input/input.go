package input

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
	"sync"
)

type Button byte

const (
	ButtonA      = Button(0)
	ButtonB      = Button(1)
	ButtonSelect = Button(2)
	ButtonStart  = Button(3)

	ButtonRight = Button(10)
	ButtonLeft  = Button(11)
	ButtonUp    = Button(12)
	ButtonDown  = Button(13)
)

type Select byte

const (
	SelectButtons = Select(0)
	SelectDPad    = Select(10)
)

type ButtonState map[Button]bool

type State struct {
	Buttons ButtonState
	mu      *sync.RWMutex
}

func NewStateFrom(buttonState ButtonState) *State {
	return &State{
		Buttons: buttonState,
		mu:      &sync.RWMutex{},
	}
}

func NewState() *State {
	return &State{
		Buttons: ButtonState{},
		mu:      &sync.RWMutex{},
	}
}

func (s *State) Press(button Button) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Buttons[button] = true
}

func (s *State) Release(button Button) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Buttons[button] = false
}

func (s *State) Get(sel Select) byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res := byte(0b1111)

	for i := Button(0); i < 4; i++ {
		if v, ok := s.Buttons[i+Button(sel)]; ok && v {
			res = bits.Reset(res, byte(i))
		}
	}

	return res
}
