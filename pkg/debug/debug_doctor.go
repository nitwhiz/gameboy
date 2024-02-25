package debug

import (
	"fmt"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"os"
)

var logFile *os.File

func CurrentState(g *gb.GameBoy) string {
	return fmt.Sprintf(
		"A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n",
		g.CPU.AF.Hi(),
		g.CPU.AF.Lo(),
		g.CPU.BC.Hi(),
		g.CPU.BC.Lo(),
		g.CPU.DE.Hi(),
		g.CPU.DE.Lo(),
		g.CPU.HL.Hi(),
		g.CPU.HL.Lo(),
		g.CPU.SP.Val(),
		g.CPU.PC.Val(),
		g.MMU.Read(g.CPU.PC.Val()),
		g.MMU.Read(g.CPU.PC.Val()+1),
		g.MMU.Read(g.CPU.PC.Val()+2),
		g.MMU.Read(g.CPU.PC.Val()+3),
	)
}

func LogCurrentState(g *gb.GameBoy) {
	if logFile == nil {
		f, err := os.Create("/tmp/gb.log")

		if err != nil {
			panic(err)
		}

		logFile = f
	}

	_, err := logFile.WriteString(CurrentState(g))

	if err != nil {
		panic(err)
	}
}
