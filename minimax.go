package main

import "math"

func hasMaxTile(board Board) bool {
	maxTile := 0
	for _, row := range board {
		for _, value := range row {
			if value > maxTile {
				maxTile = value
			}
		}
	}
	return maxTile >= 1024
}

func evaluateWithAlphaBeta(node Board, depth int, alpha, beta int, maximizingPlayer bool) int {
	if depth == 0 || isGameOver(node) {
		return evaluateBoard(node)
	}

	if maximizingPlayer {
		maxEval := math.MinInt64
		for _, move := range []Move{Up, Down, Left, Right} {
			if isValidMove(node, move) {
				newBoard := makeMove(node, move)
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
		for _, move := range []Move{Up, Down, Left, Right} {
			if isValidMove(node, move) {
				newBoard := makeMove(node, move)
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

func isCloseTo2048(board Board) bool {
	for _, row := range board {
		for _, value := range row {
			if value >= 512 {
				return true
			}
		}
	}
	return false
}

func getBestMoveDynamicDepth(board Board) Move {
	dynamicDepth := 6
	if isCloseTo2048(board) {
		dynamicDepth = 8
	}

	bestMove := Up
	bestScore := math.MinInt64

	for _, move := range []Move{Up, Down, Left, Right} {
		if isValidMove(board, move) {
			newBoard := makeMove(board, move)
			score := evaluateWithAlphaBeta(newBoard, dynamicDepth, math.MinInt64, math.MaxInt64, true)

			if score > bestScore {
				bestScore = score
				bestMove = move
			}
		}
	}

	return bestMove
}

func evaluateBoard(board [Size][Size]int) int {
	sumOfBoard := 0
	nonEmptyTiles := 0
	mergedTilesSum := 0

	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			value := board[i][j]

			// Sum of all numbers on the board
			sumOfBoard += value

			// Sum of merged tiles
			if j < Size-1 && board[i][j] == board[i][j+1] {
				mergedTilesSum += 2 * value
			}

			// Number of non-empty tiles
			if board[i][j] != 0 {
				nonEmptyTiles++
			}

		}
	}

	utility := (mergedTilesSum + sumOfBoard) / nonEmptyTiles

	return utility
}
