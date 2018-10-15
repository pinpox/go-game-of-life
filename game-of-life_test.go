package main

import (
	"fmt"
	// "reflect"
	"testing"
)

func TestGameBoard_Neighbours(t *testing.T) {

	gb0 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb1 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, false, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb2 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb3 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb4 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb5 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb6 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, false, false, false,
		false, false, false, false, false,
	}}

	gb7 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, false, false,
		false, false, false, false, false,
	}}

	gb8 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, false, false, false, false,
	}}

	//corner top left

	gb9 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb10 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, true, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb11 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, true, false, false, false,
		true, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb12 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, true, false, false, false,
		true, true, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	//corner bottom right
	gb13 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, true,
	}}

	gb14 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, true, false,
		false, false, false, false, true,
	}}

	gb15 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, true, true,
		false, false, false, false, true,
	}}

	gb16 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, true, true,
		false, false, false, true, true,
	}}

	tests := []struct {
		name string
		gb   *GameBoard
		x    int
		y    int
		want int
	}{
		{"0 Neighbours", gb0, 2, 2, 0},
		{"1 Neighbours", gb1, 2, 2, 1},
		{"2 Neighbours", gb2, 2, 2, 2},
		{"3 Neighbours", gb3, 2, 2, 3},
		{"4 Neighbours", gb4, 2, 2, 4},
		{"5 Neighbours", gb5, 2, 2, 5},
		{"6 Neighbours", gb6, 2, 2, 6},
		{"7 Neighbours", gb7, 2, 2, 7},
		{"8 Neighbours", gb8, 2, 2, 8},

		{"0 Neighbours, corner top left", gb9, 0, 0, 0},
		{"1 Neighbours, corner top left", gb10, 0, 0, 1},
		{"2 Neighbours, corner top left", gb11, 0, 0, 2},
		{"3 Neighbours, corner top left", gb12, 0, 0, 3},

		{"0 Neighbours, corner bottom right", gb13, 4, 4, 0},
		{"1 Neighbours, corner bottom right", gb14, 4, 4, 1},
		{"2 Neighbours, corner bottom right", gb15, 4, 4, 2},
		{"3 Neighbours, corner bottom right", gb16, 4, 4, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gb.Neighbours(tt.x, tt.y); got != tt.want {
				tt.gb.Print()
				t.Errorf("GameBoard.Neighbours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBoard_InBounds(t *testing.T) {

	gb := NewGameBoard(5, 5)
	tests := []struct {
		name string
		gb   *GameBoard
		x    int
		y    int
		want bool
	}{
		{"Top Left", gb, 0, 0, true},
		{"Top Right", gb, 4, 0, true},
		{"Bottom Left", gb, 0, 4, true},
		{"Bottom Right", gb, 4, 4, true},

		{"Out Left", gb, -1, 0, false},
		{"Out Right", gb, 5, 0, false},
		{"Out Top", gb, 0, -1, false},
		{"Out Bottom", gb, 5, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gb := tt.gb
			if got := gb.InBounds(tt.x, tt.y); got != tt.want {
				t.Errorf("GameBoard.InBounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBoard_Iterate(t *testing.T) {

	// Any live cell with fewer than two live neighbors dies, as if by underpopulation
	gb0 := NewGameBoard(3, 3)
	/*
		0 0 0
		0 1 0
		0 0 1
	*/

	gb0.Set(1, 1, true)
	gb0.Set(2, 2, true)

	gb0want := NewGameBoard(3, 3)

	/*
		0 0 0
		0 0 0
		0 0 0
	*/

	// Any live cell with two or three live neighbors lives on to the next generation
	gb1 := NewGameBoard(3, 3)

	/*
		0 0 0
		1 1 0
		0 1 1
	*/

	gb1.Set(0, 1, true)
	gb1.Set(1, 1, true)
	gb1.Set(1, 2, true)
	gb1.Set(2, 2, true)

	gb1want := NewGameBoard(3, 3)

	/*
		0 0 0
		1 1 1
		1 1 1
	*/

	gb1want.Set(0, 1, true)
	gb1want.Set(0, 2, true)
	gb1want.Set(1, 1, true)
	gb1want.Set(1, 2, true)
	gb1want.Set(2, 1, true)
	gb1want.Set(2, 2, true)

	// Any live cell with more than three live neighbors dies, as if by overpopulation
	gb2 := NewGameBoard(3, 3)

	/*
		1 1 1
		1 1 1
		1 1 1
	*/

	gb2.Set(0, 0, true)
	gb2.Set(1, 0, true)
	gb2.Set(2, 0, true)
	gb2.Set(0, 1, true)
	gb2.Set(1, 1, true)
	gb2.Set(2, 1, true)
	gb2.Set(0, 2, true)
	gb2.Set(1, 2, true)
	gb2.Set(2, 2, true)

	gb2want := NewGameBoard(3, 3)

	/*
		1 0 1
		0 0 0
		1 0 1
	*/

	gb2want.Set(0, 0, true)
	gb2want.Set(2, 0, true)
	gb2want.Set(0, 2, true)
	gb2want.Set(2, 2, true)

	// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction
	gb3 := NewGameBoard(3, 3)

	/*
		0 0 0
		0 0 0
		1 1 1
	*/

	gb3.Set(0, 2, true)
	gb3.Set(1, 2, true)
	gb3.Set(2, 2, true)

	gb3want := NewGameBoard(3, 3)

	/*
		0 0 0
		0 1 0
		0 1 0
	*/

	gb3want.Set(1, 1, true)
	gb3want.Set(1, 2, true)

	tests := []struct {
		name string
		gb   *GameBoard
		want *GameBoard
	}{
		{"Any live cell with fewer than two live neighbors dies, as if by underpopulation.", gb0, gb0want},
		{"Any live cell with two or three live neighbors lives on to the next generation.", gb1, gb1want},
		{"Any live cell with more than three live neighbors dies, as if by overpopulation.", gb2, gb2want},
		{"Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.", gb3, gb3want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.gb.Iterate()

			if !tt.gb.Equal(tt.want) {
				t.Errorf("GameBoard.Iterate() = %v, want %v", tt.gb, tt.want)
				fmt.Println("Got:")
				tt.gb.Print()
				fmt.Println("Want:")
				tt.want.Print()
			}
		})
	}
}

func TestGameBoard_Equal(t *testing.T) {

	// Equal, all dead
	gb0a := NewGameBoard(5, 5)
	gb0b := NewGameBoard(5, 5)

	// Different xSize
	gb1a := NewGameBoard(4, 5)
	gb1b := NewGameBoard(5, 5)

	// Different ySize
	gb2a := NewGameBoard(5, 5)
	gb2b := NewGameBoard(5, 6)

	// Different xSize and ySize
	gb3a := NewGameBoard(5, 3)
	gb3b := NewGameBoard(3, 5)

	// Different cell length
	gb4a := NewGameBoard(5, 5)
	gb4b := NewGameBoard(5, 5)
	gb4b.cells = []bool{false, false, false}

	// Different cell(s)
	gb5a := NewGameBoard(5, 5)
	gb5b := NewGameBoard(5, 5)
	gb5b.Set(0, 0, true)

	// Equal, some alive
	gb6a := NewGameBoard(5, 5)
	gb6a.Set(0, 0, true)
	gb6a.Set(1, 0, true)
	gb6a.Set(2, 4, true)
	gb6a.Set(3, 2, true)

	gb6b := NewGameBoard(5, 5)
	gb6b.Set(0, 0, true)
	gb6b.Set(1, 0, true)
	gb6b.Set(2, 4, true)
	gb6b.Set(3, 2, true)

	tests := []struct {
		name string
		gba  *GameBoard
		gbb  *GameBoard
		want bool
	}{
		{"Equal, all dead", gb0a, gb0b, true},
		{"Different xSize", gb1a, gb1b, false},
		{"Different ySize", gb2a, gb2b, false},
		{"Different xSize and ySize", gb3a, gb3b, false},
		{"Different cell length", gb4a, gb4b, false},
		{"Different cells", gb5a, gb5b, false},
		{"Equal, some alive ", gb6a, gb6b, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gba.Equal(tt.gbb); got != tt.want {
				t.Errorf("GameBoard.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBoard_RandInit(t *testing.T) {
	tests := []struct {
		name       string
		gb         *GameBoard
		percentage int
		wantAlive  int
	}{
		{"0 percent", NewGameBoard(10, 5), 0, 0},
		{"30 percent", NewGameBoard(10, 5), 30, 15},
		{"50 percent", NewGameBoard(10, 5), 50, 25},
		{"100 percent", NewGameBoard(10, 5), 100, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.gb.RandInit(tt.percentage)

			count := 0
			for _, v := range tt.gb.cells {
				if v {
					count++
				}
			}

			if count != tt.wantAlive {
				t.Errorf("GameBoard.RandInit() ran, %v, cells alive,  want %v", count, tt.wantAlive)
			}
		})
	}
}

func TestGameBoard_Set(t *testing.T) {

	// Set dead cell to alive
	gb0 := NewGameBoard(3, 3)
	gb0.cells = []bool{
		false, false, false,
		false, false, false,
		false, false, false,
	}

	gb0want := NewGameBoard(3, 3)
	gb0want.cells = []bool{
		false, false, false,
		false, true, false,
		false, false, false,
	}

	// Set dead cell to dead
	gb1 := NewGameBoard(3, 3)
	gb1.cells = []bool{
		false, false, false,
		false, false, false,
		false, false, false,
	}

	gb1want := NewGameBoard(3, 3)
	gb1want.cells = []bool{
		false, false, false,
		false, false, false,
		false, false, false,
	}
	// Set living cell to alive
	gb2 := NewGameBoard(3, 3)
	gb2.cells = []bool{
		false, false, false,
		false, true, false,
		false, false, false,
	}

	gb2want := NewGameBoard(3, 3)
	gb2want.cells = []bool{
		false, false, false,
		false, true, false,
		false, false, false,
	}
	// Set living cell to dead
	gb3 := NewGameBoard(3, 3)
	gb3.cells = []bool{
		false, false, false,
		false, true, false,
		false, false, false,
	}

	gb3want := NewGameBoard(3, 3)
	gb3want.cells = []bool{
		false, false, false,
		false, false, false,
		false, false, false,
	}

	tests := []struct {
		name string
		gb   *GameBoard
		want *GameBoard
		x    int
		y    int
		val  bool
	}{
		{"Set dead cell to alive", gb0, gb0want, 1, 1, true},
		{"Set dead cell to dead", gb1, gb1want, 1, 1, false},
		{"Set living cell to alive", gb2, gb2want, 1, 1, true},
		{"Set living cell to dead", gb3, gb3want, 1, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gb := tt.gb
			gb.Set(tt.x, tt.y, tt.val)

			if !gb.Equal(tt.want) {
				t.Errorf("GameBoard.Set() = %v, want %v", gb, tt.want)
			}
		})
	}
}

func TestGameBoard_Get(t *testing.T) {

	gb0 := NewGameBoard(5, 3)
	gb0.cells = []bool{
		false, true, true, true, true,
		true, true, true, true, true,
		true, true, true, true, true,
	}

	gb1 := NewGameBoard(5, 3)
	gb1.cells = []bool{
		true, true, true, true, false,
		true, true, true, true, true,
		true, true, true, true, true,
	}

	gb2 := NewGameBoard(5, 3)
	gb2.cells = []bool{
		true, true, true, true, true,
		true, true, true, true, true,
		false, true, true, true, true,
	}

	gb3 := NewGameBoard(5, 3)
	gb3.cells = []bool{
		true, true, true, true, true,
		true, true, true, true, true,
		true, true, true, true, false,
	}

	gb4 := NewGameBoard(5, 3)
	gb4.cells = []bool{
		true, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}

	gb5 := NewGameBoard(5, 3)
	gb5.cells = []bool{
		false, false, false, false, true,
		false, false, false, false, false,
		false, false, false, false, false,
	}

	gb6 := NewGameBoard(5, 3)
	gb6.cells = []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		true, false, false, false, false,
	}

	gb7 := NewGameBoard(5, 3)
	gb7.cells = []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, true,
	}

	tests := []struct {
		name string
		gb   *GameBoard
		x    int
		y    int
		want bool
	}{
		{"Top left, dead", gb0, 0, 0, false},
		{"Top right, dead", gb1, 4, 0, false},
		{"Bottom left, dead", gb2, 0, 2, false},
		{"Bottom right, dead", gb3, 4, 2, false},
		{"Top left, alive", gb4, 0, 0, true},
		{"Top right, alive", gb5, 4, 0, true},
		{"Bottom left, alive", gb6, 0, 2, true},
		{"Bottom right, alive", gb7, 4, 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gb := tt.gb
			if got := gb.Get(tt.x, tt.y); got != tt.want {
				t.Errorf("GameBoard.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
