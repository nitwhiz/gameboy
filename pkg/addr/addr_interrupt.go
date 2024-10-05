package addr

import "github.com/nitwhiz/gameboy/pkg/types"

const (
	InterruptVBlank = types.InterruptType(0)
	InterruptLCD    = types.InterruptType(1)
	InterruptTimer  = types.InterruptType(2)
	InterruptSerial = types.InterruptType(3)
	InterruptJoypad = types.InterruptType(4)
)
