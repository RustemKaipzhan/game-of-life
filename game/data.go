package game

type cell struct { // stucture of cell
	Value     int
	Live      bool
	Color     string
	IsVisited bool
}

var (
	tick          int
	row           int
	column        int
	numLive       int
	delay         int
	showVer       bool
	footPrints    bool
	colored       bool
	randomized    bool
	fullscreen    bool
	portal        bool
	fileRead      bool
	err           error
	continuePrint bool
)
