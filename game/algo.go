package game

// Change cells according to rules at the next iteration
func evolve(cells [][]cell) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			cell := cells[i][j]
			if !cell.Live && cell.Value == 3 {
				if footPrints {
					cells[i][j].IsVisited = true
				}
				cells[i][j].Live = true
			}
			if cell.Live && cell.Value < 2 {
				cells[i][j].Live = false
			}
			if cell.Live && cell.Value > 3 {
				cells[i][j].Live = false
			}
		}
	}
}

// Count all neighbour live cells
func countLiveNeighbors(cells [][]cell) {
	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			cells[i][j].Value = 0
		}
	}

	// Count for live neighbors
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			for _, cor := range directions {
				x, y := i+cor[0], j+cor[1]

				// adjust x and y for portal
				if portal {
					x = (x + row) % row
					y = (y + column) % column
				}
				if x >= 0 && x < row && y >= 0 && y < column {
					if cells[x][y].Live {
						cells[i][j].Value++
					}
				}
			}
		}
	}
}

// Check grid for any live cell
func isEmpty(cells [][]cell) bool {
	for _, row := range cells {
		for _, cell := range row {
			if cell.Live {
				return false
			}
		}
	}
	return true
}

// Count all live cells
func countLive(cells [][]cell) int {
	out := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if cells[i][j].Live {
				out++
			}
		}
	}
	return out
}

// Initialize cells
func initCells() [][]cell {
	cells := make([][]cell, row)

	for idx := range cells {
		cells[idx] = make([]cell, column)
	}

	return cells
}
