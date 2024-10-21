package game

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) [][]cell {
	file, err := os.Open(path)
	cells := initCells()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	len1 := 0
	i := 0
	// Check if file is empty
	info, err := file.Stat()
	if err != nil {
		fmt.Println("error getting file info")
		return nil
	}
	if info.Size() == 0 {
		fmt.Println("file is empty")
		return nil
	}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row1 []cell

		if i == 0 {
			len1 = len(line)
		}

		if len1 != len(line) {
			fmt.Println("The columns are different in file!")
			return nil
		} else if len(line) < 3 || len(line) > 30 {
			fmt.Println("The column size must be at least 3 nad less than 30")
			return nil
		}
		for _, char := range line {
			// Determine if the cell is alive or dead based on the character
			if char == '#' {
				row1 = append(row1, cell{Live: true, IsVisited: true})
			} else if char == '.' {
				row1 = append(row1, cell{})
			} else {
				fmt.Println("Invalid symbols in file. Only . #")
				return nil
			}
		}

		cells = append(cells, row1) // Add the row to the cells slice
		i++
	}

	if i < 3 || i > 30 {
		fmt.Println("The row size must be at least 3 and less than 30")
		return nil
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil // Return nil if there's an error
	}

	row = len(cells)
	column = len(cells[0])
	return cells
}
