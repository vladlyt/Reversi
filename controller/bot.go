package controller

import (
	"Reversi/game"
	"math/rand"
)

type BotPlayer struct{}

func NewBotPlayer() *BotPlayer {
	return &BotPlayer{}
}

func (p BotPlayer) Play(board game.Board, color game.Color) game.Turn {
	turns := make([]game.Turn, 0)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			turn := game.Turn{
				Row:   i,
				Col:   j,
				Color: color,
			}
			if !board[i][j].IsFilled && board.HasEnemyCellNearThatCanTake(turn) {
				turns = append(turns, turn)
			}
		}
	}
	if len(turns) != 0 {
		return turns[rand.Int31n(int32(len(turns)))]
	}
	return game.Turn{}
}
