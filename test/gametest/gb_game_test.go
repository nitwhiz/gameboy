package gametest

import (
	"context"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"github.com/nitwhiz/gameboy/pkg/inst"
	"os"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	rom, err := os.ReadFile("../../testdata/roms/custom/gametest.gb")

	if err != nil {
		t.Fatal(err)
	}

	inst.InitHandlers()

	g, err := gb.New(
		gb.WithExecuteNextOpcodeFunc(inst.ExecuteNextOpcode),
		gb.WithRom(rom),
	)

	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			g.Update(ctx)
		}
	}
}
