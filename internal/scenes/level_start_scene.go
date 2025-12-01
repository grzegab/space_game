package scenes

import (
	"fmt"
	"image/color"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/grzegab/sample_game/internal/objects"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type LevelStartScene struct {
	Width        int
	Height       int
	CurrentLevel int
	Stars        []*objects.Star
	Duration     *helpers.Timer
}

func NewLevelStartScene(w, h int) *LevelStartScene {
	return &LevelStartScene{
		Width:    w,
		Height:   h,
		Duration: helpers.NewTimer(config.IntroDuration),
	}
}

func (sc *LevelStartScene) Update(state *state.State) error {
	sc.Duration.Update()

	return nil
}

func (sc *LevelStartScene) Draw(screen *ebiten.Image) {
	for _, s := range sc.Stars {
		s.Draw(screen)
	}

	textToDraw := fmt.Sprintf("Level %d", sc.CurrentLevel)
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(float64(sc.Width)/2, float64(sc.Height)/2)
	text.Draw(screen, textToDraw, &text.GoTextFace{
		Source: assets.LevelFont,
		Size:   48,
	}, op)
}

func (sc *LevelStartScene) IsFinished(_ *state.State) bool {
	return sc.Duration.IsReady()
}
