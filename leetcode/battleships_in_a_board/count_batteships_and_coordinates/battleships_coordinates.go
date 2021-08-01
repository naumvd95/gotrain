package main

import "fmt"

func main() {
	board := [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	fmt.Printf("Here is board: %v\n", board)

	ships := countBattleships(board)

	fmt.Printf("Here is ships: %v", ships)
}

func countBattleships(board [][]byte) map[int][][]int {
	// key: number of ship, value: [][] coordinates x,y of each ship cell
	ships := make(map[int][][]int)
	currentShipNumber := 0

	for row := 0; row < len(board); row++ {
		for cell := 0; cell < len(board[0]); cell++ {
			if board[row][cell] == 'X' {
				currentShipNumber++
				dfs(&board, row, cell, &ships, currentShipNumber)
			}
		}
	}

	return ships
}

func dfs(board *[][]byte, row, cell int, ships *map[int][][]int, currentShipNumber int) {
	lenRow := len(*board) - 1
	lenCell := len((*board)[0]) - 1

	if row < 0 || row > lenRow || cell < 0 || cell > lenCell || (*board)[row][cell] != 'X' {
		// if coordinates are not exist or there is no ship
		return
	}

	if (*board)[row][cell] == 'X' {
		// adding coordinates for our new ship
		(*ships)[currentShipNumber] = append((*ships)[currentShipNumber], []int{row, cell})
		// cleaning up board to prevent doublecheck
		(*board)[row][cell] = '.'
	}
	dfs(board, row+1, cell, ships, currentShipNumber)
	dfs(board, row, cell+1, ships, currentShipNumber)
	dfs(board, row-1, cell, ships, currentShipNumber)
	dfs(board, row, cell-1, ships, currentShipNumber)
}
