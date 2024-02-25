package quarz

const (
	// CPUSpeed of the game boy in Hz
	CPUSpeed = 4194304
	// TimerSpeed of the timer in Hz
	TimerSpeed = 16384

	FramesPerSecond = 60

	CPUTicksPerFrame = CPUSpeed / FramesPerSecond

	TimerTicksPerCPUTick = float64(TimerSpeed) / float64(CPUSpeed)
)

var TACClockTicksPerCPUTick = map[byte]float64{
	0b00: 4096.0 / float64(CPUSpeed),
	0b01: 262144.0 / float64(CPUSpeed),
	0b10: 65536.0 / float64(CPUSpeed),
	0b11: 16384.0 / float64(CPUSpeed),
}
