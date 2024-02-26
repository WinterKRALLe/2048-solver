package solver

import (
	b "Solver2048/board"
	"context"
	"fmt"
)

func PureRandom(ctx context.Context) (int, int) {
	wins, losses := 0, 0

	for round := 1; round <= b.MaxRounds; round++ {
		fmt.Printf("Round %d\n", round)

		board := b.InitializeBoard()
		b.PrintBoard(board)

		select {
		case <-ctx.Done():
			fmt.Println("\nReceived interrupt signal. Cleaning up...")
			return wins, losses
		default:
		}

		result := playPureRandomRound(board)

		if result == "win" {
			wins++
		} else {
			losses++
		}
	}

	return wins, losses
}

func playPureRandomRound(board b.Board) string {
	for !b.IsGameOver(board) {
		move := b.GetRandomValidMove(board)
		fmt.Println("Selected Move:", b.MoveNames[move])
		board = b.MakeMove(board, move)

		b.PrintBoard(board)

		if b.HasMaxTile(board) {
			return "win"
		}
	}

	return "lose"
}
