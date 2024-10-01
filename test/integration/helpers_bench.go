package integration

import (
	"bytes"
	"context"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"github.com/nitwhiz/gameboy/pkg/screen"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type romBenchmarkCase struct {
	b                  *testing.B
	gameBoy            *gb.GameBoy
	expectedScreenshot *image.Image
	ctx                context.Context
	cancel             context.CancelFunc
	maxFrames          int
}

func newRomBenchmarkCase(b *testing.B, romPath string, expectedScreenshot *image.Image, serialOutCallbackCreators []serialOutCallbackCreator, ctx context.Context) *romBenchmarkCase {
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
	)

	if err != nil {
		b.Fatal(err)
	}

	return &romBenchmarkCase{
		b:                  b,
		gameBoy:            g,
		expectedScreenshot: expectedScreenshot,
		ctx:                ctx,
		cancel:             cancel,
		maxFrames:          defaultMaxFrames,
	}
}

func (r *romBenchmarkCase) checkExpectedScreenshot() {
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

			r.b.StopTimer()
			r.checkExpectedScreenshot()
			r.b.StartTimer()
		}
	}

	r.b.Errorf("game boy ran for at least %d frames", r.maxFrames)
	return 0
}

func runRomBenchmark(b *testing.B, serialOutCallbacks []serialOutCallbackCreator, romPath string, ctx context.Context) {
	gb.InitHandlers()

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

	dirName, fileName := filepath.Split(romPath)
	expectedScreenshotPath := filepath.Join(dirName, strings.TrimSuffix(fileName, filepath.Ext(fileName))+"-expected.png")
	var expectedScreenshot *image.Image

	if _, err := os.Stat(expectedScreenshotPath); err == nil {
		bs, err := os.ReadFile(expectedScreenshotPath)

		if err != nil {
			b.Fatal(err)
		}

		img, err := png.Decode(bytes.NewReader(bs))

		if err != nil {
			b.Fatal(err)
		}

		if !img.Bounds().Eq(image.Rect(0, 0, screen.Width, screen.Height)) {
			b.Fatal("expected screenshot has wrong dimensions")
		}

		expectedScreenshot = &img
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		r := newRomBenchmarkCase(b, romPath, expectedScreenshot, callbacks, ctx)

		b.StartTimer()

		framesRendered += r.runGameBoy()
	}

	b.StopTimer()

	b.ReportMetric(float64(framesRendered)/b.Elapsed().Seconds(), "fps")

	if len(serialData) == 0 {
		b.Log("(no serial data)")
	}
}
