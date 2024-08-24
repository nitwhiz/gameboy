// Do not edit. This is auto-generated.
// Timestamp: 2024-08-24T13:48:06Z

package integration

import (
	"context"
	"log/slog"
	"testing"
)

func TestMooneyeRoms(t *testing.T) {
	t.Parallel()

	slog.SetLogLoggerLevel(slog.LevelDebug)

	t.Run("acceptance", func(t *testing.T) {
		t.Parallel()

		t.Run("add_sp_e_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/add_sp_e_timing.gb", context.Background())
		})

		t.Run("bits", func(t *testing.T) {
			t.Parallel()

			t.Run("mem_oam", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/bits/mem_oam.gb", context.Background())
			})

			t.Run("reg_f", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/bits/reg_f.gb", context.Background())
			})

			t.Run("unused_hwio-GS", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/bits/unused_hwio-GS.gb", context.Background())
			})

		})

		t.Run("boot_div-S", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/boot_div-S.gb", context.Background())
		})

		t.Run("boot_div-dmg0", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/boot_div-dmg0.gb", context.Background())
		})

		t.Run("boot_div2-S", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/boot_div2-S.gb", context.Background())
		})

		t.Run("boot_hwio-S", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/boot_hwio-S.gb", context.Background())
		})

		t.Run("boot_hwio-dmg0", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/boot_hwio-dmg0.gb", context.Background())
		})

		t.Run("boot_regs-dmg0", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/boot_regs-dmg0.gb", context.Background())
		})

		t.Run("boot_regs-dmgABC", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/boot_regs-dmgABC.gb", context.Background())
		})

		t.Run("call_cc_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/call_cc_timing.gb", context.Background())
		})

		t.Run("call_cc_timing2", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/call_cc_timing2.gb", context.Background())
		})

		t.Run("call_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/call_timing.gb", context.Background())
		})

		t.Run("call_timing2", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/call_timing2.gb", context.Background())
		})

		t.Run("di_timing-GS", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/di_timing-GS.gb", context.Background())
		})

		t.Run("div_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/div_timing.gb", context.Background())
		})

		t.Run("ei_sequence", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/ei_sequence.gb", context.Background())
		})

		t.Run("ei_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/ei_timing.gb", context.Background())
		})

		t.Run("halt_ime0_ei", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/halt_ime0_ei.gb", context.Background())
		})

		t.Run("halt_ime0_nointr_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/halt_ime0_nointr_timing.gb", context.Background())
		})

		t.Run("halt_ime1_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/halt_ime1_timing.gb", context.Background())
		})

		t.Run("halt_ime1_timing2-GS", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/halt_ime1_timing2-GS.gb", context.Background())
		})

		t.Run("if_ie_registers", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/if_ie_registers.gb", context.Background())
		})

		t.Run("instr", func(t *testing.T) {
			t.Parallel()

			t.Run("daa", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/instr/daa.gb", context.Background())
			})

		})

		t.Run("interrupts", func(t *testing.T) {
			t.Parallel()

			t.Run("ie_push", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/interrupts/ie_push.gb", context.Background())
			})

		})

		t.Run("intr_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/intr_timing.gb", context.Background())
		})

		t.Run("jp_cc_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/jp_cc_timing.gb", context.Background())
		})

		t.Run("jp_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/jp_timing.gb", context.Background())
		})

		t.Run("ld_hl_sp_e_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/ld_hl_sp_e_timing.gb", context.Background())
		})

		t.Run("oam_dma", func(t *testing.T) {
			t.Parallel()

			t.Run("basic", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/oam_dma/basic.gb", context.Background())
			})

			t.Run("reg_read", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/oam_dma/reg_read.gb", context.Background())
			})

			t.Run("sources-GS", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/oam_dma/sources-GS.gb", context.Background())
			})

		})

		t.Run("oam_dma_restart", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/oam_dma_restart.gb", context.Background())
		})

		t.Run("oam_dma_start", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/oam_dma_start.gb", context.Background())
		})

		t.Run("oam_dma_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/oam_dma_timing.gb", context.Background())
		})

		t.Run("pop_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/pop_timing.gb", context.Background())
		})

		t.Run("ppu", func(t *testing.T) {
			t.Parallel()

			t.Run("hblank_ly_scx_timing-GS", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/hblank_ly_scx_timing-GS.gb", context.Background())
			})

			t.Run("intr_1_2_timing-GS", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/intr_1_2_timing-GS.gb", context.Background())
			})

			t.Run("intr_2_0_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/intr_2_0_timing.gb", context.Background())
			})

			t.Run("intr_2_mode0_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/intr_2_mode0_timing.gb", context.Background())
			})

			t.Run("intr_2_mode0_timing_sprites", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/intr_2_mode0_timing_sprites.gb", context.Background())
			})

			t.Run("intr_2_mode3_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/intr_2_mode3_timing.gb", context.Background())
			})

			t.Run("intr_2_oam_ok_timing", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/intr_2_oam_ok_timing.gb", context.Background())
			})

			t.Run("lcdon_timing-GS", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/lcdon_timing-GS.gb", context.Background())
			})

			t.Run("lcdon_write_timing-GS", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/lcdon_write_timing-GS.gb", context.Background())
			})

			t.Run("stat_irq_blocking", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/stat_irq_blocking.gb", context.Background())
			})

			t.Run("stat_lyc_onoff", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/stat_lyc_onoff.gb", context.Background())
			})

			t.Run("vblank_stat_intr-GS", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/ppu/vblank_stat_intr-GS.gb", context.Background())
			})

		})

		t.Run("push_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/push_timing.gb", context.Background())
		})

		t.Run("rapid_di_ei", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/rapid_di_ei.gb", context.Background())
		})

		t.Run("ret_cc_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/ret_cc_timing.gb", context.Background())
		})

		t.Run("ret_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/ret_timing.gb", context.Background())
		})

		t.Run("reti_intr_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/reti_intr_timing.gb", context.Background())
		})

		t.Run("reti_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/reti_timing.gb", context.Background())
		})

		t.Run("rst_timing", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/acceptance/rst_timing.gb", context.Background())
		})

		t.Run("timer", func(t *testing.T) {
			t.Parallel()

			t.Run("div_write", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/div_write.gb", context.Background())
			})

			t.Run("rapid_toggle", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/rapid_toggle.gb", context.Background())
			})

			t.Run("tim00", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tim00.gb", context.Background())
			})

			t.Run("tim00_div_trigger", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tim00_div_trigger.gb", context.Background())
			})

			t.Run("tim01", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tim01.gb", context.Background())
			})

			t.Run("tim01_div_trigger", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tim01_div_trigger.gb", context.Background())
			})

			t.Run("tim10", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tim10.gb", context.Background())
			})

			t.Run("tim10_div_trigger", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tim10_div_trigger.gb", context.Background())
			})

			t.Run("tim11", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tim11.gb", context.Background())
			})

			t.Run("tim11_div_trigger", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tim11_div_trigger.gb", context.Background())
			})

			t.Run("tima_reload", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tima_reload.gb", context.Background())
			})

			t.Run("tima_write_reloading", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tima_write_reloading.gb", context.Background())
			})

			t.Run("tma_write_reloading", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/acceptance/timer/tma_write_reloading.gb", context.Background())
			})

		})

	})

	t.Run("emulator-only", func(t *testing.T) {
		t.Parallel()

		t.Run("mbc1", func(t *testing.T) {
			t.Parallel()

			t.Run("bits_bank1", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/bits_bank1.gb", context.Background())
			})

			t.Run("bits_bank2", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/bits_bank2.gb", context.Background())
			})

			t.Run("bits_mode", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/bits_mode.gb", context.Background())
			})

			t.Run("bits_ramg", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/bits_ramg.gb", context.Background())
			})

			t.Run("multicart_rom_8Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/multicart_rom_8Mb.gb", context.Background())
			})

			t.Run("ram_256kb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/ram_256kb.gb", context.Background())
			})

			t.Run("ram_64kb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/ram_64kb.gb", context.Background())
			})

			t.Run("rom_16Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_16Mb.gb", context.Background())
			})

			t.Run("rom_1Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_1Mb.gb", context.Background())
			})

			t.Run("rom_2Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_2Mb.gb", context.Background())
			})

			t.Run("rom_4Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_4Mb.gb", context.Background())
			})

			t.Run("rom_512kb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_512kb.gb", context.Background())
			})

			t.Run("rom_8Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc1/rom_8Mb.gb", context.Background())
			})

		})

		t.Run("mbc2", func(t *testing.T) {
			t.Parallel()

			t.Run("bits_ramg", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc2/bits_ramg.gb", context.Background())
			})

			t.Run("bits_romb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc2/bits_romb.gb", context.Background())
			})

			t.Run("bits_unused", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc2/bits_unused.gb", context.Background())
			})

			t.Run("ram", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc2/ram.gb", context.Background())
			})

			t.Run("rom_1Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc2/rom_1Mb.gb", context.Background())
			})

			t.Run("rom_2Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc2/rom_2Mb.gb", context.Background())
			})

			t.Run("rom_512kb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc2/rom_512kb.gb", context.Background())
			})

		})

		t.Run("mbc5", func(t *testing.T) {
			t.Parallel()

			t.Run("rom_16Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc5/rom_16Mb.gb", context.Background())
			})

			t.Run("rom_1Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc5/rom_1Mb.gb", context.Background())
			})

			t.Run("rom_2Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc5/rom_2Mb.gb", context.Background())
			})

			t.Run("rom_32Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc5/rom_32Mb.gb", context.Background())
			})

			t.Run("rom_4Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc5/rom_4Mb.gb", context.Background())
			})

			t.Run("rom_512kb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc5/rom_512kb.gb", context.Background())
			})

			t.Run("rom_64Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc5/rom_64Mb.gb", context.Background())
			})

			t.Run("rom_8Mb", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/emulator-only/mbc5/rom_8Mb.gb", context.Background())
			})

		})

	})

	t.Run("manual-only", func(t *testing.T) {
		t.Parallel()

		t.Run("sprite_priority", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/manual-only/sprite_priority.gb", context.Background())
		})

	})

	t.Run("misc", func(t *testing.T) {
		t.Parallel()

		t.Run("bits", func(t *testing.T) {
			t.Parallel()

			t.Run("unused_hwio-C", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/misc/bits/unused_hwio-C.gb", context.Background())
			})

		})

		t.Run("boot_div-A", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/misc/boot_div-A.gb", context.Background())
		})

		t.Run("boot_hwio-C", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/misc/boot_hwio-C.gb", context.Background())
		})

		t.Run("boot_regs-A", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/misc/boot_regs-A.gb", context.Background())
		})

		t.Run("ppu", func(t *testing.T) {
			t.Parallel()

			t.Run("vblank_stat_intr-C", func(t *testing.T) {
				t.Parallel()

				runRomTest(t, []serialOutCallbackFunc{
					mooneyeSerialCallback(),
				}, "../../testdata/roms/mooneye/misc/ppu/vblank_stat_intr-C.gb", context.Background())
			})

		})

	})

	t.Run("utils", func(t *testing.T) {
		t.Parallel()

		t.Run("dump_boot_hwio", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{
				mooneyeSerialCallback(),
			}, "../../testdata/roms/mooneye/utils/dump_boot_hwio.gb", context.Background())
		})

	})

}
