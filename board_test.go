package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchNeighborBomb(t *testing.T) {
	board := Board{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}
	assert.Equal(t, []MineState{1, 1, 1, 1}, board.fetchNeighborBomb(2, 2))
	assert.Equal(t, []MineState{1}, board.fetchNeighborBomb(0, 0))
	assert.Equal(t, []MineState{1}, board.fetchNeighborBomb(4, 0))
	assert.Equal(t, []MineState{1}, board.fetchNeighborBomb(0, 4))
	assert.Equal(t, []MineState{1}, board.fetchNeighborBomb(4, 4))
}
