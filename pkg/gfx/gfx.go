package gfx

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/quarz"
	"github.com/nitwhiz/gameboy/pkg/screen"
)

const (
	// TicksHBlank - time spent in HBlank, in ticks
	TicksHBlank = 204
	// TicksOAMScan - time spent in OAM scan, in ticks
	TicksOAMScan = 80
	// TicksSendPixels - time spent sending pixels to the LCD, in ticks
	TicksSendPixels = 172

	TicksPassedBeforeOAMScan    = TicksHBlank + TicksSendPixels
	TicksPassedBeforeSendPixels = TicksHBlank

	// ScanlineDuration - time to draw a scanline, in ticks
	ScanlineDuration = TicksHBlank + TicksOAMScan + TicksSendPixels

	VisibleScanlineCount = 144
	ScanlineCount        = 153
)

type PPUMode byte

const (
	ModeHBlank     = PPUMode(0)
	ModeVBlank     = PPUMode(1)
	ModeOAMScan    = PPUMode(2)
	ModeSendPixels = PPUMode(3)
)

type GFX struct {
	MMU       *mmu.MMU
	Interrupt *interrupt.Manager
	Screen    *screen.Screen

	Ticks *quarz.TickCounter[int]
}

func New(mmu *mmu.MMU, im *interrupt.Manager) *GFX {
	return &GFX{
		MMU:       mmu,
		Interrupt: im,
		Screen:    screen.New(),
		Ticks:     quarz.NewTickCounter[int](ScanlineDuration),
	}
}

func (g *GFX) Update(ticks int) {
	lcdEnabled := g.updateLCDStatus()

	nextScanline := false

	if lcdEnabled {
		nextScanline = g.Ticks.Increase(ticks)
	}

	if nextScanline {
		currentLine := g.MMU.IncLY()

		if currentLine > ScanlineCount {
			g.MMU.SetLY(0)
			currentLine = 0
		}

		if currentLine == VisibleScanlineCount {
			g.Interrupt.Request(interrupt.VBlank)
			g.Screen.Blit()
		}
	}
}

func (g *GFX) updateLCDStatus() bool {
	lcdc := g.MMU.Read(addr.LCDC)
	stat := g.MMU.Read(addr.STAT)

	if !bits.IsLCDEnabled(lcdc) {
		g.Ticks.Reset()
		g.MMU.SetLY(0)
		g.MMU.Write(addr.STAT, bits.SetPPUMode(stat, byte(ModeVBlank)))

		return false
	}

	ly := g.MMU.Read(addr.LY)
	currentMode := PPUMode(bits.GetPPUMode(stat))

	nextMode := currentMode
	reqInterrupt := false

	if ly >= VisibleScanlineCount {
		nextMode = ModeVBlank
		stat = bits.SetPPUMode(stat, byte(ModeVBlank))
		reqInterrupt = bits.IsLCDModeSelect(stat, byte(ModeVBlank))
	} else {
		currentTicks := g.Ticks.GetValue()

		if currentTicks >= TicksPassedBeforeOAMScan {
			nextMode = ModeOAMScan
			stat = bits.SetPPUMode(stat, byte(ModeOAMScan))
			reqInterrupt = bits.IsLCDModeSelect(stat, byte(ModeOAMScan))
		} else if currentTicks >= TicksPassedBeforeSendPixels {
			nextMode = ModeSendPixels
			stat = bits.SetPPUMode(stat, byte(ModeSendPixels))

			if nextMode != currentMode {
				g.renderScanline(ly)
			}
		} else {
			nextMode = ModeHBlank
			stat = bits.SetPPUMode(stat, byte(ModeHBlank))
			reqInterrupt = bits.IsLCDModeSelect(stat, byte(ModeHBlank))
		}
	}

	if reqInterrupt && (nextMode != currentMode) {
		g.Interrupt.Request(interrupt.LCD)
	}

	if ly == g.MMU.Read(addr.LYC) {
		stat = bits.SetLYCLY(stat, true)

		if bits.GetLYCSelect(stat) {
			g.Interrupt.Request(interrupt.LCD)
		}
	} else {
		stat = bits.SetLYCLY(stat, false)
	}

	g.MMU.Write(addr.STAT, stat)

	return true
}
