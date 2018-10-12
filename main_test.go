package main

import "testing"

func TestGameBoard_Neighbours(t *testing.T) {

	gb0 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb1 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, false, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb2 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb3 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb4 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb5 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb6 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb7 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	gb8 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	//corner top left

	gb9 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, false, false, false, false,
		false, false, false, false, false,
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
		false, false, false, false, false,
	}}

	gb11 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, true, false, false, false,
		true, false, false, false, false,
		false, false, false, false, false,
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
		false, false, false, false, false,
	}}

	//corner bottom right
	gb13 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
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
		false, false, false, false, false,
		false, false, false, true, false,
		false, false, false, false, true,
	}}

	gb15 := &GameBoard{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
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
