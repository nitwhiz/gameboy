package gb

import (
	"context"
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/memory"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/ppu"
	"github.com/nitwhiz/gameboy/pkg/screen"
	"github.com/nitwhiz/gameboy/pkg/timer"
	"github.com/nitwhiz/gameboy/pkg/types"
	"sync"
)

type GameBoy struct {
	interruptController types.InterruptController
	cpu                 types.CPU
	mmu                 types.MMU

	timer *timer.Timer
	input types.InputState

	ppu types.PPU

	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup

	pendingTicks        int
	allowedPendingTicks int

	halted     bool
	haltBug    bool
	justHalted bool
	stopped    bool

	mu *sync.Mutex
}

func New(ctx context.Context, options ...GameBoyOption) (*GameBoy, error) {
	in := input.NewState()

	i := interrupt.NewController()

	t := timer.New(i)

	m := mmu.New(in, memory.NewContainer(), t, i)

	c := cpu.New(m)

	g := ppu.New(m, screen.New(), i)

	ctx, cancel := context.WithCancel(ctx)

	gameBoy := GameBoy{
		interruptController: i,
		cpu:                 c,
		mmu:                 m,
		timer:               t,
		input:               in,
		ppu:                 g,
		mu:                  &sync.Mutex{},

		ctx:    ctx,
		wg:     &sync.WaitGroup{},
		cancel: cancel,
	}

	for _, o := range options {
		if err := o(&gameBoy); err != nil {
			return nil, err
		}
	}

	t.SetDivCounter(8)

	return &gameBoy, nil
}

func (g *GameBoy) CPU() types.CPU {
	return g.cpu
}

func (g *GameBoy) PushStack(addr uint16) {
	sp := g.CPU().SP().Val()

	g.Write(sp-1, byte(addr>>8))
	g.Write(sp-2, byte(addr))

	g.CPU().SP().Set(sp - 2)
}

func (g *GameBoy) PopStack() uint16 {
	sp := g.CPU().SP().Val()

	v := uint16(g.Read8(sp)) | (uint16(g.Read8(sp+1)) << 8)

	g.CPU().SP().Set(sp + 2)

	return v
}

func (g *GameBoy) PPU() types.PPU {
	return g.ppu
}

func (g *GameBoy) MMU() types.MMU {
	return g.mmu
}

func (g *GameBoy) Input() types.InputState {
	return g.input
}

func (g *GameBoy) Timer() types.Timer {
	return g.timer
}

func (g *GameBoy) Lock() {
	g.mu.Lock()
}

func (g *GameBoy) Unlock() {
	g.mu.Unlock()
}

func (g *GameBoy) Stopped() bool {
	return g.stopped
}

func (g *GameBoy) SetStopped(s bool) {
	g.stopped = s
}

func (g *GameBoy) Halted() bool {
	return g.halted
}

func (g *GameBoy) SetHalted(h bool) {
	g.halted = h
}

func (g *GameBoy) HaltBug() bool {
	return g.haltBug
}

func (g *GameBoy) SetHaltBug(hb bool) {
	g.haltBug = hb
}

func (g *GameBoy) JustHalted() bool {
	return g.justHalted
}

func (g *GameBoy) SetJustHalted(jh bool) {
	g.justHalted = jh
}

func (g *GameBoy) InterruptController() types.InterruptController {
	return g.interruptController
}

func (g *GameBoy) AddPendingTicks(n int) {
	g.pendingTicks += n
}

func (g *GameBoy) SetPendingTicks(n int) {
	g.pendingTicks = n
}

func (g *GameBoy) PendingTicks() int {
	return g.pendingTicks
}

func (g *GameBoy) AdvanceTicks(n int) {
	g.timer.Tick(n)
	g.ppu.Ticks(n)

	// todo: DMA
}

func (g *GameBoy) AdvancePendingTicks() {
	g.AdvanceTicks(g.pendingTicks)
}

func (g *GameBoy) FlushPendingTicks() {
	if g.pendingTicks > 0 {
		g.AdvancePendingTicks()
	}

	g.pendingTicks = 0
}

func (g *GameBoy) Read8(addr uint16) byte {
	if g.pendingTicks > 0 {
		g.AdvanceTicks(g.pendingTicks)
	}

	res := g.mmu.Read(addr)

	g.pendingTicks = 4

	return res
}

func (g *GameBoy) Fetch8() byte {
	pc := g.CPU().PC().Val()
	v := g.Read8(pc)
	g.CPU().PC().Set(pc + 1)

	return v
}

func (g *GameBoy) Read16(addr uint16) uint16 {
	return uint16(g.Read8(addr)) | (uint16(g.Read8(addr+1)) << 8)
}

func (g *GameBoy) Fetch16() uint16 {
	return uint16(g.Fetch8()) | (uint16(g.Fetch8()) << 8)
}

const (
	BehaviourReadOld = iota
	BehaviourWriteCpu
	BehaviourLCDC
	BehaviourReadNew
	BehaviourSTAT
	BehaviourPalette
	BehaviourWX
	BehaviourSCX
)

func getMemoryWriteBehaviour(address uint16) int {
	switch address {
	case addr.IF:
		return BehaviourWriteCpu
	case addr.LCDC:
		return BehaviourLCDC
	case addr.SCY:
		return BehaviourReadNew
	case addr.STAT:
		return BehaviourSTAT
	case addr.BGP:
		return BehaviourPalette
	case addr.OBP0:
		return BehaviourPalette
	case addr.OBP1:
		return BehaviourPalette
	case addr.WX:
		return BehaviourWX
	case addr.SCX:
		return BehaviourSCX
	default:
		return BehaviourReadOld
	}
}

func (g *GameBoy) Write(addr uint16, v byte) {
	switch getMemoryWriteBehaviour(addr) {
	case BehaviourWriteCpu:
		g.AdvanceTicks(g.pendingTicks + 1)
		g.mmu.Write(addr, v)
		g.pendingTicks = 3
		break
	case BehaviourReadOld:
		fallthrough
	default:
		g.AdvanceTicks(g.pendingTicks)
		g.mmu.Write(addr, v)
		g.pendingTicks = 4
	}
}

func (g *GameBoy) Cycle() {
	g.pendingTicks += 4
}

func (g *GameBoy) WriteIF(v byte) byte {
	g.AdvancePendingTicks()

	prev := g.mmu.Read(addr.IF) & 0x1F

	g.mmu.Write(addr.IF, v)

	g.pendingTicks = 4

	return prev
}

func (g *GameBoy) Step() {
	if g.stopped {
		g.AdvanceTicks(4)

		joyp := g.mmu.Read(addr.JOYP)

		if joyp&0x30 != 0x30 {
			g.Input().SetAccessed(true)
		}

		if joyp&0xF != 0xF {
			// todo: see STOP, this does more than that!

			g.stopped = false
			g.AdvanceTicks(8)
		}

		return
	}

	if g.halted && !g.justHalted {
		g.AdvanceTicks(2)
	}

	interruptQueue := g.mmu.Read(addr.IE) & g.mmu.Read(addr.IF) & 0x1F

	if g.halted {
		if g.justHalted {
			g.AdvanceTicks(4)
		} else {
			g.AdvanceTicks(2)
		}
	}

	g.justHalted = false

	ime := g.CPU().IME()

	// todo: maybe condense this?
	if g.CPU().IMEToggle() {
		g.CPU().SetIME(!g.CPU().IME())
		g.CPU().SetIMEToggle(false)
	}

	if g.halted && !ime && interruptQueue != 0 {
		g.halted = false

		// todo dma_cycles = 4
		// todo: dma run
	} else if ime && interruptQueue != 0 {
		g.halted = false

		// todo: dma cycles = 4
		// todo: dma run

		g.Fetch8()

		// todo: oam bug

		g.CPU().PC().Set(g.CPU().PC().Val() - 1)

		// todo: trigger oam bug

		g.Cycle()

		sp := g.CPU().SP().Val()
		pc := g.CPU().PC().Val()

		sp -= 1
		g.Write(sp, byte(pc>>8))

		interruptQueue = g.mmu.Read(addr.IE)

		if sp == addr.IF+1 {
			sp -= 1
			interruptQueue &= g.WriteIF(uint8(pc))
		} else {
			sp -= 1
			g.Write(sp, byte(pc))
			interruptQueue &= g.mmu.Read(addr.IF) & 0x1F
		}

		g.CPU().SP().Set(sp)

		if interruptQueue != 0 {
			currentInterrupt := types.InterruptType(0)

			for (interruptQueue & 1) == 0 {
				interruptQueue >>= 1
				currentInterrupt += 1
			}

			g.pendingTicks -= 2
			g.FlushPendingTicks()
			g.pendingTicks = 2

			g.interruptController.Flush(currentInterrupt)

			pc = interrupt.GetISR(currentInterrupt)
		} else {
			pc = 0
		}

		g.CPU().PC().Set(pc)
		g.CPU().SetIME(false)
	} else if !g.halted {
		if g.haltBug {
			g.CPU().PC().Set(g.CPU().PC().Val() - 1)
			g.haltBug = false
		}

		cpu.H.ExecuteNextOpcode(g)
	}

	g.FlushPendingTicks()

	// end of cpu run

	if g.interruptController.IsRequested(addr.InterruptJoypad) {
		g.input.SetAccessed(true)
	}
}
