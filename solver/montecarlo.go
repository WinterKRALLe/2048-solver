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

		totalScore += b.CalculateScore(simulatedBoard)
	}

	return totalScore / numSimulations
}

func selectBestMove(currentBoard b.Board) (b.Move, int, int) {
	bestMove := b.Up
	bestScore := runMonteCarloSimulation(currentBoard, bestMove)
	lowestScore := runMonteCarloSimulation(currentBoard, bestMove)

	moves := []b.Move{b.Down, b.Left, b.Right}

	for _, move := range moves {
		if b.IsValidMove(currentBoard, move) {
			score := runMonteCarloSimulation(currentBoard, move)
			if score > bestScore {
				bestScore = score
				bestMove = move
			} else if score < lowestScore {
				lowestScore = score
			}
		}
	}

	return bestMove, bestScore, lowestScore
}

func MonteCarlo(ctx context.Context) (int, int, map[string]int, int, int, map[int]int) {
	wins, losses := 0, 0
	totalMoveCounts := make(map[string]int)
	newBestScore, newLowestScore := 0, 0
	maxTileCounts := make(map[int]int)

	for round := 1; round <= b.MaxRounds; round++ {
		fmt.Printf("Round %d\n", round)

		board := b.InitializeBoard()
		//b.PrintBoard(board)

		select {
		case <-ctx.Done():
			fmt.Println("\nReceived interrupt signal. Cleaning up...")
			return wins, losses, totalMoveCounts, newBestScore, newLowestScore, maxTileCounts
		default:
		}

		result, moveCount, bestScore, lowestScore, maxTile := playMonteCarloRound(board)

		for move, count := range moveCount {
			totalMoveCounts[move] += count
		}

		if bestScore > newBestScore {
			newBestScore = bestScore
		}
		if lowestScore > newLowestScore {
			newLowestScore = lowestScore
		}

		if result == "win" {
			wins++
			maxTileCounts[maxTile]++

		} else {
			losses++
			maxTileCounts[maxTile]++

		}

	}

	return wins, losses, totalMoveCounts, newBestScore, newLowestScore, maxTileCounts
}

func playMonteCarloRound(board b.Board) (string, map[string]int, int, int, int) {
	bestScore, lowestScore, maxTile := 0, 0, 0
	moveCounts := make(map[string]int)

	for !b.IsGameOver(board) {
		moveRes, bestScoreRes, lowestScoreRes := selectBestMove(board)

		bestScore = bestScore + bestScoreRes
		lowestScore = lowestScore + lowestScoreRes
		moveCounts[b.MoveNames[moveRes]]++

		for _, row := range board {
			for _, value := range row {
				if value > maxTile {
					maxTile = value
				}
			}
		}

		board = b.MakeMove(board, moveRes)

		if b.HasMaxTile(board) {
			return "win", moveCounts, bestScore, lowestScore, 2048
		}
	}
	return "lose", moveCounts, bestScore, lowestScore, maxTile

}
