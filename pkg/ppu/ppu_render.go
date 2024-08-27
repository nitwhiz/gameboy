package ppu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/screen"
)

func getWindowTileMapArea(lcdc byte) uint16 {
	if bits.Test(lcdc, 6) {
		return 0x9C00
	}

	return 0x9800
}

func getBackgroundAndWindowTileDataArea(lcdc byte) uint16 {
	if bits.Test(lcdc, 4) {
		return 0x8000
	}

	return 0x8800
}

func getBackgroundTileMapArea(lcdc byte) uint16 {
	if bits.Test(lcdc, 3) {
		return 0x9C00
	}

	return 0x9800
}

func getPalette(attributes byte) uint16 {
	if bits.Test(attributes, 4) {
		return addr.OBP1
	}

	return addr.OBP0
}

func (p *PPU) renderBackground(lcdc, ly byte) {
	if ly >= screen.Height {
		return
	}

	tileDataAddr := getBackgroundAndWindowTileDataArea(lcdc)

	unsigned := true

	if tileDataAddr == 0x8800 {
		unsigned = false
	}

	sy := p.MMU.Read(addr.SCY)
	sx := p.MMU.Read(addr.SCX)
	wy := p.MMU.Read(addr.WY)
	wx := p.MMU.Read(addr.WX) - 7

	window := bits.IsLCDWindowEnabled(lcdc) && wy <= p.MMU.Read(addr.LY)

	backgroundDataAddr := uint16(0)

	if window {
		backgroundDataAddr = getWindowTileMapArea(lcdc)
	} else {
		backgroundDataAddr = getBackgroundTileMapArea(lcdc)
	}

	y := byte(0)

	if window {
		y = ly - wy
	} else {
		y = sy + ly
	}

	tileRow := uint16(y/8) * 32

	for pixel := byte(0); pixel < screen.Width; pixel++ {
		x := pixel + sx

		if window && pixel >= wx {
			x = pixel - wx
		}

		tileCol := uint16(x / 8)
		tileLoc := tileDataAddr
		tileAddr := backgroundDataAddr + tileRow + tileCol

		if unsigned {
			tileNum := int16(p.MMU.Read(tileAddr))
			tileLoc += uint16(tileNum * 16)
		} else {
			tileNum := int16(int8(p.MMU.Read(tileAddr)))
			tileLoc = uint16(int32(tileLoc) + int32((tileNum+128)*16))
		}

		line := (y % 8) * 2

		d1 := p.MMU.Read(tileLoc + uint16(line))
		d2 := p.MMU.Read(tileLoc + uint16(line) + 1)

		colBit := (int(x%8) - 7) * -1

		colNum := (bits.Val(d2, byte(colBit)) << 1) | bits.Val(d1, byte(colBit))

		col := p.getColor(colNum, addr.BGP)

		p.Screen.SetPixel(pixel, ly, col)
		p.Screen.SetBackground(pixel, ly, colNum)
	}
}

func (p *PPU) renderSprites(lcdc, ly byte) {
	ys := int32(8)

	if bits.IsLCDObjSize8x16(lcdc) {
		ys = 16
	}

	scanline := int32(ly)
	spritesLeft := int8(10)

	for s := uint16(0); s < 40; s++ {
		index := s * 4

		y := int32(p.MMU.Read(addr.MemOAMBegin+index)) - 16

		if scanline < y || scanline >= (y+ys) {
			continue
		}

		if spritesLeft <= 0 {
			break
		}

		spritesLeft--

		x := int32(p.MMU.Read(addr.MemOAMBegin+index+1)) - 8

		tileLoc := p.MMU.Read(addr.MemOAMBegin + index + 2)
		attributes := p.MMU.Read(addr.MemOAMBegin + index + 3)

		xFlip, yFlip, priority := bits.OAMAttributes(attributes)

		pal := getPalette(attributes)

		line := scanline - y

		if yFlip {
			line = ys - line - 1
		}

		dAddress := addr.MemVRAMBegin + (uint16(tileLoc) * 16) + uint16(line*2)

		d1 := p.MMU.Read(dAddress)
		d2 := p.MMU.Read(dAddress + 1)

		for tilePixel := byte(0); tilePixel < 8; tilePixel++ {
			pixel := int16(x) + int16(7-tilePixel)

			if pixel < 0 || pixel >= screen.Width {
				continue
			}

			colBit := tilePixel

			if xFlip {
				colBit = byte(int8(tilePixel-7) * -1)
			}

			colNum := (bits.Val(d2, colBit) << 1) | bits.Val(d1, colBit)

			if colNum == 0 {
				continue
			}

			col := p.getColor(colNum, pal)

			if !priority || p.Screen.GetBackground(byte(pixel), byte(scanline)) == 0 {
				p.Screen.SetPixel(byte(pixel), byte(scanline), col)
			}
		}
	}
}

func (p *PPU) renderScanline(ly byte) {
	lcdc := p.MMU.Read(addr.LCDC)

	if bits.IsLCDBackgroundAndWindowEnabled(lcdc) {
		p.renderBackground(lcdc, ly)
	}

	if bits.IsLCDObjEnabled(lcdc) {
		p.renderSprites(lcdc, ly)
	}
}
