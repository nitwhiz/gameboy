// Do not edit. This is auto-generated.

package integration

import (
	"context"
	"log/slog"
	"testing"
)

func TestMiscRoms(t *testing.T) {
	cleanupOutputs(t)
	t.Parallel()

	slog.SetLogLoggerLevel(slog.LevelDebug)

	t.Run("dmg-acid2", func(t *testing.T) {
		t.Parallel()

		t.Run("dmg-acid2", func(t *testing.T) {
			t.Parallel()

			runRomTest(t, []serialOutCallbackFunc{}, "../../testdata/roms/misc/dmg-acid2/dmg-acid2.gb", context.Background())
		})

	})

}
