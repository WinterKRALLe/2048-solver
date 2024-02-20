package solver

import (
	b "Solver2048/board"
	"fmt"
	"math"
	"math/rand"
)

func monteCarloSimulation(board b.Board) int {
	numSimulations := 1000
	totalScore := 0

	for i := 0; i < numSimulations; i++ {
		simulatedBoard := board
		score := 0

		for !b.IsGameOver(simulatedBoard) {
			randomMove := b.Move(rand.Intn(4))
			if b.IsValidMove(simulatedBoard, randomMove) {
				simulatedBoard = b.MakeMove(simulatedBoard, randomMove)
				score += evaluateState(simulatedBoard)
			}
		}

		totalScore += score
	}

	return totalScore / numSimulations
}

func evaluateState(board b.Board) int {
	sumOfTiles := 0
	for _, row := range board {
		for _, value := range row {
			sumOfTiles += value
		}
	}
	return sumOfTiles
}

func selectBestMove(board b.Board) b.Move {
	bestMove := b.Up
	bestScore := math.MinInt64

	for _, move := range []b.Move{b.Up, b.Down, b.Left, b.Right} {
		if b.IsValidMove(board, move) {
			newBoard := b.MakeMove(board, move)
			score := monteCarloSimulation(newBoard)

			if score > bestScore {
				bestScore = score
				bestMove = move
			}
		}
	}

	return bestMove
}

func MonteCarlo() {
	board := b.InitializeBoard()

	for !b.IsGameOver(board) {
		bestMove := selectBestMove(board)

		board = b.MakeMove(board, bestMove)

		b.PrintBoard(board)
	}

	if b.HasMaxTile(board) {
		fmt.Println("You win!")
	} else {
		fmt.Println("Game over. You lose.")
	}
}
