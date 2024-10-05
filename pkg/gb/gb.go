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
)

type GameBoy struct {
	CPU types.CPU
	MMU types.MMU

	Timer *quarz.Timer
	Input types.InputState

	IM    *interrupt.Manager
	Stack types.Stack

	PPU types.PPU

	HaltBug int

	mu *sync.Mutex
}

func New(options ...GameBoyOption) (*GameBoy, error) {
	c := cpu.New()

	in := input.NewState()

	m := mmu.New(in, memory.New())

	s := stack.NewStack(c, m)

	t := quarz.NewTimer(m)

	i := interrupt.NewManager(c, m, s)

	g := ppu.New(m, screen.New())

	gameBoy := GameBoy{
		CPU:   c,
		MMU:   m,
		Timer: t,
		Input: in,
		Stack: s,
		IM:    i,
		PPU:   g,
		mu:    &sync.Mutex{},
	}

	for _, o := range options {
		if err := o(&gameBoy); err != nil {
			return nil, err
		}
	}

	return &gameBoy, nil
}

func (g *GameBoy) Lock() {
	g.mu.Lock()
}

func (g *GameBoy) Unlock() {
	g.mu.Unlock()
}

func (g *GameBoy) Update(ctx context.Context) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.MMU.Cartridge == nil {
		slog.Warn("missing cartridge, update skipped")
		return
	}

	executedTicks := 0

	for executedTicks < quarz.CPUTicksPerFrame {
		select {
		case <-ctx.Done():
			return
		default:
		}

		ticks := g.ServiceInterrupts()

		if g.CPU.Halt() {
			// this is not accurate
			ticks += 1
		} else {
			ticks += int(h.executeNextOpcode(g))

			if g.HaltBug > 0 {
				if g.HaltBug == 1 {
					g.CPU.PC().Set(g.CPU.PC().Val() - 1)
				}

				g.HaltBug--
			}
		}

		g.Timer.Tick(ticks)
		g.PPU.Update(ticks)

		executedTicks += ticks
	}
}
