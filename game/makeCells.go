package game

import (
	"fmt"
	"math/rand"
)

// Allow user to input grid
func makeCustomCells(cells [][]cell) string {
	for i := 0; i < row; i++ {
		var line string
		fmt.Scan(&line)

		if len(line) != column {
			return "Error: invalid input."
		}

		for j, char := range line {
			if char == '#' {
				cells[i][j].Live = true
				if footPrints {
					cells[i][j].IsVisited = true
				}
			} else if char == '.' {
				cells[i][j].Live = false
			} else {
				return "Error: invalid character in input."
			}
		}

	}
	return ""
}

func makeRandomCells(cells [][]cell) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			prob := rand.Float64()

			// Adjust probability: 70-80% dead cells, 20-30% live cells
			if prob < 0.25+rand.Float64()*0.05 { // Random chance between 25% to 30%
				cells[i][j].Live = true
				if footPrints {
					cells[i][j].IsVisited = true
				}
			}
		}
	}
}
