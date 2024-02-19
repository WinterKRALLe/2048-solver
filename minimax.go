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
	return maxTile >= 2048
}

type AlphaBetaPruner struct {
	alpha, beta int
}

func (abp *AlphaBetaPruner) update(pruningValue int, maximizingPlayer bool) {
	if maximizingPlayer && pruningValue > abp.alpha {
		abp.alpha = pruningValue
	} else if !maximizingPlayer && pruningValue < abp.beta {
		abp.beta = pruningValue
	}
}

func minimaxAB(board Board, depth int, maximizingPlayer bool, pruner AlphaBetaPruner) int {
	if depth == 0 || isGameOver(board) {
		return evaluateBoard(board)
	}

	if maximizingPlayer {
		maxEval := math.MinInt64
		for _, move := range []Move{Up, Down, Left, Right} {
			if isValidMove(board, move) {
				newBoard := makeMove(board, move)
				eval := minimaxAB(newBoard, depth-1, false, pruner)
				maxEval = max(maxEval, eval)

				pruner.update(eval, true)
				if pruner.alpha >= pruner.beta {
					break
				}
			}
		}
		return maxEval
	} else {
		minEval := math.MaxInt64
		for _, move := range []Move{Up, Down, Left, Right} {
			if isValidMove(board, move) {
				newBoard := makeMove(board, move)
				eval := minimaxAB(newBoard, depth-1, true, pruner)
				minEval = min(minEval, eval)

				// Alpha-beta pruning:
				pruner.update(eval, false)
				if pruner.alpha >= pruner.beta {
					break // Prune remaining branches if beta <= alpha
				}
			}
		}
		return minEval
	}
}

func getBestMoveDynamicDepth(board Board) Move {
	dynamicDepth := 5
	if isCloseTo2048(board) {
		dynamicDepth = 9
	}

	bestMove := Up
	bestScore := math.MinInt64

	pruner := AlphaBetaPruner{alpha: math.MinInt64, beta: math.MaxInt64}

	for _, move := range []Move{Up, Down, Left, Right} {
		if isValidMove(board, move) {
			newBoard := makeMove(board, move)
			score := minimaxAB(newBoard, dynamicDepth, false, pruner) // Změna: nepotřebujeme snižovat hloubku o 1

			if score > bestScore {
				bestScore = score
				bestMove = move
			}
		}
	}

	return bestMove
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

func evaluateBoard(board [Size][Size]int) int {
	score := 0
	maxTile := 0
	emptyTiles := 0

	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			value := board[i][j]

			score += value

			if value > maxTile {
				maxTile = value
			}

			// Bonus za prázdná políčka s vysokou hodnotou
			if value == 0 {
				emptyTiles++
			}

			// Bonus za sloučení dlaždic
			if j < Size-1 && board[i][j] == board[i][j+1] {
				score += 4
			}
		}
	}

	// Bonus za blízkost k hodnotě 2048
	distanceTo2048 := math.Abs(float64(maxTile - 2048))
	score += int(10 / distanceTo2048)

	// Bonus za prázdná políčka s vysokou hodnotou
	score += emptyTiles * 2

	return score
}
