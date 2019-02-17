package main

import (
	"fmt"
	"math/rand"
)

type (
	MineState int
	Board     [][]MineState
)

const (
	stateBomb   MineState = 1 << iota
	stateFlag             // 旗的な意味なので注意
	stateOpened           // 表示ずみ
	stateEmpty  = 0
)

func newBoard(row, col int) (ret Board) {
	fmt.Println(stateEmpty, stateBomb, stateFlag, stateOpened)
	for r := 0; r < row; r++ {
		var line []MineState
		for c := 0; c < col; c++ {
			n := rand.Intn(2)
			line = append(line, MineState(n))
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
			fmt.Println(cell, stateBomb, cell&stateBomb)
			if cell&stateBomb == stateBomb {
				continue
			}
			neighbor := board.fetchNeighborBomb(x, y)
			cnt := len(neighbor)
			board[y][x] = MineState(cnt)
		}
	}
}

func (b Board) fetchNeighborBomb(x, y int) (ret []MineState) {
	for x2 := x - 1; x2 <= x+1; x2++ {
		for y2 := y - 1; y2 <= y+1; y2++ {
			if x2 < 0 || len(b[0]) <= x2 || y2 < 0 || len(b) <= y2 {
				continue
			}
			v := b[y2][x2]
			if v == stateBomb {
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
