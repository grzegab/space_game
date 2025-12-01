package ui

import (
	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type ShieldIndicator struct {
	position helpers.Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewShieldIndicator(pos helpers.Vector) *ShieldIndicator {
	return &ShieldIndicator{
		position: pos,
		rotation: 0,
		sprite:   assets.ShieldIndicator,
	}
}

func (si *ShieldIndicator) Update() {}

func (si *ShieldIndicator) Draw(screen *ebiten.Image) {
	bounds := si.sprite.Bounds()
	cx := float64(bounds.Dx()) / 2
	cy := float64(bounds.Dy()) / 2

	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Rotate(si.rotation)
	op.GeoM.Translate(cx, cy)

	cm := colorm.ColorM{}
	cm.Scale(1., 1., 1., .2)

	op.GeoM.Translate(si.position.X, si.position.Y)

	colorm.DrawImage(screen, si.sprite, cm, op)
}
