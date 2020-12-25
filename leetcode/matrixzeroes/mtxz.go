package main

import "fmt"

/*
Given an m x n matrix. If an element is 0, set its entire row and column to 0. Do it in-place.

Follow up:

A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
// threeSumFast pre-sort slice to apply 2-pointer 2sum solution and avoid additional loop and hashSet

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Input:
[[-4,-2147483648,6,-7,0],[-8,6,-8,-6,0],[2147483647,2,-9,-6,-10]]
Output:
[[0,0,0,0,0],[0,0,0,0,0],[2147483647,2,-9,-6,0]]

----------------------------------------------------------
Verbal solution w/ O(1) space, coz ideally we may change all in-place:
1. if only 1 row, just check it
2. if more than 1 row in matrix, we can use 1st row as mind marker storage to remeber columns/rows which needs zeroes
3. iterating over each row(subslice) and if its equal zero, set mind marker in shape of the crossroad in matrix

1 1 1       1 _ 1      1 0 1
1 0 1 ->>>  _ 0 1 ->>> 0 0 1
1 1 1       1 1 1      1 1 1

Mind marker means: if 1st column/row == 0 ->> whole column/row should be zeroed
NOTE: matrix[0][0] value cannot be written as mind marker, it breaks column marker, remember: MIND MARKER NEVER MATRIX ENTRYPOINT[0][0]

4. Iterate 2nd time through matrix and fill in zeroes according to markers

COMPLEXITY:
space: O(1) , we dont use additional space
time: O(M*N), because we are iterating over matrix
*/

func setZeroes(matrix [][]int) {
	// len of slice itself == num in rows amount
	numInRowAmount := len(matrix[0])

	/* we will try to use matrix[0] as column marker storage
	   Its kinda mind markers for notice, do we need to set all columns to zeros
	   But firstly we need to check initial state of matrix[0]
	*/
	matrixZeroEntrypoint := false
	// need to save such value before we start to overwrite marker storage
	for _, num := range matrix[0] {
		if num == 0 {
			matrixZeroEntrypoint = true // notice that matrix[0] should be overwritten w/ zeros after marker operations
			break
		}
	}

	// if its only 1 row, we cant use mind markers, coz we need at least 2 rows
	if len(matrix) != 1 {
		for rowPos, row := range matrix {
			for colNumPos, colNum := range row {
				if colNum == 0 {
					// set mind markers, they met condition: if 1st row/col == 0 -> all row/col should be zeros

					matrix[0][colNumPos] = 0 // column marker
					if rowPos != 0 {
						// row marker can be set only if its not initial matrix entrypoint
						matrix[rowPos][0] = 0
					}
				}
			}
			/* Matrix w/ markers
			matrix := [][]int{
				[]int{1, 0, 1},
				[]int{0, 0, 1},
				[]int{1, 1, 1},
			*/
		}

		// using matrix[0] as markers storage
		for _, row := range matrix[1:] {
			if row[0] == 0 { // row marker detected
				for i := 1; i < numInRowAmount; i++ {
					row[i] = 0
				}
				continue // dont need to check column marker
			}

			for colNumPos := range row {
				if matrix[0][colNumPos] == 0 { // column marker detected
					row[colNumPos] = 0
				}
			}
		}
	}

	// remebered that initially there was zero-in-a-row, so overwrite slice
	if matrixZeroEntrypoint {
		for i := 0; i < numInRowAmount; i++ {
			matrix[0][i] = 0
		}
	}

}

func main() {
	matrix := [][]int{
		[]int{1, 1, 1}, // answer []int{1, 0, 1}
		[]int{1, 0, 1}, //        []int{0, 0, 0}
		[]int{1, 1, 1}, //        []int{1, 0, 1}
	}
	fmt.Printf("original matrix: %v\n", matrix)

	setZeroes(matrix)
	fmt.Printf("zeroed: %v\n", matrix)

}
