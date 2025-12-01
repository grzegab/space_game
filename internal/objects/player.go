package objects

import (
	"math"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"
)

const (
	forwardKey = ebiten.KeyW
	leftKey    = ebiten.KeyA
	rightKey   = ebiten.KeyD
)

var (
	shotsFired   int
	acceleration float64
	dyingCounter int
)

type PlayerTimers struct {
	shootTimeDown  *helpers.Timer
	burstTimeDown  *helpers.Timer
	dyingTimeDown  *helpers.Timer
	shieldTimeDown *helpers.Timer
	drift          *helpers.Timer
}

type PlayerSounds struct {
	exhaustSound  *audio.Player
	shieldSound   *audio.Player
	laserOneSound *audio.Player
	laserTwoSound *audio.Player
}

type PlayerShield struct {
	sprite *ebiten.Image
	Count  int
	Active bool
}

type PlayerLasers struct {
	Lasers map[int]*Laser
	count  int
}

type Player struct {
	isDying      bool
	timers       PlayerTimers
	Position     helpers.Vector
	Rotation     float64
	velocity     float64
	driftAngle   float64
	sprite       *ebiten.Image
	explosion    []*ebiten.Image
	exhaust      *Exhaust
	playerAudio  PlayerSounds
	Shield       PlayerShield
	PlayerLasers PlayerLasers
	CollisionObj *resolv.Circle
}

func NewPlayer(x, y int, a *audio.Context) *Player {
	s := assets.PlayerImage
	b := s.Bounds()
	cx := float64(b.Dx()) / 2
	cy := float64(b.Dy()) / 2

	startingPos := helpers.Vector{
		X: float64(x)/2 - cx,
		Y: float64(y)/2 - cy,
	}

	obj := resolv.NewCircle(startingPos.X, startingPos.Y, cx/2)

	t := PlayerTimers{
		burstTimeDown: helpers.NewTimer(config.BurstCoolDown),
		shootTimeDown: helpers.NewTimer(config.ShootCoolDown),
	}

	thrustSound, _ := a.NewPlayer(assets.ThrustSound)
	shieldSound, _ := a.NewPlayer(assets.ShieldSound)
	laserOneSound, _ := a.NewPlayer(assets.LaserOneSound)
	laserTwoSound, _ := a.NewPlayer(assets.LaserTwoSound)

	pa := PlayerSounds{
		exhaustSound:  thrustSound,
		shieldSound:   shieldSound,
		laserOneSound: laserOneSound,
		laserTwoSound: laserTwoSound,
	}

	pl := PlayerLasers{
		Lasers: make(map[int]*Laser),
		count:  0,
	}

	ps := PlayerShield{
		sprite: assets.ShieldSprite,
		Count:  config.NumberOfShields,
		Active: false,
	}

	return &Player{
		sprite:       s,
		timers:       t,
		playerAudio:  pa,
		PlayerLasers: pl,
		Shield:       ps,
		Position:     startingPos,
		explosion:    assets.Explosion,
		CollisionObj: obj,
	}
}

func (p *Player) Update(state *state.State) {
	p.updateTimers()
	p.rotatePlayer(state)
	p.accelerate(state)
	p.useShield()

	p.isPlayerDrifting(state)
	p.checkForDying()
	p.fireLasers()

	p.CollisionObj.SetPosition(p.Position.X, p.Position.Y)

	if p.PlayerLasers.Lasers != nil {
		for _, laser := range p.PlayerLasers.Lasers {
			laser.Update()
		}
	}

	if p.exhaust != nil {
		p.exhaust.Update()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	if p.exhaust != nil && ebiten.IsKeyPressed(forwardKey) {
		p.exhaust.Draw(screen)
	}

	if len(p.PlayerLasers.Lasers) > 0 {
		for _, laser := range p.PlayerLasers.Lasers {
			laser.Draw(screen)
		}
	}

	bounds := p.sprite.Bounds()
	cx := float64(bounds.Dx()) / 2
	cy := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Rotate(p.Rotation)
	op.GeoM.Translate(cx, cy)

	op.GeoM.Translate(p.Position.X, p.Position.Y)

	screen.DrawImage(p.sprite, op)

	// shield
	if p.Shield.Active {
		sb := p.Shield.sprite.Bounds()
		sCx := float64(sb.Dx()) / 2
		sCy := float64(sb.Dy()) / 2

		pos := helpers.Vector{
			X: p.Position.X - sCx/2 + config.ShieldOffset,
			Y: p.Position.Y - sCy/2 + config.ShieldOffset,
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-sCx, -sCy)
		op.GeoM.Rotate(p.Rotation)
		op.GeoM.Translate(sCx, sCy)
		op.GeoM.Translate(pos.X, pos.Y)

		screen.DrawImage(p.Shield.sprite, op)
	}

}

func (p *Player) updateTimers() {
	p.timers.burstTimeDown.Update()
	p.timers.shootTimeDown.Update()
}

func (p *Player) rotatePlayer(state *state.State) {
	s := config.RotationSpeed / float64(ebiten.TPS())
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Rotation -= s
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Rotation += s
	}
}

func (p *Player) accelerate(state *state.State) {
	if ebiten.IsKeyPressed(forwardKey) {
		p.timers.drift = nil
		p.keepOnScreen(state)

		if p.velocity < config.MaxAcceleration {
			if p.velocity+config.AccelerationStep >= config.MaxAcceleration {
				p.velocity = config.MaxAcceleration
			}

			p.velocity += config.AccelerationStep
		}

		dx := math.Sin(p.Rotation) * p.velocity
		dy := math.Cos(p.Rotation) * -p.velocity

		b := p.sprite.Bounds()
		cx := float64(b.Dx()) / 2
		cy := float64(b.Dy()) / 2

		exhaustPos := helpers.Vector{
			X: p.Position.X + cx + math.Sin(p.Rotation)*config.ExhaustSpawnOffset,
			Y: p.Position.Y + cy + math.Cos(p.Rotation)*-config.ExhaustSpawnOffset,
		}

		p.exhaust = NewExhaust(exhaustPos, p.Rotation+180*math.Pi/180)

		p.Position.X += dx
		p.Position.Y += dy

		if !p.playerAudio.exhaustSound.IsPlaying() {
			_ = p.playerAudio.exhaustSound.Rewind()
			p.playerAudio.exhaustSound.Play()
		}
	}

	if inpututil.IsKeyJustReleased(forwardKey) {
		if p.playerAudio.exhaustSound.IsPlaying() {
			p.playerAudio.exhaustSound.Pause()
		}

		if p.velocity < acceleration*10 {
			p.velocity = acceleration*10 - 5.0
		}

		if p.velocity < 0 {
			p.velocity = 0
		}

		acceleration = 0.0
		p.timers.drift = helpers.NewTimer(config.DriftTime)
		p.driftAngle = p.Rotation
	}

	if !ebiten.IsKeyPressed(forwardKey) && p.exhaust != nil {
		p.exhaust = nil
	}
}

func (p *Player) checkForDying() {
	if p.isDying {
		if p.playerAudio.exhaustSound.IsPlaying() {
			p.playerAudio.exhaustSound.Pause()
		}

		p.exhaust = nil
		p.timers.dyingTimeDown.Update()
		if p.timers.dyingTimeDown.IsReady() {
			p.timers.dyingTimeDown.Reset()
			dyingCounter++

			if dyingCounter < config.DyingAnimationFrames {
				p.sprite = p.explosion[dyingCounter]
			}
		}
	}
}

func (p *Player) useShield() {
	if ebiten.IsKeyPressed(ebiten.KeyQ) && !p.Shield.Active && p.Shield.Count > 0 {
		if !p.playerAudio.shieldSound.IsPlaying() {
			_ = p.playerAudio.shieldSound.Rewind()
			p.playerAudio.shieldSound.Play()
		}

		p.Shield.Active = true
		p.Shield.Count--
		p.timers.shieldTimeDown = helpers.NewTimer(config.ShieldDuration)

		//p.game.shield = NewShield(p.Position, p.Rotation, p.game)
		//p.ShieldIndicators = p.ShieldIndicators[:len(p.ShieldIndicators)-1]
	}

	if p.timers.shieldTimeDown != nil && p.Shield.Active {
		p.timers.shieldTimeDown.Update()
	}

	if p.timers.shieldTimeDown != nil && p.timers.shieldTimeDown.IsReady() {
		p.timers.shieldTimeDown = nil
		p.Shield.Active = false
		//p.game.space.Remove(p.game.shield.obj)
		//p.game.shield = nil
	}
}

func (p *Player) isPlayerDrifting(state *state.State) {
	if p.timers.drift != nil {
		if p.timers.drift.IsReady() {
			p.timers.drift = nil
			p.velocity = 0

			return
		}

		p.keepOnScreen(state)
		p.timers.drift.Update()
		decelerationSpeed := p.velocity / float64(ebiten.TPS()) * 4
		p.Position.X += math.Sin(p.driftAngle) * decelerationSpeed
		p.Position.Y += math.Cos(p.driftAngle) * -decelerationSpeed
		p.CollisionObj.SetPosition(p.Position.X, p.Position.Y)
	}
}

func (p *Player) fireLasers() {
	if p.timers.burstTimeDown.IsReady() {
		if p.timers.shootTimeDown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
			p.timers.shootTimeDown.Reset()
			shotsFired++

			if shotsFired <= config.MaxShotsPerBurst {
				bounds := p.sprite.Bounds()
				halfW := float64(bounds.Dx()) / 2
				halfH := float64(bounds.Dy()) / 2

				spawnPos := helpers.Vector{
					X: p.Position.X + halfW + math.Sin(p.Rotation)*config.LaserSpawnOffset,
					Y: p.Position.Y + halfH + math.Cos(p.Rotation)*-config.LaserSpawnOffset,
				}

				p.PlayerLasers.count++
				laser := NewLaser(spawnPos, p.Rotation, p.PlayerLasers.count)
				p.PlayerLasers.Lasers[p.PlayerLasers.count] = laser

				switch shotsFired {
				case 1:
					if !p.playerAudio.laserOneSound.IsPlaying() {
						_ = p.playerAudio.laserOneSound.Rewind()
						p.playerAudio.laserOneSound.Play()
					}
				case 2:
					if !p.playerAudio.laserTwoSound.IsPlaying() {
						_ = p.playerAudio.laserTwoSound.Rewind()
						p.playerAudio.laserTwoSound.Play()
					}
				}
			} else {
				p.timers.burstTimeDown.Reset()
				shotsFired = 0
			}
		}
	}
}

func (p *Player) keepOnScreen(state *state.State) {
	if p.Position.X >= float64(state.ScreenWidth) {
		p.Position.X = 0
		p.CollisionObj.SetPosition(0, p.Position.Y)
	}

	if p.Position.X < 0 {
		p.Position.X = float64(state.ScreenWidth)
		p.CollisionObj.SetPosition(float64(state.ScreenWidth), p.Position.Y)
	}

	if p.Position.Y >= float64(state.ScreenHeight) {
		p.Position.Y = 0
		p.CollisionObj.SetPosition(p.Position.X, 0)
	}

	if p.Position.Y < 0 {
		p.Position.Y = float64(state.ScreenHeight)
		p.CollisionObj.SetPosition(p.Position.X, float64(state.ScreenHeight))
	}
}
