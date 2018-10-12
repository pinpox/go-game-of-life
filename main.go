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

	gbOld := gb

	for y := 0; y < gb.ySize; y++ { // Rows
		for x := 0; x < gb.xSize; x++ { // Collumns

			// Eine tote Zelle mit genau drei lebenden Nachbarn wird in der Folgegeneration neu geboren.
			if !gbOld.Get(x, y) && gbOld.Neighbours(x, y) == 3 {
				gb.Set(x, y, true)
			}

			// Lebende Zellen mit weniger als zwei lebenden Nachbarn sterben in der Folgegeneration an Einsamkeit.
			if gbOld.Get(x, y) && gbOld.Neighbours(x, y) < 2 {
				gb.Set(x, y, false)
			}

			// Eine lebende Zelle mit zwei oder drei lebenden Nachbarn bleibt in der Folgegeneration am Leben.

			// Lebende Zellen mit mehr als drei lebenden Nachbarn sterben in der Folgegeneration an Überbevölkerung.
			if gbOld.Neighbours(x, y) > 3 {
				gb.Set(x, y, false)
			}
		}
	}

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
	return (x >= 0 && x < gb.xSize && y > 0 && y < gb.ySize)
}

func (gb *GameBoard) Print() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for y := 0; y < gb.ySize; y++ { // Rows
		for x := 0; x < gb.xSize; x++ { // Collumns
			if gb.Get(x, y) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
