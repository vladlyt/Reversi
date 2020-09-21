package game

import "fmt"

type Board [8][8]Cell

func (b Board) String() string {
	out := "一一一一一一一一一一一一一一一一一\n"
	for _, row := range b {
		for _, cell := range row {
			out += fmt.Sprintf("| %s ", cell)
		}
		out += "|\n一一一一一一一一一一一一一一一一一\n"
	}
	return out
}

func NewBoard() Board {
	cells := [8][8]Cell{}
	for i := 0; i < 8; i++ {
		cells[i] = [8]Cell{}
		for j := 0; j < 8; j++ {
			cells[i][j] = Cell{}
		}
	}
	cells[3][3].Fill(WHITE)
	cells[3][4].Fill(BLACK)
	cells[4][3].Fill(BLACK)
	cells[4][4].Fill(WHITE)

	return cells
}

func (b *Board) UpdateBoard(turn Turn) error {
	err := b.CheckTurn(turn)
	if err != nil {
		return err
	}
	b.SetTurnOnBoard(turn)
	return nil
}

func (b *Board) SetTurnOnBoard(turn Turn) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			cell := b.GetCellEvenIfNotExist(turn.Row+i, turn.Col+j)
			if cell.IsFilled && cell.Color == !turn.Color {
				if b.HasOppositeCell(turn.Row, turn.Col, i, j, turn.Color) {
					b.SwapAllToOppositeCell(turn.Row, turn.Col, i, j, turn.Color)
				}
			}
		}
	}
	b[turn.Row][turn.Col].Fill(turn.Color)
}

func (b *Board) GetCellEvenIfNotExist(row, col int) Cell {
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return Cell{}
	}
	return b[row][col]
}

func IsOutOfRange(val int) bool {
	return val < 0 || val > 7
}

func (b *Board) GetOppositeCell(row, col, dRow, dCol int, color Color) (int, int) {
	for i, j := row+dRow, col+dCol; !IsOutOfRange(i) && !IsOutOfRange(j) && b[i][j].IsFilled; i, j = i+dRow, j+dCol {
		if b[i][j].Color == color {
			return i, j
		}
	}
	return -1, -1
}

func (b *Board) HasOppositeCell(row, col, dRow, dCol int, color Color) bool {
	i, j := b.GetOppositeCell(row, col, dRow, dCol, color)
	return i != -1 && j != -1
}

func (b *Board) SwapAllToOppositeCell(row, col, dRow, dCol int, color Color) {
	toRow, toCol := b.GetOppositeCell(row, col, dRow, dCol, color)
	for i, j := row+dRow, col+dCol; i != toRow || j != toCol; i, j = i+dRow, j+dCol {
		b[i][j].Swap()
	}
}

func (b *Board) HasEnemyCellNearThatCanTake(turn Turn) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			cell := b.GetCellEvenIfNotExist(turn.Row+i, turn.Col+j)
			if cell.IsFilled && cell.Color == !turn.Color {
				if b.HasOppositeCell(turn.Row, turn.Col, i, j, turn.Color) {
					return true
				}
			}
		}
	}
	return false
}

func (b *Board) CanDoTurn(color Color) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if !b[i][j].IsFilled && b.HasEnemyCellNearThatCanTake(Turn{
				Row:   i,
				Col:   j,
				Color: color,
			}) {
				return true
			}
		}
	}
	return false
}

func (b *Board) CheckTurn(turn Turn) error {
	if turn.Row < 0 || turn.Row > 7 {
		return fmt.Errorf("row is out of range")
	}
	if turn.Col < 0 || turn.Col > 7 {
		return fmt.Errorf("column is out of range")
	}
	cell := b[turn.Row][turn.Col]
	if cell.IsFilled {
		return fmt.Errorf("cell is already filled")
	}
	if !b.HasEnemyCellNearThatCanTake(turn) {
		return fmt.Errorf("no enemy cell near")
	}

	return nil
}

func (b *Board) GetCellColorCount(color Color) int {
	count := 0
	for i := range b {
		for j := range b[i] {
			if b[i][j].IsFilled && b[i][j].Color == color {
				count++
			}
		}
	}
	return count
}
