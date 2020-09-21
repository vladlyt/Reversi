package game

import (
	"fmt"
)

type Game struct {
	CurrentColor Color
	Board        Board
	events       chan<- Event
}

func NewGame(events chan<- Event) *Game {
	board := NewBoard()
	events <- Event{
		Event: BoardUpdate,
		Board: board,
	}
	return &Game{
		CurrentColor: BLACK,
		Board:        board,
		events:       events,
	}
}

func (g *Game) IsGameFinished() bool {
	return !g.CanDoTurn(BLACK) && !g.CanDoTurn(WHITE)
}

func (g *Game) CanDoTurn(color Color) bool {
	return g.Board.CanDoTurn(color)
}

func (g *Game) SetNextPlayer() {
	g.CurrentColor = !g.CurrentColor
}

func (g *Game) GetResults() (white int, black int) {
	return g.Board.GetCellColorCount(WHITE), g.Board.GetCellColorCount(BLACK)
}

func (g *Game) DoTurn(turn Turn) error {
	err := g.Board.UpdateBoard(turn)
	g.events <- Event{
		Event: BoardUpdate,
		Board: g.Board,
	}
	fmt.Println(g.Board)
	if err == nil && g.CanDoTurn(!turn.Color) {
		g.SetNextPlayer()
	}
	return err
}
