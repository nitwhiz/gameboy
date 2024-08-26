package integration

import (
	"context"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"github.com/nitwhiz/gameboy/pkg/inst"
	"os"
	"testing"
)

type romBenchmarkCase struct {
	b         *testing.B
	gameBoy   *gb.GameBoy
	ctx       context.Context
	cancel    context.CancelFunc
	maxFrames int
}

func newRomBenchmarkCase(b *testing.B, romPath string, serialOutCallbackCreators []serialOutCallbackCreator, ctx context.Context) *romBenchmarkCase {
	rom, err := os.ReadFile(romPath)

	if err != nil {
		b.Fatal(err)
	}

	ctx, cancel := context.WithCancel(ctx)

	var serialOutCallbacks []serialOutCallbackFunc

	for _, cc := range serialOutCallbackCreators {
		serialOutCallbacks = append(serialOutCallbacks, cc())
	}

	g, err := gb.New(
		gb.WithSerialReceiver(func(d byte) {
			b.StopTimer()

			for _, serialOutCallback := range serialOutCallbacks {
				cont, ok := serialOutCallback(d)

				if !ok {
					b.Fail()
				}

				if !ok || !cont {
					cancel()
					return
				}
			}

			b.StartTimer()
		}),
		gb.WithRom(rom),
		gb.WithExecuteNextOpcodeFunc(inst.ExecuteNextOpcode),
	)

	if err != nil {
		b.Fatal(err)
	}

	return &romBenchmarkCase{
		b:         b,
		gameBoy:   g,
		ctx:       ctx,
		cancel:    cancel,
		maxFrames: defaultMaxFrames,
	}
}

func (r *romBenchmarkCase) runGameBoy() int {
	framesRendered := 0
	frames := r.maxFrames

	for ; frames > 0; frames-- {
		select {
		case <-r.ctx.Done():
			return framesRendered
		default:
			r.gameBoy.Update(r.ctx)
			framesRendered++
		}
	}

	r.b.Errorf("game boy ran for at least %d frames", r.maxFrames)
	return 0
}

func runRomBenchmark(b *testing.B, serialOutCallbacks []serialOutCallbackCreator, romPath string, ctx context.Context) {
	inst.InitHandlers()

	var serialData []byte

	callbacks := append(
		serialOutCallbacks,
		func() serialOutCallbackFunc {
			return func(d byte) (bool, bool) {
				serialData = append(serialData, d)

				return true, true
			}
		},
	)

	framesRendered := 0

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		r := newRomBenchmarkCase(b, romPath, callbacks, ctx)

		b.StartTimer()

		framesRendered += r.runGameBoy()
	}

	b.StopTimer()

	b.ReportMetric(float64(framesRendered)/b.Elapsed().Seconds(), "fps")

	if len(serialData) == 0 {
		b.Log("(no serial data)")
	}
}
