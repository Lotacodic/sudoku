# 🧩 Sudoku Solver — Pseudocode

## 🟦 1. isValid

```text
FUNCTION isValid(grid, row, col, num):

    FOR i FROM 0 TO 8:
        IF grid[row][i] == num:
            RETURN false

        IF grid[i][col] == num:
            RETURN false

    startRow = (row DIV 3) * 3
    startCol = (col DIV 3) * 3

    FOR i FROM 0 TO 2:
        FOR j FROM 0 TO 2:
            IF grid[startRow + i][startCol + j] == num:
                RETURN false

    RETURN true

------

## 🟦 2. isValid

```text
FUNCTION isGridValid(grid):

    FOR i FROM 0 TO 8:
        FOR j FROM 0 TO 8:

            IF grid[i][j] != 0:

                num = grid[i][j]
                grid[i][j] = 0

                IF isValid(grid, i, j, num) == false:
                    RETURN false

                grid[i][j] = num

    RETURN true