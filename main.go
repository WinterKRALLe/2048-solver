package main

import (
	"fmt"
)

const (
	MaxRounds = 30
)

func main() {
	wins, losses := 0, 0

	for round := 1; round <= MaxRounds; round++ {
		fmt.Printf("Round %d\n", round)

		board := initializeBoard()
		printBoard(board)

		for !isGameOver(board) {
			move := getBestMoveDynamicDepth(board)

			fmt.Println("Move:", moveNames[move])
			board = makeMove(board, move)
			printBoard(board)
		}

		if hasMaxTile(board) {
			fmt.Println("You win!")
			wins++
		} else {
			fmt.Println("Game over. You lose.")
			losses++
		}

		fmt.Printf("Wins: %d, Losses: %d\n", wins, losses)
	}
}
