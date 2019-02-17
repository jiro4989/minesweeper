package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchNeighborBomb(t *testing.T) {
	board := Board{
		{Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}},
		{Mine{hasBomb: false}, Mine{hasBomb: true}, Mine{hasBomb: false}, Mine{hasBomb: true}, Mine{hasBomb: false}},
		{Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}},
		{Mine{hasBomb: false}, Mine{hasBomb: true}, Mine{hasBomb: false}, Mine{hasBomb: true}, Mine{hasBomb: false}},
		{Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}, Mine{hasBomb: false}},
	}
	assert.Equal(t, []Mine{Mine{hasBomb: true}, Mine{hasBomb: true}, Mine{hasBomb: true}, Mine{hasBomb: true}}, board.fetchNeighborBomb(2, 2))
	assert.Equal(t, []Mine{Mine{hasBomb: true}}, board.fetchNeighborBomb(0, 0))
	assert.Equal(t, []Mine{Mine{hasBomb: true}}, board.fetchNeighborBomb(4, 0))
	assert.Equal(t, []Mine{Mine{hasBomb: true}}, board.fetchNeighborBomb(0, 4))
	assert.Equal(t, []Mine{Mine{hasBomb: true}}, board.fetchNeighborBomb(4, 4))
}
