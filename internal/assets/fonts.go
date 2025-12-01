package assets

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var IntroCompanyFontFace = mustLoadFontFace("fonts/intro.ttf")
var TitleFont = mustLoadFontFace("fonts/title.ttf")
var ScoreFont = mustLoadFontFace("fonts/score.ttf")
var LevelFont = mustLoadFontFace("fonts/score.ttf")

func mustLoadFontFace(path string) *text.GoTextFaceSource {
	f, err := assets.ReadFile(path)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(f)
	ts, err := text.NewGoTextFaceSource(r)
	if err != nil {
		panic(err)
	}

	return ts
}
