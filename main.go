package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

/*
Coordinates start top left
	   x ->
	y
	|
	V
*/

func main() {

	g := NewGameBoard(80, 30)
	// g.RandInit(600)
	g.Set(10, 10, true)
	g.Set(10, 11, true)
	g.Set(10, 11, true)

	for i := 0; i < 10; i++ {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		g.Print()
		fmt.Print("Generation: ")
		fmt.Println(i)
		g.Iterate()
		time.Sleep(1000 * time.Millisecond)
	}
}

type GameBoard struct {
	generation   int
	xSize, ySize int
	cells        []bool
}

func NewGameBoard(x, y int) *GameBoard {
	cells := make([]bool, x*y)
	return &GameBoard{generation: 0, xSize: x, ySize: y, cells: cells}
}

func (gb *GameBoard) RandInit(numAlive int) {

	//TODO check for doubles
	rand.Seed(time.Now().UnixNano())
	// rand.Seed(1)

	for i := 0; i < numAlive; i++ {
		gb.cells[rand.Intn(len(gb.cells))] = true
	}
}

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

func (gb *GameBoard) Set(x, y int, val bool) {
	if !gb.InBounds(x, y) {

	}
	if x > gb.xSize-1 || x < 0 || y > gb.ySize-1 || y < 0 {
		log.Fatal("Invalid Coordinate")
	}

	gb.cells[y*(gb.xSize)+x] = val
}

func (gb *GameBoard) Get(x, y int) bool {
	if x > gb.xSize-1 || x < 0 || y > gb.ySize-1 || y < 0 {
		log.Fatal("Invalid Coordinate")
	}

	// fmt.Println(y*(gb.xSize) + x)
	return gb.cells[y*(gb.xSize)+x]
}

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

func (gb *GameBoard) InBounds(x int, y int) bool {
	return (x >= 0 &&
		x < gb.xSize &&
		y >= 0 &&
		y < gb.ySize)
}

func (gb *GameBoard) Print() {
	//Top margin
	fmt.Print("╔")
	for x := 1; x <= gb.xSize; x++ {
		fmt.Print("═")
	}
	fmt.Println("╗")

	// Rows
	for y := 0; y < gb.ySize; y++ {
		fmt.Print("║")
		// Collumns
		for x := 0; x < gb.xSize; x++ {
			if gb.Get(x, y) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("║")
	}

	//Bottom margin
	fmt.Print("╚")
	for x := 1; x <= gb.xSize; x++ {
		fmt.Print("═")
	}
	fmt.Println("╝")
}
