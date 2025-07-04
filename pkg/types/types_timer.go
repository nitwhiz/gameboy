package types

type Timer interface {
	GetValue() uint16
	SetValue(v uint16)
	Inc()
	Tick(ticks int)
}
