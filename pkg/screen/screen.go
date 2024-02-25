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
	Hot  Data
	Cold Data
}

func New() *Screen {
	s := Screen{}

	s.Clear()

	return &s
}

func (s *Screen) SetPixel(x, y, v byte) {
	s.Hot[int(x)+int(y)*Width] = v
}

func (s *Screen) GetPixel(x, y byte) byte {
	return s.Hot[int(x)+int(y)*Width]
}

func (s *Screen) Blit() {
	copy(s.Cold[:], s.Hot[:])
	s.Clear()
}

func (s *Screen) Clear() {
	for i := range Pixels {
		s.Hot[i] = 0
	}
}

func (s *Screen) Display() Data {
	return s.Cold
}

func (s *Screen) ColorModel() color.Model {
	return color.RGBAModel
}

func (s *Screen) Bounds() image.Rectangle {
	return image.Rect(0, 0, Width, Height)
}

func (s *Screen) At(x, y int) color.Color {
	c := s.Cold[x+y*Width]

	return color.RGBA{
		R: c,
		G: c,
		B: c,
		A: 0xFF,
	}
}
