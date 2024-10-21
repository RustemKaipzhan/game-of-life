# Game of Life Simulation

## Overview

The Game of Life is a cellular automaton devised by mathematician John Conway. This simulation allows users to interactively play the Game of Life using a grid of cells, with live cells represented by `#` and dead cells by `.`. The program implements various features and flags for customization.

## Features

- Accepts grid input to define the initial state of the game.
- Visualizes the grid using `×` for live cells and `·` for dead cells.
- Supports several command-line flags for enhanced functionality:
  - `--help`: Display help information.
  - `--verbose`: Show detailed information about the simulation at each tick.
  - `--delay-ms=X`: Set the animation speed in milliseconds (default is 2500 ms).
  - `--file=X`: Load the initial grid from a specified file.
  - `--edges-portal`: Enable portal edges, wrapping cells that exit the grid to the opposite side.
  - `--random=WxH`: Generate a random grid of specified dimensions.
  - `--fullscreen`: Adjust the grid to fit the terminal size.
  - `--footprints`: Add traces of visited cells, displayed as `∘`.
  - `--colored`: Add color to live cells and traces if footprints are enabled.
  - `--continuePrint`: Prints each tick of game one after another.
## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/game-of-life.git
    cd game-of-life

Input Format
The input consists of a grid of characters followed by optional flags. Each grid should be formatted as follows:

The first line specifies the dimensions (width and height) of the grid.
The following lines represent the grid, with # for live cells and . for dead cells.
Example:


    4 4
    ....
    .#..
    .##.
    ....


Running the Program
To run the simulation with various flags, use the following commands:


Basic run:


    go run main.go


Help:

    go run main.go --help

Verbose mode:


    go run main.go --verbose
Set animation speed:


    go run main.go --delay-ms=1000
Load from file:


    go run main.go --file=path/to/grid.txt
Enable portal edges:


    go run main.go --edges-portal
Generate a random grid:


    go run main.go --random=5x5
Fullscreen mode:


    go run main.go --fullscreen
Enable footprints:


    go run main.go --footprints
Enable colored output:

    go run main.go --colored

Enable continued print output:

    go run main.go --continuePrint

Game Rules
The simulation evolves according to the following rules:

A live cell with fewer than two live neighbors dies (underpopulation).
A live cell with two or three live neighbors lives on to the next generation.
A live cell with more than three live neighbors dies (overpopulation).
A dead cell with exactly three live neighbors becomes a live cell (reproduction).


The simulation continues until there are no living cells left.
>>>>>>> master
