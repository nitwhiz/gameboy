package types

import "image"

type Screen interface {
	image.Image
	SetBackground(x, y, colNum, color byte)
	SetSprite(x, y byte, priority bool, colNum, color byte)
	ClearScanline(y byte)
	GetBackground(x, y byte) uint16
	GetSprite(x, y byte) uint16
	BlitScanline(y byte)
}
