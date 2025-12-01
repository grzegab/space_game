package objects

import (
	"math"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

type Exhaust struct {
	position helpers.Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewExhaust(pos helpers.Vector, rotation float64) *Exhaust {
	s := assets.ExhaustSprite
	b := s.Bounds()
	hW := float64(b.Dx()) / 2
	hF := float64(b.Dy()) / 2

	pos.X -= hW
	pos.Y -= hF

	return &Exhaust{
		position: pos,
		rotation: rotation,
		sprite:   s,
	}
}

func (e *Exhaust) Draw(screen *ebiten.Image) {
	b := e.sprite.Bounds()
	hW := float64(b.Dx()) / 2
	hF := float64(b.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-hW, -hF)
	op.GeoM.Rotate(e.rotation)
	op.GeoM.Translate(hW, hF)
	op.GeoM.Translate(e.position.X, e.position.Y)

	screen.DrawImage(e.sprite, op)
}

func (e *Exhaust) Update() {
	speed := config.MaxAcceleration / float64(ebiten.TPS())
	e.position.X += math.Sin(e.rotation) * speed
	e.position.Y += math.Cos(e.rotation) * -speed
}
