package scenes

import (
	"github.com/grzegab/sample_game/internal/objects"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
)

//
//import (
//	"image/color"
//	"os"
//
//	"github.com/grzegab/sample_game/internal/assets"
//	"github.com/hajimehoshi/ebiten/v2"
//	"github.com/hajimehoshi/ebiten/v2/text/v2"
//)
//

type GameOverScene struct {
	Width       int
	Height      int
	game        *GameScene
	meteors     map[int]*objects.Meteor
	meteorCount int
	stars       []*objects.Star
}

func NewGameOverScene(w, h int) *GameOverScene {
	return &GameOverScene{
		Width:  w,
		Height: h,
	}
}

func (sc *GameOverScene) Update(state *state.State) error {
	return nil
}

func (sc *GameOverScene) Draw(screen *ebiten.Image) {
	for _, s := range sc.stars {
		s.Draw(screen)
	}
}

func (sc *GameOverScene) IsFinished(_ *state.State) bool {
	return false
}

//
//func (gos *GameOverScene) Draw(screen *ebiten.Image) {
//	for _, s := range gos.stars {
//		s.Draw(screen)
//	}
//
//	for _, m := range gos.meteors {
//		m.Draw(screen)
//	}
//
//	textToDraw := "Game Over"
//	op := &text.DrawOptions{
//		LayoutOptions: text.LayoutOptions{
//			PrimaryAlign: text.AlignCenter,
//		},
//	}
//
//	op.ColorScale.ScaleWithColor(color.White)
//	op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2+100)
//	text.Draw(screen, textToDraw, &text.GoTextFace{
//		Source: assets.TitleFont,
//		Size:   48,
//	}, op)
//
//	if gos.game.score > originalHighScore {
//		hsToDraw := "New High Score!"
//		op := &text.DrawOptions{
//			LayoutOptions: text.LayoutOptions{
//				PrimaryAlign: text.AlignCenter,
//			},
//		}
//
//		op.ColorScale.ScaleWithColor(color.White)
//		op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2-200)
//		text.Draw(screen, hsToDraw, &text.GoTextFace{
//			Source: assets.TitleFont,
//			Size:   48,
//		}, op)
//	}
//}
//
//func (gos *GameOverScene) Update(state *State) error {
//	if len(gos.meteors) < 10 {
//		m := NewMeteor(0.25, gos.game, len(gos.meteors)-1)
//		gos.meteorCount++
//		gos.meteors[gos.meteorCount] = m
//	}
//
//	for _, m := range gos.meteors {
//		m.Update()
//	}
//
//	if ebiten.IsKeyPressed(ebiten.KeySpace) {
//		gos.game.Reset()
//		state.SceneManager.GoToScene(gos.game)
//	}
//
//	if ebiten.IsKeyPressed(ebiten.KeyQ) {
//		os.Exit(0)
//	}
//
//	return nil
//}
