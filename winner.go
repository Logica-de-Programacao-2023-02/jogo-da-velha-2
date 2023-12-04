package main

type WinnableBoard interface {
	Value(i int) CellType
}

func Winner(board WinnableBoard) CellType {
	winner := checkWinnerColumns(board)
	if winner != Empty {
		return winner
	}

	winner = checkWinnerDiagonal(board)
	if winner != Empty {
		return winner
	}

	winner = checkWinnerLines(board)
	if winner != Empty {
		return winner
	}

	return checkDraw(board)
}

func checkDraw(board WinnableBoard) CellType {
	for i := 0; i < 9; i++ {
		if board.Value(i) == Empty {
			return Empty
		}
	}

	return CellDraw
}

func checkWinnerColumns(board WinnableBoard) CellType {
	for i := 0; i < 3; i++ {
		if board.Value(i) == Empty {
			continue
		}
		if equals(board, i, i+3, i+6) {
			return board.Value(i)
		}
	}
	return Empty
}

func checkWinnerDiagonal(board WinnableBoard) CellType {
	center := board.Value(4)
	if center == Empty {
		return Empty
	}
	if equals(board, 0, 4, 8) {
		return center
	}
	if equals(board, 2, 4, 6) {
		return center
	}
	return Empty
}

func checkWinnerLines(board WinnableBoard) CellType {
	for i := 0; i < 9; i += 3 {
		if board.Value(i) == Empty {
			continue
		}
		if equals(board, i, i+1, i+2) {
			return board.Value(i)
		}
	}
	return Empty
}

func equals(board WinnableBoard, i, j, k int) bool {
	valI := board.Value(i)
	valJ := board.Value(j)
	valK := board.Value(k)

	switch {
	case valI == valJ && valI == valK:
		return true
	case valI == valJ && valK == CellDraw:
		return true
	case valI == valK && valJ == CellDraw:
		return true
	case valJ == valK && valI == CellDraw:
		return true
	default:
		return false
	}
}
