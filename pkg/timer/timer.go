package timer

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/types"
)

const (
	// cpuSpeed of the Game Boy in Hz
	cpuSpeed          = 4194304
	fps               = 60
	cpuTicksPerSecond = cpuSpeed / fps
)

func GetTACMask(tacClockSelect byte) uint16 {
	switch tacClockSelect {
	case 0b00:
		return 512
	case 0b01:
		return 8
	case 0b10:
		return 32
	case 0b11:
		return 128
	default:
		panic("missing tac mask")
	}
}

type Timer struct {
	divCounter uint16
	divState   int
	divCycles  int

	// todo: check me in mmu -> move to mmu
	timaState int

	tac  byte
	tima byte
	tma  byte

	interruptController *interrupt.Controller
}

func (t *Timer) GetDiv() byte {
	return byte(t.divCounter >> 8)
}

func (t *Timer) SetDivState(v int) {
	t.divState = v
}

func (t *Timer) DivCycles() int {
	return t.divCycles
}

func (t *Timer) SetDivCycles(v int) {
	t.divCycles = v
}

func (t *Timer) SetDivCounter(v uint16) {
	triggers := t.divCounter & ^v

	if (t.tac&4 != 0) && triggers&GetTACMask(t.tac&3) != 0 {
		t.incTIMA()
	}

	// todo: serial master edge

	t.divCounter = v
}

func (t *Timer) TIMAState() int {
	return t.timaState
}

func (t *Timer) incTIMA() {
	t.tima += 1

	if t.tima == 0 {
		t.tima = t.tma
		t.timaState = types.TimaReloading
	}
}

func (t *Timer) advanceTIMAState() {
	if t.timaState == types.TimaReloaded {
		t.timaState = types.TimaRunning
	} else if t.timaState == types.TimaReloading {
		t.interruptController.Request(addr.InterruptTimer)
		t.timaState = types.TimaReloaded
	}
}

func (t *Timer) Tick(n int) {
	t.divCycles += n

	if t.divCycles <= 0 {
		return
	}

	if t.divState != 1 && t.divState != 2 {
		t.divCycles -= 3

		if t.divCycles <= 0 {
			t.divState = 1
			return
		}
	}

	for {
		t.advanceTIMAState()
		t.SetDivCounter(t.divCounter + 4)

		t.divCycles -= 4

		if t.divCycles <= 0 {
			t.divState = 2
			return
		}
	}
}

func (t *Timer) TIMA() byte {
	return t.tima
}

func (t *Timer) SetTIMA(v byte) {
	t.tima = v
}

func (t *Timer) TAC() byte {
	return t.tac
}

func (t *Timer) emulateTacBehaviour(v byte) {
	if t.tac&4 == 0 {
		return
	}

	currClock := GetTACMask(t.tac & 3)
	newClocks := GetTACMask(v & 3)

	if (t.divCounter & currClock) != 0 {
		if (v&4) == 0 || (t.divCounter&newClocks) == 0 {
			t.incTIMA()
		}
	}
}

func (t *Timer) SetTAC(v byte) {
	t.emulateTacBehaviour(v)
	t.tac = v
}

func (t *Timer) TMA() byte {
	return t.tma
}

func (t *Timer) SetTMA(v byte) {
	t.tma = v
}

func New(interruptController *interrupt.Controller) *Timer {
	return &Timer{
		interruptController: interruptController,
	}
}
