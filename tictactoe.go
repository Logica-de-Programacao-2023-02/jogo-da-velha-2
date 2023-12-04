package main

import (
	"errors"
	"fmt"
)

type SmallBoard [9]CellType

func (b SmallBoard) Value(i int) CellType {
	return b[i]
}

type TicTacToe struct {
	Cells       SmallBoard
	CurrentTurn CellType
}

func NewTicTacToe() *TicTacToe {
	return &TicTacToe{
		Cells:       [9]CellType{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		CurrentTurn: CellO,
	}
}

func (t TicTacToe) Print() {
	for i := 6; i >= 0; i -= 3 {
		t.PrintLine(i, ColorReset)
		fmt.Println()
	}
}

func (t TicTacToe) PrintLine(i int, color Color) {
	winner := t.Winner()
	for j := i; j < i+3; j++ {
		fmt.Print(color.Paint("|"))
		cell := t.Cells[j]
		if winner != Empty {
			cell = winner
		}
		fmt.Print(color.Paint(cell.String()))
	}
	fmt.Print(color.Paint("|"))
}

func (t *TicTacToe) Play() {
	winner := Empty
	for ; winner == Empty; winner = t.Winner() {
		fmt.Printf("Vez do jogador (%s)\nUtilize o NumPad:\n", t.CurrentTurn)
		t.Print()
		for {
			if _, err := t.play(t.CurrentTurn); err != nil {
				fmt.Printf("Houve um erro ao processar a jogada: %s\n", err)
			} else {
				break
			}
		}
		if t.CurrentTurn == CellX {
			t.CurrentTurn = CellO
		} else {
			t.CurrentTurn = CellX
		}
	}
	t.Print()
	fmt.Println("O vencedor é", winner)
}

func (t *TicTacToe) play(player CellType) (int, error) {
	var position int
	if _, err := fmt.Scan(&position); err != nil {
		return 0, fmt.Errorf("entrada inválida: %w", err)
	}

	position--

	if t.Cells[position] != Empty {
		return 0, errors.New("não é possível jogar em uma célula não-vazia")
	}

	t.Cells[position] = player
	return position, nil
}

func (t TicTacToe) Winner() CellType {
	return Winner(t.Cells)
}
