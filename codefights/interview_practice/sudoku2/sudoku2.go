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
	return &board{
		make([]row, 0),
		make([]column, 0),
		make([]grid3x3, 0),
	}
}

func newRow() *row {
	valueMap := make(map[string]int)
	return &row{valueMap}
}

func newColumn() *column {
	valueMap := make(map[string]int)
	return &column{valueMap}
}

func newGrid3x3() *grid3x3 {
	valueMap := make(map[string]int)
	return &grid3x3{valueMap}
}

func sudoku2(grid [][]string) bool {
	// cells := getCells(grid)
	// var rows []row
	// var columns []column
	// var grid3x3s []grid3x3

	board := newBoard()

	rowMap := make(map[int]map[string]int)
	columnMap := make(map[int]map[string]int)
	gridMap := make(map[int]map[string]int)

	for _, currentRow := range grid {
		for j, val := range currentRow {
			if val != "." {
				valMap := map[string]int{
					val: 1,
				}
				if _, exists := board.columns[j].values[val]; exists {
					return false
				}
				board.columns = append(board.columns, board.columns[j])
				fmt.Println(board.columns)
				// board.columns[j].values = valMap
				// columnMap[j][val]++

				// currentGrid3x3 := calculateGrid(i, j)
				// gridMap[currentGrid3x3][val]++
				// rows = append(rows, row{values: })
				// columns = append(columns, column{j, val})
				// cells = append(cells, cell{
				// 	row:    i,
				// 	val: j,
				// 	grid:   calculateGrid(i, j),
				// 	value:  val,

				// fmt.Printf("cells: %v", cells)
			}
			// fmt.Printf("val %d: %s\n", j, val)
		}
	}
	fmt.Printf("Entire DSC: %v\n", rowMap)
	fmt.Printf("Entire DSC: %v\n", columnMap)
	fmt.Printf("Entire DSC: %v\n", gridMap)

	return false
}

func calculateGrid(i, j int) int {
	var grid int
	if i >= 0 && i <= 2 && j >= 0 && j <= 2 {
		grid = 1
	} else if i >= 0 && i <= 2 && j >= 3 && j <= 5 {
		grid = 2
	} else if i >= 0 && i <= 2 && j >= 6 && j <= 8 {
		grid = 3
	} else if i >= 3 && i <= 5 && j >= 0 && j <= 2 {
		grid = 4
	} else if i >= 3 && i <= 5 && j >= 3 && j <= 5 {
		grid = 5
	} else if i >= 3 && i <= 5 && j >= 6 && j <= 8 {
		grid = 6
	} else if i >= 6 && i <= 8 && j >= 0 && j <= 2 {
		grid = 7
	} else if i >= 6 && i <= 8 && j >= 3 && j <= 5 {
		grid = 8
	} else if i >= 6 && i <= 8 && j >= 6 && j <= 8 {
		grid = 9
	}
	return grid
}

// func sudoku2(grid [][]string) bool {
// 	for i := 0; i < 9; i++ {
// 		existing := make(map[string]int)
// 		// existing := make(map[string]map[string]int)
// 		for j := 0; j < 9; j++ {
// 			//check all
// 			if grid[i][j] != "." {
// 				if val, exists := existing[grid[i][j]]; exists && val > 1 {
// 					return false
// 				}
// 				existing[grid[i][j]]++
// 				fmt.Printf("(%d,%d) = %s ... %v\n", i, j, grid[i][j], existing)
// 			} else {
// 				fmt.Printf("(%d,%d) = %s\n", i, j, grid[i][j])
// 			}
// 		}
// 		fmt.Print("NEXT\n")
// 	}

// 	for j := 0; j < 9; j++ {
// 		existing := make(map[string]int)
// 		for i := 0; i < 9; i++ {
// 			//check all
// 			if grid[i][j] != "." {
// 				if val, exists := existing[grid[i][j]]; exists && val > 1 {
// 					return false
// 				}
// 				existing[grid[i][j]]++
// 				fmt.Printf("(%d,%d) = %s ... %v\n", i, j, grid[i][j], existing)
// 			} else {
// 				fmt.Printf("(%d,%d) = %s\n", i, j, grid[i][j])
// 			}
// 		}
// 		fmt.Print("NEXT\n")
// 	}

// 	return true
// }

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
