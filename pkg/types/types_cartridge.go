package types

type CartridgeType byte

type CartridgeMode byte

type BankingController interface {
	Read(address uint16) byte
	WriteROM(address uint16, v byte)
	WriteRAM(address uint16, v byte)
	GetRAM() []byte
	SetRAM(data []byte)
}

type Cartridge interface {
	BankingController() BankingController
	Type() CartridgeType
}
