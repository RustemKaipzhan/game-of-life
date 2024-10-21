package game

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func InitGame() {
	readInput()
}

func runGame(cells [][]cell) {
	for {
		if !continuePrint { // clear console at each iteration
			clearScreen()
		}
		tick++
		numLive = countLive(cells)
		if showVer {
			fmt.Println("Tick: ", tick)
			fmt.Println("Grid Size: ", row, "x", column)
			fmt.Println("Live Cells: ", numLive)
			fmt.Println("Default Delay: ", delay, "ms")
		}
		if colored {
			changeColor(cells)
		}
		drawCells(cells)
		if isEmpty(cells) { // Stop loop if there is no live cell
			break
		}

		countLiveNeighbors(cells)
		evolve(cells)
		fmt.Println()

		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout // Set the output to the terminal
	cmd.Run()
}
