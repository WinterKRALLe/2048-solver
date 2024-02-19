package main

import "fmt"

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
