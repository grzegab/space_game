package scenes

import (
	"image/color"

	"github.com/grzegab/sample_game/internal/assets"
	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/helpers"
	"github.com/grzegab/sample_game/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type CompanyScene struct {
	Width    int
	Height   int
	Duration *helpers.Timer
}

func NewCompanyScene(w, h int) *CompanyScene {
	return &CompanyScene{
		Width:    w,
		Height:   h,
		Duration: helpers.NewTimer(config.IntroDuration),
	}
}

func (sc *CompanyScene) Update(state *state.State) error {
	sc.Duration.Update()

	return nil
}

func (sc *CompanyScene) Draw(screen *ebiten.Image) {
	textToDraw := "GrzeGab Productions"
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}

	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(float64(sc.Width/2), float64(sc.Height-400))
	text.Draw(screen, textToDraw, &text.GoTextFace{
		Source: assets.IntroCompanyFontFace,
		Size:   48,
	}, op)

	imgToDraw := assets.IntroComp
	bounds := imgToDraw.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	imgOp := &ebiten.DrawImageOptions{}
	imgOp.GeoM.Translate(-halfW, -halfH)
	imgOp.GeoM.Rotate(10.1)
	imgOp.GeoM.Translate(halfW, halfH)
	imgOp.GeoM.Translate(float64(sc.Width/5+30), float64(sc.Height/2-350))
	imgOp.GeoM.Scale(2, 2)
	screen.DrawImage(imgToDraw, imgOp)
}

func (sc *CompanyScene) IsFinished(_ *state.State) bool {
	return sc.Duration.IsReady()
}
