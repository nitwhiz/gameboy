package gametest

import (
	"context"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"github.com/nitwhiz/gameboy/pkg/inst"
	"os"
	"testing"
)

func getGameBoy(b *testing.B) *gb.GameBoy {
	rom, err := os.ReadFile("../../testdata/roms/custom/gametest.gb")

	if err != nil {
		b.Fatal(err)
	}

	inst.InitHandlers()

	g, err := gb.New(
		gb.WithExecuteNextOpcodeFunc(inst.ExecuteNextOpcode),
		gb.WithRom(rom),
	)

	if err != nil {
		b.Fatal(err)
	}

	return g
}

func BenchmarkGame(b *testing.B) {
	b.ReportAllocs()

	ctx := context.Background()
	g := getGameBoy(b)

	framesRendered := 0

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Update(ctx)
		framesRendered++
	}

	b.StopTimer()

	b.ReportMetric(float64(framesRendered)/b.Elapsed().Seconds(), "fps")
}
