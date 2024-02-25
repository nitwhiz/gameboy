package cartridge

type ROM struct {
	ROM []byte
}

func NewROM(data []byte, romSize int) *ROM {
	rom := make([]byte, romSize)

	copy(rom, data)

	return &ROM{rom}
}

func (c *ROM) Read(address uint16) byte {
	return c.ROM[address]
}

func (c *ROM) WriteROM(uint16, byte) {}

func (c *ROM) WriteRAM(uint16, byte) {}

func (c *ROM) GetRAM() []byte {
	return nil
}

func (c *ROM) SetRAM([]byte) {}
