package game

const ( // color for value of cell
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	white  = "\033[37m"
	gray   = "\033[37m"
)

// Change color of cell according to condition
func changeColor(cells [][]cell) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if cells[i][j].Live {
				cells[i][j].Color = green
			} else if cells[i][j].IsVisited && footPrints {
				cells[i][j].Color = blue
			} else {
				cells[i][j].Color = white
			}
		}
	}
}
