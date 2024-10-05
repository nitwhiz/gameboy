package ppu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/types"
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

	// ScanlineDuration - time to draw a scanline, in ticks
	ScanlineDuration = TicksHBlank + TicksOAMScan + TicksSendPixels

	VisibleScanlineCount = 144
	ScanlineCount        = 153
	ScanlineWidth        = 160
)

const (
	ModeHBlank     = types.PPUMode(0)
	ModeVBlank     = types.PPUMode(1)
	ModeOAMScan    = types.PPUMode(2)
	ModeSendPixels = types.PPUMode(3)
)

const TileBaseAddrUnsigned = uint16(0x8000)
const TileBaseAddrSigned = uint16(0x8800)

const BgWindowTileMapAddr9C00 = uint16(0x9C00)
const BgWindowTileMapAddr9800 = uint16(0x9800)

type PPU struct {
	mmu    types.MMU
	screen types.Screen

	ticks             int
	windowLineCounter uint16
}

func New(mmu types.MMU, s types.Screen) *PPU {
	return &PPU{
		mmu:               mmu,
		screen:            s,
		ticks:             0,
		windowLineCounter: 0,
	}
}

func (p *PPU) Screen() types.Screen {
	return p.screen
}

func (p *PPU) Update(ticks int) {
	stat := p.mmu.Read(addr.STAT)
	lcdc := p.mmu.Read(addr.LCDC)

	if !bits.IsLCDEnabled(lcdc) {
		p.mmu.Write(addr.STAT, bits.SetPPUMode(stat, byte(ModeVBlank)))
		p.ticks = 0

		return
	}

	p.ticks += ticks

	ly := p.mmu.Read(addr.LY)
	mode := types.PPUMode(bits.GetPPUMode(stat))

	nextMode := ModeOAMScan
	intr := false

	if ly >= VisibleScanlineCount {
		nextMode = ModeVBlank
		intr = bits.Test(stat, addr.STAT_VBLANK_INTERRUPT_ENABLE)
	} else {
		if p.ticks >= TicksPassedBeforeHBlank {
			nextMode = ModeHBlank
			intr = bits.Test(stat, addr.STAT_HBLANK_INTERRUPT_ENABLE)
		} else if p.ticks >= TicksPassedBeforeSendPixels {
			nextMode = ModeSendPixels
		} else {
			nextMode = ModeOAMScan
			intr = bits.Test(stat, addr.STAT_OAM_SCAN_INTERRUPT_ENABLE)
		}
	}

	stat = bits.SetPPUMode(stat, byte(nextMode))

	p.mmu.Write(addr.STAT, stat)

	if intr && (mode != nextMode) {
		p.mmu.RequestInterrupt(addr.InterruptLCD)
	}

	if p.ticks >= ScanlineDuration {
		p.ticks -= ScanlineDuration

		nextLY := p.mmu.IncLY()

		if nextLY > ScanlineCount {
			p.mmu.ResetLY()
		}

		if nextLY <= VisibleScanlineCount {
			p.renderScanline(lcdc, ly)
		}

		if nextLY == VisibleScanlineCount {
			p.mmu.RequestInterrupt(addr.InterruptVBlank)
			p.windowLineCounter = 0
		}

		p.mmu.CheckLYCLY()
	}
}
