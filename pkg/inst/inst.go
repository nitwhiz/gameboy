package inst

import (
	"errors"
	"fmt"
	"github.com/nitwhiz/gameboy/pkg/gb"
	"log"
)

type handler func(g *gb.GameBoy) (ticks byte)

type table [0x100]handler

// h - instruction handler table
var h = &table{}

// p - prefixed instruction handler table
var p = &table{}

func (i *table) add(code byte, inst handler) {
	if foundI := i[code]; foundI != nil {
		log.Printf("code %2X is already defined\n", code)
		return
	}

	i[code] = inst
}

func (i *table) handler(code byte) handler {
	return i[code]
}

func (i *table) executeNextOpcode(g *gb.GameBoy) (ticks byte) {
	code := g.Fetch8()
	hand := i.handler(code)

	if hand == nil {
		panic(errors.New(fmt.Sprintf("missing handler for opcode 0x%02X", code)))
	}

	return hand(g)
}

func InitHandlers() {
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
}

// ExecuteNextOpcode executes the next opcode.
// returns the amount of system clock ticks it consumed.
func ExecuteNextOpcode(g *gb.GameBoy) (ticks byte) {
	return h.executeNextOpcode(g)
}
