package gb

import (
	"context"
	"errors"
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/gfx"
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/memory"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/quarz"
	"github.com/nitwhiz/gameboy/pkg/stack"
	"sync"
)

type ExecuteNextOpcodeFunc func(g *GameBoy) (ticks byte)

type GameBoy struct {
	CPU *cpu.CPU
	MMU *mmu.MMU

	Input *input.State

	IM    *interrupt.Manager
	Stack *stack.Stack

	GFX *gfx.GFX

	DIVTimerTicks  *quarz.TickCounter[float64]
	TIMATimerTicks *quarz.TickCounter[float64]

	ExecuteNextOpcodeFunc ExecuteNextOpcodeFunc

	mu *sync.Mutex
}

func New(options ...GameBoyOption) (*GameBoy, error) {
	c := cpu.New().Init()

	in := input.NewState()

	m := mmu.MMU{
		Cartridge:      nil,
		Memory:         memory.New().Init(),
		Input:          in,
		SerialReceiver: nil,
	}

	s := stack.Stack{
		CPU: c,
		MMU: &m,
	}

	i := interrupt.Manager{
		CPU:   c,
		MMU:   &m,
		Stack: &s,
	}

	g := gfx.New(&m, &i)

	gameBoy := GameBoy{
		CPU:            c,
		MMU:            &m,
		Input:          in,
		Stack:          &s,
		IM:             &i,
		GFX:            g,
		DIVTimerTicks:  quarz.NewTickCounter[float64](1),
		TIMATimerTicks: quarz.NewTickCounter[float64](1),
		mu:             &sync.Mutex{},
	}

	for _, o := range options {
		if err := o(&gameBoy); err != nil {
			return nil, err
		}
	}

	if g.MMU.Cartridge == nil {
		return nil, errors.New("missing cartridge")
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

	executedTicks := 0

	for executedTicks < quarz.CPUTicksPerFrame {
		select {
		case <-ctx.Done():
			return
		default:
		}

		ticks := g.ServiceInterrupts()

		if g.CPU.Halt {
			// this is not accurate
			ticks += 1
		} else {
			ticks += int(g.ExecuteNextOpcodeFunc(g))
		}

		g.TickTimers(ticks)
		g.UpdateGFX(ticks)

		executedTicks += ticks
	}
}
