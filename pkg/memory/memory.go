package memory

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/mmu"
)

type Memory struct {
	vram [0x2000]byte
	wram [0x2000]byte
	oam  [0x0100]byte
	hram [0x0080]byte
	io   [0x0080]byte

	timerCounter uint16
}

func New() *Memory {
	m := Memory{
		vram: [0x2000]byte{},
		wram: [0x2000]byte{},
		oam:  [0x0100]byte{},
		hram: [0x0080]byte{},
		io:   [0x0080]byte{},

		timerCounter: 0xAC00,
	}

	for i := range uint16(0x0080) {
		m.WriteIO(i, 0xFF)
	}

	m.WriteHRAM(0x00, 0xCF)
	m.WriteHRAM(0x01, 0x00)
	m.WriteHRAM(0x02, 0x7E)
	m.WriteHRAM(0x04, 0xAB)
	m.WriteHRAM(0x05, 0x00)
	m.WriteHRAM(0x06, 0x00)
	m.WriteHRAM(0x07, 0xF8)
	m.WriteHRAM(0x0F, 0xE1)
	m.WriteHRAM(0x10, 0x80)
	m.WriteHRAM(0x11, 0xBF)
	m.WriteHRAM(0x12, 0xF3)
	m.WriteHRAM(0x13, 0xFF)
	m.WriteHRAM(0x14, 0xBF)
	m.WriteHRAM(0x16, 0x3F)
	m.WriteHRAM(0x17, 0x00)
	m.WriteHRAM(0x18, 0xFF)
	m.WriteHRAM(0x19, 0xBF)
	m.WriteHRAM(0x1A, 0x7F)
	m.WriteHRAM(0x1B, 0xFF)
	m.WriteHRAM(0x1C, 0x9F)
	m.WriteHRAM(0x1D, 0xFF)
	m.WriteHRAM(0x1E, 0xBF)
	m.WriteHRAM(0x20, 0xFF)
	m.WriteHRAM(0x21, 0x00)
	m.WriteHRAM(0x22, 0x00)
	m.WriteHRAM(0x23, 0xBF)
	m.WriteHRAM(0x24, 0x77)
	m.WriteHRAM(0x25, 0xF3)
	m.WriteHRAM(0x26, 0xF1)
	m.WriteHRAM(0x40, 0x91)
	m.WriteHRAM(0x41, 0x85)
	m.WriteHRAM(0x42, 0x00)
	m.WriteHRAM(0x43, 0x00)
	m.WriteHRAM(0x44, 0x00)
	m.WriteHRAM(0x45, 0x00)
	m.WriteHRAM(0x46, 0xFF)
	m.WriteHRAM(0x47, 0xFC)
	m.WriteHRAM(0x4A, 0x00)
	m.WriteHRAM(0x4B, 0x00)

	// JOYP
	m.WriteIO(0x00, 0xCF)
	// SB
	m.WriteIO(0x01, 0x00)
	// SC
	m.WriteIO(0x02, 0x7E)
	// TIMA
	m.WriteIO(0x05, 0x00)
	// TMA
	m.WriteIO(0x06, 0x00)
	// TAC
	m.WriteIO(0x07, 0xF8)
	// IF
	m.WriteIO(0x0F, 0xE1)

	// something audio
	m.WriteIO(0x11, 0xBF)
	m.WriteIO(0x12, 0xF3)
	m.WriteIO(0x13, 0xFF)
	m.WriteIO(0x14, 0xBF)
	m.WriteIO(0x15, 0xFF)
	m.WriteIO(0x16, 0x3F)
	m.WriteIO(0x17, 0x00)
	m.WriteIO(0x18, 0xFF)
	m.WriteIO(0x19, 0xBF)
	m.WriteIO(0x1A, 0x7F)
	m.WriteIO(0x1B, 0xFF)
	m.WriteIO(0x1C, 0x9F)
	m.WriteIO(0x1D, 0xFF)
	m.WriteIO(0x1E, 0xBF)
	m.WriteIO(0x1F, 0xFF)
	m.WriteIO(0x20, 0xFF)
	m.WriteIO(0x21, 0x00)
	m.WriteIO(0x22, 0x00)
	m.WriteIO(0x23, 0xBF)
	m.WriteIO(0x24, 0x77)
	m.WriteIO(0x25, 0xF3)
	m.WriteIO(0x26, 0xF1)

	// misc?
	m.WriteIO(0x2A, 0xFF)
	m.WriteIO(0x2B, 0xFF)
	m.WriteIO(0x2C, 0xFF)
	m.WriteIO(0x2D, 0xFF)
	m.WriteIO(0x2E, 0xFF)
	m.WriteIO(0x2F, 0xFF)

	// display
	m.WriteIO(0x40, 0x91)
	m.WriteIO(0x41, 0x00)
	m.WriteIO(0x42, 0x00)
	m.WriteIO(0x43, 0x00)
	m.WriteIO(0x44, 0x0A)
	m.WriteIO(0x45, 0x00)
	m.WriteIO(0x46, 0xFF)
	m.WriteIO(0x47, 0xFC)

	for i := range uint16(0x0080) {
		m.WriteIO(i, m.ReadIO(i)|mmu.GetUnusedBits(addr.MemIOBegin+i))
	}

	return &m
}

func (m *Memory) AddrVRAM(address uint16) uint16 {
	return address - addr.MemVRAMBegin
}

func (m *Memory) ReadVRAM(address uint16) byte {
	return m.vram[address]
}

func (m *Memory) WriteVRAM(address uint16, v byte) {
	m.vram[address] = v
}

func (m *Memory) AddrWRAM(address uint16) uint16 {
	return address - addr.MemWRAMBegin
}

func (m *Memory) ReadWRAM(address uint16) byte {
	return m.wram[address]
}

func (m *Memory) WriteWRAM(address uint16, v byte) {
	m.wram[address] = v
}

func (m *Memory) AddrOAM(address uint16) uint16 {
	return address - addr.MemOAMBegin
}

func (m *Memory) ReadOAM(address uint16) byte {
	return m.oam[address]
}

func (m *Memory) WriteOAM(address uint16, v byte) {
	m.oam[address] = v
}

func (m *Memory) AddrHRAM(address uint16) uint16 {
	return address - addr.MemHRAMBegin
}

func (m *Memory) ReadHRAM(address uint16) byte {
	return m.hram[address]
}

func (m *Memory) WriteHRAM(address uint16, v byte) {
	m.hram[address] = v
}

func (m *Memory) AddrIO(address uint16) uint16 {
	return address - addr.MemIOBegin
}

func (m *Memory) ReadIO(address uint16) byte {
	return m.io[address]
}

func (m *Memory) WriteIO(address uint16, v byte) {
	m.io[address] = v
}

func (m *Memory) Div() byte {
	return byte((m.timerCounter & 0xFF00) >> 8)
}

func (m *Memory) TimerCounter() uint16 {
	return m.timerCounter
}

func (m *Memory) IncTimerCounter() {
	m.timerCounter++
}

func (m *Memory) ResetTimerCounter() {
	m.timerCounter = 0
}
