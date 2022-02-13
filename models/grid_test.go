package models

import (
	"github.com/stretchr/testify/assert"
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
				Nx:             1,
				Ny:             1,
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
				aliveCellCoord: []Coord{{0,0}},
			},
			want: [][]Cell{
				{true},
			},
		},
		{
			name: "3x3 cells false",
			args: args{
				Nx:             3,
				Ny:             3,
			},
			want: [][]Cell{
				{false, false, false},
				{false, false, false},
				{false, false, false},
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