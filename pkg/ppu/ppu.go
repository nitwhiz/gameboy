package ppu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/interrupt_bus"
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

const BG_WINDOW_TILE_MAP_ADDR_9C00 = uint16(0x9C00)
const BG_WINDOW_TILE_MAP_ADDR_9800 = uint16(0x9800)

type PPU struct {
	MMU    *mmu.MMU
	IMBus  *interrupt_bus.Bus
	Screen *screen.Screen

	Ticks             int
	WindowLineCounter uint16
}

func New(mmu *mmu.MMU, imbus *interrupt_bus.Bus) *PPU {
	return &PPU{
		MMU:               mmu,
		IMBus:             imbus,
		Screen:            screen.New(),
		Ticks:             0,
		WindowLineCounter: 0,
	}
}

func (p *PPU) Update(ticks int) {
	stat := p.MMU.Read(addr.STAT)
	lcdc := p.MMU.Read(addr.LCDC)

	if !bits.IsLCDEnabled(lcdc) {
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
		intr = bits.Test(stat, addr.STAT_VBLANK_INTERRUPT_ENABLE)
	} else {
		if p.Ticks >= TicksPassedBeforeHBlank {
			nextMode = ModeHBlank
			intr = bits.Test(stat, addr.STAT_HBLANK_INTERRUPT_ENABLE)
		} else if p.Ticks >= TicksPassedBeforeSendPixels {
			nextMode = ModeSendPixels
		} else {
			nextMode = ModeOAMScan
			intr = bits.Test(stat, addr.STAT_OAM_SCAN_INTERRUPT_ENABLE)
		}
	}

	stat = bits.SetPPUMode(stat, byte(nextMode))

	p.MMU.Write(addr.STAT, stat)

	if intr && (mode != nextMode) {
		p.IMBus.Request(interrupt_bus.LCD)
	}

	if p.Ticks >= ScanlineDuration {
		p.Ticks -= ScanlineDuration

		nextLY := p.MMU.IncLY()

		if nextLY > ScanlineCount {
			p.MMU.SetLY(0)
		}

		if nextLY <= VisibleScanlineCount {
			p.renderScanline(lcdc, ly)
		}

		if nextLY == VisibleScanlineCount {
			p.IMBus.Request(interrupt_bus.VBlank)
			p.WindowLineCounter = 0
		}

		p.MMU.CheckLYCLY()
	}
}
