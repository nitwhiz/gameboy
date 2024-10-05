package cartridge

import (
	"encoding/json"
	"errors"
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/types"
)

const (
	CGBAndDMGMode = types.CartridgeMode(0x80)
	CGBOnlyMode   = types.CartridgeMode(0xC0)
)

const (
	TypeROM                        = types.CartridgeType(0x00)
	TypeMBC1                       = types.CartridgeType(0x01)
	TypeMBC1RAM                    = types.CartridgeType(0x02)
	TypeMBC1RAMBattery             = types.CartridgeType(0x03)
	TypeMBC2                       = types.CartridgeType(0x05)
	TypeMBC2Battery                = types.CartridgeType(0x06)
	TypeROMRAM                     = types.CartridgeType(0x08)
	TypeROMRAMBattery              = types.CartridgeType(0x09)
	TypeMmm01                      = types.CartridgeType(0x0B)
	TypeMmm01RAM                   = types.CartridgeType(0x0C)
	TypeMmm01RAMBattery            = types.CartridgeType(0x0D)
	TypeMBC3TimerBattery           = types.CartridgeType(0x0F)
	TypeMBC3TimerRAMBattery        = types.CartridgeType(0x10)
	TypeMBC3                       = types.CartridgeType(0x11)
	TypeMBC3RAM                    = types.CartridgeType(0x12)
	TypeMBC3RAMBattery             = types.CartridgeType(0x13)
	TypeMBC5                       = types.CartridgeType(0x19)
	TypeMBC5RAM                    = types.CartridgeType(0x1A)
	TypeMBC5RAMBattery             = types.CartridgeType(0x1B)
	TypeMBC5Rumble                 = types.CartridgeType(0x1C)
	TypeMBC5RumbleRAM              = types.CartridgeType(0x1D)
	TypeMBC5RumbleRAMBattery       = types.CartridgeType(0x1E)
	TypeMBC6                       = types.CartridgeType(0x20)
	TypeMBC7SensorRumbleRAMBattery = types.CartridgeType(0x22)
	TypePocket                     = types.CartridgeType(0xFC)
	TypeBandai                     = types.CartridgeType(0xFD)
	TypeHuc3                       = types.CartridgeType(0xFE)
	TypeHuc1RAMBattery             = types.CartridgeType(0xFF)
)

type Cartridge struct {
	typ               types.CartridgeType
	bankingController types.BankingController
}

func (c *Cartridge) Type() types.CartridgeType {
	return c.typ
}

func (c *Cartridge) BankingController() types.BankingController {
	return c.bankingController
}

func (c *Cartridge) UnmarshalJSON(bs []byte) error {
	// todo
	var jsonCartridge struct {
		Type              types.CartridgeType
		BankingController json.RawMessage
	}

	if err := json.Unmarshal(bs, &jsonCartridge); err != nil {
		return err
	}

	c.typ = jsonCartridge.Type

	switch c.typ {
	case TypeROM:
		c.bankingController = &ROM{}
	case TypeMBC1:
		fallthrough
	case TypeMBC1RAM, TypeMBC1RAMBattery:
		c.bankingController = &MBC1{}
	case TypeMBC3, TypeMBC3TimerBattery:
		fallthrough
	case TypeMBC3RAM, TypeMBC3RAMBattery, TypeMBC3TimerRAMBattery:
		c.bankingController = &MBC3{}
	default:
		return errors.New("cartridge type is not supported")
	}

	return json.Unmarshal(jsonCartridge.BankingController, c.bankingController)
}

func getRamSize(romData []byte) int {
	ramSizeModifier := 0

	switch romData[addr.CartridgeRamSize] {
	case 0x02:
		ramSizeModifier = 1
	case 0x03:
		ramSizeModifier = 4
	case 0x04:
		ramSizeModifier = 16
	case 0x05:
		ramSizeModifier = 8
	}

	return 0x2000 * ramSizeModifier
}

func New(romData []byte) (*Cartridge, error) {
	c := Cartridge{}

	if types.CartridgeMode(romData[addr.CartridgeCGBMode]) == CGBOnlyMode {
		return nil, errors.New("CGB only is not supported")
	}

	cType := types.CartridgeType(romData[addr.CartridgeType])
	romSize := 0x8000 * (1 << romData[addr.CartridgeRomSize])

	c.typ = cType

	switch c.typ {
	case TypeROM:
		c.bankingController = NewROM(romData, romSize)
	case TypeMBC1:
		c.bankingController = NewMBC1(romData, romSize, 0)
	case TypeMBC1RAM, TypeMBC1RAMBattery:
		c.bankingController = NewMBC1(romData, romSize, getRamSize(romData))
	case TypeMBC3, TypeMBC3TimerBattery:
		c.bankingController = NewMBC3(romData, romSize, 0)
	case TypeMBC3RAM, TypeMBC3RAMBattery, TypeMBC3TimerRAMBattery:
		c.bankingController = NewMBC3(romData, romSize, getRamSize(romData))
	default:
		return nil, errors.New("cartridge type is not supported")
	}

	return &c, nil
}
