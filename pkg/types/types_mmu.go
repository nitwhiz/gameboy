package types

type InterruptType byte

type MMU interface {
	Read(address uint16) byte
	Write(address uint16, v byte)
	RequestInterrupt(typ InterruptType)
	Cartridge() Cartridge
	SetCartridge(cartridge Cartridge)
	SetSerialReceiver(receiver func(byte))
	// todo: only quarz_timer and some instruction uses this. move calls to mmu
	Memory() Memory
	// todo: ppu functions - maybe solve differently?
	IncLY() byte
	ResetLY()
	CheckLYCLY()
}
