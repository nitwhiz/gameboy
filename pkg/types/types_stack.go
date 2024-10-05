package types

type Stack interface {
	Push(v uint16)
	Pop() uint16
}
