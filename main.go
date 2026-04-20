package main

// Import required packages:
// fmt → for printing output
// os  → for accessing command-line arguments
import (
	"fmt"
	"os"
)

// ==============================
// FUNCTION 1: CHECK VALIDITY
// ==============================

// This function checks if placing a number in a specific position is allowed
// It ensures the Sudoku rules are not broken
func isValid(grid [][]int, row, col, num int) bool {
	// Loop through all columns in the same row
	// Also loop through all rows in the same column
	for i := 0; i < 9; i++ {

		// Check if the number already exists in the same row
		if grid[row][i] == num {
			return false // Not valid
		}

		// Check if the number already exists in the same column
		if grid[i][col] == num {
			return false // Not valid
		}
	}

	// Calculate the starting position of the 3x3 box
	// Example: row=5 → (5/3)*3 = 3 → box starts at row 3
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	// Loop through the 3x3 box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Check if number exists in the box
			if grid[startRow+i][startCol+j] == num {
				return false // Not valid
			}
		}
	}

	// If no rule is broken, the number is valid
	return true
}

// ==============================
// FUNCTION 2: VALIDATE INITIAL GRID
// ==============================

// This checks if the initial Sudoku input is already invalid
// (before solving begins)
func isGridValid(grid [][]int) bool {
	// Loop through every cell in the grid
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Only check cells that already have numbers
			if grid[i][j] != 0 {

				// Store the number
				num := grid[i][j]

				// Temporarily remove it
				// (to avoid comparing the number with itself)
				grid[i][j] = 0

				// Check if placing it again is valid
				if !isValid(grid, i, j, num) {
					return false // Invalid Sudoku
				}

				// Restore the number
				grid[i][j] = num
			}
		}
	}

	// If all cells are valid
	return true
}

// ==============================
// FUNCTION 3: SOLVE SUDOKU
// ==============================

// This function uses BACKTRACKING to solve the Sudoku
func solve(grid [][]int) bool {
	// Loop through the entire grid
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Find an empty cell (represented by 0)
			if grid[i][j] == 0 {

				// Try numbers from 1 to 9
				for num := 1; num <= 9; num++ {
					// Check if the number is valid
					if isValid(grid, i, j, num) {

						// Place the number
						grid[i][j] = num

						// Recursively try to solve the rest of the grid
						if solve(grid) {
							return true // Solved
						}

						// If it fails, BACKTRACK (undo the move)
						grid[i][j] = 0
					}
				}

				// If no number works, return false (dead end)
				return false
			}
		}
	}

	// If no empty cell is found → puzzle is solved
	return true
}

// ==============================
// MAIN FUNCTION (PROGRAM ENTRY)
// ==============================

func main() {
	// Check if the number of arguments is exactly 10
	// (1 program name + 9 rows of Sudoku)
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	// Create a 2D slice (9x9 grid)
	grid := make([][]int, 9)

	// Loop through each row input
	for i := 0; i < 9; i++ {

		// Check if each row has exactly 9 characters
		if len(os.Args[i+1]) != 9 {
			fmt.Println("Error")
			return
		}

		// Initialize columns for this row
		grid[i] = make([]int, 9)

		// Loop through each character in the row
		for j, ch := range os.Args[i+1] {
			// If character is '.', it's an empty cell
			if ch == '.' {
				grid[i][j] = 0

				// If character is a number between '1' and '9'
			} else if ch >= '1' && ch <= '9' {
				// Convert character to integer
				grid[i][j] = int(ch - '0')

				// If character is invalid
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	// Validate the initial Sudoku grid
	if !isGridValid(grid) {
		fmt.Println("Error")
		return
	}

	// Attempt to solve the Sudoku
	if !solve(grid) {
		fmt.Println("Error")
		return
	}

	// Print the solved Sudoku
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			// Print number
			fmt.Print(grid[i][j])

			// Print space between numbers (except last column)
			if j < 8 {
				fmt.Print(" ")
			}
		}

		// Move to next line after each row
		fmt.Println()
	}
}
