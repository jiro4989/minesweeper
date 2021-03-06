package main

import (
	"fmt"
	"math/rand"
)

type (
	Mine struct {
		neighborMineCount int
		hasBomb           bool
		isEmpty           bool
		hasFlag           bool
		isOpened          bool
	}
	Board [][]Mine
)

func newBoard(row, col int) (ret Board) {
	for r := 0; r < row; r++ {
		var line []Mine
		for c := 0; c < col; c++ {
			n := rand.Intn(2)
			mine := Mine{}
			if n == 0 {
				mine.hasBomb = true
			}
			line = append(line, mine)
		}
		ret = append(ret, line)
	}
	ret.updateNeighborBombCount()
	return
}

func (b *Board) updateNeighborBombCount() {
	board := (*b)
	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			cell := board[y][x]
			if !cell.hasBomb {
				continue
			}
			neighbor := board.fetchNeighborBomb(x, y)
			cnt := len(neighbor)
			board[y][x].neighborMineCount = cnt
		}
	}
}

func (b Board) fetchNeighborBomb(x, y int) (ret []Mine) {
	for x2 := x - 1; x2 <= x+1; x2++ {
		for y2 := y - 1; y2 <= y+1; y2++ {
			if x2 < 0 || len(b[0]) <= x2 || y2 < 0 || len(b) <= y2 {
				continue
			}
			v := b[y2][x2]
			if v.hasBomb {
				ret = append(ret, v)
			}
		}
	}
	return
}

func (b Board) show() {
	for _, row := range b {
		fmt.Println(row)
	}
}
