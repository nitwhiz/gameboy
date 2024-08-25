package integration

import (
	"context"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"github.com/nitwhiz/gameboy/pkg/inst"
	"image/png"
	"os"
	"path"
	"reflect"
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

type serialOutCallbackFunc func(b byte) (bool, bool)

func newRomTestCase(t *testing.T, romPath string, serialOutCallbacks []serialOutCallbackFunc, ctx context.Context) *romTestCase {
	rom, err := os.ReadFile(romPath)

	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(ctx)

	g, err := gb.New(
		gb.WithSerialReceiver(func(b byte) {
			for _, serialOutCallback := range serialOutCallbacks {
				cont, ok := serialOutCallback(b)

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

func cleanupOutputs(t *testing.T) {
	if err := os.RemoveAll(path.Join("../../testdata/output", t.Name())); err != nil {
		t.Fatal(err)
	}
}

func runRomTest(t *testing.T, serialOutCallbacks []serialOutCallbackFunc, romPath string, ctx context.Context) {
	inst.InitHandlers()

	var serialData []byte

	callbacks := append(
		serialOutCallbacks,
		func(b byte) (bool, bool) {
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

	if t.Failed() {
		name := t.Name()

		if err := os.MkdirAll(path.Join("../../testdata/output", path.Dir(name)), 0775); err != nil {
			panic(err)
		}

		f, err := os.Create(path.Join("../../testdata/output", name+".png"))

		if err != nil {
			t.Fatal(err)
		}

		defer func(f *os.File) {
			err := f.Close()

			if err != nil {
				t.Fatal(err)
			}
		}(f)

		if err := png.Encode(f, r.gameBoy.GFX.Screen); err != nil {
			t.Fatal(err)
		}
	}
}

func blarggSerialCallback() serialOutCallbackFunc {
	var serialData []byte

	return func(b byte) (bool, bool) {
		serialData = append(serialData, b)

		if strings.HasSuffix(string(serialData), "Passed") {
			return false, true
		}

		if strings.HasSuffix(string(serialData), "Failed") {
			return false, false
		}

		return true, true
	}
}

func mooneyeSerialCallback() serialOutCallbackFunc {
	var serialData []byte

	successData := []byte{3, 5, 8, 13, 21, 34}
	failData := []byte{66, 66, 66, 66, 66, 66}

	return func(b byte) (bool, bool) {
		serialData = append(serialData, b)

		if reflect.DeepEqual(serialData, successData) {
			return false, true
		}

		if reflect.DeepEqual(serialData, failData) {
			return false, false
		}

		return true, true
	}
}
