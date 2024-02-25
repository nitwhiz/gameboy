package cartridge

import (
	"encoding/json"
	"errors"
	"github.com/nitwhiz/gameboy/pkg/addr"
)

type Mode byte

const (
	CGBAndDMGMode = Mode(0x80)
	CGBOnlyMode   = Mode(0xC0)
)

type Cartridge struct {
	Type Type
	BankingController
}

func (c *Cartridge) UnmarshalJSON(bs []byte) error {
	var jsonCartridge struct {
		Type              Type
		BankingController json.RawMessage
	}

	if err := json.Unmarshal(bs, &jsonCartridge); err != nil {
		return err
	}

	c.Type = jsonCartridge.Type

	switch c.Type {
	case TypeROM:
		c.BankingController = &ROM{}
	case TypeMBC1:
		fallthrough
	case TypeMBC1RAM, TypeMBC1RAMBattery:
		c.BankingController = &MBC1{}
	case TypeMBC3, TypeMBC3TimerBattery:
		fallthrough
	case TypeMBC3RAM, TypeMBC3RAMBattery, TypeMBC3TimerRAMBattery:
		c.BankingController = &MBC3{}
	default:
		return errors.New("cartridge type is not supported")
	}

	return json.Unmarshal(jsonCartridge.BankingController, c.BankingController)
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

	if Mode(romData[addr.CartridgeCGBMode]) == CGBOnlyMode {
		return nil, errors.New("CGB only is not supported")
	}

	cType := Type(romData[addr.CartridgeType])
	romSize := 0x8000 * (1 << romData[addr.CartridgeRomSize])

	c.Type = cType

	switch c.Type {
	case TypeROM:
		c.BankingController = NewROM(romData, romSize)
	case TypeMBC1:
		c.BankingController = NewMBC1(romData, romSize, 0)
	case TypeMBC1RAM, TypeMBC1RAMBattery:
		c.BankingController = NewMBC1(romData, romSize, getRamSize(romData))
	case TypeMBC3, TypeMBC3TimerBattery:
		c.BankingController = NewMBC3(romData, romSize, 0)
	case TypeMBC3RAM, TypeMBC3RAMBattery, TypeMBC3TimerRAMBattery:
		c.BankingController = NewMBC3(romData, romSize, getRamSize(romData))
	default:
		return nil, errors.New("cartridge type is not supported")
	}

	return &c, nil
}
