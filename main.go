package main

import (
	"fmt"
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Cursor struct {
	x int
	y int
}

var (
	board  = newBoard(5, 5)
	cursor = Cursor{}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	drawUI()
	waitKeyInput()
}

func drawUI() {
	const dc = termbox.ColorDefault
	termbox.Clear(dc, dc)
	for y, row := range board {
		for x, cell := range row {
			color := termbox.ColorWhite
			cnt := cell.neighborMineCount
			s := fmt.Sprintf("%d", cnt)
			c := []rune(s)[0]
			if !cell.isOpened {
				c = ' '
			}
			if cell.hasFlag {
				c = '?'
			}
			if cursor.x == x && cursor.y == y {
				color = termbox.ColorGreen
			}
			termbox.SetCell(2*x, y, c, dc, color)
			termbox.SetCell(2*x+1, y, ' ', dc, color)
		}
	}
	termbox.Flush()
}

func waitKeyInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC, termbox.KeyCtrlD:
				return
			case termbox.KeySpace:
				x, y := cursor.x, cursor.y
				board[y][x].isOpened = true
				if board[y][x].hasBomb {
					fmt.Println("GAME OVER")
				}
			}
			switch ev.Ch {
			case 'q':
				return
			case 'k':
				cursor.y--
			case 'h':
				cursor.x--
			case 'j':
				cursor.y++
			case 'l':
				cursor.x++
			}
		}
		drawUI()
	}
}
