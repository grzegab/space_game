package ui

import (
	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type LifeIndicator struct {
	position helpers.Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewLifeIndicator(pos helpers.Vector) *LifeIndicator {
	return &LifeIndicator{
		position: pos,
		rotation: 0,
		sprite:   assets.LifeIndicator,
	}
}

func (li *LifeIndicator) Update() {}

func (li *LifeIndicator) Draw(screen *ebiten.Image) {
	bounds := li.sprite.Bounds()
	cx := float64(bounds.Dx()) / 2
	cy := float64(bounds.Dy()) / 2

	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(cx, cy)
	op.GeoM.Translate(li.position.X, li.position.Y)
	cm := colorm.ColorM{}
	cm.Scale(1., 1., 1., .2)
	colorm.DrawImage(screen, li.sprite, cm, op)

}
