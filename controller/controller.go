package controller

import (
	"Reversi/game"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

type Players map[game.Color]Player

type GameController struct {
	win       *pixelgl.Window
	gameBoard [8][8]pixel.Vec
	h         float64
	w         float64
}

type Controller interface {
	Run(boards chan<- game.Event)
}

func NewGameController(
	win *pixelgl.Window,
	gameBoard [8][8]pixel.Vec,
	spriteHeight, spriteWidth float64,
) *GameController {
	return &GameController{
		win:       win,
		gameBoard: gameBoard,
		h:         spriteHeight,
		w:         spriteWidth,
	}
}

func (controller *GameController) initPlayers() Players {
	players := make(Players)

	var playerFirst Player
	var playerSecond Player
	for {
		input := controller.win.Typed()
		if input == "1" {
			playerFirst = NewRealPlayer(controller.win, controller.gameBoard, controller.h, controller.w)
			playerSecond = NewBotPlayer()
			break
		} else if input == "2" {
			playerFirst = NewRealPlayer(controller.win, controller.gameBoard, controller.h, controller.w)
			playerSecond = NewRealPlayer(controller.win, controller.gameBoard, controller.h, controller.w)
			break
		}
		time.Sleep(time.Millisecond)
	}

	players[game.BLACK] = playerFirst
	players[game.WHITE] = playerSecond

	return players
}

func (controller *GameController) Run(events chan<- game.Event) {
	for {
		events <- game.Event{
			Event: game.GameStarted,
		}
		players := controller.initPlayers()
		gameModel := game.NewGame(events)
		for !gameModel.IsGameFinished() {
			currentColor := gameModel.CurrentColor
			turn := players[currentColor].Play(gameModel.Board, currentColor)
			err := gameModel.DoTurn(turn)
			if err != nil {
				fmt.Println("ERROR", err)
			}
		}
		whiteRes, blackRes := gameModel.GetResults()
		events <- game.Event{
			Event:       game.WinnerScreen,
			WhiteResult: whiteRes,
			BlackResult: blackRes,
		}
		for !controller.win.JustPressed(pixelgl.KeyEnter) {
			time.Sleep(time.Millisecond)
		}
	}
}
