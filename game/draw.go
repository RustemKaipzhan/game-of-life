package game

import "fmt"

// Draw grid
func drawCells(cells [][]cell) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			cell := cells[i][j]

			fmt.Print(cell.Color)

			if cell.Live {
				fmt.Print("\033[1m") // bold text
				fmt.Print("×")
			} else if cells[i][j].IsVisited {
				fmt.Print("\033[1m")
				fmt.Print("∘")
			} else {
				fmt.Print("·")
			}
			fmt.Print(reset)
			fmt.Print(" ")
		}

		fmt.Print("\n")
	}
}
