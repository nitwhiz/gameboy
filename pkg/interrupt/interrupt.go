package interrupt

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/types"
)

func GetISR(i types.InterruptType) uint16 {
	switch i {
	case addr.InterruptVBlank:
		return addr.ISRVBlank
	case addr.InterruptLCD:
		return addr.ISRLCD
	case addr.InterruptTimer:
		return addr.ISRTimer
	case addr.InterruptSerial:
		return addr.ISRSerial
	case addr.InterruptJoypad:
		return addr.ISRJoypad
	default:
		return 0x00
	}
}

// todo: use this to request and check for interrupts

type Controller struct {
	flag   byte
	enable byte
}

func NewController() *Controller {
	return &Controller{
		flag:   0,
		enable: 0,
	}
}

func (c *Controller) Request(i types.InterruptType) {
	c.flag = bits.Set(c.flag, byte(i))
}

func (c *Controller) IsRequested(i types.InterruptType) bool {
	return bits.Val(c.flag, byte(i)) != 0
}

func (c *Controller) Flush(i types.InterruptType) {
	c.flag = bits.Reset(c.flag, byte(i))
}

func (c *Controller) IF() byte {
	return c.flag
}

func (c *Controller) IE() byte {
	return c.enable
}

func (c *Controller) SetIF(flag byte) {
	c.flag = flag
}

func (c *Controller) SetIE(enable byte) {
	c.enable = enable
}
