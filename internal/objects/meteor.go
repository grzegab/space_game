package objects

import (
	"math"
	"math/rand/v2"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Meteor struct {
	size          config.MeteorSize
	Position      helpers.Vector
	Movement      helpers.Vector
	rotation      float64
	angle         float64
	rotationSpeed float64
	sprite        *ebiten.Image
	CollisionObj  *resolv.Circle
}

func NewMeteor(size config.MeteorSize, v float64, i int, w int, h int) *Meteor {
	angle := rand.Float64() * 2 * math.Pi
	target := helpers.Vector{
		X: float64(w) / 2,
		Y: float64(h) / 2,
	}
	r := float64(w/2.0 + 500)
	pos := helpers.Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}
	velocity := v + rand.Float64()*1.5
	d := helpers.Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}
	normalizeDirection := d.Normalize()

	movement := helpers.Vector{
		X: normalizeDirection.X * velocity,
		Y: normalizeDirection.Y * velocity,
	}

	rs := rand.Float64()*(config.RotationSpeedMax-config.RotationSpeedMin) + config.RotationSpeedMin

	var sprite *ebiten.Image
	switch size {
	case config.MeteorSmall:
		sprite = assets.SmallMeteorSprite[rand.IntN(len(assets.SmallMeteorSprite))]
	case config.MeteorLarge:
		sprite = assets.BigMeteorSprites[rand.IntN(len(assets.BigMeteorSprites))]
	}
	collisionObj := resolv.NewCircle(pos.X, pos.Y, float64(sprite.Bounds().Dx())/2)

	m := &Meteor{
		size:          size,
		Position:      pos,
		Movement:      movement,
		rotation:      r,
		angle:         angle,
		rotationSpeed: rs,
		sprite:        sprite,
		CollisionObj:  collisionObj,
	}

	m.CollisionObj.SetPosition(pos.X, pos.Y)
	m.CollisionObj.Tags().Set(helpers.TagMeteor | helpers.TagSmallMeteor)
	m.CollisionObj.SetData(&helpers.ObjectData{Index: i})

	return m
}

func (m *Meteor) Update(state *state.State) {
	m.keepOnScreen(state.ScreenWidth, state.ScreenHeight)

	m.Position.X += m.Movement.X
	m.Position.Y += m.Movement.Y
	m.rotation += m.rotationSpeed
	m.CollisionObj.SetPosition(m.Position.X, m.Position.Y)
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	cx := float64(bounds.Dx()) / 2
	cy := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(cx, cy)

	op.GeoM.Translate(m.Position.X, m.Position.Y)

	screen.DrawImage(m.sprite, op)
}

func (m *Meteor) keepOnScreen(w, h int) {
	if m.Position.X >= float64(w) {
		m.Position.X = 0
		m.CollisionObj.SetPosition(0, m.Position.Y)
	}
	if m.Position.X < 0 {
		m.Position.X = float64(w)
		m.CollisionObj.SetPosition(float64(w), m.Position.Y)
	}

	if m.Position.Y >= float64(h) {
		m.Position.Y = 0
		m.CollisionObj.SetPosition(m.Position.X, 0)
	}
	if m.Position.Y < 0 {
		m.Position.Y = float64(h)
		m.CollisionObj.SetPosition(m.Position.X, float64(h))
	}
}
