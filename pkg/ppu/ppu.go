package ppu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/screen"
)

const (
	// TicksOAMScan - time spent in OAM scan, in ticks
	TicksOAMScan = 80
	// TicksSendPixels - time spent sending pixels to the LCD, in ticks
	TicksSendPixels = 172
	// TicksHBlank - time spent in HBlank, in ticks
	TicksHBlank = 204

	TicksPassedBeforeSendPixels = TicksOAMScan
	TicksPassedBeforeHBlank     = TicksOAMScan + TicksSendPixels

	// deprecated
	TicksPassedBeforeOAMScan = TicksHBlank + TicksSendPixels

	// ScanlineDuration - time to draw a scanline, in ticks
	ScanlineDuration = TicksHBlank + TicksOAMScan + TicksSendPixels

	VisibleScanlineCount = 144
	ScanlineCount        = 153
	ScanlineWidth        = 160
)

type PPUMode byte

const (
	ModeHBlank     = PPUMode(0)
	ModeVBlank     = PPUMode(1)
	ModeOAMScan    = PPUMode(2)
	ModeSendPixels = PPUMode(3)
)

const TILE_BASE_ADDR_UNSIGNED = uint16(0x8000)
const TILE_BASE_ADDR_SIGNED = uint16(0x8800)

const WINDOW_TILE_MAP_ADDR_1 = uint16(0x9C00)
const WINDOW_TILE_MAP_ADDR_0 = uint16(0x9800)

type PPU struct {
	MMU       *mmu.MMU
	Interrupt *interrupt.Manager
	Screen    *screen.Screen

	Ticks int
}

func New(mmu *mmu.MMU, im *interrupt.Manager) *PPU {
	return &PPU{
		MMU:       mmu,
		Interrupt: im,
		Screen:    screen.New(),
		Ticks:     0,
	}
}

func (p *PPU) Update(ticks int) {
	stat := p.MMU.Read(addr.STAT)
	lcdc := p.MMU.Read(addr.LCDC)

	if !bits.IsLCDEnabled(lcdc) {
		p.MMU.SetLY(0)
		p.MMU.Write(addr.STAT, bits.SetPPUMode(stat, byte(ModeVBlank)))
		p.Ticks = 0

		return
	}

	p.Ticks += ticks

	ly := p.MMU.Read(addr.LY)
	mode := PPUMode(bits.GetPPUMode(stat))

	nextMode := ModeOAMScan
	intr := false

	if ly >= VisibleScanlineCount {
		nextMode = ModeVBlank
		intr = bits.Test(stat, bits.STAT_VBLANK_INTERRUPT_ENABLE)
	} else {
		if p.Ticks >= TicksPassedBeforeHBlank {
			nextMode = ModeHBlank
			intr = bits.Test(stat, bits.STAT_HBLANK_INTERRUPT_ENABLE)
		} else if p.Ticks >= TicksPassedBeforeSendPixels {
			nextMode = ModeSendPixels
		} else {
			nextMode = ModeOAMScan
			intr = bits.Test(stat, bits.STAT_OAM_SCAN_INTERRUPT_ENABLE)
		}
	}

	stat = bits.SetPPUMode(stat, byte(nextMode))

	if intr && (mode != nextMode) {
		p.Interrupt.Request(interrupt.LCD)
	}

	if ly == p.MMU.Read(addr.LYC) {
		stat = bits.SetLYCLY(stat, true)

		if bits.Test(stat, bits.STAT_LYCLY_INTERRUPT_ENABLE) {
			p.Interrupt.Request(interrupt.LCD)
		}
	} else {
		stat = bits.SetLYCLY(stat, false)
	}

	p.MMU.Write(addr.STAT, stat)

	if p.Ticks >= ScanlineDuration {
		p.Ticks -= ScanlineDuration

		nextLY := p.MMU.IncLY()

		if nextLY == VisibleScanlineCount {
			p.Interrupt.Request(interrupt.VBlank)
		}

		if nextLY > ScanlineCount {
			p.MMU.SetLY(0)
		}

		if nextLY <= VisibleScanlineCount {
			p.DrawScanline(lcdc, ly)
		}
	}
}

func (p *PPU) RenderBackground(lcdc byte, ly byte) {
	scx := p.MMU.Read(addr.SCX)
	scy := p.MMU.Read(addr.SCY)

	wx := p.MMU.Read(addr.WX) - 7
	wy := p.MMU.Read(addr.WY)

	tileAddr := TILE_BASE_ADDR_SIGNED

	if bits.Test(lcdc, bits.LCDC_TILE_DATA_SELECT) {
		tileAddr = TILE_BASE_ADDR_UNSIGNED
	}

	win := bits.Test(lcdc, bits.LCDC_WINDOW_DISPLAY_ENABLE) && wy <= ly

	bgMapAddr := WINDOW_TILE_MAP_ADDR_0

	if win {
		if bits.Test(lcdc, bits.LCDC_WINDOW_TILE_MAP_SELECT) {
			bgMapAddr = WINDOW_TILE_MAP_ADDR_1
		}
	} else {
		if bits.Test(lcdc, bits.LCDC_BG_TILE_MAP_SELECT) {
			bgMapAddr = WINDOW_TILE_MAP_ADDR_1
		}
	}

	y := byte(0)

	if win {
		y = ly - wy
	} else {
		y = scy + ly
	}

	tileRow := uint16(y/8) * 32

	for px := byte(0); px < ScanlineWidth; px++ {
		x := px + scx

		if win && px >= wx {
			x = px - wx
		}

		tileCol := uint16(x / 8)
		tileLoc := tileAddr
		tileNumAddr := bgMapAddr + tileRow + tileCol

		if tileAddr == TILE_BASE_ADDR_UNSIGNED {
			tileLoc += uint16(p.MMU.Read(tileNumAddr)) * 16
		} else {
			tileLoc += uint16((int(int8(p.MMU.Read(tileNumAddr))) + 128) * 16)
		}

		line := (y % 8) * 2

		d1 := p.MMU.Read(tileLoc + uint16(line))
		d2 := p.MMU.Read(tileLoc + uint16(line) + 1)

		cBit := byte((int(x%8) - 7) * -1)
		cNum := (bits.Val(d2, cBit) << 1) | bits.Val(d1, cBit)

		col := p.getColor(cNum, addr.BGP)

		p.Screen.SetPixel(px, ly, col)
		p.Screen.SetBackground(px, ly, cNum)
	}
}

func (p *PPU) DrawScanline(lcdc byte, ly byte) {
	p.RenderBackground(lcdc, ly)
	// todo: renderSprites
}

func (p *PPU) Update_deprecated(ticks int) {
	lcdEnabled := p.updateLCDStatus()
	p.Ticks += ticks

	nextScanline := false

	if lcdEnabled && p.Ticks > ScanlineDuration {
		nextScanline = true
		// todo: is this correct?
		p.Ticks = 0
	}

	if nextScanline {
		currentLine := p.MMU.IncLY()

		if currentLine > ScanlineCount {
			p.MMU.SetLY(0)
			currentLine = 0
		}

		if currentLine == VisibleScanlineCount {
			p.Interrupt.Request(interrupt.VBlank)
		}
	}
}

func (p *PPU) updateLCDStatus() bool {
	lcdc := p.MMU.Read(addr.LCDC)
	stat := p.MMU.Read(addr.STAT)

	if !bits.IsLCDEnabled(lcdc) {
		p.Ticks = 0
		p.MMU.SetLY(0)
		p.MMU.Write(addr.STAT, bits.SetPPUMode(stat, byte(ModeVBlank)))

		return false
	}

	ly := p.MMU.Read(addr.LY)
	currentMode := PPUMode(bits.GetPPUMode(stat))

	nextMode := currentMode
	reqInterrupt := false

	if ly >= VisibleScanlineCount {
		nextMode = ModeVBlank
		stat = bits.SetPPUMode(stat, byte(ModeVBlank))
		reqInterrupt = bits.IsLCDModeSelect(stat, byte(ModeVBlank))
	} else {
		currentTicks := p.Ticks

		if currentTicks >= TicksPassedBeforeOAMScan {
			nextMode = ModeOAMScan
			stat = bits.SetPPUMode(stat, byte(ModeOAMScan))
			reqInterrupt = bits.IsLCDModeSelect(stat, byte(ModeOAMScan))
		} else if currentTicks >= TicksPassedBeforeSendPixels {
			nextMode = ModeSendPixels
			stat = bits.SetPPUMode(stat, byte(ModeSendPixels))

			if nextMode != currentMode {
				p.renderScanline(ly)
			}
		} else {
			// todo: this might be too early
			nextMode = ModeHBlank
			stat = bits.SetPPUMode(stat, byte(ModeHBlank))
			reqInterrupt = bits.IsLCDModeSelect(stat, byte(ModeHBlank))
		}
	}

	if reqInterrupt && (nextMode != currentMode) {
		p.Interrupt.Request(interrupt.LCD)
	}

	if ly == p.MMU.Read(addr.LYC) {
		stat = bits.SetLYCLY(stat, true)

		if bits.GetLYCSelect(stat) {
			p.Interrupt.Request(interrupt.LCD)
		}
	} else {
		stat = bits.SetLYCLY(stat, false)
	}

	p.MMU.Write(addr.STAT, stat)

	return true
}
