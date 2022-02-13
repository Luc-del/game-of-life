package main

import (
	"game-of-life/models"
)

const (
	Nx = 5
	Ny = 10
)

func main() {
	grid := models.NewGrid(Nx, Ny)
	grid.Iter()
	grid.Print()
}
