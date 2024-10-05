package types

type Flag byte

const (
	// FlagZ - Zero Flag
	FlagZ = Flag(7)
	// FlagN - Subtraction Flag
	FlagN = Flag(6)
	// FlagH - Half Carry Flag
	FlagH = Flag(5)
	// FlagC - Carry Flag
	FlagC = Flag(4)
)

type CPU interface {
	AF() Register
	BC() Register
	DE() Register
	HL() Register
	SP() Register
	PC() Register
	IME() bool
	SetIME(ime bool)
	Halt() bool
	SetHalt(halt bool)
	Flag(f Flag) bool
	SetFlag(f Flag, v bool)
	IncPC() uint16
}

type Register interface {
	Set(v uint16)
	Val() uint16
	SetLo(v byte)
	SetHi(v byte)
	Lo() byte
	Hi() byte
}
