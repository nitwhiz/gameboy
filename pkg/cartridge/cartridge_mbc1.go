package cartridge

type MBC1 struct {
	ROM     []byte
	ROMBank uint32

	RAM     []byte
	RAMBank uint32

	RAMEnabled bool
	HasRAM     bool
	ROMBanking bool
}

func NewMBC1(data []byte, romSize int, ramSize int) *MBC1 {
	rom := make([]byte, romSize)
	ram := make([]byte, ramSize)

	copy(rom, data)

	return &MBC1{
		ROM:        rom,
		ROMBank:    1,
		RAM:        ram,
		RAMBank:    0,
		RAMEnabled: false,
		HasRAM:     ramSize > 0,
		ROMBanking: false,
	}
}

func (c *MBC1) setROMBank(bank uint32) {
	if bank == 0x00 || bank == 0x20 || bank == 0x40 || bank == 0x60 {
		c.ROMBank = bank + 1
	} else {
		c.ROMBank = bank
	}
}

func (c *MBC1) Read(address uint16) byte {
	switch {
	case address < 0x4000:
		return c.ROM[address]
	case address < 0x8000:
		// ROM Bank
		return c.ROM[uint32(address-0x4000)+(c.ROMBank*0x4000)]
	default:
		if !c.HasRAM {
			return 0xFF
		}

		// RAM Bank
		return c.RAM[(0x2000*c.RAMBank)+uint32(address-0xA000)]
	}
}

func (c *MBC1) WriteROM(address uint16, v byte) {
	switch {
	case address < 0x2000:
		// RAM Enable

		c.RAMEnabled = (v & 0xA) != 0
	case address < 0x4000:
		// ROM Bank Number

		c.setROMBank((c.ROMBank & 0xE0) | uint32(v&0x1F))
	case address < 0x6000:
		// RAM Bank Number or Upper Bits of ROM Bank Number

		if c.ROMBanking {
			c.setROMBank((c.ROMBank & 0x1F) | uint32(v&0xE0))
		} else {
			c.RAMBank = uint32(v & 0x3)
		}
	case address < 0x8000:
		// Banking Mode Select

		c.ROMBanking = v&0x1 == 0x00

		if c.ROMBanking {
			c.RAMBank = 0
		} else {
			c.ROMBank = c.ROMBank & 0x1F
		}
	}
}

func (c *MBC1) WriteRAM(address uint16, v byte) {
	if c.HasRAM && c.RAMEnabled {
		c.RAM[(0x2000*c.RAMBank)+uint32(address-0xA000)] = v
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
