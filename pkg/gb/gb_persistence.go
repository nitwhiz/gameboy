package gb

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/nitwhiz/gameboy/pkg/cartridge"
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/gfx"
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/interrupt"
	"github.com/nitwhiz/gameboy/pkg/memory"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"github.com/nitwhiz/gameboy/pkg/quarz"
	"github.com/nitwhiz/gameboy/pkg/screen"
	"github.com/nitwhiz/gameboy/pkg/stack"
)

type SaveState struct {
	CPU              *cpu.CPU
	Cartridge        *cartridge.Cartridge
	Memory           *memory.Memory
	InputButtonState input.ButtonState
	Screen           *screen.Screen
	GFXTicks         int
	DIVTimerTicks    float64
	TIMATimerTicks   float64
}

func (g *GameBoy) PersistState() ([]byte, error) {
	g.Lock()
	defer g.Unlock()

	buf := bytes.Buffer{}

	gzWriter, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)

	if err != nil {
		return nil, err
	}

	jsonEncoder := json.NewEncoder(gzWriter)

	if err := jsonEncoder.Encode(SaveState{
		CPU:              g.CPU,
		Cartridge:        g.MMU.Cartridge,
		Memory:           g.MMU.Memory,
		InputButtonState: g.Input.Buttons,
		Screen:           g.GFX.Screen,
		GFXTicks:         g.GFX.Ticks.GetValue(),
		DIVTimerTicks:    g.DIVTimerTicks.GetValue(),
		TIMATimerTicks:   g.TIMATimerTicks.GetValue(),
	}); err != nil {
		return nil, err
	}

	if err := gzWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (g *GameBoy) LoadState(bs []byte) error {
	g.Lock()
	defer g.Unlock()

	var s SaveState

	buf := bytes.NewBuffer(bs)

	gzReader, err := gzip.NewReader(buf)

	if err != nil {
		return err
	}

	jsonDecoder := json.NewDecoder(gzReader)

	if err := jsonDecoder.Decode(&s); err != nil {
		return err
	}

	if err := gzReader.Close(); err != nil {
		return err
	}

	inp := input.NewStateFrom(s.InputButtonState)

	m := mmu.MMU{
		Cartridge:      s.Cartridge,
		Memory:         s.Memory,
		Input:          inp,
		SerialReceiver: nil,
	}

	stck := stack.Stack{
		CPU: s.CPU,
		MMU: &m,
	}

	in := interrupt.Manager{
		CPU:   s.CPU,
		MMU:   &m,
		Stack: &stck,
	}

	g.CPU = s.CPU
	g.MMU = &m
	g.Input = inp
	g.IM = &in
	g.Stack = &stck
	g.GFX = &gfx.GFX{
		MMU:       &m,
		Interrupt: &in,
		Screen:    s.Screen,
		Ticks:     quarz.NewTickCounter[int](gfx.ScanlineDuration),
	}
	g.DIVTimerTicks = quarz.NewTickCounter[float64](1)
	g.TIMATimerTicks = quarz.NewTickCounter[float64](1)

	g.GFX.Ticks.SetValue(s.GFXTicks)
	g.DIVTimerTicks.SetValue(s.DIVTimerTicks)
	g.TIMATimerTicks.SetValue(s.TIMATimerTicks)

	return nil
}
