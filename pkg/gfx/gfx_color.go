package gfx

import "github.com/nitwhiz/gameboy/pkg/bits"

const (
	ColorWhite     = byte(0xFF)
	ColorLightGrey = byte(0xCC)
	ColorDarkGrey  = byte(0x77)
	ColorBlack     = byte(0x00)
)

func (g *GFX) getColor(c byte, address uint16) byte {
	palette := g.MMU.Read(address)

	hi := (c << 1) | 1
	lo := c << 1

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
