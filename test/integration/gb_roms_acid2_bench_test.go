// Do not edit. This is auto-generated.

package integration

import (
	"context"
	"testing"
)

func BenchmarkAcid2Roms(b *testing.B) {

	b.Run("dmg-acid2", func(b *testing.B) {

		b.Run("dmg-acid2", func(b *testing.B) {
			runRomBenchmark(b, []serialOutCallbackCreator{}, "../../testdata/roms/misc/dmg-acid2/dmg-acid2.gb", context.Background())
		})

	})

}
