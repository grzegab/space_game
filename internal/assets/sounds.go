package assets

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

var ThrustSound = mustLoadOggSound("audio/thrust.ogg")
var LaserOneSound = mustLoadOggSound("audio/fire.ogg")
var LaserTwoSound = mustLoadOggSound("audio/fire.ogg")
var ExplosionSound = mustLoadOggSound("audio/explosion.ogg")
var BeatOneSound = mustLoadOggSound("audio/beat1.ogg")
var BeatTwoSound = mustLoadOggSound("audio/beat2.ogg")
var ShieldSound = mustLoadOggSound("audio/shield.ogg")
var AlienSound = mustLoadOggSound("audio/alien-sound.ogg")
var AlienLaserSound = mustLoadOggSound("audio/alien-laser.ogg")

func mustLoadOggSound(path string) *vorbis.Stream {
	f, err := assets.ReadFile(path)
	if err != nil {
		panic(err)
	}
	stream, err := vorbis.DecodeWithoutResampling(bytes.NewReader(f))
	if err != nil {
		panic(err)
	}

	return stream
}
