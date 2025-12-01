package core

import (
	"fmt"

	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/objects"
	"github.com/grzegab/sample_game/internal/scenes"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var (
	EndGame = fmt.Errorf("end of game")
)

type Game struct {
	input             Input
	config            config.BaseConfig
	state             *state.State
	fadingInCount     int
	fadingOutCount    int
	currentSceneImage *ebiten.Image
	nextSceneImage    *ebiten.Image
}

func NewGame(cfg config.BaseConfig) *Game {
	meteors := make(map[int]*objects.Meteor)
	stars := objects.GenerateStars(config.StarsNumber, cfg.ScreenWidth, cfg.ScreenHeight)
	st := state.NewState(cfg.ScreenWidth, cfg.ScreenHeight)
	st.AudioContext = audio.NewContext(48000)
	st.ShieldCount = config.NumberOfShields

	g := &Game{
		config:         cfg,
		state:          st,
		fadingInCount:  config.FadingInLength,
		fadingOutCount: 0,
	}

	company := scenes.NewCompanyScene(cfg.ScreenWidth, cfg.ScreenHeight)
	g.state.AddScene(company)
	intro := &scenes.IntroScene{Meteors: meteors, Stars: stars}
	g.state.AddScene(intro)

	g.currentSceneImage = ebiten.NewImage(cfg.ScreenWidth, cfg.ScreenHeight)
	g.nextSceneImage = ebiten.NewImage(cfg.ScreenWidth, cfg.ScreenHeight)

	return g
}

func (g *Game) Update() error {
	g.input.Update()

	// Update current scene screen
	if g.fadingInCount == config.FadingInLength && g.fadingOutCount == 0 {
		err := g.state.CurrentScene.Update(g.state)
		if err != nil {
			return fmt.Errorf("error while updating scene: %w", err)
		}
	}

	// Make transitions between scenes
	if g.state.CurrentScene.IsFinished(g.state) {
		if g.fadingOutCount < config.FadingOutLength {
			g.fadingOutCount++
			return nil
		}

		if g.fadingInCount > 0 {
			g.fadingInCount--
			return nil
		}

		// Quit if no more scenes
		if len(g.state.Scenes) == 0 {
			return fmt.Errorf("game has ended: %w", EndGame)
		}

		// Take another scene from heap
		g.state.Scenes = g.state.Scenes[1:]
		if len(g.state.Scenes) > 1 {
			g.state.NextScene = g.state.Scenes[1]
		} else {
			g.state.NextScene = nil
		}

		if len(g.state.Scenes) > 0 {
			g.state.CurrentScene = g.state.Scenes[0]
			g.resetCounters()
		}
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	// Print current scene screen
	if g.fadingInCount == config.FadingInLength && g.fadingOutCount == 0 {
		g.state.CurrentScene.Draw(screen)

		return
	}

	g.currentSceneImage.Clear()
	g.nextSceneImage.Clear()

	if g.fadingOutCount < config.FadingOutLength {
		g.state.CurrentScene.Draw(g.currentSceneImage)
		currentOp := &ebiten.DrawImageOptions{}
		alpha := 1 - float32(g.fadingOutCount)/float32(config.FadingOutLength)
		currentOp.ColorScale.ScaleAlpha(alpha)
		screen.DrawImage(g.currentSceneImage, currentOp)
	}

	if g.fadingInCount < config.FadingInLength && g.state.NextScene != nil {
		g.state.NextScene.Draw(g.nextSceneImage)
		nextOp := &ebiten.DrawImageOptions{}
		alpha := 1 - float32(g.fadingInCount)/float32(config.FadingInLength)
		nextOp.ColorScale.ScaleAlpha(alpha)
		screen.DrawImage(g.nextSceneImage, nextOp)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.config.ScreenWidth, g.config.ScreenHeight
}

func (g *Game) generateTransitionImage() *ebiten.Image {
	return ebiten.NewImage(g.config.ScreenWidth, g.config.ScreenHeight)
}

func (g *Game) resetCounters() {
	g.fadingInCount = config.FadingInLength
	g.fadingOutCount = 0
}
