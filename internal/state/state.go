package state

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Scene interface {
	Update(state *State) error
	Draw(screen *ebiten.Image)
	IsFinished(state *State) bool
}

type State struct {
	AudioContext *audio.Context
	//PlayerCurrentPosition helpers.Vector
	//PlayerBounds          helpers.Vector
	//PlayerRotation        float64
	ShieldCount  int
	ScreenWidth  int
	ScreenHeight int
	Scenes       []Scene
	CurrentScene Scene
	NextScene    Scene
}

func NewState(screenWidth, screenHeight int) *State {
	return &State{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
}

func (s *State) AddScene(scene Scene) {
	s.Scenes = append(s.Scenes, scene)

	if s.CurrentScene == nil {
		s.CurrentScene = scene
	} else if s.CurrentScene != nil && s.NextScene == nil {
		s.NextScene = scene
	}
}
