package types

type PPUMode byte

type PPU interface {
	Update(ticks int)
	Screen() Screen
}
