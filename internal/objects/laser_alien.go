package objects

import (
	"math"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type LaserAlien struct {
	position     helpers.Vector
	rotation     float64
	sprite       *ebiten.Image
	collisionObj *resolv.ConvexPolygon
}

func NewAlienLaser(pos helpers.Vector, rotation float64) *LaserAlien {
	s := assets.AlienLaserSprite
	b := s.Bounds()
	cx := float64(b.Dx()) / 2
	cy := float64(b.Dy()) / 2

	pos.X -= cx
	pos.Y -= cy

	al := &LaserAlien{
		position:     pos,
		rotation:     rotation,
		sprite:       s,
		collisionObj: resolv.NewRectangle(pos.X, pos.Y, float64(s.Bounds().Dx()), float64(s.Bounds().Dy())),
	}

	al.collisionObj.SetPosition(pos.X, pos.Y)
	al.collisionObj.Tags().Set(helpers.TagLaser)

	return al
}

func (l *LaserAlien) Update() {
	speed := config.AlienLaserSpeed / float64(ebiten.TPS())
	l.position.X += math.Sin(l.rotation) * speed
	l.position.Y += math.Cos(l.rotation) * -speed

	l.collisionObj.SetPosition(l.position.X, l.position.Y)
}

func (l *LaserAlien) Draw(screen *ebiten.Image) {
	b := l.sprite.Bounds()
	cx := float64(b.Dx()) / 2
	cy := float64(b.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Rotate(l.rotation)
	op.GeoM.Translate(cx, cy)

	op.GeoM.Translate(l.position.X, l.position.Y)
	screen.DrawImage(l.sprite, op)
}
