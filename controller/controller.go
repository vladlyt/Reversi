package controller

import (
	"Reversi/game"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

type Players map[game.Color]Player

type GameRunner struct {
	win       *pixelgl.Window
	gameBoard [8][8]pixel.Vec
	h         float64
	w         float64
}

type Controller interface {
	Run(boards chan<- game.Event)
}

func NewGameRunner(
	win *pixelgl.Window,
	gameBoard [8][8]pixel.Vec,
	spriteHeight, spriteWidth float64,
) *GameRunner {
	return &GameRunner{
		win:       win,
		gameBoard: gameBoard,
		h:         spriteHeight,
		w:         spriteWidth,
	}
}

func (runner *GameRunner) initPlayers() Players {
	players := make(Players)

	var playerFirst Player
	var playerSecond Player
	for {
		input := runner.win.Typed()
		if input == "1" {
			playerFirst = NewRealPlayer(runner.win, runner.gameBoard, runner.h, runner.w)
			playerSecond = NewBotPlayer()
			break
		} else if input == "2" {
			playerFirst = NewRealPlayer(runner.win, runner.gameBoard, runner.h, runner.w)
			playerSecond = NewRealPlayer(runner.win, runner.gameBoard, runner.h, runner.w)
			break
		}
		time.Sleep(time.Millisecond)
	}

	players[game.BLACK] = playerFirst
	players[game.WHITE] = playerSecond

	return players
}

func (runner *GameRunner) Run(events chan<- game.Event) {
	for {
		events <- game.Event{
			Event: game.GameStarted,
		}
		players := runner.initPlayers()
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
		for !runner.win.JustPressed(pixelgl.KeyEnter) {
			time.Sleep(time.Millisecond)
		}
	}
}
