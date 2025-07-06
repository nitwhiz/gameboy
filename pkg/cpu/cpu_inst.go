package cpu

import (
	"fmt"
	"github.com/nitwhiz/gameboy/pkg/types"
	"log/slog"
	"os"
	"sync"
)

type handler func(g types.GameBoy)

type table [0x100]handler

// H - instruction handler table
var H = &table{}

// p - prefixed instruction handler table
var p = &table{}

var initialized = false
var initLock = &sync.Mutex{}

func (i *table) add(code byte, inst handler) {
	if foundI := i[code]; foundI != nil {
		slog.Warn("code is already defined", "code", fmt.Sprintf("%2X", code))
		return
	}

	i[code] = inst
}

func (i *table) handler(code byte) handler {
	return i[code]
}

var logFile, _ = os.Create("/tmp/gb.log")

func WriteGameBoyDoctorLog(g types.GameBoy) {
	pc := g.CPU().PC().Val()

	pcMem0 := g.MMU().Read(pc)
	pcMem1 := g.MMU().Read(pc + 1)
	pcMem2 := g.MMU().Read(pc + 2)
	pcMem3 := g.MMU().Read(pc + 3)

	_, _ = logFile.WriteString(
		fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n",
			g.CPU().AF().Hi(),
			g.CPU().AF().Lo(),
			g.CPU().BC().Hi(),
			g.CPU().BC().Lo(),
			g.CPU().DE().Hi(),
			g.CPU().DE().Lo(),
			g.CPU().HL().Hi(),
			g.CPU().HL().Lo(),
			g.CPU().SP().Val(),
			pc,
			pcMem0, pcMem1, pcMem2, pcMem3,
		),
	)
}

// ExecuteNextOpcode
// Deprecated: PREFIX needs this here, but it reads HORRIBLE to call cpu.H.ExecuteNextOpcode(gb) within GameBoy
func (i *table) ExecuteNextOpcode(g types.GameBoy) {
	op := g.Fetch8()
	hand := i.handler(op)

	if hand != nil {
		hand(g)
	}

	if op != 0xCB {
		WriteGameBoyDoctorLog(g)
	}
}

func InitHandlers() {
	initLock.Lock()
	defer initLock.Unlock()

	if initialized {
		return
	}

	addADDHandlers()
	addADCHandlers()
	addINCHandlers()

	addSUBHandlers()
	addSBCHandlers()
	addDECHandlers()

	addANDHandlers()
	addXORHandlers()
	addORHandlers()

	addCPHandlers()

	addLDHandlers()

	addControlHandlers()

	addPUSHHandlers()
	addPOPHandlers()

	addJRHandlers()
	addJPHandlers()
	addCALLHandlers()
	addRETHandlers()
	AddRSTHandlers()

	addBitInstructions()

	initPHandlers()

	initialized = true
}
