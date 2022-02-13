package models

import "fmt"

type Cell bool

type Grid struct{
	Nx, Ny int
	Cells [][]Cell
}

type Coord struct {
	x, y int
}

func NewGrid(Nx, Ny int, aliveCellCoord ...Coord) Grid {
	g := Grid{
		Nx:    Nx,
		Ny:    Ny,
	}

	g.Cells = make([][]Cell, Nx)
	for i := range g.Cells {
		g.Cells[i] = make([]Cell, Ny)
	}

	for _, c := range aliveCellCoord {
		g.Cells[c.x][c.y] = true
	}

	return g
}

func (g Grid) Print() {
	for _, k := range g.Cells {
		fmt.Println(k)
	}
}

func (g Grid) CountAlive() (count int) {
	for _, x := range g.Cells {
		for _, y := range x {
			if y {
				count++
			}
		}
	}

	return count
}

func (g *Grid) Iter() {

}