package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		Nx             int
		Ny             int
		aliveCellCoord []Coord
	}

	tests := []struct {
		name string
		args args
		want [][]Cell
	}{
		{
			name: "1 cell false",
			args: args{
				Nx: 1,
				Ny: 1,
			},
			want: [][]Cell{
				{false},
			},
		},
		{
			name: "1 cell true",
			args: args{
				Nx:             1,
				Ny:             1,
				aliveCellCoord: []Coord{{0, 0}},
			},
			want: [][]Cell{
				{true},
			},
		},
		{
			name: "3x3 cells false",
			args: args{
				Nx: 3,
				Ny: 3,
			},
			want: [][]Cell{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
		},
		{
			name: "3x3 cells with trues",
			args: args{
				Nx: 3,
				Ny: 3,
				aliveCellCoord: []Coord{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 2,
						Y: 1,
					},
					{
						X: 0,
						Y: 2,
					},
				},
			},
			want: [][]Cell{
				{true, false, true},
				{false, false, false},
				{false, true, false},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grid := NewGrid(test.args.Nx, test.args.Ny, test.args.aliveCellCoord...)
			assert.Equal(t, test.want, grid.Cells)
		})
	}

}

func TestGrid_CountAlive(t *testing.T) {
	type args struct {
		Nx             int
		Ny             int
		aliveCellCoord []Coord
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1x1, 0 cell alive",
			args: args{
				Nx: 1,
				Ny: 1,
			},
			want: 0,
		},
		{
			name: "1x1, 1 cell alive",
			args: args{
				Nx:             1,
				Ny:             1,
				aliveCellCoord: []Coord{{0, 0}},
			},
			want: 1,
		},
		{
			name: "3x3, 0 cells false",
			args: args{
				Nx: 3,
				Ny: 3,
			},
			want: 0,
		},
		{
			name: "3x3, 4cells with trues",
			args: args{
				Nx: 3,
				Ny: 3,
				aliveCellCoord: []Coord{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 2,
						Y: 1,
					},
					{
						X: 0,
						Y: 2,
					},
					{
						X: 1,
						Y: 2,
					},
				},
			},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGrid(test.args.Nx, test.args.Ny, test.args.aliveCellCoord...)
			require.Equal(t, test.want, g.CountAlive())
		})
	}
}
