// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20

		initVx, initVy = 5, 2
	)

	var (
		px, py   int              // ball position
		ppx, ppy int              // previous ball position
		vx, vy   = initVx, initVy // velocities

		cell rune // current cell (for caching)
	)

	// you can get the width and height using the screen package easily:
	width, height := screen.Size()

	// get the rune width of the ball emoji
	ballWidth := runewidth.RuneWidth(cellBall)

	// adjust the width and height
	width /= ballWidth
	height-- // there is a 1 pixel border in my terminal

	// create a single-dimensional board
	board := make([]bool, width*height)

	// create a drawing buffer
	buf := make([]rune, 0, width*height)

	// clear the screen once
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-initVx {
			vx *= -1
		}
		if py <= 0 || py >= height-initVy {
			vy *= -1
		}

		// check whether the ball goes beyond the borders
		if px < width && py < height {
			// calculate the new and the previous ball positions
			pos := py*width + px
			ppos := ppy*width + ppx

			// remove the previous ball and put the new ball
			board[pos], board[ppos] = true, false

			// save the previous positions
			ppx, ppy = px, py
		}

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				cell = cellEmpty

				if board[y*width+x] {
					cell = cellBall
				}

				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		// print the buffer
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		// slow down the animation
		time.Sleep(speed)
	}
}
