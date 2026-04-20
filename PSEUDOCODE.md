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

```

## 🟦 2. isGridValid

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
```

## 🟦 3. solve (Backtracking Core)


    ```
    FUNCTION solve(grid):

    FOR i FROM 0 TO 8:
        FOR j FROM 0 TO 8:

            IF grid[i][j] == 0:

                FOR num FROM 1 TO 9:

                    IF isValid(grid, i, j, num):

                        grid[i][j] = num

                        IF solve(grid):
                            RETURN true

                        grid[i][j] = 0

                RETURN false

    RETURN true
    ```

## 🟦 4. Main Program Flow

```
FUNCTION main:

    IF argc != 10:
        PRINT "Error"
        STOP

    CREATE 9x9 grid

    FOR i FROM 0 TO 8:

        row = input[i + 1]

        IF length(row) != 9:
            PRINT "Error"
            STOP

        FOR j FROM 0 TO 8:

            IF row[j] == '.':
                grid[i][j] = 0

            ELSE IF '1' <= row[j] <= '9':
                grid[i][j] = digit value

            ELSE:
                PRINT "Error"
                STOP

    IF isGridValid(grid) == false:
        PRINT "Error"
        STOP

    IF solve(grid) == false:
        PRINT "Error"
        STOP

    FOR i FROM 0 TO 8:
        FOR j FROM 0 TO 8:

            PRINT grid[i][j]

            IF j < 8:
                PRINT " "

        PRINT newline
        ```

```
## Master Flow

```
Input → Validate → Build Grid → Check Rules → Backtracking Solve → Output
```

## Core Idea

- Find empty cell
- Try numbers 1–9
- Check validity
- Place number
- Recurse
- If fail → backtrack

## 💡 Summary
The solver is based on:

1. Constraint checking (isValid)
2. Grid validation (isGridValid)
3. Backtracking solver (solve)
