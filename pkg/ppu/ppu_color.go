package ppu

import "github.com/nitwhiz/gameboy/pkg/bits"

const (
	ColorWhite     = byte(0xFF)
	ColorLightGrey = byte(0xAA)
	ColorDarkGrey  = byte(0x55)
	ColorBlack     = byte(0x00)
)

func (p *PPU) getColor(c byte, address uint16) byte {
	palette := p.MMU.Read(address)

	hi := c*2 + 1
	lo := c * 2

	col := (bits.Val(palette, hi) << 1) | bits.Val(palette, lo)

	switch col {
	case 1:
		return ColorLightGrey
	case 2:
		return ColorDarkGrey
	case 3:
		return ColorBlack
	case 0:
		fallthrough
	default:
		return ColorWhite
	}
}
