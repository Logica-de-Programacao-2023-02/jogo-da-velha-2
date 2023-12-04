package main

import (
	"errors"
	"fmt"
)

type BigBoard [9]*TicTacToe

func (b BigBoard) Value(i int) CellType {
	return b[i].Winner()
}

type TicTacToe2 struct {
	Cells        BigBoard
	CurrentTurn  CellType
	plays        []int
	currentBoard int
}

func NewTicTacToe2() *TicTacToe2 {
	return &TicTacToe2{
		Cells: [9]*TicTacToe{
			NewTicTacToe(),
			NewTicTacToe(),
			NewTicTacToe(),
			NewTicTacToe(),
			NewTicTacToe(),
			NewTicTacToe(),
			NewTicTacToe(),
			NewTicTacToe(),
			NewTicTacToe(),
		},
		CurrentTurn:  CellO,
		currentBoard: -1,
	}
}

func (t TicTacToe2) Play() {
	winner := Empty
	for ; winner == Empty; winner = t.Winner() {
		for {
			if err := t.pickBoard(); err != nil {
				fmt.Printf("Houve um erro ao processar a jogada: %s\n", err)
			} else {
				break
			}
		}
		t.Print()
		fmt.Printf("Vez do jogador (%s)\nUtilize o NumPad:\n", t.CurrentTurn)
		t.play()
		if t.CurrentTurn == CellX {
			t.CurrentTurn = CellO
		} else {
			t.CurrentTurn = CellX
		}
	}
	t.Print()
	fmt.Println("O vencedor é", winner)
}

func (t *TicTacToe2) play() {
	for {
		if innerPos, err := t.Cells[t.currentBoard].play(t.CurrentTurn); err != nil {
			fmt.Printf("Houve um erro ao processar a jogada: %s\n", err)
		} else {
			t.plays = append(t.plays, innerPos)
			break
		}
	}
}

func (t *TicTacToe2) pickBoard() error {
	if len(t.plays) > 0 {
		last := t.plays[len(t.plays)-1]
		if t.Cells.Value(last) == Empty {
			t.currentBoard = last
			return nil
		}
	}

	t.currentBoard = -1
	t.Print()
	fmt.Printf("Vez do jogador (%s)\nEscolha um tabuleiro a ser jogado\nUtilize o NumPad:\n", t.CurrentTurn)

	var position int
	if _, err := fmt.Scan(&position); err != nil {
		return fmt.Errorf("entrada inválida: %w", err)
	}

	position--

	if t.Cells.Value(position) != Empty {
		return errors.New("não é possível jogar em um tabuleiro finalizado")
	}

	t.currentBoard = position
	return nil
}

func (t TicTacToe2) Print() {
	for i := 8; i >= 0; i-- {
		t.PrintLine(i)
		fmt.Println()
		if i%3 == 0 {
			fmt.Println(ColorBlue.Paint("========================="))
		}
	}
}

func (t TicTacToe2) PrintLine(i int) {
	i2 := start(i)
	for j := i2; j < i2+3; j++ {
		fmt.Print(ColorCyan.Paint("I"))
		if j == t.currentBoard {
			t.Cells[j].PrintLine((i%3)*3, ColorYellow)
		} else {
			t.Cells[j].PrintLine((i%3)*3, ColorReset)
		}
	}
	fmt.Print(ColorCyan.Paint("I"))
}

func (t TicTacToe2) Winner() CellType {
	return Winner(t.Cells)
}

func start(i int) int {
	if i >= 6 {
		return 6
	}
	if i >= 3 {
		return 3
	}
	return 0
}
