package quarz

const (
	// CPUSpeed of the game boy in Hz
	CPUSpeed = 4194304

	FramesPerSecond = 60

	CPUTicksPerFrame = CPUSpeed / FramesPerSecond
)

var TACMask = map[byte]uint16{
	0b00: 1 << 9,
	0b01: 1 << 3,
	0b10: 1 << 5,
	0b11: 1 << 7,
}
