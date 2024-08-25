package quarz

const (
	// CPUSpeed of the game boy in Hz
	CPUSpeed = 4194304

	FramesPerSecond = 60

	CPUTicksPerFrame = CPUSpeed / FramesPerSecond
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
		return 0
	}
}
