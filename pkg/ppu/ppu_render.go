package ppu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/screen"
	"slices"
)

func (p *PPU) renderBackground(lcdc byte, ly byte) {
	hadWindowPixels := false

	scx := p.MMU.Read(addr.SCX)
	scy := p.MMU.Read(addr.SCY)

	wx := p.MMU.Read(addr.WX) - 7
	wy := p.MMU.Read(addr.WY)

	if ly == wy {
		p.WYLYInFrame = true
	}

	tileDataBaseAddr := TILE_BASE_ADDR_SIGNED

	if bits.Test(lcdc, addr.LCDC_TILE_DATA_SELECT) {
		tileDataBaseAddr = TILE_BASE_ADDR_UNSIGNED
	}

	for tileX := uint16(0); tileX < 20; tileX++ {
		win := bits.Test(lcdc, addr.LCDC_WINDOW_DISPLAY_ENABLE) && p.WYLYInFrame && byte(tileX) >= wx

		if win {
			hadWindowPixels = true
		}

		mapAddr := WINDOW_TILE_MAP_ADDR_0

		if win {
			if bits.Test(lcdc, addr.LCDC_WINDOW_TILE_MAP_SELECT) {
				mapAddr = WINDOW_TILE_MAP_ADDR_1
			}
		} else {
			if bits.Test(lcdc, addr.LCDC_BG_TILE_MAP_SELECT) {
				mapAddr = WINDOW_TILE_MAP_ADDR_1
			}
		}

		mapOffset := uint16(0)

		if win {
			mapOffset += (32 * p.WindowLineCounter / 8) & uint16(0x3FF)
		} else {
			mapOffset += (32 * (uint16((ly+scy)&0xFF) / 8)) & uint16(0x3FF)
		}

		tileMapOffset := mapOffset

		if !win {
			tileMapOffset += (tileX + uint16((scx/8)&0x1F)) & uint16(0x3FF)
		}

		tileNo := p.MMU.Read(mapAddr + tileMapOffset)

		tileDataAddr := tileDataBaseAddr

		if tileDataBaseAddr == TILE_BASE_ADDR_UNSIGNED {
			tileDataAddr += uint16(tileNo) * 16
		} else {
			tileDataAddr = uint16(int32(tileDataAddr) + int32(int16(tileNo)*16))
		}

		tileDataLoAddr := uint16(0)

		if win {
			tileDataLoAddr = tileDataAddr + (2 * (p.WindowLineCounter % 8))
		} else {
			tileDataLoAddr = tileDataAddr + (2 * ((uint16(ly) + uint16(scy)) % 8))
		}

		tileDataLo := p.MMU.Read(tileDataLoAddr)
		tileDataHi := p.MMU.Read(tileDataLoAddr + 1)

		for pixelX := byte(0); pixelX < 8; pixelX++ {
			lo := bits.Val(tileDataLo, pixelX)
			hi := bits.Val(tileDataHi, pixelX)

			colNum := (hi << 1) | lo

			col := p.getColor(colNum, addr.BGP)

			p.Screen.SetBackground(byte(tileX)*8+pixelX, ly, colNum, col)
		}
	}

	if hadWindowPixels {
		p.WindowLineCounter++
	}
}

func (p *PPU) renderSprites(lcdc byte, ly byte) {
	ys := uint32(8)

	if bits.Test(lcdc, addr.LCDC_SPRITE_SIZE) {
		ys = 16
	}

	scanline := uint32(ly)
	spritesLeft := int8(10)

	var objects []uint32

	for s := uint16(0); s < 40; s++ {
		index := s * 4

		y := uint32(p.MMU.Read(addr.MemOAMBegin + index))
		x := uint32(p.MMU.Read(addr.MemOAMBegin + index + 1))

		if x > 0 && scanline+16 >= y && scanline+16 < y+ys {
			skip := false

			for _, o := range objects {
				ox := (o >> 24) & 0xFF

				if ox == x {
					// skip objects with same x, but higher oam index
					skip = true
				}
			}

			if !skip {
				objects = append(objects, (x<<24)|(uint32(index)<<8)|y)
				spritesLeft--

				if spritesLeft <= 0 {
					break
				}
			}
		}
	}

	slices.SortStableFunc(objects, func(a, b uint32) int {
		return int(a) - int(b)
	})

	for _, o := range objects {
		index := uint16((o >> 8) & 0xFFFF)
		x := int32((o>>24)&0xFF) - 8
		y := int32(o&0xFF) - 16

		//		x := int32(p.MMU.Read(addr.MemOAMBegin+index+1)) - 8

		tileLoc := p.MMU.Read(addr.MemOAMBegin + index + 2)

		if ys == 16 {
			// bit 0 is ignored for 8x16
			tileLoc = tileLoc & 0xFE
		}

		attributes := p.MMU.Read(addr.MemOAMBegin + index + 3)

		xFlip, yFlip, bgPriority, pal := bits.OAMAttributes(attributes)

		line := int32(scanline) - y

		if yFlip {
			line = int32(ys) - line - 1
		}

		dAddress := addr.MemVRAMBegin + (uint16(tileLoc) * 16) + uint16(line*2)

		d1 := p.MMU.Read(dAddress)
		d2 := p.MMU.Read(dAddress + 1)

		for tilePixel := 0; tilePixel < 8; tilePixel++ {
			pixel := int16(x) + int16(7-tilePixel)

			if pixel < 0 || pixel >= screen.Width {
				continue
			}

			colBit := byte(tilePixel)

			if xFlip {
				colBit = byte((int8(tilePixel) - 7) * -1)
			}

			colNum := (bits.Val(d2, colBit) << 1) | bits.Val(d1, colBit)

			col := p.getColor(colNum, pal)

			prioByte := byte(0)

			if bgPriority {
				prioByte = 1
			}

			p.Screen.SetSprite(byte(pixel), byte(scanline), prioByte, colNum, col)
		}
	}
}

func (p *PPU) renderScanline(lcdc byte, ly byte) {
	if bits.Test(lcdc, addr.LCDC_BG_WINDOW_ENABLE) {
		p.renderBackground(lcdc, ly)
	}

	if bits.Test(lcdc, addr.LCDC_SPRITE_ENABLE) {
		p.renderSprites(lcdc, ly)
	}

	p.Screen.BlitScanline(ly)
}
