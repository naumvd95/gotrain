package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// time: O(n) n=cells | space: O(1)
func countBattleships(board [][]byte) int {
	// algo have we met such ship earlier??
	var res int

	for row := 0; row < len(board); row++ {
		for cell := 0; cell < len(board[0]); cell++ {
			if board[row][cell] != 'X' {
				// gate style, skip empty cells
				continue
			}
			// otherwise board[row][cell] == 'X'
			if row == 0 && cell == 0 {
				// we are at the beginning - its defenitely new ship
				res++
			}

			if row == 0 && cell > 0 && board[row][cell-1] != 'X' {
				// we are at the first row, check if prev cell is not a ship
				res++
			}

			if cell == 0 && row > 0 && board[row-1][cell] != 'X' {
				// we are at the X row, check if cell from prev row is not a ship
				res++
			}

			if cell > 0 && row > 0 && board[row-1][cell] != 'X' && board[row][cell-1] != 'X' {
				// we are completely in the middle, check if cell from prev row is not a ship
				// and check if prev cell is not a ship
				res++
			}

		}
	}

	return res
}
