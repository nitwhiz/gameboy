package types

type PPUMode byte

type PPU interface {
	Ticks(ticks int)
	Screen() Screen
}
