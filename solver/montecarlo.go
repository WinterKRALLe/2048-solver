package solver

import (
	b "Solver2048/board"
	"context"
	"fmt"
)

func runMonteCarloSimulation(initialBoard b.Board, move b.Move) int {
	totalScore := 0
	numSimulations := 1000

	for i := 0; i < numSimulations; i++ {
		simulatedBoard := b.CopyBoard(initialBoard)

		simulatedBoard = b.MakeMove(simulatedBoard, move)

		for !b.IsGameOver(simulatedBoard) {
			randomMove := b.GetRandomValidMove(simulatedBoard)
			simulatedBoard = b.MakeMove(simulatedBoard, randomMove)
		}

		totalScore += calculateScore(simulatedBoard)
	}

	return totalScore / numSimulations
}

func calculateScore(boardState b.Board) int {
	totalScore := 0
	for _, row := range boardState {
		for _, value := range row {
			totalScore += value
		}
	}
	return totalScore
}

func selectBestMove(currentBoard b.Board) b.Move {
	bestMove := b.Up
	bestScore := runMonteCarloSimulation(currentBoard, bestMove)

	moves := []b.Move{b.Down, b.Left, b.Right}

	for _, move := range moves {
		if b.IsValidMove(currentBoard, move) {
			score := runMonteCarloSimulation(currentBoard, move)
			if score > bestScore {
				bestScore = score
				bestMove = move
			}
		}
	}

	return bestMove
}

func MonteCarlo(ctx context.Context) (int, int) {
	wins, losses := 0, 0

	for round := 1; round <= b.MaxRounds; round++ {
		fmt.Printf("Round %d\n", round)

		board := b.InitializeBoard()
		//b.PrintBoard(board)

		select {
		case <-ctx.Done():
			fmt.Println("\nReceived interrupt signal. Cleaning up...")
			return wins, losses
		default:
		}

		result := playMonteCarloRound(board)

		if result == "win" {
			wins++
		} else {
			losses++
		}
	}

	return wins, losses
}

func playMonteCarloRound(board b.Board) string {
	for !b.IsGameOver(board) {
		move := selectBestMove(board)
		//fmt.Println("Move:", b.MoveNames[b.Up], "Average Score:", runMonteCarloSimulation(board, b.Up))
		//fmt.Println("Move:", b.MoveNames[b.Down], "Average Score:", runMonteCarloSimulation(board, b.Down))
		//fmt.Println("Move:", b.MoveNames[b.Left], "Average Score:", runMonteCarloSimulation(board, b.Left))
		//fmt.Println("Move:", b.MoveNames[b.Right], "Average Score:", runMonteCarloSimulation(board, b.Right))
		//fmt.Println("Selected Move:", b.MoveNames[move])
		board = b.MakeMove(board, move)
		//b.PrintBoard(board)

		if b.HasMaxTile(board) {
			return "win"
		}
	}

	return "lose"
}
