package main

import (
	"fmt"
	"game-of-life/models"
)

const (
	Nx = 5
	Ny = 10
)

var (
	initState = []models.Coord{
		{X: 2, Y: 2},
		{X: 2, Y: 3},
		{X: 2, Y: 1},
	}
)

func main() {
	grid := models.NewGrid(Nx, Ny, initState...)
	fmt.Println("Before:", grid.CountAlive())
	grid.Print()

	grid = grid.Iter()

	fmt.Println("After:", grid.CountAlive())
	grid.Print()
}
