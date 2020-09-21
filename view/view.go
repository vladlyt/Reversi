package view

import (
	"Reversi/game"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"image"
	"os"
)

type GameViewBoard [8][8]pixel.Vec

type GameView struct {
	win             *pixelgl.Window
	board           GameViewBoard
	rectSprite      *pixel.Sprite
	blackDiskSprite *pixel.Sprite
	whiteDickSprite *pixel.Sprite
}

type View interface {
	Run(events <-chan game.Event)
}

func getSpriteFromFilepath(path string) (*pixel.Sprite, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	picture := pixel.PictureDataFromImage(img)
	return pixel.NewSprite(picture, picture.Bounds()), nil
}

func NewGameView(win *pixelgl.Window, spriteHeight int, spriteWidth int) *GameView {

	blackDiskSprite, err := getSpriteFromFilepath("static/black.png")
	if err != nil {
		panic(err)
	}

	whiteDiskSprite, err := getSpriteFromFilepath("static/white.png")
	if err != nil {
		panic(err)
	}

	rectSprite, err := getSpriteFromFilepath("static/rect.png")
	if err != nil {
		panic(err)
	}

	// board init
	board := GameViewBoard{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			// TODO optimize?
			board[i][j] = pixel.V(
				float64((i-4)*spriteHeight)+win.Bounds().Center().X+float64(i),
				float64((j-4)*spriteWidth)+win.Bounds().Center().Y+float64(j),
			)
		}
	}

	return &GameView{
		win:             win,
		board:           board,
		rectSprite:      rectSprite,
		blackDiskSprite: blackDiskSprite,
		whiteDickSprite: whiteDiskSprite,
	}
}

func (view *GameView) GetBoard() GameViewBoard {
	return view.board
}

func (view *GameView) drawWelcomeScreen() {
	view.win.Clear(colornames.Skyblue)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(50, 250), basicAtlas)

	fmt.Fprintln(basicTxt, "Welcome to Reversi!")
	fmt.Fprintln(basicTxt, "Please, press '1' to play against bot")
	fmt.Fprintln(basicTxt, "And '2' to play against real player")

	basicTxt.Draw(view.win, pixel.IM.Scaled(basicTxt.Orig, 2))
}

func (view *GameView) drawBoard(board game.Board) {
	view.win.Clear(colornames.Skyblue)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			view.rectSprite.Draw(view.win, pixel.IM.Moved(view.board[i][j]))
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j].IsFilled {
				if board[i][j].Color == game.BLACK {
					view.blackDiskSprite.Draw(
						view.win,
						pixel.IM.Moved(view.board[j][7-i]),
					)
				} else {
					view.whiteDickSprite.Draw(
						view.win,
						pixel.IM.Moved(view.board[j][7-i]),
					)
				}
			}
		}
	}
}
func (view *GameView) drawWinnerScreen(whiteResult, blackResult int) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(50, 400), basicAtlas)

	if whiteResult == blackResult {
		fmt.Fprintln(basicTxt, "IT IS DRAW!")
	} else {
		if whiteResult > blackResult {
			fmt.Fprintln(basicTxt, fmt.Sprintf("WHITE PLAYER WINS, SCORE: %d", whiteResult))
		} else {
			fmt.Fprintln(basicTxt, fmt.Sprintf("BLACK PLAYER WINS, SCORE: %d", blackResult))
		}
	}
	fmt.Fprintln(basicTxt, "Press ENTER to play again")

	basicTxt.Draw(view.win, pixel.IM.Scaled(basicTxt.Orig, 2))
}

func (view *GameView) Run(events <-chan game.Event) {
	for !view.win.Closed() {
		select {
		case event := <-events:
			switch event.Event {
			case game.GameStarted:
				view.drawWelcomeScreen()
			case game.BoardUpdate:
				view.drawBoard(event.Board)
			case game.WinnerScreen:
				view.drawWinnerScreen(event.WhiteResult, event.BlackResult)
			}
		default:
		}
		view.win.Update()
	}
}
