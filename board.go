package main

import (
	"fmt"
	"math/rand"
)

const (
	Size = 4
)

type Board [Size][Size]int

func initializeBoard() Board {
	board := Board{}

	for i := 0; i < 2; i++ {
		placeRandomTile(&board)
	}

	return board
}

func printBoard(board Board) {
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println()
}

func placeRandomTile(board *Board) {
	var emptyCells []int

	for i, row := range board {
		for j, value := range row {
			if value == 0 {
				emptyCells = append(emptyCells, i*Size+j)
			}
		}
	}

	if len(emptyCells) > 0 {
		randomIndex := emptyCells[rand.Intn(len(emptyCells))]
		row, col := randomIndex/Size, randomIndex%Size

		value := 2
		if rand.Intn(2) == 1 {
			value = 4
		}

		board[row][col] = value
	}
}
