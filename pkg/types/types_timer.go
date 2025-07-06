package types

const (
	TimaRunning = iota
	TimaReloading
	TimaReloaded
)

type Timer interface {
	GetDiv() byte
	SetDivState(v int)
	DivCycles() int
	SetDivCycles(v int)
	SetDivCounter(v uint16)
	TIMAState() int
	Tick(ticks int)

	TIMA() byte
	SetTIMA(v byte)
	TAC() byte
	SetTAC(v byte)
	TMA() byte
	SetTMA(v byte)
}
