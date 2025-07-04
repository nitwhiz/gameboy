package types

type InterruptType byte

type MMU interface {
	Read(address uint16) byte
	Write(address uint16, v byte)

	SetCartridge(cartridge Cartridge)
	HasCartridge() bool

	RequestInterrupt(typ InterruptType)
	SetSerialReceiver(receiver func(byte))

	CheckLYCLY()

	Timer() Timer
}
