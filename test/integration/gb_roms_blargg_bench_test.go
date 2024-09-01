// Do not edit. This is auto-generated.
// Timestamp: 2024-09-01T11:47:01Z

package integration

import (
	"context"
	"testing"
)

func BenchmarkBlarggCpuInstrsRoms(b *testing.B) {

	b.Run("cpu_instrs", func(b *testing.B) {

		b.Run("cpu_instrs", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				blarggSerialCallback,
			}, "../../testdata/roms/blargg/cpu_instrs/cpu_instrs.gb", context.Background())
		})

		b.Run("individual", func(b *testing.B) {

			b.Run("01-special", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/01-special.gb", context.Background())
			})

			b.Run("02-interrupts", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/02-interrupts.gb", context.Background())
			})

			b.Run("03-op sp,hl", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/03-op sp,hl.gb", context.Background())
			})

			b.Run("04-op r,imm", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/04-op r,imm.gb", context.Background())
			})

			b.Run("05-op rp", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/05-op rp.gb", context.Background())
			})

			b.Run("06-ld r,r", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/06-ld r,r.gb", context.Background())
			})

			b.Run("07-jr,jp,call,ret,rst", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/07-jr,jp,call,ret,rst.gb", context.Background())
			})

			b.Run("08-misc instrs", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/08-misc instrs.gb", context.Background())
			})

			b.Run("09-op r,r", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/09-op r,r.gb", context.Background())
			})

			b.Run("10-bit ops", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/10-bit ops.gb", context.Background())
			})

			b.Run("11-op a,(hl)", func(b *testing.B) {
				runRomBenchmark(b, []serialOutCallbackCreator{
					blarggSerialCallback,
				}, "../../testdata/roms/blargg/cpu_instrs/individual/11-op a,(hl).gb", context.Background())
			})

		})

	})

}
