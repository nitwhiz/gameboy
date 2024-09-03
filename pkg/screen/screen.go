package screen

import (
	"image"
	"image/color"
)

const (
	Width  = 160
	Height = 144
	Pixels = Width * Height
)

type PixelData = [Pixels]byte
type BackgroundData = [Pixels]uint16
type SpriteData = [Pixels]uint16

type Screen struct {
	Background BackgroundData
	Sprite     SpriteData

	Buffer PixelData
}

func New() *Screen {
	return &Screen{}
}

func (s *Screen) SetBackground(x, y, colNum, color byte) {
	s.Background[int(x)+int(y)*Width] = uint16(color) | ((uint16(colNum) | (1 << 2)) << 8)
}

func (s *Screen) SetSprite(x, y, priority, colNum, color byte) {
	sprite := s.Sprite[int(x)+int(y)*Width]
	spriteInfo := byte((sprite >> 8) & 0xFF)

	if colNum == 0 || spriteInfo&0b1000 == 0 {
		s.Sprite[int(x)+int(y)*Width] = uint16(color) |
			((uint16(colNum) | (uint16(priority) << 2) | (1 << 3)) << 8)
	}
}

func (s *Screen) ClearScanline(y byte) {
	for x := range Width {
		s.Sprite[x+int(y)*Width] = 0
		s.Background[x+int(y)*Width] = 0
	}
}

func (s *Screen) GetBackground(x, y byte) uint16 {
	return s.Background[int(x)+int(y)*Width]
}

func (s *Screen) GetSprite(x, y byte) uint16 {
	return s.Sprite[int(x)+int(y)*Width]
}

func (s *Screen) Display() PixelData {
	return s.Buffer
}

func (s *Screen) BlitScanline(y byte) {
	for x := range Width {
		background := s.Background[x+int(y)*Width]
		backgroundColor := byte(background & 0xFF)
		backgroundInfo := byte((background >> 8) & 0xFF)

		if backgroundInfo&0b100 == 0 {
			backgroundColor = 0xFF
		}

		sprite := s.Sprite[x+int(y)*Width]
		spriteColor := byte(sprite & 0xFF)
		spriteInfo := byte((sprite >> 8) & 0xFF)

		pixelColor := spriteColor

		if spriteInfo&0b11 == 0 {
			// sprite pixel color num is 0
			pixelColor = backgroundColor
		} else if spriteInfo&0b100 != 0 && backgroundInfo&0b11 != 0 {
			// priority is 1 and background is not 0
			pixelColor = backgroundColor
		}

		s.Buffer[x+int(y)*Width] = pixelColor
	}

	s.ClearScanline(y)
}

func (s *Screen) ColorModel() color.Model {
	return color.GrayModel
}

func (s *Screen) Bounds() image.Rectangle {
	return image.Rect(0, 0, Width, Height)
}

func (s *Screen) At(x, y int) color.Color {
	return color.Gray{
		Y: s.Buffer[x+y*Width],
	}
}
