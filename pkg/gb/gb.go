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
	cpu types.CPU
	mmu types.MMU

	Timer *quarz.Timer
	Input types.InputState

	IM    *interrupt.Manager
	stack types.Stack

	ppu types.PPU

	HaltBug int

	mu *sync.Mutex
}

func New(options ...GameBoyOption) (*GameBoy, error) {
	in := input.NewState()

	m := mmu.New(in, memory.New())

	c := cpu.New(context.Background(), m)

	s := stack.NewStack(c, m)

	t := quarz.NewTimer(m)

	i := interrupt.NewManager(c, m, s)

	g := ppu.New(m, screen.New())

	gameBoy := GameBoy{
		cpu:   c,
		mmu:   m,
		Timer: t,
		Input: in,
		stack: s,
		IM:    i,
		ppu:   g,
		mu:    &sync.Mutex{},
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
	return g.MMU()
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
	if g.MMU().Cartridge == nil {
		slog.Warn("missing cartridge, not starting")
		return
	}

	g.cpu.Start()
}

func (g *GameBoy) Stop() {
	g.cpu.Stop()
}

// todo: using ticks now, break this down
func (g *GameBoy) Update(ctx context.Context) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.MMU().Cartridge == nil {
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

		if g.cpu.Halt() {
			// this is not accurate
			ticks += 1
		} else {
			//ticks += int(cpu.h.executeNextOpcode(g))
		}

		g.Timer.Tick(ticks)
		g.ppu.Update(ticks)

		executedTicks += ticks
	}
}
