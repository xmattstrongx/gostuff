package sudoku2

import "fmt"

type board struct {
	rows     []row
	columns  []column
	grid3x3s []grid3x3
}

type row struct {
	values map[string]int
}

type column struct {
	values map[string]int
}

type grid3x3 struct {
	values map[string]int
}

func newBoard() *board {
	rows := make([]row, 9)
	for i := range rows {
		rows[i].values = make(map[string]int)
	}
	columns := make([]column, 9)
	for i := range columns {
		columns[i].values = make(map[string]int)
	}
	grid3x3s := make([]grid3x3, 9)
	for i := range grid3x3s {
		grid3x3s[i].values = make(map[string]int)
	}
	return &board{
		rows,
		columns,
		grid3x3s,
	}
}

func sudoku2(grid [][]string) bool {
	return validate(grid, newBoard())
}

func validate(grid [][]string, b *board) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] != "." {
				if _, exists := b.columns[j].values[grid[i][j]]; exists {
					fmt.Printf("(%d,%d) = %s. THIS ALREADY EXISTS IN COLUMN %d. GAME OVER \n", i, j, grid[i][j], j)
					return false
				}
				b.columns[j].values[grid[i][j]] = 1

				if _, exists := b.rows[i].values[grid[i][j]]; exists {
					fmt.Printf("(%d,%d) = %s. THIS ALREADY EXISTS IN ROW %d. GAME OVER \n", i, j, grid[i][j], i)
					return false
				}
				b.rows[i].values[grid[i][j]] = 1

				currentGrid3x3 := calculateGrid(i, j)
				if _, exists := b.grid3x3s[currentGrid3x3].values[grid[i][j]]; exists {
					fmt.Printf("(%d,%d) = %s. THIS ALREADY EXISTS IN GRID %d. GAME OVER \n", i, j, grid[i][j], currentGrid3x3)
					return false
				}
				b.grid3x3s[currentGrid3x3].values[grid[i][j]] = 1
			}
		}
	}
	return true
}

func calculateGrid(i, j int) int {
	var grid int
	if i >= 0 && i <= 2 && j >= 0 && j <= 2 {
		grid = 0
	} else if i >= 0 && i <= 2 && j >= 3 && j <= 5 {
		grid = 1
	} else if i >= 0 && i <= 2 && j >= 6 && j <= 8 {
		grid = 2
	} else if i >= 3 && i <= 5 && j >= 0 && j <= 2 {
		grid = 3
	} else if i >= 3 && i <= 5 && j >= 3 && j <= 5 {
		grid = 4
	} else if i >= 3 && i <= 5 && j >= 6 && j <= 8 {
		grid = 5
	} else if i >= 6 && i <= 8 && j >= 0 && j <= 2 {
		grid = 6
	} else if i >= 6 && i <= 8 && j >= 3 && j <= 5 {
		grid = 7
	} else if i >= 6 && i <= 8 && j >= 6 && j <= 8 {
		grid = 8
	}
	return grid
}

// 2 dimensional array reminder

// grid = [[(0,0), (0,1), (0,2), (0,3), (0,4), (0,5), (0,6), (0,7), (0,8)],
//         [(1,0), (1,1), (1,2), (1,3), (1,4), (1,5), (1,6), (1,7), (1,8)],
//         [(2,0), (2,1), (2,2), (2,3), (2,4), (2,5), (2,6), (2,7), (2,8)],
//         [(3,0), (3,1), (3,2), (3,3), (3,4), (3,5), (3,6), (3,7), (3,8)],
//         [(4,0), (4,1), (4,2), (4,3), (4,4), (4,5), (4,6), (4,7), (4,8)],
//         [(5,0), (5,1), (5,2), (5,3), (5,4), (5,5), (5,6), (5,7), (5,8)],
//         [(6,0), (6,1), (6,2), (6,3), (6,4), (6,5), (6,6), (6,7), (6,8)],
//         [(7,0), (7,1), (7,2), (7,3), (7,4), (7,5), (7,6), (7,7), (7,8)],
//         [(8,0), (8,1), (8,2), (8,3), (8,4), (8,5), (8,6), (8,7), (8,8)]

// Sudoku is a number-placement puzzle. The objective is to fill a 9 × 9 grid with numbers in such a way that each column, each row, and each of the nine 3 × 3 sub-grids that compose the grid all contain all of the numbers from 1 to 9 one time.

// Implement an algorithm that will check whether the given grid of numbers represents a valid Sudoku puzzle according to the layout rules described above. Note that the puzzle represented by grid does not have to be solvable.

// Example

// For

// grid = [['.', '.', '.', '1', '4', '.', '.', '2', '.'],
//         ['.', '.', '6', '.', '.', '.', '.', '.', '.'],
//         ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
//         ['.', '.', '1', '.', '.', '.', '.', '.', '.'],
//         ['.', '6', '7', '.', '.', '.', '.', '.', '9'],
//         ['.', '.', '.', '.', '.', '.', '8', '1', '.'],
//         ['.', '3', '.', '.', '.', '.', '.', '.', '6'],
//         ['.', '.', '.', '.', '.', '7', '.', '.', '.'],
//         ['.', '.', '.', '5', '.', '.', '.', '7', '.']]
// the output should be
// sudoku2(grid) = true;

// For

// grid = [['.', '.', '.', '.', '2', '.', '.', '9', '.'],
//         ['.', '.', '.', '.', '6', '.', '.', '.', '.'],
//         ['7', '1', '.', '.', '7', '5', '.', '.', '.'],
//         ['.', '7', '.', '.', '.', '.', '.', '.', '.'],
//         ['.', '.', '.', '.', '8', '3', '.', '.', '.'],
//         ['.', '.', '8', '.', '.', '7', '.', '6', '.'],
//         ['.', '.', '.', '.', '.', '2', '.', '.', '.'],
//         ['.', '1', '.', '2', '.', '.', '.', '.', '.'],
//         ['.', '2', '.', '.', '3', '.', '.', '.', '.']]
// the output should be
// sudoku2(grid) = false.

// The given grid is not correct because there are two 1s in the second column. Each column, each row, and each 3 × 3 subgrid can only contain the numbers 1 through 9 one time.

// Input/Output

// [time limit] 4000ms (go)
// [input] array.array.char grid

// A 9 × 9 array of characters, in which each character is either a digit from '1' to '9' or a period '.'.

// [output] boolean

// Return true if grid represents a valid Sudoku puzzle, otherwise return false.
