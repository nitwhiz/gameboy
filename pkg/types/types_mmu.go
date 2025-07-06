package types

type MMU interface {
	Read(address uint16) byte
	Write(address uint16, v byte)

	SetCartridge(cartridge Cartridge)
	Cartridge() Cartridge

	SetSerialReceiver(receiver func(byte))

	CheckLYCLY()

	Timer() Timer
}
