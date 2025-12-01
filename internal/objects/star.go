package objects

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	white = 0xff
	red   = 0xbb
	green = 0xdd
)

type Star struct {
	x          float32
	y          float32
	radius     float32
	brightness float32
}

func GenerateStars(count int, x, y int) []*Star {
	stars := make([]*Star, count)

	for i := 0; i < count; i++ {
		stars[i] = NewStar(x, y)
	}

	return stars
}

func NewStar(x, y int) *Star {
	return &Star{
		x:          rand.Float32() * float32(x),
		y:          rand.Float32() * float32(y),
		radius:     rand.Float32()*(3-1) + 1,
		brightness: rand.Float32() * white,
	}
}

func (s *Star) Update() {}

func (s *Star) Draw(screen *ebiten.Image) {
	rgba := color.RGBA{
		R: uint8(red * s.brightness / white),
		G: uint8(green * s.brightness / white),
		B: uint8(white * s.brightness / white),
		A: white,
	}

	vector.FillCircle(screen, s.x, s.y, s.radius, rgba, true)
}
