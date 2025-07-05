package timer

const (
	// cpuSpeed of the game boy in Hz
	cpuSpeed          = 4194304
	fps               = 60
	cpuTicksPerSecond = cpuSpeed / fps
)

func GetTACMask(tacClockSelect byte) uint16 {
	switch tacClockSelect {
	case 0b00:
		return 1 << 9
	case 0b01:
		return 1 << 3
	case 0b10:
		return 1 << 5
	case 0b11:
		return 1 << 7
	default:
		panic("missing tac mask")
	}
}

type Timer struct {
	value       uint16
	tCycleCount int
}

func (t *Timer) GetValue() uint16 {
	return t.value
}

func (t *Timer) SetValue(v uint16) {
	t.value = v
}

func (t *Timer) Inc() {
	t.value += 1
}

func (t *Timer) Tick(n int) {
	t.tCycleCount += n

	for t.tCycleCount >= 4 {
		t.Inc()
		t.tCycleCount -= 4
	}
}

func New() *Timer {
	return &Timer{
		value: 0,
	}
}
