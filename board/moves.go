package board

func IsValidMove(board Board, move Move) bool {
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
			for row := 0; row < Size-1; row++ {
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
			for col := 0; col < Size-1; col++ {
				if board[row][col] != 0 && (board[row][col+1] == 0 || board[row][col+1] == board[row][col]) {
					return true
				}
			}
		}
	}

	return false
}

func MakeMove(board Board, move Move) Board {
	newBoard := CopyBoard(board)

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

func IsGameOver(board Board) bool {
	return !IsValidMove(board, Up) && !IsValidMove(board, Down) && !IsValidMove(board, Left) && !IsValidMove(board, Right)
}
