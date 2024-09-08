package mmu

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/bits"
	"github.com/nitwhiz/gameboy/pkg/cartridge"
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/memory"
)

type MMU struct {
	Cartridge *cartridge.Cartridge
	Memory    *memory.Memory
	Input     *input.State

	TimerCounter uint16
	TimerLock    bool

	SerialReceiver func(byte)
}

func New(in *input.State) *MMU {
	return &MMU{
		Cartridge:      nil,
		Memory:         memory.New().Init(),
		Input:          in,
		TimerCounter:   0xAC00,
		TimerLock:      false,
		SerialReceiver: nil,
	}
}

func inRange(a, l, u uint16) bool {
	return a >= l && a <= u
}

func (m *MMU) Read(address uint16) byte {
	switch {
	case address == addr.IE:
		return m.Memory.IE
	case address == addr.IF:
		return m.Memory.IF
	case address == addr.DIV:
		return byte((m.TimerCounter & 0xFF00) >> 8)
	case address == addr.JOYP:
		v := m.Memory.IO[address-addr.MemIOBegin] & 0xF0

		if bits.IsJOYPSelectButtons(v) {
			return v | m.Input.Get(input.SelectButtons)
		} else if bits.IsJOYPSelectDPad(v) {
			return v | m.Input.Get(input.SelectDPad)
		}

		return v | 0x0F | memory.GetUnusedBits(addr.JOYP)
	case inRange(address, addr.MemROMBegin, addr.MemROMEnd):
		if m.Cartridge == nil {
			return 0xFF
		}

		return m.Cartridge.Read(address)
	case inRange(address, addr.MemVRAMBegin, addr.MemVRAMEnd):
		return m.Memory.VRAM[address-addr.MemVRAMBegin]
	case inRange(address, addr.MemCartridgeRAMBegin, addr.MemCartridgeRAMEnd):
		if m.Cartridge == nil {
			return 0xFF
		}

		return m.Cartridge.Read(address)
	case inRange(address, addr.MemWRAMBegin, addr.MemWRAMEnd):
		return m.Memory.WRAM[address-addr.MemWRAMBegin]
	case inRange(address, addr.MemOAMBegin, addr.MemOAMEnd):
		return m.Memory.OAM[address-addr.MemOAMBegin]
	case inRange(address, addr.MemIOBegin, addr.MemIOEnd):
		return m.Memory.IO[address-addr.MemIOBegin]
	case inRange(address, addr.MemHRAMBegin, addr.MemHRAMEnd):
		return m.Memory.HRAM[address-addr.MemHRAMBegin]
	default:
		return 0xFF
	}
}

func (m *MMU) Write(address uint16, v byte) {
	switch {
	case address == addr.IE:
		m.Memory.IE = v
	case address == addr.IF:
		m.Memory.IF = v | memory.GetUnusedBits(addr.IF)
	case address == addr.DIV:
		m.ResetTimer()
	case inRange(address, addr.MemAudioBegin, addr.MemAudioEnd):
		// not implemented
		return
	case inRange(address, addr.MemWaveBegin, addr.MemWaveEnd):
		// not implemented
		return
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
	}
}

func (m *MMU) RequestInterrupt(typ addr.InterruptType) {
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
	if memory.IsUnmapped(address) {
		return
	}

	v |= memory.GetUnusedBits(address)

	switch {
	case address == addr.JOYP:
		m.Memory.IO[address-addr.MemIOBegin] = v & 0b00110000
		return
	case address == addr.SC:
		if bits.Test(v, 7) && bits.Test(v, 0) {
			if m.SerialReceiver != nil {
				m.SerialReceiver(m.Read(addr.SB))
			}
		}
	case address == addr.LCDC:
		if !bits.Test(v, addr.LCDC_ENABLE) {
			m.SetLY(0)
			return
		}
	case address == addr.DMA:
		m.dmaTransfer(v)
		return
	case address == addr.TIMA:
		if m.TimerLock {
			return
		}
	case address == addr.LYC:
		m.Memory.IO[address-addr.MemIOBegin] = v
		m.CheckLYCLY()
		return
	default:
	}

	m.Memory.IO[address-addr.MemIOBegin] = v
}

// SetTIMA does not respect the timer lock
func (m *MMU) SetTIMA(tima byte) {
	m.Memory.IO[addr.TIMA-addr.MemIOBegin] = tima
}

func (m *MMU) IncLY() byte {
	v := m.Memory.IO[addr.LY-addr.MemIOBegin] + 1

	m.SetLY(v)

	return v
}

func (m *MMU) SetLY(v byte) {
	m.Memory.IO[addr.LY-addr.MemIOBegin] = v
}

func (m *MMU) dmaTransfer(v byte) {
	address := uint16(v) << 8

	for i := uint16(0); i < 0xA0; i++ {
		m.Write(0xFE00+i, m.Read(address+i))
	}
}

func (m *MMU) ResetTimer() {
	m.TimerCounter = 0
}
