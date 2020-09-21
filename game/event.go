package game

const (
	GameStarted = iota
	BoardUpdate
	WinnerScreen
)

type Event struct {
	Event       int
	Board       Board
	WhiteResult int
	BlackResult int
}
