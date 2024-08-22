package cartridge

import (
	"github.com/nitwhiz/gameboy/pkg/bits"
	"slices"
)

type MBC1 struct {
	ROM Memory
	RAM Memory

	Is1M bool

	Bank1      byte
	Bank1Width byte
	Bank2      byte
	Mode       byte

	RomBankMask byte
	RamBankMask byte

	RAMEnabled bool
	HasRAM     bool
}

func NewMBC1(data []byte, romSize int, ramSize int) *MBC1 {
	rom := make([]byte, romSize)
	ram := make([]byte, ramSize)

	copy(rom, data)

	romBankMask := bits.GetAllOnes(bits.GetCountIn(romSize/0x4000) - 1)
	ramBankMask := bits.GetAllOnes(bits.GetCountIn(ramSize/0x2000) - 1)

	bank1Width := byte(5)

	if romSize >= 0x44000 && slices.Equal(rom[0x104:0x104+0x30], rom[0x40104:0x40104+0x30]) {
		bank1Width = 4
	}

	return &MBC1{
		ROM:         rom,
		RAM:         ram,
		Bank1:       1,
		Bank2:       0,
		Bank1Width:  bank1Width,
		Mode:        0,
		RomBankMask: romBankMask,
		RamBankMask: ramBankMask,
		RAMEnabled:  false,
		HasRAM:      ramSize > 0,
	}
}

func (c *MBC1) Read(address uint16) byte {
	switch {
	case address < 0x4000:
		bank := byte(0)

		if c.Mode == 1 {
			bank = c.Bank2 << c.Bank1Width
		}

		bank &= c.RomBankMask

		return c.ROM.Read(int(address) + (int(bank) * 0x4000))
	case address < 0x8000:
		bank := ((c.Bank2 << c.Bank1Width) | c.Bank1) & c.RomBankMask

		return c.ROM.Read(int(address-0x4000) + (int(bank) * 0x4000))
	case address >= 0xA000 && address < 0xC000:
		if !c.HasRAM || !c.RAMEnabled {
			return 0xFF
		}

		bank := byte(0)

		if c.Mode == 1 {
			bank = c.Bank2
		}

		bank &= c.RamBankMask

		return c.RAM.Read(int(address-0xA000) + (int(bank) * 0x2000))
	default:
		return 0xFF
	}
}

func (c *MBC1) WriteROM(address uint16, v byte) {
	switch {
	case address < 0x2000:
		// RAM Enable

		ln := v & 0x0F

		if ln == 0x0A {
			c.RAMEnabled = true
		} else if ln == 0 {
			c.RAMEnabled = false
		}
	case address < 0x4000:
		// Bank1

		b := v & 0b11111

		if b == 0 {
			b = 1
		}

		c.Bank1 = (b << (8 - c.Bank1Width)) >> (8 - c.Bank1Width)
	case address < 0x6000:
		// Bank2

		c.Bank2 = v & 0b11
	case address < 0x8000:
		// Banking Mode Select

		c.Mode = v & 0b1
	}
}

func (c *MBC1) WriteRAM(address uint16, v byte) {
	if c.HasRAM && c.RAMEnabled {
		bank := byte(0)

		if c.Mode == 1 {
			bank = c.Bank2
		}

		bank &= c.RamBankMask

		c.RAM.Write(int(address-0xA000)+(int(bank)*0x2000), v)
	}
}

func (c *MBC1) GetRAM() []byte {
	data := make([]byte, len(c.RAM))

	copy(data, c.RAM)

	return data
}

func (c *MBC1) SetRAM(data []byte) {
	copy(c.RAM, data)
}
