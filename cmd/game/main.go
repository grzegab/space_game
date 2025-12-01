package main

import (
	"errors"
	"os"

	"github.com/grzegab/sample_game/internal/config"
	"github.com/grzegab/sample_game/internal/core"

	//"github.com/grzegab/sample_game/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	cfg := config.BaseConfig{
		ScreenWidth:  1920,
		ScreenHeight: 1080,
	}

	game := core.NewGame(cfg)

	ebiten.SetWindowTitle("Asteroids GO Game")
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)

	if err := ebiten.RunGame(game); err != nil {
		if errors.Is(err, core.EndGame) {
			os.Exit(0)
		}

		panic(err)
	}
}
