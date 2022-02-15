package models

import "fmt"

type Grid struct {
	Nx, Ny int
	Cells  [][]Cell
}

type Coord struct {
	X, Y int
}

func NewGrid(Nx, Ny int, aliveCellCoord ...Coord) Grid {
	g := Grid{
		Nx: Nx,
		Ny: Ny,
	}

	g.Cells = make([][]Cell, Nx)
	for i := range g.Cells {
		g.Cells[i] = make([]Cell, Ny)
	}

	for _, c := range aliveCellCoord {
		g.Cells[c.X][c.Y] = true
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

func (g Grid) Iter() Grid {
	var alive []Coord
	for i := range g.Cells {
		for j, c := range g.Cells[i] {
			count := g.aliveNeighbours(i, j)
			if c.birthRule(count) || c.comfortRule(count) {
				alive = append(alive, Coord{i, j})
			}
		}
	}

	return NewGrid(g.Nx, g.Ny, alive...)
}

func (g Grid) aliveNeighbours(x, y int) (count int) {
	for i := max(0, x-1); i <= min(x+1, g.Nx-1); i++ {
		for j := max(0, y-1); j <= min(y+1, g.Ny-1); j++ {
			if g.Cells[i][j] {
				count++
			}
		}
	}

	//retrieve self
	if g.Cells[x][y] {
		count--
	}

	return count
}
