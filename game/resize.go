package game

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// Resize grid and copy old to it
func resizeGrid(oldGrid [][]cell, newWidth, newHeight int) [][]cell {
	newGrid := make([][]cell, newHeight)
	for i := range newGrid {
		newGrid[i] = make([]cell, newWidth)
		for j := range newGrid[i] {

			// If old grid fits into new grid copy old grid
			if i < len(oldGrid) && j < len(oldGrid[i]) {
				newGrid[i][j] = oldGrid[i][j]
			} else {
				newGrid[i][j] = cell{IsVisited: false, Live: false} // else make cell dead
			}
		}
	}
	return newGrid
}

// Get terminal height and width
func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Printf("Error on getting size: %v\n", err)
		fmt.Println("Defoult size: 80x24")
		return 80, 24
	}
	fmt.Println(width, height)
	return width / 2, height
}
