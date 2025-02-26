package cpu

import (
	"fmt"
	"github.com/nitwhiz/gameboy/pkg/types"
	"log/slog"
	"sync"
)

type handler func(g types.GameBoy) (ticks byte)

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

func (i *table) ExecuteNextOpcode(g types.GameBoy) (ticks byte) {
	code := g.CPU().Fetch8()
	hand := i.handler(code)

	if hand == nil {
		return 0
	}

	return hand(g)
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
