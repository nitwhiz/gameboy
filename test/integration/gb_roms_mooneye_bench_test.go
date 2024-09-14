// Do not edit. This is auto-generated.

package integration

import (
	"context"
	"testing"
)

func BenchmarkMooneyeMBC1Roms(b *testing.B) {

	b.Run("mbc1", func(b *testing.B) {

		b.Run("bits_bank1", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/bits_bank1.gb", context.Background())
		})

		b.Run("bits_bank2", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/bits_bank2.gb", context.Background())
		})

		b.Run("bits_mode", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/bits_mode.gb", context.Background())
		})

		b.Run("bits_ramg", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/bits_ramg.gb", context.Background())
		})

		b.Run("multicart_rom_8Mb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/multicart_rom_8Mb.gb", context.Background())
		})

		b.Run("ram_256kb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/ram_256kb.gb", context.Background())
		})

		b.Run("ram_64kb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/ram_64kb.gb", context.Background())
		})

		b.Run("rom_16Mb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_16Mb.gb", context.Background())
		})

		b.Run("rom_1Mb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_1Mb.gb", context.Background())
		})

		b.Run("rom_2Mb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_2Mb.gb", context.Background())
		})

		b.Run("rom_4Mb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_4Mb.gb", context.Background())
		})

		b.Run("rom_512kb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_512kb.gb", context.Background())
		})

		b.Run("rom_8Mb", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{
				mooneyeSerialCallback,
			}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_8Mb.gb", context.Background())
		})

	})

}
