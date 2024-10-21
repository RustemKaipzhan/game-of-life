package game

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printHelp() {
	fmt.Println(`Game of Life simulation.
--verbose      : Show additional information.
--delay-ms=X   : Set animation speed in milliseconds (default 2500).
--file=X       : Load the initial grid from a specified file.
--edges-portal : Enable portal edges where cells that exit the grid appear on the opposite side.
--random=WxH   : Generate a random grid of the specified width (W) and height (H).
--fullscreen   : Adjust the grid to fit the terminal size with empty cells.
--footprints   : Add traces of visited cells, displayed as '∘'.
--colored      : Add color to live cells and traces if footprints are enabled.
--continuePrint: Prints each tick of the game one after another.`)
}
func readInput() {
	delay = 2500
	cells := [][]cell{}
	for _, arg := range os.Args[1:] { // take all args end compare it with cases
		switch {
		case arg == "--continuePrint":
			continuePrint = true
		case arg == "--help":
			printHelp()
			return
		case arg == "--edges-portal":
			portal = true
		case arg == "--verbose":
			showVer = true
		case arg == "--footprints":
			footPrints = true
		case arg == "--colored":
			colored = true
		case arg == "--fullscreen":
			fullscreen = true
		case strings.HasPrefix(arg, "--delay-ms="):
			value := strings.TrimPrefix(arg, "--delay-ms=")
			if d, err := strconv.Atoi(value); err == nil {
				delay = d
			} else {
				fmt.Println("Error: invalid delay value.")
				return
			}
		case strings.HasPrefix(arg, "--random="):
			if !randomized {
				if fileRead {
					fmt.Println("Error: cannot make random cells while reading file!")
					return
				}
				sizeGrid := strings.TrimPrefix(arg, "--random=")
				dimensions := strings.Split(sizeGrid, "x")

				if len(dimensions) != 2 {
					fmt.Println("Error: invalid random value. Use format AxB (e.g., 3x3).")
					return
				}

				if row, err = strconv.Atoi(dimensions[0]); err != nil {
					fmt.Println("Error: invalid row value.")
					return
				}
				if column, err = strconv.Atoi(dimensions[1]); err != nil {
					fmt.Println("Error: invalid column value.")
					return
				}
				if row < 3 || column < 3 || row > 30 || column > 30 {
					fmt.Println("Error: invalid value.")
					return
				}
				randomized = true
			}

		case strings.HasPrefix(arg, "--file="):

			if randomized {
				fmt.Println("Error: cannot read file while making random cells!")
				return
			}
			if !fileRead {
				path := strings.TrimPrefix(arg, "--file=")
				cells = readFile(path)

				if cells == nil {
					return
				}

				fileRead = true
			}
		default:
			fmt.Println("Error: unrecognized flag:", arg)
			return
		}
	}

	// Handle conflict flags
	if !randomized && !fileRead {
		fmt.Print("Enter the number of rows and columns (minimum 3 x 3 and maximum 30x30): ")
		_, err = fmt.Scan(&row, &column)
		if err != nil || row < 3 || column < 3 || row > 30 || column > 30 {
			fmt.Println("Error: invalid input.")
			return
		}

		cells = initCells()
		message := makeCustomCells(cells)
		if message != "" {
			fmt.Println(message)
			return
		}
	} else if !fileRead {
		cells = initCells()
		makeRandomCells(cells)
	}

	if fullscreen {
		ncolumn, nrow := getTerminalSize()
		row = nrow
		column = ncolumn
		if ncolumn != len(cells[0]) || nrow != len(cells) {
			cells = resizeGrid(cells, ncolumn, nrow)
		}

	}

	runGame(cells)
}
