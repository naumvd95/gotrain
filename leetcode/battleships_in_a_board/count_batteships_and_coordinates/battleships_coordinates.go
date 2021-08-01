package main

import "fmt"

func getShips(board *Board) map[int][][]int {
	// algo: loop over matrix, if 'X' notified = start DFS traversing up/down/left/right to check is that ship has more cells
	// do not forget to mark cells as visited in DFS to prevent doublecheck (set cell to '.')
	ships := make(map[int][][]int)
	shipCounter := 0

	for row := 0; row < len(*board); row++ {
		for cell := 0; cell < len((*board)[0]); cell++ {
			if (*board)[row][cell] != 'X' {
				//gating style
				continue
			}

			// otherwise we found a ship
			shipCounter++
			// running ship sniffer!
			board.dfs(row, cell, shipCounter, &ships)
		}
	}
	return ships
}

type Board [][]byte

func (b *Board) dfs(row, cell, shipCounter int, ships *map[int][][]int) {
	if row < 0 || row > len(*b)-1 || cell < 0 || cell > len((*b)[0])-1 || (*b)[row][cell] != 'X' {
		// if coordinates out of range or there is no cell
		return
	}

	// otherwise we found ship tail, save coordinates
	(*ships)[shipCounter] = append((*ships)[shipCounter], []int{row, cell})
	// spoil current cell to prevent doublecheck
	(*b)[row][cell] = '.'

	// sniff deeper
	b.dfs(row-1, cell, shipCounter, ships)
	b.dfs(row+1, cell, shipCounter, ships)
	b.dfs(row, cell-1, shipCounter, ships)
	b.dfs(row, cell+1, shipCounter, ships)
}

func main() {
	board := Board{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	expectedShips := 2

	fmt.Printf("Current board: %v\n", board)
	ships := getShips(&board)
	fmt.Printf("ships amount: %v , expected: %v\n", len(ships), expectedShips)
	fmt.Printf("ships coordinates: %v\n", ships)
}
