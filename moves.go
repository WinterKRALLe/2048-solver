package main

import "math/rand"

type Move int

const (
	Up Move = iota
	Down
	Left
	Right
)

var moveNames = map[Move]string{
	Up:    "Up",
	Down:  "Down",
	Left:  "Left",
	Right: "Right",
}

func isValidMove(board Board, move Move) bool {
	switch move {
	case Up:
		for col := 0; col < Size; col++ {
			for row := 1; row < Size; row++ {
				if board[row][col] != 0 && (board[row-1][col] == 0 || board[row-1][col] == board[row][col]) {
					return true
				}
			}
		}
	case Down:
		for col := 0; col < Size; col++ {
			for row := Size - 2; row >= 0; row-- {
				if board[row][col] != 0 && (board[row+1][col] == 0 || board[row+1][col] == board[row][col]) {
					return true
				}
			}
		}
	case Left:
		for row := 0; row < Size; row++ {
			for col := 1; col < Size; col++ {
				if board[row][col] != 0 && (board[row][col-1] == 0 || board[row][col-1] == board[row][col]) {
					return true
				}
			}
		}
	case Right:
		for row := 0; row < Size; row++ {
			for col := Size - 2; col >= 0; col-- {
				if board[row][col] != 0 && (board[row][col+1] == 0 || board[row][col+1] == board[row][col]) {
					return true
				}
			}
		}
	}

	return false
}

func makeMove(board Board, move Move) Board {
	newBoard := copyBoard(board)

	switch move {
	case Up:
		for col := 0; col < Size; col++ {
			merged := make(map[int]bool)
			for row := 1; row < Size; row++ {
				if newBoard[row][col] != 0 {
					currentRow := row
					for currentRow > 0 && (newBoard[currentRow-1][col] == 0 || (newBoard[currentRow-1][col] == newBoard[currentRow][col] && !merged[currentRow-1])) {
						if newBoard[currentRow-1][col] == 0 {
							newBoard[currentRow-1][col] = newBoard[currentRow][col]
							newBoard[currentRow][col] = 0
						} else {
							if merged[currentRow-1] {
								break // Předčasný odchod, sloučení již proběhlo
							}
							newBoard[currentRow-1][col] *= 2
							newBoard[currentRow][col] = 0
							merged[currentRow-1] = true
							break
						}
						currentRow--
					}
				}
			}
		}
	case Down:
		for col := 0; col < Size; col++ {
			merged := make(map[int]bool)
			for row := Size - 2; row >= 0; row-- {
				if newBoard[row][col] != 0 {
					currentRow := row
					for currentRow < Size-1 && (newBoard[currentRow+1][col] == 0 || (newBoard[currentRow+1][col] == newBoard[currentRow][col] && !merged[currentRow+1])) {
						if newBoard[currentRow+1][col] == 0 {
							newBoard[currentRow+1][col] = newBoard[currentRow][col]
							newBoard[currentRow][col] = 0
						} else {
							if merged[currentRow+1] {
								break // Předčasný odchod, sloučení již proběhlo
							}
							newBoard[currentRow+1][col] *= 2
							newBoard[currentRow][col] = 0
							merged[currentRow+1] = true
							break
						}
						currentRow++
					}
				}
			}
		}
	case Left:
		for row := 0; row < Size; row++ {
			merged := make(map[int]bool)
			for col := 1; col < Size; col++ {
				if newBoard[row][col] != 0 {
					currentCol := col
					for currentCol > 0 && (newBoard[row][currentCol-1] == 0 || (newBoard[row][currentCol-1] == newBoard[row][currentCol] && !merged[currentCol-1])) {
						if newBoard[row][currentCol-1] == 0 {
							newBoard[row][currentCol-1] = newBoard[row][currentCol]
							newBoard[row][currentCol] = 0
						} else {
							if merged[currentCol-1] {
								break // Předčasný odchod, sloučení již proběhlo
							}
							newBoard[row][currentCol-1] *= 2
							newBoard[row][currentCol] = 0
							merged[currentCol-1] = true
							break
						}
						currentCol--
					}
				}
			}
		}
	case Right:
		for row := 0; row < Size; row++ {
			merged := make(map[int]bool)
			for col := Size - 2; col >= 0; col-- {
				if newBoard[row][col] != 0 {
					currentCol := col
					for currentCol < Size-1 && (newBoard[row][currentCol+1] == 0 || (newBoard[row][currentCol+1] == newBoard[row][currentCol] && !merged[currentCol+1])) {
						if newBoard[row][currentCol+1] == 0 {
							newBoard[row][currentCol+1] = newBoard[row][currentCol]
							newBoard[row][currentCol] = 0
						} else {
							if merged[currentCol+1] {
								break // Předčasný odchod, sloučení již proběhlo
							}
							newBoard[row][currentCol+1] *= 2
							newBoard[row][currentCol] = 0
							merged[currentCol+1] = true
							break
						}
						currentCol++
					}
				}
			}
		}
	}

	placeRandomTile(&newBoard)
	return newBoard
}

func copyBoard(board Board) Board {
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

func isGameOver(board Board) bool {
	return !isValidMove(board, Up) && !isValidMove(board, Down) && !isValidMove(board, Left) && !isValidMove(board, Right)
}
