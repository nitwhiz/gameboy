package cartridge

type Type byte

const (
	TypeROM                        = Type(0x00)
	TypeMBC1                       = Type(0x01)
	TypeMBC1RAM                    = Type(0x02)
	TypeMBC1RAMBattery             = Type(0x03)
	TypeMBC2                       = Type(0x05)
	TypeMBC2Battery                = Type(0x06)
	TypeROMRAM                     = Type(0x08)
	TypeROMRAMBattery              = Type(0x09)
	TypeMmm01                      = Type(0x0B)
	TypeMmm01RAM                   = Type(0x0C)
	TypeMmm01RAMBattery            = Type(0x0D)
	TypeMBC3TimerBattery           = Type(0x0F)
	TypeMBC3TimerRAMBattery        = Type(0x10)
	TypeMBC3                       = Type(0x11)
	TypeMBC3RAM                    = Type(0x12)
	TypeMBC3RAMBattery             = Type(0x13)
	TypeMBC5                       = Type(0x19)
	TypeMBC5RAM                    = Type(0x1A)
	TypeMBC5RAMBattery             = Type(0x1B)
	TypeMBC5Rumble                 = Type(0x1C)
	TypeMBC5RumbleRAM              = Type(0x1D)
	TypeMBC5RumbleRAMBattery       = Type(0x1E)
	TypeMBC6                       = Type(0x20)
	TypeMBC7SensorRumbleRAMBattery = Type(0x22)
	TypePocket                     = Type(0xFC)
	TypeBandai                     = Type(0xFD)
	TypeHuc3                       = Type(0xFE)
	TypeHuc1RAMBattery             = Type(0xFF)
)

type BankingController interface {
	Read(address uint16) byte
	WriteROM(address uint16, v byte)
	WriteRAM(address uint16, v byte)
	GetRAM() []byte
	SetRAM(data []byte)
}
