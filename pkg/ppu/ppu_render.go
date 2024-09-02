package ppu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/screen"
	"slices"
)

func (p *PPU) renderBackground(lcdc byte, ly byte) {
	scx := p.MMU.Read(addr.SCX)
	scy := p.MMU.Read(addr.SCY)

	wx := uint16(p.MMU.Read(addr.WX)) - 7
	wy := p.MMU.Read(addr.WY)

	tileDataBaseAddr := TILE_BASE_ADDR_SIGNED

	if bits.Test(lcdc, addr.LCDC_TILE_DATA_SELECT) {
		tileDataBaseAddr = TILE_BASE_ADDR_UNSIGNED
	}

	lywy := ly >= wy

	hadWindowPixels := false

	for pixelX := uint16(0); pixelX < ScanlineWidth; pixelX++ {
		win := bits.Test(lcdc, addr.LCDC_WINDOW_DISPLAY_ENABLE) && lywy && pixelX >= wx

		if win {
			hadWindowPixels = true
		}

		mapAddr := BG_WINDOW_TILE_MAP_ADDR_9800

		if win {
			if bits.Test(lcdc, addr.LCDC_WINDOW_TILE_MAP_SELECT) {
				mapAddr = BG_WINDOW_TILE_MAP_ADDR_9C00
			}
		} else {
			if bits.Test(lcdc, addr.LCDC_BG_TILE_MAP_SELECT) {
				mapAddr = BG_WINDOW_TILE_MAP_ADDR_9C00
			}
		}

		tileCol := pixelX / 8
		mapOffset := uint16(0)

		if win {
			mapOffset = ((p.WindowLineCounter / 8) * 32) + (tileCol & 0x1F)
		} else {
			mapOffset = ((((uint16(ly) + uint16(scy)) & 0xFF) / 8) * 32) + ((uint16(scx/8) + tileCol) & 0x1F)
		}

		mapOffset &= uint16(0x03FF)

		tileNo := p.MMU.Read(mapAddr + mapOffset)

		tileDataAddr := uint16(0)

		if tileDataBaseAddr == TILE_BASE_ADDR_UNSIGNED {
			tileDataAddr = tileDataBaseAddr + (uint16(tileNo) * 16)
		} else {
			tileDataAddr = uint16(int32(tileDataBaseAddr) + int32(int8(tileNo))*16)
		}

		tileDataLoAddr := tileDataAddr

		if win {
			tileDataLoAddr += 2 * (p.WindowLineCounter % 8)
		} else {
			tileDataLoAddr += 2 * ((uint16(ly) + uint16(scy)) % 8)
		}

		tileDataLo := p.MMU.Read(tileDataLoAddr)
		tileDataHi := p.MMU.Read(tileDataLoAddr + 1)

		tileDataColorBit := byte(7 - (pixelX % 8))

		lo := bits.Val(tileDataLo, tileDataColorBit)
		hi := bits.Val(tileDataHi, tileDataColorBit)

		colNum := (hi << 1) | lo

		col := p.getColor(colNum, addr.BGP)

		p.Screen.SetBackground(byte(pixelX), ly, colNum, col)
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
