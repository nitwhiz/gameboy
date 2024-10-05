package types

// todo: address translation should not happen here, but in mmu

type Memory interface {
	AddrVRAM(address uint16) uint16
	ReadVRAM(address uint16) byte
	WriteVRAM(address uint16, v byte)

	AddrWRAM(address uint16) uint16
	ReadWRAM(address uint16) byte
	WriteWRAM(address uint16, v byte)

	AddrOAM(address uint16) uint16
	ReadOAM(address uint16) byte
	WriteOAM(address uint16, v byte)

	AddrHRAM(address uint16) uint16
	ReadHRAM(address uint16) byte
	WriteHRAM(address uint16, v byte)

	AddrIO(address uint16) uint16
	ReadIO(address uint16) byte
	WriteIO(address uint16, v byte)

	Div() byte
	TimerCounter() uint16
	IncTimerCounter()
	ResetTimerCounter()
}
