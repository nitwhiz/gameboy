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

type Data = [Pixels]byte

type Screen struct {
	Background Data
	Buffer     Data
}

func New() *Screen {
	return &Screen{}
}

func (s *Screen) SetBackground(x, y, v byte) {
	s.Background[int(x)+int(y)*Width] = v
}

func (s *Screen) SetPixel(x, y, v byte) {
	s.Buffer[int(x)+int(y)*Width] = v
}

func (s *Screen) GetBackground(x, y byte) byte {
	return s.Background[int(x)+int(y)*Width]
}

func (s *Screen) GetPixel(x, y byte) byte {
	return s.Buffer[int(x)+int(y)*Width]
}

func (s *Screen) Display() Data {
	return s.Buffer
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
