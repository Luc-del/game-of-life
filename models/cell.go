package models

type Cell bool

// Here we can find the rules of the game
// It only covers scenarios were a cell is created or stays alive
// because "false" is the default boolean value, so we only review "true" cases
// If a given cell is not covered by any rule, it is dead

// a cell is created or stays alive if it has exactly 3 neighbours
func (c Cell) birthRule(aliveNeighbours int) bool {
	return aliveNeighbours == 3
}

// a cell stays alive if it has exactly 2 neighbours (has to be alive in the first place)
func (c Cell) comfortRule(aliveNeighbours int) bool {
	return bool(c) && aliveNeighbours == 2
}
