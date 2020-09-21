package main

import (
	"Reversi/controller"
	"Reversi/game"
	"Reversi/view"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	SpriteHeight = 32
	SpriteWidth  = 32
)

func StartGame(gameController controller.Controller, gameView view.View) {
	events := make(chan game.Event)
	go gameController.Run(events)
	gameView.Run(events)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Reversi",
		Bounds: pixel.R(0, 0, 600, 450),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	gameView := view.NewGameView(
		win,
		SpriteHeight,
		SpriteWidth,
	)
	gameController := controller.NewGameController(
		win,
		gameView.GetBoard(),
		float64(SpriteHeight),
		float64(SpriteWidth),
	)

	StartGame(gameController, gameView)
}

func main() {
	pixelgl.Run(run)
}
