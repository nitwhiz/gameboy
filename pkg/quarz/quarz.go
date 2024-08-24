package quarz

const (
	// CPUSpeed of the game boy in Hz
	CPUSpeed = 4194304
	// TimerSpeed of the timer in Hz
	TimerSpeed = 16384

	FramesPerSecond = 60

	CPUTicksPerFrame = CPUSpeed / FramesPerSecond
)

var TACClockSpeed = map[byte]float64{
	0b00: 4096.0,
	0b01: 262144.0,
	0b10: 65536.0,
	0b11: 16384.0,
}

var TACMask = map[byte]uint16{
	0b00: 1 << 9,
	0b01: 1 << 3,
	0b10: 1 << 5,
	0b11: 1 << 7,
}
