package addr

// overall memory layout
const (
	MemROMBegin = uint16(0x0000)
	MemROMEnd   = uint16(0x7FFF)

	MemVRAMBegin = uint16(0x8000)
	MemVRAMEnd   = uint16(0x9FFF)

	MemCartridgeRAMBegin = uint16(0xA000)
	MemCartridgeRAMEnd   = uint16(0xBFFF)

	MemWRAMBegin = uint16(0xC000)
	MemWRAMEnd   = uint16(0xDFFF)

	MemERAMBegin = uint16(0xE000)
	MemERAMEnd   = uint16(0xFDFF)

	MemOAMBegin = uint16(0xFE00)
	MemOAMEnd   = uint16(0xFE9F)

	MemIOBegin = uint16(0xFF00)
	MemIOEnd   = uint16(0xFF7F)

	MemHRAMBegin = uint16(0xFF80)
	MemHRAMEnd   = uint16(0xFFFE)

	MemAudioBegin = uint16(0xFF10)
	MemAudioEnd   = uint16(0xFF26)

	MemWaveBegin = uint16(0xFF30)
	MemWaveEnd   = uint16(0xFF3F)
)

// specific registers
const (
	// JOYP - Joypad
	JOYP = uint16(0xFF00)

	// SB - Serial transfer data
	SB = uint16(0xFF01)
	// SC - Serial transfer control
	SC = uint16(0xFF02)

	// DIV - Divider
	DIV = uint16(0xFF04)
	// TIMA - Timer counter
	TIMA = uint16(0xFF05)
	// TMA - Timer modula
	TMA = uint16(0xFF06)
	// TAC - Timer control
	TAC = uint16(0xFF07)

	// note: audio registers are left out
	// note: wave pattern registers are left out

	// IF - IM flag
	IF = uint16(0xFF0F)

	// LCDC - LCD control
	LCDC = uint16(0xFF40)
	// STAT - LCD status
	STAT = uint16(0xFF41)
	// SCY - Background viewport Y position
	SCY = uint16(0xFF42)
	// SCX - Background viewport X position
	SCX = uint16(0xFF43)
	// LY - LCD Y coordinate [read-only]
	LY = uint16(0xFF44)
	// LYC - LY compare
	LYC = uint16(0xFF45)
	// DMA - OAM DMA source address & start
	DMA = uint16(0xFF46)
	// BGP - BG palette data
	BGP = uint16(0xFF47)
	// OBP0 - OBJ palette 0 data
	OBP0 = uint16(0xFF48)
	// OBP1 - OBJ palette 1 data
	OBP1 = uint16(0xFF49)
	// WY - Window Y position
	WY = uint16(0xFF4A)
	// WX - Window X position
	WX = uint16(0xFF4B)

	DisableBootRom = uint16(0xFF50)

	// IE - IM Enable
	IE = uint16(0xFFFF)
)
