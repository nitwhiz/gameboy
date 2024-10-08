package memory

import "github.com/nitwhiz/gameboy/pkg/addr"

type Memory struct {
	VRAM [0x2000]byte
	WRAM [0x2000]byte
	OAM  [0x0100]byte
	HRAM [0x0080]byte

	IO [0x0080]byte

	// IE - Interrupt Enable
	IE byte
	// IF - Interrupt Flags
	IF byte
	// IME - Interrupt Master Enable
	IME byte

	TimerCounter uint16
}

func New() *Memory {
	m := Memory{
		VRAM: [0x2000]byte{},
		WRAM: [0x2000]byte{},
		OAM:  [0x0100]byte{},
		HRAM: [0x0080]byte{},
		IO:   [0x0080]byte{},

		IE: 0,
	}

	return &m
}

func (m *Memory) Init() *Memory {
	for i := range len(m.IO) {
		m.IO[i] = 0xFF
	}

	m.HRAM[0x00] = 0xCF
	m.HRAM[0x01] = 0x00
	m.HRAM[0x02] = 0x7E
	m.HRAM[0x04] = 0xAB
	m.HRAM[0x05] = 0x00
	m.HRAM[0x06] = 0x00
	m.HRAM[0x07] = 0xF8
	m.HRAM[0x0F] = 0xE1
	m.HRAM[0x10] = 0x80
	m.HRAM[0x11] = 0xBF
	m.HRAM[0x12] = 0xF3
	m.HRAM[0x13] = 0xFF
	m.HRAM[0x14] = 0xBF
	m.HRAM[0x16] = 0x3F
	m.HRAM[0x17] = 0x00
	m.HRAM[0x18] = 0xFF
	m.HRAM[0x19] = 0xBF
	m.HRAM[0x1A] = 0x7F
	m.HRAM[0x1B] = 0xFF
	m.HRAM[0x1C] = 0x9F
	m.HRAM[0x1D] = 0xFF
	m.HRAM[0x1E] = 0xBF
	m.HRAM[0x20] = 0xFF
	m.HRAM[0x21] = 0x00
	m.HRAM[0x22] = 0x00
	m.HRAM[0x23] = 0xBF
	m.HRAM[0x24] = 0x77
	m.HRAM[0x25] = 0xF3
	m.HRAM[0x26] = 0xF1
	m.HRAM[0x40] = 0x91
	m.HRAM[0x41] = 0x85
	m.HRAM[0x42] = 0x00
	m.HRAM[0x43] = 0x00
	m.HRAM[0x44] = 0x00
	m.HRAM[0x45] = 0x00
	m.HRAM[0x46] = 0xFF
	m.HRAM[0x47] = 0xFC
	m.HRAM[0x4A] = 0x00
	m.HRAM[0x4B] = 0x00

	// JOYP
	m.IO[0x00] = 0xCF
	// SB
	m.IO[0x01] = 0x00
	// SC
	m.IO[0x02] = 0x7E
	// TIMA
	m.IO[0x05] = 0x00
	// TMA
	m.IO[0x06] = 0x00
	// TAC
	m.IO[0x07] = 0xF8
	// IF
	m.IO[0x0F] = 0xE1

	// something audio
	m.IO[0x11] = 0xBF
	m.IO[0x12] = 0xF3
	m.IO[0x13] = 0xFF
	m.IO[0x14] = 0xBF
	m.IO[0x15] = 0xFF
	m.IO[0x16] = 0x3F
	m.IO[0x17] = 0x00
	m.IO[0x18] = 0xFF
	m.IO[0x19] = 0xBF
	m.IO[0x1A] = 0x7F
	m.IO[0x1B] = 0xFF
	m.IO[0x1C] = 0x9F
	m.IO[0x1D] = 0xFF
	m.IO[0x1E] = 0xBF
	m.IO[0x1F] = 0xFF
	m.IO[0x20] = 0xFF
	m.IO[0x21] = 0x00
	m.IO[0x22] = 0x00
	m.IO[0x23] = 0xBF
	m.IO[0x24] = 0x77
	m.IO[0x25] = 0xF3
	m.IO[0x26] = 0xF1

	// misc?
	m.IO[0x2A] = 0xFF
	m.IO[0x2B] = 0xFF
	m.IO[0x2C] = 0xFF
	m.IO[0x2D] = 0xFF
	m.IO[0x2E] = 0xFF
	m.IO[0x2F] = 0xFF

	// display
	m.IO[0x40] = 0x91
	m.IO[0x41] = 0x00
	m.IO[0x42] = 0x00
	m.IO[0x43] = 0x00
	m.IO[0x44] = 0x0A
	m.IO[0x45] = 0x00
	m.IO[0x46] = 0xFF
	m.IO[0x47] = 0xFC

	for i := range m.IO {
		m.IO[i] |= GetUnusedBits(addr.MemIOBegin + uint16(i))
	}

	return m
}
