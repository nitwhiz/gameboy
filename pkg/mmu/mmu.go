package mmu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/memory"
	"github.com/nitwhiz/gameboy/pkg/types"
)

type MMU struct {
	cartridge types.Cartridge
	input     types.InputState
	memory    *memory.Container
	timer     types.Timer

	timerLock      bool
	serialReceiver func(byte)
}

func New(in types.InputState, memory *memory.Container, timer types.Timer) *MMU {
	return &MMU{
		cartridge: nil,
		input:     in,
		memory:    memory,
		timer:     timer,

		timerLock:      false,
		serialReceiver: nil,
	}
}

func inRange(a, l, u uint16) bool {
	return a >= l && a <= u
}

func addrVRAM(address uint16) uint16 {
	return address - addr.MemVRAMBegin
}

func addrWRAM(address uint16) uint16 {
	return address - addr.MemWRAMBegin
}

func addrOAM(address uint16) uint16 {
	return address - addr.MemOAMBegin
}

func addrHRAM(address uint16) uint16 {
	return address - addr.MemHRAMBegin
}

func addrIO(address uint16) uint16 {
	return address - addr.MemIOBegin
}

func (m *MMU) mappedWrite(address uint16, v byte) {
	switch {
	case address == addr.DIV:
		m.timer.SetValue(0)
	case inRange(address, addr.MemAudioBegin, addr.MemAudioEnd):
		// not implemented
		return
	case inRange(address, addr.MemWaveBegin, addr.MemWaveEnd):
		// not implemented
		return
	case inRange(address, addr.MemROMBegin, addr.MemROMEnd):
		m.cartridge.BankingController().WriteROM(address, v)
	case inRange(address, addr.MemVRAMBegin, addr.MemVRAMEnd):
		m.memory.VRAM.Write(addrVRAM(address), v)
	case inRange(address, addr.MemCartridgeRAMBegin, addr.MemCartridgeRAMEnd):
		m.cartridge.BankingController().WriteRAM(address, v)
	case inRange(address, addr.MemWRAMBegin, addr.MemWRAMEnd):
		m.memory.WRAM.Write(addrWRAM(address), v)
	case inRange(address, addr.MemOAMBegin, addr.MemOAMEnd):
		m.memory.OAM.Write(addrOAM(address), v)
	case inRange(address, addr.MemIOBegin, addr.MemIOEnd):
		m.writeIO(address, v)
	case inRange(address, addr.MemHRAMBegin, addr.MemHRAMEnd):
		m.memory.HRAM.Write(addrHRAM(address), v)
	}
}

func (m *MMU) mappedRead(address uint16) byte {
	switch {
	case address == addr.DIV:
		return byte((m.timer.GetValue() & 0xFF00) >> 8)
	case address == addr.JOYP:
		v := m.memory.IO.Read(addrIO(address)) & 0xF0

		if bits.IsJOYPSelectButtons(v) {
			return v | m.input.Value(types.InputSelectButtons)
		} else if bits.IsJOYPSelectDPad(v) {
			return v | m.input.Value(types.InputSelectDPad)
		}

		return v | 0x0F | memory.GetUnusedBits(addr.JOYP)
	case inRange(address, addr.MemROMBegin, addr.MemROMEnd):
		return m.cartridge.BankingController().Read(address)
	case inRange(address, addr.MemVRAMBegin, addr.MemVRAMEnd):
		return m.memory.VRAM.Read(addrVRAM(address))
	case inRange(address, addr.MemCartridgeRAMBegin, addr.MemCartridgeRAMEnd):
		return m.cartridge.BankingController().Read(address)
	case inRange(address, addr.MemWRAMBegin, addr.MemWRAMEnd):
		return m.memory.WRAM.Read(addrWRAM(address))
	case inRange(address, addr.MemOAMBegin, addr.MemOAMEnd):
		return m.memory.OAM.Read(addrOAM(address))
	case inRange(address, addr.MemIOBegin, addr.MemIOEnd):
		return m.memory.IO.Read(addrIO(address))
	case inRange(address, addr.MemHRAMBegin, addr.MemHRAMEnd):
		return m.memory.HRAM.Read(addrHRAM(address))
	default:
		return 0xFF
	}
}

func (m *MMU) Read(address uint16) byte {
	return m.mappedRead(address)
}

func (m *MMU) Write(address uint16, v byte) {
	m.mappedWrite(address, v)
}

func (m *MMU) HasCartridge() bool {
	return m.cartridge != nil
}

func (m *MMU) SetCartridge(cartridge types.Cartridge) {
	m.cartridge = cartridge
}

func (m *MMU) SetSerialReceiver(receiver func(byte)) {
	m.serialReceiver = receiver
}

func (m *MMU) RequestInterrupt(typ types.InterruptType) {
	m.Write(addr.IF, bits.Set(m.Read(addr.IF), byte(typ)))
}

func (m *MMU) CheckLYCLY() {
	ly := m.Read(addr.LY)
	lyc := m.Read(addr.LYC)
	stat := m.Read(addr.STAT)

	if ly == lyc {
		stat = bits.Set(stat, addr.STAT_COINCIDENCE_FLAG)

		if bits.Test(stat, addr.STAT_LYCLY_INTERRUPT_ENABLE) {
			m.RequestInterrupt(addr.InterruptLCD)
		}
	} else {
		stat = bits.Reset(stat, addr.STAT_COINCIDENCE_FLAG)
	}

	m.Write(addr.STAT, stat)
}

func (m *MMU) Timer() types.Timer {
	return m.timer
}

func (m *MMU) writeIO(address uint16, v byte) {
	if isUnmapped(address) {
		return
	}

	v |= memory.GetUnusedBits(address)

	switch {
	case address == addr.JOYP:
		m.memory.IO.Write(addrIO(address), v&0b00110000)
		return
	case address == addr.SC:
		if bits.Test(v, 7) && bits.Test(v, 0) {
			if m.serialReceiver != nil {
				m.serialReceiver(m.Read(addr.SB))
			}
		}
	case address == addr.LCDC:
		if !bits.Test(v, addr.LCDC_ENABLE) {
			m.Write(addr.LY, 0)
			return
		}
	case address == addr.DMA:
		m.dmaTransfer(v)
		return
	case address == addr.TIMA:
		if m.timerLock {
			return
		}
	case address == addr.LYC:
		m.memory.IO.Write(addrIO(address), v)
		m.CheckLYCLY()
		return
	default:
	}

	m.memory.IO.Write(addrIO(address), v)
}

func (m *MMU) dmaTransfer(v byte) {
	address := uint16(v) << 8

	for i := uint16(0); i < 0xA0; i++ {
		m.Write(0xFE00+i, m.Read(address+i))
	}
}

func isUnmapped(address uint16) bool {
	if address < 0xFF03 || address > 0xFF7F {
		return false
	}

	if address >= 0xFF4C {
		return true
	}

	return address == 0xFF03 ||
		address == 0xFF08 ||
		address == 0xFF09 ||
		address == 0xFF0A ||
		address == 0xFF0B ||
		address == 0xFF0C ||
		address == 0xFF0D ||
		address == 0xFF0E ||
		address == 0xFF15 ||
		address == 0xFF1F ||
		address == 0xFF27 ||
		address == 0xFF28 ||
		address == 0xFF29
}
