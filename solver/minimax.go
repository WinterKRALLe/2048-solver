package solver

import (
	b "Solver2048/board"
	"context"
	"fmt"
	"math"
)

func evaluateWithAlphaBeta(node b.Board, depth int, alpha, beta int, maximizingPlayer bool) int {
	if depth == 0 || b.IsGameOver(node) {
		return evaluateBoard(node)
	}

	if maximizingPlayer {
		maxEval := math.MinInt64
		for _, move := range []b.Move{b.Up, b.Down, b.Left, b.Right} {
			if b.IsValidMove(node, move) {
				newBoard := b.MakeMove(node, move)
				eval := evaluateWithAlphaBeta(newBoard, depth-1, alpha, beta, true)
				maxEval = max(maxEval, eval)

				alpha = max(alpha, eval)
				if beta <= alpha {
					break
				}
			}
		}
		return maxEval
	} else {
		minEval := math.MaxInt64
		for _, move := range []b.Move{b.Up, b.Down, b.Left, b.Right} {
			if b.IsValidMove(node, move) {
				newBoard := b.MakeMove(node, move)
				eval := evaluateWithAlphaBeta(newBoard, depth-1, alpha, beta, false)
				minEval = min(minEval, eval)

				beta = min(beta, eval)
				if beta <= alpha {
					break
				}
			}
		}
		return minEval
	}
}

func isCloseTo2048(board b.Board) bool {
	for _, row := range board {
		for _, value := range row {
			if value >= 512 {
				return true
			}
		}
	}
	return false
}

func GetBestMoveDynamicDepth(board b.Board) b.Move {
	dynamicDepth := 6
	if isCloseTo2048(board) {
		dynamicDepth = 8
	}

	bestMove := b.Up
	bestScore := math.MinInt64

	for _, move := range []b.Move{b.Up, b.Down, b.Left, b.Right} {
		if b.IsValidMove(board, move) {
			newBoard := b.MakeMove(board, move)
			score := evaluateWithAlphaBeta(newBoard, dynamicDepth, math.MinInt64, math.MaxInt64, true)

			if score > bestScore {
				bestScore = score
				bestMove = move
			}
		}
	}

	return bestMove
}

func evaluateBoard(board b.Board) int {
	sumOfBoard := 0
	nonEmptyTiles := 0
	mergedTilesSum := 0

	newBoard := b.CopyBoard(board)

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			value := newBoard[i][j]

			// Sum of all numbers on the board
			sumOfBoard += value

			// Sum of merged tiles
			if j < b.Size-1 && newBoard[i][j] == newBoard[i][j+1] {
				mergedTilesSum += 2 * value
			}

			// Number of non-empty tiles
			if newBoard[i][j] != 0 {
				nonEmptyTiles++
			}

		}
	}

	utility := (mergedTilesSum + sumOfBoard) / nonEmptyTiles

	return utility
}

func Minimax(ctx context.Context) (int, int) {
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

		result := playRound(board)

		if result == "win" {
			wins++
		} else {
			losses++
		}
	}

	return wins, losses
}

func playRound(board b.Board) string {
	for !b.IsGameOver(board) {
		move := GetBestMoveDynamicDepth(board)

		fmt.Println("Move:", b.MoveNames[move])
		board = b.MakeMove(board, move)
		b.PrintBoard(board)
	}

	if b.HasMaxTile(board) {
		return "win"
	}

	return "lose"
}
