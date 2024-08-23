// Do not edit. This is auto-generated.
// Timestamp: 2024-08-23T14:21:22Z

package integration

import (
	"context"
	"testing"
)

func TestBlarggRoms(t *testing.T) {
	t.Parallel()

	t.Run("cpu_instrs", func(t *testing.T) {
		t.Parallel()

		t.Run("cpu_instrs", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				blarggSerialCallback(),
			}, "../../testdata/roms/blargg/cpu_instrs/cpu_instrs.gb", context.Background())
		})

		t.Run("individual", func(t *testing.T) {
			t.Parallel()

			t.Run("01-special", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/01-special.gb", context.Background())
			})

			t.Run("02-interrupts", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/02-interrupts.gb", context.Background())
			})

			t.Run("03-op sp,hl", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/03-op sp,hl.gb", context.Background())
			})

			t.Run("04-op r,imm", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/04-op r,imm.gb", context.Background())
			})

			t.Run("05-op rp", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/05-op rp.gb", context.Background())
			})

			t.Run("06-ld r,r", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/06-ld r,r.gb", context.Background())
			})

			t.Run("07-jr,jp,call,ret,rst", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/07-jr,jp,call,ret,rst.gb", context.Background())
			})

			t.Run("08-misc instrs", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/08-misc instrs.gb", context.Background())
			})

			t.Run("09-op r,r", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/09-op r,r.gb", context.Background())
			})

			t.Run("10-bit ops", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/10-bit ops.gb", context.Background())
			})

			t.Run("11-op a,(hl)", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/cpu_instrs/individual/11-op a,(hl).gb", context.Background())
			})

		})

	})

	t.Run("halt_bug", func(t *testing.T) {
		t.Parallel()

		runRomTest(t, []serialOutCallbackFunc{
			blarggSerialCallback(),
		}, "../../testdata/roms/blargg/halt_bug.gb", context.Background())
	})

	t.Run("instr_timing", func(t *testing.T) {
		t.Parallel()

		t.Run("instr_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				blarggSerialCallback(),
			}, "../../testdata/roms/blargg/instr_timing/instr_timing.gb", context.Background())
		})

	})

	t.Run("mem_timing-2", func(t *testing.T) {
		t.Parallel()

		t.Run("individual", func(t *testing.T) {
			t.Parallel()

			t.Run("01-read_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/mem_timing-2/individual/01-read_timing.gb", context.Background())
			})

			t.Run("02-write_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/mem_timing-2/individual/02-write_timing.gb", context.Background())
			})

			t.Run("03-modify_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/mem_timing-2/individual/03-modify_timing.gb", context.Background())
			})

		})

		t.Run("mem_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				blarggSerialCallback(),
			}, "../../testdata/roms/blargg/mem_timing-2/mem_timing.gb", context.Background())
		})

	})

	t.Run("mem_timing", func(t *testing.T) {
		t.Parallel()

		t.Run("individual", func(t *testing.T) {
			t.Parallel()

			t.Run("01-read_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/mem_timing/individual/01-read_timing.gb", context.Background())
			})

			t.Run("02-write_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/mem_timing/individual/02-write_timing.gb", context.Background())
			})

			t.Run("03-modify_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/mem_timing/individual/03-modify_timing.gb", context.Background())
			})

		})

		t.Run("mem_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				blarggSerialCallback(),
			}, "../../testdata/roms/blargg/mem_timing/mem_timing.gb", context.Background())
		})

	})

	t.Run("oam_bug", func(t *testing.T) {
		t.Parallel()

		t.Run("individual", func(t *testing.T) {
			t.Parallel()

			t.Run("1-lcd_sync", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/oam_bug/individual/1-lcd_sync.gb", context.Background())
			})

			t.Run("2-causes", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/oam_bug/individual/2-causes.gb", context.Background())
			})

			t.Run("3-non_causes", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/oam_bug/individual/3-non_causes.gb", context.Background())
			})

			t.Run("4-scanline_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/oam_bug/individual/4-scanline_timing.gb", context.Background())
			})

			t.Run("5-timing_bug", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/oam_bug/individual/5-timing_bug.gb", context.Background())
			})

			t.Run("6-timing_no_bug", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/oam_bug/individual/6-timing_no_bug.gb", context.Background())
			})

			t.Run("7-timing_effect", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/oam_bug/individual/7-timing_effect.gb", context.Background())
			})

			t.Run("8-instr_effect", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					blarggSerialCallback(),
				}, "../../testdata/roms/blargg/oam_bug/individual/8-instr_effect.gb", context.Background())
			})

		})

		t.Run("oam_bug", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				blarggSerialCallback(),
			}, "../../testdata/roms/blargg/oam_bug/oam_bug.gb", context.Background())
		})

	})

}
