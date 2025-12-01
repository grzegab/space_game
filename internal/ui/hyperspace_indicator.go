package ui

import (
	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type HyperspaceIndicator struct {
	position helpers.Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewHyperspaceIndicator(pos helpers.Vector) *HyperspaceIndicator {
	return &HyperspaceIndicator{
		position: pos,
		rotation: 0,
		sprite:   assets.HyperSpaceIndicator,
	}
}

func (hsi *HyperspaceIndicator) Update() {}

func (hsi *HyperspaceIndicator) Draw(screen *ebiten.Image) {
	bounds := hsi.sprite.Bounds()
	cx := float64(bounds.Dx()) / 2
	cy := float64(bounds.Dy()) / 2

	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Rotate(hsi.rotation)
	op.GeoM.Translate(cx, cy)

	op.GeoM.Translate(hsi.position.X, hsi.position.Y)

	cm := colorm.ColorM{}
	cm.Scale(1., 1., 1., .2)

	colorm.DrawImage(screen, hsi.sprite, cm, op)
}
