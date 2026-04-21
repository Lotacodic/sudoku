# 🧩 Sudoku Solver in Go

A command-line Sudoku solver written in Go that reads a 9×9 Sudoku puzzle, validates it, and solves it using a **backtracking algorithm**.

---

## 📌 Features

* ✅ Validates input format (9 rows, 9 characters each)
* ✅ Ensures only valid characters (`1-9` and `.`)
* ✅ Checks Sudoku rules before solving (rows, columns, 3×3 boxes)
* ✅ Solves Sudoku using **recursion + backtracking**
* ✅ Handles invalid or unsolvable puzzles
* ✅ Clean and readable output format

---

## 🧠 How It Works

This solver is built around **three core concepts**:

### 1. Constraint Checking

Checks whether placing a number violates Sudoku rules:

* No duplicates in a row
* No duplicates in a column
* No duplicates in a 3×3 box

---

### 2. Grid Validation

Before solving, the program ensures the given puzzle is **already valid**:

* No conflicting numbers
* Proper structure

---

### 3. Backtracking Algorithm (Core Engine)

The solver:

1. Finds an empty cell (`.` → `0`)
2. Tries numbers from `1` to `9`
3. Checks if valid
4. Recursively continues
5. If stuck → **backtracks** (undo and try another number)

---

## 📂 Project Structure

```text
.
├── main.go   # Full implementation
└── README.md # Documentation
```

---

## 🚀 Usage

### 🔧 Run the program

```bash
go run . "ROW1" "ROW2" "ROW3" "ROW4" "ROW5" "ROW6" "ROW7" "ROW8" "ROW9"
```

---

## ✅ Example (Valid Sudoku)

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

### ✔ Output

```text
3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

---

## ❌ Error Cases

The program prints:

```text
Error
```

### When?

* ❌ Less or more than 9 rows
* ❌ Row length is not 9
* ❌ Invalid characters (not `1-9` or `.`)
* ❌ Duplicate numbers in row/column/box
* ❌ Unsolvable Sudoku

---

## 🧪 Example (Invalid Input)

```bash
go run . 1 2 3 4
```

```text
Error
```

---

## 🧪 Example (Duplicate Violation)

```bash
go run . "11......." "........." "........." "........." "........." "........." "........." "........." "........."
```

```text
Error
```

---

## 🧠 Algorithm Breakdown

### 🔁 Backtracking Flow

```text
Find empty cell
    ↓
Try numbers 1–9
    ↓
Check validity
    ↓
Place number
    ↓
Recurse
    ↓
If fail → undo (backtrack)
```

---

## ⚙️ Key Functions

### `isValid(grid, row, col, num)`

Checks if a number can be placed safely.

---

### `isGridValid(grid)`

Ensures the initial grid has no rule violations.

---

### `solve(grid)`

Uses recursion and backtracking to solve the puzzle.

---

### `main()`

Controls:

* Input parsing
* Validation
* Solving
* Output

---

## 🧱 Data Structure

The Sudoku grid is stored as:

```go
[][]int
```

* `0` represents empty cells
* `1–9` represent filled cells

---

## 📈 Complexity

* **Time Complexity:** Exponential (worst case)
* **Space Complexity:** O(1) (in-place solving)

---

## 💡 Concepts Learned

* Recursion
* Backtracking
* 2D slices in Go
* Input validation
* Problem decomposition

---

## 🚀 Future Improvements

* 🔹 GUI version
* 🔹 Solve multiple puzzles from file
* 🔹 Optimize with heuristics (e.g., least possibilities first)
* 🔹 Add step-by-step solver visualization

---

## 🏁 Conclusion

This project demonstrates how a complex problem like Sudoku can be solved using:

* Logical constraints
* Recursive thinking
* Backtracking

It’s a foundational algorithm useful in many real-world problems like:

* Pathfinding
* Puzzle solving
* AI search problems

---

## 👨‍💻 Author

Built as part of learning **Go and algorithmic problem-solving**.

---

## ⭐ If you like this project

Give it a star ⭐ and keep building!