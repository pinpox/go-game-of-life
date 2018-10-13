package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var (
	inputX = kingpin.Flag("xsize", "The width of the grid").Short('x').Default("80").Int()
	inputY = kingpin.Flag("ysize", "The height of the grid").Short('y').Default("15").Int()
	inputI = kingpin.Flag("iterations", "Number of iterations. Any negative number will use the default, infinity").Short('i').Default("-1").Int()
	inputF = kingpin.Flag("fps", "Frames per second, how log to wait until the next iteration is displayed").Short('f').Default("25").Int()
	inputP = kingpin.Flag("percentage", "Percentage of living cells at the start").Short('p').Default("35").Int()
)

/*
Coordinates start top left
	   x ->
	y
	|
	V
*/

func main() {

	kingpin.Version("1.0.0")
	kingpin.Parse()

	g := NewGameBoard(*inputX, *inputY)
	g.RandInit(*inputP)
	sleepTime := time.Duration(1000 / *inputF) * time.Millisecond

	//Main game loop
	i := *inputI
	for {
		if i == 0 {
			break
		}
		i--
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		g.Print()
		fmt.Print("Generation: ")
		fmt.Println(i)
		g.Iterate()
		time.Sleep(sleepTime)
	}
}

// GameBoard holds the grid on which the cells are placed as well as a few of it's attributes.
// It provides the necessary methods to run the initialize and start the game
type GameBoard struct {
	generation   int
	xSize, ySize int
	cells        []bool
}

// NewGameBoard initializes a new game (constructor) with an all-dead grid
func NewGameBoard(x, y int) *GameBoard {
	cells := make([]bool, x*y)
	return &GameBoard{generation: 0, xSize: x, ySize: y, cells: cells}
}

// RandInit sets a given percentage of the cells in the grid to "alive".
// The grid is then randomized, leaving a random grid with a defined amount of
// living cells.
func (gb *GameBoard) RandInit(percentage int) {

	// Calculate number of living cells
	numAlive := percentage * len(gb.cells) / 100

	// Insert living cells at the beginning
	for i := 0; i < numAlive; i++ {
		gb.cells[i] = true
	}

	// Randomize slice
	vals := gb.cells
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(vals); n > 0; n-- {
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
	}

	gb.cells = vals
}

// Iterate implements the rules for Conway's Game of Life. It takes the current
// grid, applies the 4 rules and sets the grid to it's new state.
func (gb *GameBoard) Iterate() {

	gbOld := NewGameBoard(gb.xSize, gb.ySize)
	gbOld.cells = gb.cells

	for y := 0; y < gb.ySize; y++ { // Rows
		for x := 0; x < gb.xSize; x++ { // Collumns

			// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
			if !gbOld.Get(x, y) && gbOld.Neighbours(x, y) == 3 {
				gb.Set(x, y, true)
				continue
			}

			// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
			if gbOld.Get(x, y) && gbOld.Neighbours(x, y) < 2 {
				gb.Set(x, y, false)
				continue
			}

			// Any live cell with two or three live neighbors lives on to the next generation.
			if gbOld.Get(x, y) && ((gbOld.Neighbours(x, y) == 2) || (gbOld.Neighbours(x, y) == 3)) {
				//No need to set, already alive
				continue
			}

			// Any live cell with more than three live neighbors dies, as if by overpopulation.
			if gbOld.Get(x, y) && (gbOld.Neighbours(x, y) > 3) {
				gb.Set(x, y, false)
				continue
			}
		}
	}
}

//Equal is a helper function to compare a GameBoard to another. It returns true
//if size, dimensions and the cell's states are identical
func (gb *GameBoard) Equal(gb2 *GameBoard) bool {
	if gb.xSize != gb2.xSize || gb.ySize != gb2.ySize {
		return false
	}

	if len(gb.cells) != len(gb2.cells) {
		return false
	}

	for k := range gb.cells {
		if gb.cells[k] != gb2.cells[k] {
			return false
		}
	}

	return true
}

// Set sets a cell defined by it's x and y coordinates to a given state (alive:
// true, dead: false)
func (gb *GameBoard) Set(x, y int, val bool) {
	if !gb.InBounds(x, y) {

	}
	if x > gb.xSize-1 || x < 0 || y > gb.ySize-1 || y < 0 {
		log.Fatal("Invalid Coordinate")
	}

	gb.cells[y*(gb.xSize)+x] = val
}

// Get retrieves a cell's state by it's x and y coordinates (alive: true, dead:
// false)
func (gb *GameBoard) Get(x, y int) bool {
	if x > gb.xSize-1 || x < 0 || y > gb.ySize-1 || y < 0 {
		log.Fatal("Invalid Coordinate")
	}

	// fmt.Println(y*(gb.xSize) + x)
	return gb.cells[y*(gb.xSize)+x]
}

//Neighbours returns the number of alive neighbours of a given cell
func (gb *GameBoard) Neighbours(x, y int) int {
	count := 0
	arr := []int{-1, 0, 1}

	for _, v1 := range arr {
		for _, v2 := range arr {
			if gb.InBounds(x+v1, y+v2) {
				if gb.Get(x+v1, y+v2) && !(v1 == 0 && v2 == 0) {
					count++
				}
			}
		}
	}
	return count
}

//InBounds is a helper function to check if a coordinate tupel is inside the grid.
func (gb *GameBoard) InBounds(x int, y int) bool {
	return (x >= 0 &&
		x < gb.xSize &&
		y >= 0 &&
		y < gb.ySize)
}

// Print displays a ASCII representation of the grid to stout
func (gb *GameBoard) Print() {
	//Top margin
	fmt.Print("╔")
	for x := 1; x <= gb.xSize; x++ {
		fmt.Print("══")
	}
	fmt.Println("╗")

	// Rows
	for y := 0; y < gb.ySize; y++ {
		fmt.Print("║")
		// Collumns
		for x := 0; x < gb.xSize; x++ {
			if gb.Get(x, y) {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("║")
	}

	//Bottom margin
	fmt.Print("╚")
	for x := 1; x <= gb.xSize; x++ {
		fmt.Print("══")
	}
	fmt.Println("╝")
}
