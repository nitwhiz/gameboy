package types

type InterruptType byte

type InterruptController interface {
	Request(i InterruptType)
	IsRequested(i InterruptType) bool
	Flush(i InterruptType)
	IF() byte
	IE() byte
	SetIF(flag byte)
	SetIE(enable byte)
}
