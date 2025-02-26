package gb

import (
	"context"
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/memory"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/ppu"
	"github.com/nitwhiz/gameboy/pkg/quarz"
	"github.com/nitwhiz/gameboy/pkg/screen"
	"github.com/nitwhiz/gameboy/pkg/stack"
	"github.com/nitwhiz/gameboy/pkg/types"
	"log/slog"
	"sync"
	"time"
)

type GameBoy struct {
	cpu types.CPU
	mmu types.MMU

	Timer *quarz.Timer
	Input types.InputState

	IM    *interrupt.Manager
	stack types.Stack

	ppu types.PPU

	HaltBug int

	ticker *time.Ticker
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup

	mu *sync.Mutex
}

func New(ctx context.Context, options ...GameBoyOption) (*GameBoy, error) {
	in := input.NewState()

	m := mmu.New(in, memory.New())

	c := cpu.New(m)

	s := stack.NewStack(c, m)

	t := quarz.NewTimer(m)

	i := interrupt.NewManager(c, m, s)

	g := ppu.New(m, screen.New())

	ctx, cancel := context.WithCancel(ctx)

	gameBoy := GameBoy{
		cpu:   c,
		mmu:   m,
		Timer: t,
		Input: in,
		stack: s,
		IM:    i,
		ppu:   g,
		mu:    &sync.Mutex{},

		ticker: quarz.Ticker,
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

func (g *GameBoy) MMU() types.MMU {
	return g.mmu
}

func (g *GameBoy) Stack() types.Stack {
	return g.stack
}

func (g *GameBoy) PPU() types.PPU {
	return g.ppu
}

func (g *GameBoy) Lock() {
	g.mu.Lock()
}

func (g *GameBoy) Unlock() {
	g.mu.Unlock()
}

func (g *GameBoy) Start() {
	if g.MMU().Cartridge() == nil {
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

				instructionTicks = int(cpu.H.ExecuteNextOpcode(g))

				ticksLeft += instructionTicks

			end:
				g.Timer.Tick(1)
				g.ppu.Update(1)

				break
			}
		}
	}()
}

func (g *GameBoy) Stop() {
	g.cancel()
	g.wg.Wait()
}
