package mmu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/cartridge"
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/memory"
)

type Writer interface {
}

type MMU struct {
	Cartridge *cartridge.Cartridge
	Memory    *memory.Memory
	Input     *input.State

	SerialReceiver func(byte)
}

func inRange(a, l, u uint16) bool {
	return a >= l && a <= u
}

func (m *MMU) Read(address uint16) byte {
	switch {
	case address == addr.IE:
		return m.Memory.IE
	case inRange(address, addr.MemROMBegin, addr.MemROMEnd):
		return m.Cartridge.Read(address)
	case inRange(address, addr.MemVRAMBegin, addr.MemVRAMEnd):
		return m.Memory.VRAM[address-addr.MemVRAMBegin]
	case inRange(address, addr.MemCartridgeRAMBegin, addr.MemCartridgeRAMEnd):
		return m.Cartridge.Read(address)
	case inRange(address, addr.MemWRAMBegin, addr.MemWRAMEnd):
		return m.Memory.WRAM[address-addr.MemWRAMBegin]
	case inRange(address, addr.MemOAMBegin, addr.MemOAMEnd):
		return m.Memory.OAM[address-addr.MemOAMBegin]
	case inRange(address, addr.MemIOBegin, addr.MemIOEnd):
		return m.readIO(address)
	case inRange(address, addr.MemHRAMBegin, addr.MemHRAMEnd):
		return m.Memory.HRAM[address-addr.MemHRAMBegin]
	default:
		return 0xFF
	}
}

func (m *MMU) readIO(address uint16) byte {
	switch {
	case address == addr.JOYP:
		v := m.Memory.IO[address-addr.MemIOBegin] & 0xF0

		if bits.IsJOYPSelectButtons(v) {
			return v | m.Input.Get(input.SelectButtons)
		} else if bits.IsJOYPSelectDPad(v) {
			return v | m.Input.Get(input.SelectDPad)
		}

		return v | 0x0F
	default:
		return m.Memory.IO[address-addr.MemIOBegin]
	}
}

func (m *MMU) Write(address uint16, v byte) {
	switch {
	case inRange(address, addr.MemROMBegin, addr.MemROMEnd):
		m.Cartridge.WriteROM(address, v)
	case inRange(address, addr.MemVRAMBegin, addr.MemVRAMEnd):
		m.Memory.VRAM[address-addr.MemVRAMBegin] = v
	case inRange(address, addr.MemCartridgeRAMBegin, addr.MemCartridgeRAMEnd):
		m.Cartridge.WriteRAM(address, v)
	case inRange(address, addr.MemWRAMBegin, addr.MemWRAMEnd):
		m.Memory.WRAM[address-addr.MemWRAMBegin] = v
	case inRange(address, addr.MemOAMBegin, addr.MemOAMEnd):
		m.Memory.OAM[address-addr.MemOAMBegin] = v
	case inRange(address, addr.MemIOBegin, addr.MemIOEnd):
		m.writeIO(address, v)
	case inRange(address, addr.MemHRAMBegin, addr.MemHRAMEnd):
		m.Memory.HRAM[address-addr.MemHRAMBegin] = v
	case address == addr.IE:
		m.Memory.IE = v
	}
}

func (m *MMU) writeIO(address uint16, v byte) {
	switch {
	case inRange(address, addr.MemAudioBegin, addr.MemAudioEnd):
		// not implemented
	case inRange(address, addr.MemWaveBegin, addr.MemWaveEnd):
		// not implemented
	case address == addr.JOYP:
		m.Memory.IO[address-addr.MemIOBegin] = v & 0b00110000
	case address == addr.SC:
		if v == 0x81 {
			if m.SerialReceiver != nil {
				m.SerialReceiver(m.Read(addr.SB))
			}
		}
	case address == addr.DIV:
		m.SetDIV(0)
	case address == addr.STAT:
		m.Memory.IO[address-addr.MemIOBegin] = v | 0x80
	case address == addr.LY:
		m.SetLY(0)
	case address == addr.DMA:
		m.dmaTransfer(v)
	default:
		m.Memory.IO[address-addr.MemIOBegin] = v
	}
}

func (m *MMU) SetDIV(v byte) {
	m.Memory.IO[addr.DIV-addr.MemIOBegin] = v
}

func (m *MMU) IncLY() byte {
	v := m.Memory.IO[addr.LY-addr.MemIOBegin] + 1

	m.SetLY(v)

	return v
}

func (m *MMU) SetLY(v byte) {
	m.Memory.IO[addr.LY-addr.MemIOBegin] = v
}

func (m *MMU) SetJOYP(v byte) {
	m.Memory.IO[addr.JOYP-addr.MemIOBegin] = v
}

func (m *MMU) dmaTransfer(v byte) {
	address := uint16(v) << 8

	for i := uint16(0); i < 0xA0; i++ {
		m.Write(0xFE00+i, m.Read(address+i))
	}
}
