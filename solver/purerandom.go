package solver

import (
	b "Solver2048/board"
	"context"
	"fmt"
)

func PureRandom(ctx context.Context) (int, int, map[string]int, int, int, map[int]int) {
	wins, losses := 0, 0
	totalMoveCounts := make(map[string]int)
	newBestScore, newLowestScore := 0, 0
	maxTileCounts := make(map[int]int)

	for round := 1; round <= b.MaxRounds; round++ {
		fmt.Printf("Round %d\n", round)

		board := b.InitializeBoard()

		select {
		case <-ctx.Done():
			fmt.Println("\nReceived interrupt signal. Cleaning up...")
			return wins, losses, totalMoveCounts, newBestScore, newLowestScore, maxTileCounts
		default:
		}

		result, moveCount, score, maxTile := playPureRandomRound(board)

		for move, count := range moveCount {
			totalMoveCounts[move] += count
		}

		if score > newBestScore {
			newBestScore = score
		}

		if score < newLowestScore || score < newBestScore {
			newLowestScore = score
		}

		maxTileCounts[maxTile]++

		if result == "win" {
			wins++
		} else {
			losses++
		}
	}

	return wins, losses, totalMoveCounts, newBestScore, newLowestScore, maxTileCounts
}

func playPureRandomRound(board b.Board) (string, map[string]int, int, int) {
	score, maxTile := 0, 0
	moveCounts := make(map[string]int)

	for !b.IsGameOver(board) {
		move := b.GetRandomValidMove(board)

		moveCounts[b.MoveNames[move]]++

		for _, row := range board {
			for _, value := range row {
				if value > maxTile {
					maxTile = value
				}
			}
		}

		score = b.CalculateScore(board)

		board = b.MakeMove(board, move)

		if b.HasMaxTile(board) {
			return "win", moveCounts, score, 2048
		}
	}

	return "lose", moveCounts, score, maxTile
}
