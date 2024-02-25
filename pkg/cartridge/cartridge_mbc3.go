package cartridge

type MBC3 struct {
	ROM     []byte
	ROMBank uint32

	RAM     []byte
	RAMBank uint32

	RAMEnabled bool

	RTC        [5]byte
	RTCLatched bool
	RTCLatch   byte
}

func NewMBC3(data []byte, romSize int, ramSize int) *MBC3 {
	rom := make([]byte, romSize)
	ram := make([]byte, ramSize)

	copy(rom, data)

	return &MBC3{
		ROM:        rom,
		ROMBank:    1,
		RAM:        ram,
		RAMBank:    0,
		RAMEnabled: false,
		RTC:        [5]byte{},
		RTCLatched: false,
		RTCLatch:   0x0,
	}
}

func isRTCRegister(bank uint32) bool {
	return bank >= 0x08 && bank <= 0x0C
}

func (c *MBC3) readRTC() byte {
	return c.RTC[c.RAMBank-0x08]
}

func (c *MBC3) writeRTC(v byte) {
	if !c.RTCLatched {
		c.RTC[c.RAMBank-0x08] = v
	}
}

func (c *MBC3) Read(address uint16) byte {
	switch {
	case address < 0x4000:
		return c.ROM[address]
	case address < 0x8000:
		// ROM Bank
		return c.ROM[uint32(address-0x4000)+(c.ROMBank*0x4000)]
	default:
		if isRTCRegister(c.RAMBank) {
			return c.readRTC()
		}

		// RAM Bank
		return c.RAM[(0x2000*c.RAMBank)+uint32(address-0xA000)]
	}
}

func (c *MBC3) WriteROM(address uint16, v byte) {
	switch {
	case address < 0x2000:
		// RAM Enable

		c.RAMEnabled = (v & 0xA) != 0
	case address < 0x4000:
		// ROM Bank Number

		b := uint32(v & 0x7F)

		if b == 0 {
			b += 1
		}

		c.ROMBank = b
	case address < 0x6000:
		c.RAMBank = uint32(v)
	case address < 0x8000:
		previousValue := c.RTCLatch

		if previousValue == 0x0 && v == 0x1 {
			c.RTCLatched = !c.RTCLatched
		}

		c.RTCLatch = v
	}
}

func (c *MBC3) WriteRAM(address uint16, v byte) {
	if c.RAMEnabled {
		if isRTCRegister(c.RAMBank) {
			c.writeRTC(v)
		} else {
			c.RAM[(0x2000*c.RAMBank)+uint32(address-0xA000)] = v
		}
	}
}

func (c *MBC3) GetRAM() []byte {
	data := make([]byte, len(c.RAM))

	copy(data, c.RAM)

	return data
}

func (c *MBC3) SetRAM(data []byte) {
	copy(c.RAM, data)
}
