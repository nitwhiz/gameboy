package integration

import (
	"bytes"
	"context"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"github.com/nitwhiz/gameboy/pkg/inst"
	"github.com/nitwhiz/gameboy/pkg/screen"
	"image"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

const defaultMaxFrames = 5000

type romTestCase struct {
	t                  *testing.T
	gameBoy            *gb.GameBoy
	expectedScreenshot *image.Image
	ctx                context.Context
	cancel             context.CancelFunc
	maxFrames          int
}

type serialOutCallbackFunc func(b byte) (bool, bool)
type serialOutCallbackCreator func() serialOutCallbackFunc

func newRomTestCase(t *testing.T, romPath string, expectedScreenshot *image.Image, serialOutCallbacks []serialOutCallbackFunc, ctx context.Context) *romTestCase {
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
		t:                  t,
		gameBoy:            g,
		expectedScreenshot: expectedScreenshot,
		ctx:                ctx,
		cancel:             cancel,
		maxFrames:          defaultMaxFrames,
	}
}

func (r *romTestCase) runGameBoy() {
	frames := r.maxFrames

	for ; frames > 0; frames-- {
		select {
		case <-r.ctx.Done():
			return
		default:
			r.gameBoy.Update(r.ctx)
			r.checkExpectedScreenshot()
		}
	}

	r.t.Errorf("game boy ran for at least %d frames", r.maxFrames)
}

func cleanupOutputs(t *testing.T) {
	if err := os.RemoveAll(filepath.Join("../../testdata/output", t.Name())); err != nil {
		t.Fatal(err)
	}
}

func (r *romTestCase) checkExpectedScreenshot() {
	if r.expectedScreenshot == nil {
		return
	}

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			if (*r.expectedScreenshot).At(x, y) != r.gameBoy.PPU.Screen.At(x, y) {
				return
			}
		}
	}

	r.cancel()
}

func (r *romTestCase) screenshot() {
	name := r.t.Name()

	if err := os.MkdirAll(filepath.Join("../../testdata/output", path.Dir(name)), 0775); err != nil {
		r.t.Fatal(err)
	}

	f, err := os.Create(filepath.Join("../../testdata/output", name+".png"))

	if err != nil {
		r.t.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()

		if err != nil {
			r.t.Fatal(err)
		}
	}(f)

	if err := png.Encode(f, r.gameBoy.PPU.Screen); err != nil {
		r.t.Fatal(err)
	}
}

func runRomTest(t *testing.T, serialOutCallbacks []serialOutCallbackFunc, romPath string, ctx context.Context) {
	inst.InitHandlers()

	var serialData []byte

	serialCallbacks := append(
		serialOutCallbacks,
		func(b byte) (bool, bool) {
			serialData = append(serialData, b)

			return true, true
		},
	)

	dirName, fileName := filepath.Split(romPath)
	expectedScreenshotPath := filepath.Join(dirName, strings.TrimSuffix(fileName, filepath.Ext(fileName))+"-expected.png")
	var expectedScreenshot *image.Image

	if _, err := os.Stat(expectedScreenshotPath); err == nil {
		bs, err := os.ReadFile(expectedScreenshotPath)

		if err != nil {
			t.Fatal(err)
		}

		img, err := png.Decode(bytes.NewReader(bs))

		if err != nil {
			t.Fatal(err)
		}

		if !img.Bounds().Eq(image.Rect(0, 0, screen.Width, screen.Height)) {
			t.Fatal("expected screenshot has wrong dimensions")
		}

		expectedScreenshot = &img
	}

	r := newRomTestCase(t, romPath, expectedScreenshot, serialCallbacks, ctx)

	defer r.cancel()

	r.runGameBoy()

	if len(serialData) == 0 {
		t.Log("(no serial data)")
	} else {
		t.Logf("serial data: %v", serialData)
	}

	if t.Failed() {
		r.screenshot()
	}
}
