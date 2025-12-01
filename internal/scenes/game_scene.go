package scenes

import (
	"math/rand"

	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/grzegab/sample_game/internal/objects"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/solarlune/resolv"
)

type GameTimers struct {
	hyperspaceCooldown *helpers.Timer
	shieldWorkTime     *helpers.Timer
	meteorSpawn        *helpers.Timer
	alienSpawn         *helpers.Timer
	derbiesCleanup     *helpers.Timer
	velocitySpeed      *helpers.Timer
	beat               *helpers.Timer
}

type GameMeteors struct {
	velocity     float64
	levelCount   int
	currentCount int
	list         map[int]*objects.Meteor
}

type Lasers struct {
	playerLasers     map[int]*objects.Laser
	playerLaserCount int
	alienLasers      map[int]*objects.Laser
}

type Audio struct {
	hyperspaceSound *audio.Player
}

type GameScene struct {
	lives   int
	player  *objects.Player
	space   *resolv.Space
	stars   []*objects.Star
	meteors GameMeteors
	level   int
	timers  GameTimers
	lasers  map[int]*objects.Laser
	audio   Audio
}

func NewGameScene(w, h int, a *audio.Context) *GameScene {
	p := objects.NewPlayer(w, h, a)
	s := resolv.NewSpace(w, h, 16, 16)
	s.Add(p.CollisionObj)

	t := GameTimers{
		meteorSpawn:        helpers.NewTimer(config.BaseMeteorSpawnTime),
		hyperspaceCooldown: helpers.NewTimer(config.HyperspaceCooldown),
	}

	m := GameMeteors{
		levelCount:   3,
		velocity:     config.BaseVelocity,
		list:         make(map[int]*objects.Meteor),
		currentCount: 0,
	}

	g := &GameScene{
		level:   1,
		timers:  t,
		meteors: m,
		lasers:  make(map[int]*objects.Laser),
		lives:   config.NumberOfLives,
		player:  p,
		space:   s,
		stars:   objects.GenerateStars(config.StarsNumber, w, h),
	}

	return g
}

func (sc *GameScene) Update(state *state.State) error {
	sc.spawnMeteors()

	for _, m := range sc.meteors.list {
		m.Update(state)
	}

	sc.hyperspace(state)
	sc.player.Update(state)

	sc.updateLasers(state)
	_ = sc.checkForCollisions(state)

	return nil
}

func (sc *GameScene) Draw(screen *ebiten.Image) {
	for _, s := range sc.stars {
		s.Draw(screen)
	}

	for _, m := range sc.meteors.list {
		m.Draw(screen)
	}

	sc.player.Draw(screen)
}

func (sc *GameScene) IsFinished(state *state.State) bool {
	if sc.lives <= 0 {
		state.AddScene(NewGameOverScene(state.ScreenWidth, state.ScreenHeight))
		return true
	}

	if len(sc.meteors.list) == sc.meteors.levelCount && sc.meteors.currentCount <= 0 {
		state.AddScene(NewLevelStartScene(state.ScreenWidth, state.ScreenHeight))
		return true
	}

	return false
}

func (sc *GameScene) spawnMeteors() {
	sc.timers.meteorSpawn.Update()
	if sc.timers.meteorSpawn.IsReady() {
		sc.timers.meteorSpawn.Reset()
		if len(sc.meteors.list) <= sc.meteors.levelCount && sc.meteors.currentCount < sc.meteors.levelCount {
			m := objects.NewMeteor(config.MeteorLarge, sc.meteors.velocity, len(sc.meteors.list)-1, sc.space.Width(), sc.space.Height())
			sc.space.Add(m.CollisionObj)
			sc.meteors.currentCount++
			sc.meteors.list[sc.meteors.currentCount] = m
		}
	}
}

func (sc *GameScene) updateLasers(state *state.State) {
	for i, l := range sc.player.PlayerLasers.Lasers {
		if sc.lasers[i] == nil {
			sc.lasers[i] = l
			sc.space.Add(l.CollisionObj)

			continue
		}

		if l.CollisionObj.Position().X < 0 {
			sc.removeOffscreenLasers(i, l)
		}

		if l.CollisionObj.Position().X > float64(state.ScreenWidth) {
			sc.removeOffscreenLasers(i, l)
		}

		if l.CollisionObj.Position().Y < 0 {
			sc.removeOffscreenLasers(i, l)
		}

		if l.CollisionObj.Position().Y > float64(state.ScreenHeight) {
			sc.removeOffscreenLasers(i, l)
		}
	}
}

func (sc *GameScene) removeOffscreenLasers(i int, l *objects.Laser) {
	delete(sc.player.PlayerLasers.Lasers, i)
	delete(sc.lasers, i)
	sc.space.Remove(l.CollisionObj)
}

func (sc *GameScene) hyperspace(state *state.State) {
	sc.timers.hyperspaceCooldown.Update()

	if ebiten.IsKeyPressed(ebiten.KeyE) && (sc.timers.hyperspaceCooldown == nil || sc.timers.hyperspaceCooldown.IsReady()) {
		var randX, randY int
		for {
			randX = rand.Intn(state.ScreenWidth)
			randY = rand.Intn(state.ScreenHeight)

			collision := sc.checkForCollisions(state)
			if !collision {
				break
			}
		}

		sc.space.Remove(sc.player.CollisionObj)
		sc.player.Position.X = float64(randX)
		sc.player.Position.Y = float64(randY)
		sc.player.CollisionObj.SetPosition(sc.player.Position.X, sc.player.Position.Y)
		sc.space.Add(sc.player.CollisionObj)
		sc.timers.hyperspaceCooldown.Reset()
	}
}

func (sc *GameScene) checkForCollisions(state *state.State) bool {
	//for i, m := range sc.meteors.list {
	//	fmt.Println(i, m.CollisionObj)
	//}

	//fmt.Println(sc.player.CollisionObj)

	//for i, l := range sc.player.PlayerLasers.Lasers {
	//	fmt.Println(i, l)
	//}

	// player vs meteor
	// player vs alien
	// player vs alien laser
	// meteor vs alien laser
	// meteor vs meteor
	//for _, m1 := range sc.meteors.list {
	//	for _, m2 := range sc.meteors.list {
	//		if m1.CollisionObj.IsIntersecting(m2.CollisionObj) {
	//			directionM1 := helpers.Vector{
	//				X: (float64(state.ScreenWidth/2) - m1.Position.X) * -1,
	//				Y: (float64(state.ScreenHeight/2) - m1.Position.Y) * -1,
	//			}
	//
	//			normalizedDirectionM1 := directionM1.Normalize()
	//			velocity := config.BaseVelocity * 1.2
	//		}
	//	}
	//}

	return false
}
