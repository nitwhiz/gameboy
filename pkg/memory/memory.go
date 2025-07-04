package memory

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/types"
)

type VRAM [0x2000]byte

func NewVRAM() *VRAM {
	return &VRAM{}
}

func (m *VRAM) Read(addr uint16) byte {
	return m[addr]
}

func (m *VRAM) Write(addr uint16, v byte) {
	m[addr] = v
}

type WRAM [0x2000]byte

func NewWRAM() *WRAM {
	return &WRAM{}
}

func (m *WRAM) Read(addr uint16) byte {
	return m[addr]
}

func (m *WRAM) Write(addr uint16, v byte) {
	m[addr] = v
}

type OAM [0x0100]byte

func NewOAM() *OAM {
	return &OAM{}
}

func (m *OAM) Read(addr uint16) byte {
	return m[addr]
}

func (m *OAM) Write(addr uint16, v byte) {
	m[addr] = v
}

type HRAM [0x0080]byte

func NewHRAM() *HRAM {
	m := HRAM{}

	m.Write(0x00, 0xCF)
	m.Write(0x01, 0x00)
	m.Write(0x02, 0x7E)
	m.Write(0x04, 0xAB)
	m.Write(0x05, 0x00)
	m.Write(0x06, 0x00)
	m.Write(0x07, 0xF8)
	m.Write(0x0F, 0xE1)
	m.Write(0x10, 0x80)
	m.Write(0x11, 0xBF)
	m.Write(0x12, 0xF3)
	m.Write(0x13, 0xFF)
	m.Write(0x14, 0xBF)
	m.Write(0x16, 0x3F)
	m.Write(0x17, 0x00)
	m.Write(0x18, 0xFF)
	m.Write(0x19, 0xBF)
	m.Write(0x1A, 0x7F)
	m.Write(0x1B, 0xFF)
	m.Write(0x1C, 0x9F)
	m.Write(0x1D, 0xFF)
	m.Write(0x1E, 0xBF)
	m.Write(0x20, 0xFF)
	m.Write(0x21, 0x00)
	m.Write(0x22, 0x00)
	m.Write(0x23, 0xBF)
	m.Write(0x24, 0x77)
	m.Write(0x25, 0xF3)
	m.Write(0x26, 0xF1)
	m.Write(0x40, 0x91)
	m.Write(0x41, 0x85)
	m.Write(0x42, 0x00)
	m.Write(0x43, 0x00)
	m.Write(0x44, 0x00)
	m.Write(0x45, 0x00)
	m.Write(0x46, 0xFF)
	m.Write(0x47, 0xFC)
	m.Write(0x4A, 0x00)
	m.Write(0x4B, 0x00)

	return &m
}

func (m *HRAM) Read(addr uint16) byte {
	return m[addr]
}

func (m *HRAM) Write(addr uint16, v byte) {
	m[addr] = v
}

type IO [0x0080]byte

func NewIO() *IO {
	m := IO{}

	for i := range uint16(0x0080) {
		m.Write(i, 0xFF)
	}

	// JOYP
	m.Write(0x00, 0xCF)
	// SB
	m.Write(0x01, 0x00)
	// SC
	m.Write(0x02, 0x7E)
	// TIMA
	m.Write(0x05, 0x00)
	// TMA
	m.Write(0x06, 0x00)
	// TAC
	m.Write(0x07, 0xF8)
	// IF
	m.Write(0x0F, 0xE1)

	// something audio
	m.Write(0x10, 0x80)
	m.Write(0x11, 0xBF)
	m.Write(0x12, 0xF3)
	m.Write(0x13, 0xFF)
	m.Write(0x14, 0xBF)
	m.Write(0x15, 0xFF)
	m.Write(0x16, 0x3F)
	m.Write(0x17, 0x00)
	m.Write(0x18, 0xFF)
	m.Write(0x19, 0xBF)
	m.Write(0x1A, 0x7F)
	m.Write(0x1B, 0xFF)
	m.Write(0x1C, 0x9F)
	m.Write(0x1D, 0xFF)
	m.Write(0x1E, 0xBF)
	m.Write(0x1F, 0xFF)
	m.Write(0x20, 0xFF)
	m.Write(0x21, 0x00)
	m.Write(0x22, 0x00)
	m.Write(0x23, 0xBF)
	m.Write(0x24, 0x77)
	m.Write(0x25, 0xF3)
	m.Write(0x26, 0xF1)

	// misc?
	m.Write(0x2A, 0xFF)
	m.Write(0x2B, 0xFF)
	m.Write(0x2C, 0xFF)
	m.Write(0x2D, 0xFF)
	m.Write(0x2E, 0xFF)
	m.Write(0x2F, 0xFF)

	// display
	m.Write(0x40, 0x91)
	m.Write(0x41, 0x00)
	m.Write(0x42, 0x00)
	m.Write(0x43, 0x00)
	m.Write(0x44, 0x0A)
	m.Write(0x45, 0x00)
	m.Write(0x46, 0xFF)
	m.Write(0x47, 0xFC)

	for i := range uint16(0x0080) {
		m.Write(i, m.Read(i)|GetUnusedBits(addr.MemIOBegin+i))
	}

	return &m
}

func (m *IO) Read(addr uint16) byte {
	return m[addr]
}

func (m *IO) Write(addr uint16, v byte) {
	m[addr] = v
}

type Container struct {
	VRAM types.Memory
	WRAM types.Memory
	OAM  types.Memory
	HRAM types.Memory
	IO   types.Memory
}

func NewContainer() *Container {
	return &Container{
		VRAM: NewVRAM(),
		WRAM: NewWRAM(),
		OAM:  NewOAM(),
		HRAM: NewHRAM(),
		IO:   NewIO(),
	}
}

// GetUnusedBits returns a byte with 1's for unused bits
func GetUnusedBits(address uint16) byte {
	switch address {
	case addr.JOYP:
		return 0b11000000
	case addr.SC:
		return 0b01111110
	case addr.TAC:
		return 0b11111000
	case addr.IF:
		return 0b11100000
	case 0xFF10:
		return 0b10000000
	case 0xFF1A:
		return 0b01111111
	case 0xFF1C:
		return 0b10011111
	case 0xFF20:
		return 0b11000000
	case 0xFF23:
		return 0b00111111
	case 0xFF26:
		return 0b01110000
	case addr.STAT:
		return 0b10000000
	default:
		return 0
	}
}
