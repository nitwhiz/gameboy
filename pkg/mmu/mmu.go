package mmu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/types"
)

type MMU struct {
	cartridge      types.Cartridge
	memory         types.Memory
	input          types.InputState
	timerLock      bool
	serialReceiver func(byte)
}

func New(in types.InputState, mem types.Memory) *MMU {
	return &MMU{
		cartridge:      nil,
		memory:         mem,
		input:          in,
		timerLock:      false,
		serialReceiver: nil,
	}
}

func inRange(a, l, u uint16) bool {
	return a >= l && a <= u
}

func (m *MMU) mappedWrite(address uint16, v byte) {
	switch {
	case address == addr.DIV:
		m.memory.ResetTimerCounter()
	case inRange(address, addr.MemAudioBegin, addr.MemAudioEnd):
		// not implemented
		return
	case inRange(address, addr.MemWaveBegin, addr.MemWaveEnd):
		// not implemented
		return
	case inRange(address, addr.MemROMBegin, addr.MemROMEnd):
		m.cartridge.BankingController().WriteROM(address, v)
	case inRange(address, addr.MemVRAMBegin, addr.MemVRAMEnd):
		m.memory.WriteVRAM(m.memory.AddrVRAM(address), v)
	case inRange(address, addr.MemCartridgeRAMBegin, addr.MemCartridgeRAMEnd):
		m.cartridge.BankingController().WriteRAM(address, v)
	case inRange(address, addr.MemWRAMBegin, addr.MemWRAMEnd):
		m.memory.WriteWRAM(address-addr.MemWRAMBegin, v)
	case inRange(address, addr.MemOAMBegin, addr.MemOAMEnd):
		m.memory.WriteOAM(m.memory.AddrOAM(address), v)
	case inRange(address, addr.MemIOBegin, addr.MemIOEnd):
		m.writeIO(address, v)
	case inRange(address, addr.MemHRAMBegin, addr.MemHRAMEnd):
		m.memory.WriteHRAM(address-addr.MemHRAMBegin, v)
	}
}

func (m *MMU) mappedRead(address uint16) byte {
	switch {
	case address == addr.DIV:
		return m.memory.Div()
	case address == addr.JOYP:
		v := m.memory.ReadIO(m.memory.AddrIO(address)) & 0xF0

		if bits.IsJOYPSelectButtons(v) {
			return v | m.input.Value(types.InputSelectButtons)
		} else if bits.IsJOYPSelectDPad(v) {
			return v | m.input.Value(types.InputSelectDPad)
		}

		return v | 0x0F | GetUnusedBits(addr.JOYP)
	case inRange(address, addr.MemROMBegin, addr.MemROMEnd):
		return m.cartridge.BankingController().Read(address)
	case inRange(address, addr.MemVRAMBegin, addr.MemVRAMEnd):
		return m.memory.ReadVRAM(address - addr.MemVRAMBegin)
	case inRange(address, addr.MemCartridgeRAMBegin, addr.MemCartridgeRAMEnd):
		return m.cartridge.BankingController().Read(address)
	case inRange(address, addr.MemWRAMBegin, addr.MemWRAMEnd):
		return m.memory.ReadWRAM(m.memory.AddrWRAM(address))
	case inRange(address, addr.MemOAMBegin, addr.MemOAMEnd):
		return m.memory.ReadOAM(m.memory.AddrOAM(address))
	case inRange(address, addr.MemIOBegin, addr.MemIOEnd):
		return m.memory.ReadIO(m.memory.AddrIO(address))
	case inRange(address, addr.MemHRAMBegin, addr.MemHRAMEnd):
		return m.memory.ReadHRAM(m.memory.AddrHRAM(address))
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

func (m *MMU) Cartridge() types.Cartridge {
	return m.cartridge
}

func (m *MMU) SetCartridge(cartridge types.Cartridge) {
	m.cartridge = cartridge
}

func (m *MMU) SetSerialReceiver(receiver func(byte)) {
	m.serialReceiver = receiver
}

func (m *MMU) Memory() types.Memory {
	return m.memory
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

func (m *MMU) writeIO(address uint16, v byte) {
	if IsUnmapped(address) {
		return
	}

	v |= GetUnusedBits(address)

	switch {
	case address == addr.JOYP:
		m.memory.WriteIO(address-addr.MemIOBegin, v&0b00110000)
		return
	case address == addr.SC:
		if bits.Test(v, 7) && bits.Test(v, 0) {
			if m.serialReceiver != nil {
				m.serialReceiver(m.Read(addr.SB))
			}
		}
	case address == addr.LCDC:
		if !bits.Test(v, addr.LCDC_ENABLE) {
			m.ResetLY()
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
		m.memory.WriteIO(m.memory.AddrIO(address), v)
		m.CheckLYCLY()
		return
	default:
	}

	m.memory.WriteIO(address-addr.MemIOBegin, v)
}

func (m *MMU) IncLY() byte {
	ly := m.memory.ReadIO(m.memory.AddrIO(addr.LY)) + 1
	m.memory.WriteIO(m.memory.AddrIO(addr.LY), ly)

	return ly
}

func (m *MMU) ResetLY() {
	m.memory.WriteIO(addr.LY-addr.MemIOBegin, 0)
}

func (m *MMU) dmaTransfer(v byte) {
	address := uint16(v) << 8

	for i := uint16(0); i < 0xA0; i++ {
		m.Write(0xFE00+i, m.Read(address+i))
	}
}

func IsUnmapped(address uint16) bool {
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

// GetUnusedBits returns a byte with 1's for unused bits
func GetUnusedBits(address uint16) byte {
	switch address {
	case addr.JOYP:
		return 0b11000000
	case addr.SC:
		return 0b01111110
	case addr.TAC:
		return 0b11111000
	case addr.IF:
		return 0b11100000
	case 0xFF10:
		return 0b10000000
	case 0xFF1A:
		return 0b01111111
	case 0xFF1C:
		return 0b10011111
	case 0xFF20:
		return 0b11000000
	case 0xFF23:
		return 0b00111111
	case 0xFF26:
		return 0b01110000
	case addr.STAT:
		return 0b10000000
	default:
		return 0
	}
}
