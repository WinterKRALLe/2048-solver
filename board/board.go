package board

import (
	"fmt"
	"math/rand"
)

func InitializeBoard() Board {
	board := Board{}

	for i := 0; i < 2; i++ {
		placeRandomTile(&board)
	}

	return board
}

func PrintBoard(board Board) {
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println()
}

func CopyBoard(board Board) Board {
	var newBoard Board
	for i, row := range board {
		for j, value := range row {
			newBoard[i][j] = value
		}
	}
	return newBoard
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

func HasMaxTile(board Board) bool {
	maxTile := 0
	for _, row := range board {
		for _, value := range row {
			if value > maxTile {
				maxTile = value
			}
		}
	}
	return maxTile >= 2048
}

func CalculateScore(boardState Board) int {
	totalScore := 0
	for _, row := range boardState {
		for _, value := range row {
			totalScore += value
		}
	}
	return totalScore
}
