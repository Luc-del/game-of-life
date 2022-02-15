package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCell_birthRule(t *testing.T) {
	c := new(Cell)
	for i := 0; i < 9; i++ {
		t.Run(fmt.Sprintf("%d neighbours", i), func(t *testing.T) {
			if i == 3 {
				assert.True(t, c.birthRule(i))
			} else {
				assert.False(t, c.birthRule(i))
			}
		})
	}
}

func TestCell_comfortRule(t *testing.T) {
	t.Run("dead cell", func(t *testing.T) {
		c := Cell(false)
		for i := 0; i < 9; i++ {
			t.Run(fmt.Sprintf("%d neighbours", i), func(t *testing.T) {
				assert.False(t, c.comfortRule(i))
			})
		}
	})

	t.Run("living cell", func(t *testing.T) {
		c := Cell(true)
		for i := 0; i < 9; i++ {
			t.Run(fmt.Sprintf("%d neighbours", i), func(t *testing.T) {
				if i == 2 {
					assert.True(t, c.comfortRule(i))
				} else {
					assert.False(t, c.comfortRule(i))
				}

			})
		}
	})
}
