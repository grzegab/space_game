package scenes

import (
	"image/color"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/objects"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type IntroScene struct {
	Meteors     map[int]*objects.Meteor
	MeteorCount int
	Stars       []*objects.Star
	IsEnded     bool
}

func (sc *IntroScene) Update(state *state.State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		gs := NewGameScene(state.ScreenWidth, state.ScreenHeight, state.AudioContext)
		state.AddScene(gs)
		sc.IsEnded = true
		return nil
	}

	if len(sc.Meteors) < 10 {
		m := objects.NewMeteor(config.MeteorLarge, config.BaseVelocity, len(sc.Meteors)-1, state.ScreenWidth, state.ScreenHeight)
		sc.MeteorCount++
		sc.Meteors[sc.MeteorCount] = m
	}

	for _, m := range sc.Meteors {
		m.Update(state)
	}

	return nil
}

func (sc *IntroScene) Draw(screen *ebiten.Image) {
	for _, s := range sc.Stars {
		s.Draw(screen)
	}

	for _, m := range sc.Meteors {
		m.Draw(screen)
	}

	textToDraw := "Spacebar to play"
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}

	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(float64(screen.Bounds().Dx()/2), float64(screen.Bounds().Dy()/2))
	text.Draw(screen, textToDraw, &text.GoTextFace{
		Source: assets.TitleFont,
		Size:   48,
	}, op)
}

func (sc *IntroScene) IsFinished(_ *state.State) bool {
	return sc.IsEnded
}
