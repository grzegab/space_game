package assets

import (
	"fmt"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

var IntroComp = mustLoadImage("images/intro/comp.png")
var PlayerImage = mustLoadImage("images/player.png")
var BigMeteorSprites = mustLoadImages("images/meteors/*.png")
var SmallMeteorSprite = mustLoadImages("images/meteors-small/*.png")
var LaserSprite = mustLoadImage("images/laser.png")
var ExplosionSprite = mustLoadImage("images/explosion.png")
var ExplosionSmallSprite = mustLoadImage("images/explosion-small.png")
var Explosion = createSequence("explosion")
var ExhaustSprite = mustLoadImage("images/fire.png")
var LifeIndicator = mustLoadImage("images/life-indicator.png")
var ShieldSprite = mustLoadImage("images/shield.png")
var ShieldIndicator = mustLoadImage("images/shield-indicator.png")
var HyperSpaceIndicator = mustLoadImage("images/hyperspace.png")
var AlienSprites = mustLoadImages("images/aliens/*.png")
var AlienLaserSprite = mustLoadImage("images/red-laser.png")

func createSequence(s string) []*ebiten.Image {
	var frames []*ebiten.Image
	for i := 1; i <= 12; i++ {
		frame := mustLoadImage(fmt.Sprintf("images/%s/%d.png", s, i))
		frames = append(frames, frame)
	}

	return frames
}

func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}

func mustLoadImage(path string) *ebiten.Image {
	f, err := assets.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
