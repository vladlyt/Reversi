package controller

import (
	"Reversi/game"
	"Reversi/view"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

type RealPlayer struct {
	win       *pixelgl.Window
	gameBoard view.GameViewBoard
	h         float64
	w         float64
}

func NewRealPlayer(win *pixelgl.Window, gameBoard view.GameViewBoard, h float64, w float64) *RealPlayer {
	return &RealPlayer{
		win:       win,
		gameBoard: gameBoard,
		h:         h,
		w:         w,
	}
}

func (p RealPlayer) Play(board game.Board, color game.Color) game.Turn {
	for {
		if !p.win.JustPressed(pixelgl.MouseButtonLeft) {
			time.Sleep(time.Millisecond)
			continue
		}
		pos := p.win.MousePosition()
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if p.isVectorsOverlap(pos, p.gameBoard[j][7-i]) {
					return game.Turn{
						Row:   i,
						Col:   j,
						Color: color,
					}
				}
			}
		}
	}
}

func (p RealPlayer) isVectorsOverlap(first, second pixel.Vec) bool {
	return first.X >= second.X-p.h/2 && first.X < second.X+p.h/2 && first.Y >= second.Y-p.w/2 && first.Y < second.Y+p.w/2
}
