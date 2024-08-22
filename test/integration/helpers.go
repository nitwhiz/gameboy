package integration

import (
	"context"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"github.com/nitwhiz/gameboy/pkg/inst"
	"os"
	"path"
	"strings"
	"testing"
)

const defaultMaxTicks = 5000

type romTestCase struct {
	t        *testing.T
	gameBoy  *gb.GameBoy
	ctx      context.Context
	cancel   context.CancelFunc
	maxTicks int
}

type serialOutCallbackFunc func(testName string, b byte) (bool, bool)

func newRomTestCase(t *testing.T, romPath string, serialOutCallbacks []serialOutCallbackFunc, ctx context.Context) *romTestCase {
	rom, err := os.ReadFile(romPath)

	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(ctx)

	g, err := gb.New(
		gb.WithSerialReceiver(func(b byte) {
			for _, serialOutCallback := range serialOutCallbacks {
				cont, ok := serialOutCallback(t.Name(), b)

				if !ok {
					t.Fail()
				}

				if !ok || !cont {
					cancel()
					return
				}
			}

		}),
		gb.WithRom(rom),
		gb.WithExecuteNextOpcodeFunc(inst.ExecuteNextOpcode),
	)

	if err != nil {
		t.Fatal(err)
	}

	return &romTestCase{
		t:        t,
		gameBoy:  g,
		ctx:      ctx,
		cancel:   cancel,
		maxTicks: defaultMaxTicks,
	}
}

func (r *romTestCase) runGameBoy() {
	ticks := r.maxTicks

	for ; ticks > 0; ticks-- {
		select {
		case <-r.ctx.Done():
			return
		default:
		}

		r.gameBoy.Update(r.ctx)
	}

	r.t.Errorf("game boy ran for at least %d ticks", r.maxTicks)
}

func testRomsRecursive(t *testing.T, root string, serialOutCallbacks []serialOutCallbackFunc) {
	dir, err := os.ReadDir(root)

	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range dir {
		name := entry.Name()
		fullPath := path.Join(root, name)

		if !entry.IsDir() {
			if strings.HasSuffix(name, ".gb") {
				t.Run(name, func(tt *testing.T) {
					tt.Parallel()
					runRomTest(tt, serialOutCallbacks, fullPath, context.Background())
				})
			}
		} else {
			t.Run(name, func(tt *testing.T) {
				tt.Parallel()
				testRomsRecursive(tt, fullPath, serialOutCallbacks)
			})
		}
	}
}

func runRomTest(t *testing.T, serialOutCallbacks []serialOutCallbackFunc, romPath string, ctx context.Context) {
	var serialData []byte

	callbacks := append(
		serialOutCallbacks,
		func(testName string, b byte) (bool, bool) {
			serialData = append(serialData, b)

			return true, true
		},
	)

	r := newRomTestCase(t, romPath, callbacks, ctx)

	defer r.cancel()

	r.runGameBoy()

	if len(serialData) == 0 {
		t.Log("(no serial data)")
	} else {
		t.Logf("serial data: %v", serialData)
	}
}
