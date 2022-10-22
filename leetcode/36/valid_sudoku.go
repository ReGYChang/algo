package main

func isValidSudoku(board [][]byte) bool {
	var freq [128]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if freq[board[i][j]] == 0 || board[i][j] == '.' {
				freq[board[i][j]]++
			} else {
				return false
			}
		}
		freq = [128]int{}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if freq[board[j][i]] == 0 || board[j][i] == '.' {
				freq[board[j][i]]++
			} else {
				return false
			}
		}
		freq = [128]int{}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if freq[board[i+k][j+l]] == 0 || board[i+k][j+l] == '.' {
						freq[board[i+k][j+l]]++
					} else {
						return false
					}
				}
			}
			freq = [128]int{}
		}
	}

	return true
}
