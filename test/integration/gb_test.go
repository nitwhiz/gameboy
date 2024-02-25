package integration

import (
	"context"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"github.com/nitwhiz/gameboy/pkg/inst"
	"os"
	"strings"
	"testing"
)

const MaxFrames = 1500

func TestBlarggRoms(t *testing.T) {
	testRoms := []string{
		"cpu_instrs/individual/01-special",
		"cpu_instrs/individual/02-interrupts",
		"cpu_instrs/individual/03-op sp,hl",
		"cpu_instrs/individual/04-op r,imm",
		"cpu_instrs/individual/05-op rp",
		"cpu_instrs/individual/06-ld r,r",
		"cpu_instrs/individual/07-jr,jp,call,ret,rst",
		"cpu_instrs/individual/08-misc instrs",
		"cpu_instrs/individual/09-op r,r",
		"cpu_instrs/individual/10-bit ops",
		"cpu_instrs/individual/11-op a,(hl)",
	}

	inst.InitHandlers()

	for _, r := range testRoms {
		t.Run(r, func(t *testing.T) {
			romData, err := os.ReadFile("../../testdata/roms/blargg/" + r + ".gb")

			if err != nil {
				t.Fatal(err)
			}

			testOutput := ""
			testResult := false

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			g, err := gb.New(
				gb.WithSerialReceiver(func(b byte) {
					testOutput += string(b)

					if strings.Contains(testOutput, "Passed") {
						testResult = true
						cancel()
					} else if strings.Contains(testOutput, "Failed") {
						cancel()
					}
				}),
				gb.WithRom(romData),
				gb.WithExecuteNextOpcodeFunc(inst.ExecuteNextOpcode),
			)

			if err != nil {
				t.Fatal(err)
			}

			for range MaxFrames {
				select {
				case <-ctx.Done():
					goto end
				default:
					g.Update(ctx)
				}
			}

			t.Errorf("test took >= %d frames", MaxFrames)

		end:
			if !testResult {
				t.Fail()
			}

			t.Log(testOutput)
		})
	}
}
