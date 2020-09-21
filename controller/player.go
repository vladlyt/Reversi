package controller

import "Reversi/game"

type Player interface {
	Play(game.Board, game.Color) game.Turn
}
