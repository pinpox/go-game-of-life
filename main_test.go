package main

import (
	"fmt"
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
		{"0 Neigbours", gb0, 2, 2, 0},
		{"1 Neigbours", gb1, 2, 2, 1},
		{"2 Neigbours", gb2, 2, 2, 2},
		{"3 Neigbours", gb3, 2, 2, 3},
		{"4 Neigbours", gb4, 2, 2, 4},
		{"5 Neigbours", gb5, 2, 2, 5},
		{"6 Neigbours", gb6, 2, 2, 6},
		{"7 Neigbours", gb7, 2, 2, 7},
		{"8 Neigbours", gb8, 2, 2, 8},

		{"0 Neigbours, corner top left", gb9, 0, 0, 0},
		{"1 Neigbours, corner top left", gb10, 0, 0, 1},
		{"2 Neigbours, corner top left", gb11, 0, 0, 2},
		{"3 Neigbours, corner top left", gb12, 0, 0, 3},

		{"0 Neigbours, corner bottom right", gb13, 4, 4, 0},
		{"1 Neigbours, corner bottom right", gb14, 4, 4, 1},
		{"2 Neigbours, corner bottom right", gb15, 4, 4, 2},
		{"3 Neigbours, corner bottom right", gb16, 4, 4, 3},
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

	gb0 := NewGameBoard(5, 5)
	gb1 := NewGameBoard(5, 5)
	gb2 := NewGameBoard(5, 5)
	gb3 := NewGameBoard(5, 5)

	gb0want := NewGameBoard(5, 5)
	gb1want := NewGameBoard(5, 5)
	gb2want := NewGameBoard(5, 5)
	gb3want := NewGameBoard(5, 5)

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
