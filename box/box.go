package box

import "math/rand"

var Grid, Target int

type Box [][]int

func New(g int, t int) Box {
	Grid = g
	Target = t
	return make(Box, Grid)
}

// Add element will random posisition and check if it's possible to filled
func (b Box) AddElement() {
	filled := false

	for !filled {
		posX := rand.Intn(Grid)
		posY := rand.Intn(Grid)

		if b[posX][posY] == 0 {
			b[posX][posY] = 2
			filled = true
		}
	}
}

func (b Box) isFull() bool {
	empty := 0
	for _, row := range b {
		for _, col := range row {
			if col == 0 {
				empty++
			}
		}
	}
	return empty == 0
}
