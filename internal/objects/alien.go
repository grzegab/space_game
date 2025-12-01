package objects

import (
	"math/rand/v2"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Alien struct {
	position     helpers.Vector
	movement     helpers.Vector
	angle        float64
	iq           int
	sprite       *ebiten.Image
	collisionObj *resolv.Circle
}

func NewAlien(state *state.State, p *Player) *Alien {
	iq := rand.IntN(250)
	a := &Alien{}

	switch {
	case iq > 200:
		a.sprite = assets.AlienSprites[4]
		x := float64(rand.IntN(state.ScreenWidth-300)) + 300
		y := float64(rand.IntN(state.ScreenHeight-400)) + 400
		t := p.Position
		p := helpers.Vector{X: x, Y: y}
		v := config.BaseVelocity + rand.Float64()*2
		direction := helpers.Vector{X: t.X - p.X, Y: t.Y - p.Y}
		nd := direction.Normalize()
		a.movement = helpers.Vector{X: nd.X * v, Y: nd.Y * v}
		a.position = p
		a.collisionObj = resolv.NewCircle(p.X, p.Y, float64(a.sprite.Bounds().Dx())/2)
		a.angle = 0
		a.iq = iq
	case iq > 150:
		a.sprite = assets.AlienSprites[3]
		x := float64(rand.IntN(state.ScreenWidth-100)) + 100
		y := float64(rand.IntN(state.ScreenHeight-100)) + 100
		t := helpers.Vector{X: 0, Y: y}
		p := helpers.Vector{X: x, Y: y}
		v := config.BaseVelocity + rand.Float64()*1.5
		direction := helpers.Vector{X: t.X - p.X, Y: t.Y - p.Y}
		nd := direction.Normalize()
		a.movement = helpers.Vector{X: nd.X * v, Y: nd.Y * v}
		a.position = p
		a.collisionObj = resolv.NewCircle(p.X, p.Y, float64(a.sprite.Bounds().Dx())/2)
		a.angle = 0
		a.iq = iq
	case iq > 100:
		a.sprite = assets.AlienSprites[2]
		x := float64(rand.IntN(state.ScreenWidth-200)) + 200
		y := float64(rand.IntN(state.ScreenHeight-200)) + 200
		t := helpers.Vector{X: 0, Y: y}
		p := helpers.Vector{X: x, Y: y}
		v := config.BaseVelocity + rand.Float64()
		direction := helpers.Vector{X: t.X - p.X, Y: t.Y - p.Y}
		nd := direction.Normalize()
		a.movement = helpers.Vector{X: nd.X * v, Y: nd.Y * v}
		a.position = p
		a.collisionObj = resolv.NewCircle(p.X, p.Y, float64(a.sprite.Bounds().Dx())/2)
		a.angle = 0
		a.iq = iq
	case iq > 50:
		a.sprite = assets.AlienSprites[1]
		x := float64(state.ScreenWidth + 100)
		y := float64(rand.IntN(state.ScreenHeight-100)) + 100
		t := helpers.Vector{X: 0, Y: y}
		p := helpers.Vector{X: x, Y: y}
		v := config.BaseVelocity + rand.Float64()*0.5
		m := helpers.Vector{X: t.X - v, Y: 0}
		a.position = p
		a.collisionObj = resolv.NewCircle(p.X, p.Y, float64(a.sprite.Bounds().Dx())/2)
		a.angle = 0
		a.movement = m
		a.iq = iq
	default:
		a.sprite = assets.AlienSprites[0]
		x := float64(state.ScreenWidth + 100)
		y := float64(rand.IntN(state.ScreenHeight-100)) + 100
		t := helpers.Vector{X: 0, Y: y}
		p := helpers.Vector{X: x, Y: y}
		v := config.BaseVelocity + rand.Float64()*0.2
		m := helpers.Vector{X: t.X - v, Y: 0}
		a.position = p
		a.collisionObj = resolv.NewCircle(p.X, p.Y, float64(a.sprite.Bounds().Dx())/2)
		a.angle = 0
		a.movement = m
		a.iq = iq
	}

	a.collisionObj.SetPosition(a.position.X, a.position.Y)
	a.collisionObj.Tags().Set(helpers.TagAlien)

	return a
}

func (a *Alien) Update() {
	a.position.X += a.movement.X
	a.position.Y += a.movement.Y
	a.collisionObj.SetPosition(a.position.X, a.position.Y)
}

func (a *Alien) Draw(screen *ebiten.Image) {
	b := a.sprite.Bounds()
	cx := float64(b.Dx()) / 2
	cy := float64(b.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Rotate(a.angle)
	op.GeoM.Translate(cx, cy)
	op.GeoM.Translate(a.position.X, a.position.Y)
	screen.DrawImage(a.sprite, op)
}
