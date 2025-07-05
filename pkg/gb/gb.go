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
	"log/slog"
	"sync"
)

type GameBoy struct {
	cpu types.CPU
	mmu types.MMU

	Timer *timer.Timer
	Input types.InputState

	IM *interrupt.Manager

	ppu types.PPU

	HaltBug int

	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup

	pendingTicks        int
	allowedPendingTicks int

	mu *sync.Mutex
}

func New(ctx context.Context, options ...GameBoyOption) (*GameBoy, error) {
	in := input.NewState()

	t := timer.New()

	m := mmu.New(in, memory.NewContainer(), t)

	c := cpu.New(m)

	i := interrupt.NewManager(c, m)

	g := ppu.New(m, screen.New())

	ctx, cancel := context.WithCancel(ctx)

	gameBoy := GameBoy{
		cpu:   c,
		mmu:   m,
		Timer: t,
		Input: in,
		IM:    i,
		ppu:   g,
		mu:    &sync.Mutex{},

		ctx:    ctx,
		wg:     &sync.WaitGroup{},
		cancel: cancel,
	}

	for _, o := range options {
		if err := o(&gameBoy); err != nil {
			return nil, err
		}
	}

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

func (g *GameBoy) Lock() {
	g.mu.Lock()
}

func (g *GameBoy) Unlock() {
	g.mu.Unlock()
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
	g.Timer.Tick(n)
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

func (g *GameBoy) Write(addr uint16, v byte) {
	// todo: conflicts

	// this is GB_CONFLICT_READ_OLD; the default

	g.AdvanceTicks(g.pendingTicks)
	g.mmu.Write(addr, v)
	g.pendingTicks = 4
}

func (g *GameBoy) Cycle() {
	g.pendingTicks += 4
}

func (g *GameBoy) Start() {
	if !g.mmu.HasCartridge() {
		slog.Warn("missing cartridge, not starting")
		return
	}

	cpu.WriteGameBoyDoctorLog(g)

	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		for {
			select {
			case <-g.ctx.Done():
				return
			default:
			}

			// todo: debug

			halted := false
			justHalted := false

			// todo: if stopped; continue here

			interruptQueue := g.mmu.Read(addr.IE) & g.mmu.Read(addr.IF) & 0x1F

			if halted {
				if justHalted {
					g.AdvanceTicks(4)
				} else {
					g.AdvanceTicks(2)
				}
			}

			justHalted = false

			ime := g.CPU().IME()

			// todo: maybe condense this?
			if g.CPU().IMEToggle() {
				g.CPU().SetIME(!ime)
				g.CPU().SetIMEToggle(false)
			}

			if halted && !ime && interruptQueue != 0 {
				// todo: wake up from halt, without calling interrupt
			} else if ime && interruptQueue != 0 {
				// todo: call interrupt
				// todo: dma

				//pc := g.CPU().PC().Val()
				//callAddr := pc
				//
				//g.Read8(pc)
				//
				//pc += 1
				//
				//g.CPU().PC().Set(pc)
				//
				//// todo: oam bug
				//
				//pc -= 1
				//
				//g.CPU().PC().Set(pc)
				//
				//// todo: trigger oam bug
				//
				//g.Cycle()
				//
				//g.PushStack()
			} else if !halted {
				// todo: if halt bug; do halt bug stuff

				cpu.H.ExecuteNextOpcode(g)
			}

			g.FlushPendingTicks()
		}
	}()
}

func (g *GameBoy) Start2() {
	if !g.mmu.HasCartridge() {
		slog.Warn("missing cartridge, not starting")
		return
	}

	g.wg.Add(1)

	ticksLeft := 0

	interruptTicks := 0
	instructionTicks := 0

	go func() {
		defer g.wg.Done()

		for {
			select {
			case <-g.ctx.Done():
				return
			default:
				interruptTicks = 0
				instructionTicks = 0

				if ticksLeft > 0 {
					ticksLeft--
					goto end
				}

				interruptTicks = g.ServiceInterrupts()

				if interruptTicks > 0 {
					ticksLeft += interruptTicks
					goto end
				}

				if g.cpu.Halt() {
					goto end
				}

				cpu.H.ExecuteNextOpcode(g)

				// todo: rewrite me to new pending ticks system

				ticksLeft += instructionTicks

			end:
				g.Timer.Tick(1)
				g.ppu.Ticks(1)

				break
			}
		}
	}()
}

func (g *GameBoy) Stop() {
	g.cancel()
	g.wg.Wait()
}
