package main

import (
	"fmt"
	"log"
)

/*
Coordinates start top left
	   x ->
	y
	|
	V
*/

func main() {

	g := NewGameBoard(3, 5)
	g.Set(1, 0, true)
	g.Print()
}

type GameBoard struct {
	xSize, ySize int
	cells        []bool
}

func NewGameBoard(x, y int) *GameBoard {
	cells := make([]bool, x*y)
	return &GameBoard{xSize: x, ySize: y, cells: cells}
}

func (gb *GameBoard) Iterate() {

}

func (gb *GameBoard) Set(x, y int, val bool) {
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

func (gb *GameBoard) Print() {
	for y := 0; y < gb.ySize; y++ { // Rows
		for x := 0; x < gb.xSize; x++ { // Collumns
			if gb.Get(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
