package types

type GameBoy interface {
	Lock()
	Unlock()

	CPU() CPU
	MMU() MMU
	Stack() Stack
	PPU() PPU

	Start()
	Stop()

	PressButton(button ButtonType)
	ReleaseButton(button ButtonType)

	ServiceInterrupts() (ticks int)
}
