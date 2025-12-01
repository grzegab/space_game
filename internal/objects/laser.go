package objects

import (
	"math"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Laser struct {
	position     helpers.Vector
	rotation     float64
	sprite       *ebiten.Image
	CollisionObj *resolv.ConvexPolygon
}

func NewLaser(pos helpers.Vector, rotation float64, index int) *Laser {
	sprite := assets.LaserSprite

	b := sprite.Bounds()
	cx := float64(b.Dx()) / 2
	cy := float64(b.Dy()) / 2

	pos.X -= cx
	pos.Y -= cy

	l := &Laser{
		position:     pos,
		rotation:     rotation,
		sprite:       sprite,
		CollisionObj: resolv.NewRectangle(pos.X, pos.Y, float64(sprite.Bounds().Dx()), float64(sprite.Bounds().Dy())),
	}

	l.CollisionObj.SetPosition(pos.X, pos.Y)
	l.CollisionObj.Tags().Set(helpers.TagLaser)
	l.CollisionObj.SetData(&helpers.ObjectData{Index: index})

	return l

}

func (l *Laser) Update() {
	speed := config.LaserSpeed / float64(ebiten.TPS())
	dx := math.Sin(l.rotation) * speed
	dy := math.Cos(l.rotation) * -speed

	l.position.X += dx
	l.position.Y += dy

	l.CollisionObj.SetPosition(l.position.X, l.position.Y)
}

func (l *Laser) Draw(screen *ebiten.Image) {
	bounds := l.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(l.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(l.position.X, l.position.Y)

	screen.DrawImage(l.sprite, op)
}
