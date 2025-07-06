package types

type GameBoy interface {
	Lock()
	Unlock()

	CPU() CPU
	PPU() PPU
	MMU() MMU
	InterruptController() InterruptController
	Input() InputState
	Timer() Timer

	PushStack(addr uint16)
	PopStack() uint16

	Start()
	Stop()

	Halted() bool
	SetHalted(h bool)
	HaltBug() bool
	SetHaltBug(hb bool)
	JustHalted() bool
	SetJustHalted(jh bool)
	Stopped() bool
	SetStopped(s bool)

	PressButton(button ButtonType)
	ReleaseButton(button ButtonType)

	Fetch8() byte
	Fetch16() uint16

	Read8(addr uint16) byte
	Read16(addr uint16) uint16
	Write(addr uint16, v byte)
	Cycle()

	AddPendingTicks(n int)
	SetPendingTicks(n int)
	PendingTicks() int
	AdvanceTicks(n int)
	AdvancePendingTicks()
	FlushPendingTicks()
}
