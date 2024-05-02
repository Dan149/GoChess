package chessboard

func (c *Chessboard) movePiece(fromPosition [2]uint8, toPosition [2]uint8) {
	if c.pieceExists(fromPosition) {
		p := c.matrix[fromPosition[0]][fromPosition[1]]
		p.position = toPosition
		c.matrix[toPosition[0]][toPosition[1]] = p
		c.refreshMovesGrid(p.position)
		c.cleanCell(fromPosition)
		if c.playingSide == 1 {
			c.turn++
		}
	}
}

func (c *Chessboard) TryMovingPiece(fromPosition [2]uint8, toPosition [2]uint8) {
	if c.canMove(fromPosition, toPosition) {
		c.movePiece(fromPosition, toPosition)
		if c.playingSide == 0 {
			c.playingSide = 1
		} else if c.playingSide == 1 {
			c.playingSide = 0
		}
	}
}

func (c *Chessboard) cleanCell(position [2]uint8) {
	c.matrix[position[0]][position[1]] = Piece{}
}

// refreshing Piece.movesGrid

func (c *Chessboard) refreshMovesGrid(position [2]uint8) {
	p := c.matrix[position[0]][position[1]]
	switch p.kind {
	case 'p':
		p.movesGrid = refreshPawnMovesGrid([2]int8{int8(position[0]), int8(position[1])}, p.side)

	case 'n':
		p.movesGrid = refreshKnightMovesGrid([2]int8{int8(position[0]), int8(position[1])})
	case 'b':
		p.movesGrid = refreshBishopMovesGrid([2]int8{int8(position[0]), int8(position[1])})
	case 'r':
		p.movesGrid = refreshRookMovesGrid([2]int8{int8(position[0]), int8(position[1])})

	case 'q':
		p.movesGrid = refreshQueenMovesGrid([2]int8{int8(position[0]), int8(position[1])})
	case 'k':
		p.movesGrid = refreshKingMovesGrid([2]int8{int8(position[0]), int8(position[1])})
	}
	c.matrix[position[0]][position[1]] = p
}

func refreshPawnMovesGrid(position [2]int8, side uint8) [8][8]bool {
	var freshGrid [8][8]bool
	var movement int8
	switch side {
	case 0:
		movement = 1
		if position[0] == 1 {
			freshGrid[3][position[1]] = true
		}
	case 1:
		movement = -1
		if position[0] == 6 {
			freshGrid[4][position[1]] = true
		}
	}
	if position[0] < 7 {
		freshGrid[position[0]+movement][position[1]] = true
		if position[1] < 7 {
			freshGrid[position[0]+movement][position[1]+1] = true
		}
		if position[1] > 0 {
			freshGrid[position[0]+movement][position[1]-1] = true
		}
	}
	return freshGrid
}
func refreshKnightMovesGrid(position [2]int8) [8][8]bool {
	var freshGrid [8][8]bool
	movements := [8][2]int8{{1, 2}, {2, 1}, {-1, -2}, {-2, -1}, {1, -2}, {-2, 1}, {-1, 2}, {2, -1}}
	for x := range movements {
		mov_pos0 := position[0] + movements[x][0]
		mov_pos1 := position[1] + movements[x][1]
		if 0 <= mov_pos0 && mov_pos0 <= 7 && 0 <= mov_pos1 && mov_pos1 <= 7 {
			freshGrid[mov_pos0][mov_pos1] = true
		}
	}
	return freshGrid
}

func refreshBishopMovesGrid(position [2]int8) [8][8]bool {
	var freshGrid [8][8]bool
	for i := int8(1); i <= 7; i++ {
		movements := [4][2]int8{{i, i}, {-i, -i}, {-i, i}, {i, -i}}
		for x := range movements {
			mov_pos0 := position[0] + movements[x][0]
			mov_pos1 := position[1] + movements[x][1]
			if 0 <= mov_pos0 && mov_pos0 <= 7 && 0 <= mov_pos1 && mov_pos1 <= 7 {
				freshGrid[mov_pos0][mov_pos1] = true
			}
		}
	}
	return freshGrid
}

func refreshRookMovesGrid(position [2]int8) [8][8]bool {
	var freshGrid [8][8]bool
	for i := int8(1); i <= 7; i++ {
		movements := [4][2]int8{{i, 0}, {-i, 0}, {0, i}, {0, -i}}
		for x := range movements {
			mov_pos0 := position[0] + movements[x][0]
			mov_pos1 := position[1] + movements[x][1]
			if 0 <= mov_pos0 && mov_pos0 <= 7 && 0 <= mov_pos1 && mov_pos1 <= 7 {
				freshGrid[mov_pos0][mov_pos1] = true
			}
		}
	}
	return freshGrid
}

func refreshQueenMovesGrid(position [2]int8) [8][8]bool {
	var freshGrid [8][8]bool
	for i := int8(1); i <= 7; i++ {
		movements := [8][2]int8{{i, 0}, {-i, 0}, {0, i}, {0, -i}, {i, i}, {-i, -i}, {-i, i}, {i, -i}}
		for x := range movements {
			mov_pos0 := position[0] + movements[x][0]
			mov_pos1 := position[1] + movements[x][1]
			if 0 <= mov_pos0 && mov_pos0 <= 7 && 0 <= mov_pos1 && mov_pos1 <= 7 {
				freshGrid[mov_pos0][mov_pos1] = true
			}
		}
	}
	return freshGrid
}

func refreshKingMovesGrid(position [2]int8) [8][8]bool {
	var freshGrid [8][8]bool
	movements := [8][2]int8{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	for x := range movements {
		mov_pos0 := position[0] + movements[x][0]
		mov_pos1 := position[1] + movements[x][1]
		if 0 <= mov_pos0 && mov_pos0 <= 7 && 0 <= mov_pos1 && mov_pos1 <= 7 {
			freshGrid[mov_pos0][mov_pos1] = true
		}
	}
	return freshGrid
}
