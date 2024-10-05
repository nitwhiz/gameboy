package types

type ButtonType byte

type SelectType byte

const (
	InputSelectButtons = SelectType(0)
	InputSelectDPad    = SelectType(1)
)

type InputState interface {
	Press(button ButtonType)
	Release(button ButtonType)
	Value(sel SelectType) byte
}
