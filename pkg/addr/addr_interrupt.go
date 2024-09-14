package addr

type InterruptType byte

const (
	InterruptVBlank = InterruptType(0)
	InterruptLCD    = InterruptType(1)
	InterruptTimer  = InterruptType(2)
	InterruptSerial = InterruptType(3)
	InterruptJoypad = InterruptType(4)
)
