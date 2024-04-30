package chessboard

func (c *Chessboard) movePiece(fromPosition [2]uint8, toPosition [2]uint8) {
	if c.pieceExists(fromPosition) {
		p := c.matrix[fromPosition[0]][fromPosition[1]]
		p.position = toPosition
		c.matrix[toPosition[0]][toPosition[1]] = p
		p.refreshMovesGrid()
		c.cleanCell(fromPosition)
	}
}

func (c *Chessboard) TryMovingPiece(fromPosition [2]uint8, toPosition [2]uint8) {
	if c.canMove(fromPosition, toPosition) {
		c.movePiece(fromPosition, toPosition)
	}
}

func (c *Chessboard) cleanCell(position [2]uint8) {
	c.matrix[position[0]][position[1]] = Piece{}
}

// refreshing Piece.movesGrid

func (p *Piece) refreshMovesGrid() {
	switch p.kind {
	case 'p':
		p.movesGrid = refreshPawnMovesGrid([2]int8{int8(p.position[0]), int8(p.position[1])}, p.side)

	case 'n':
		p.movesGrid = refreshKnightMovesGrid([2]int8{int8(p.position[0]), int8(p.position[1])})
	}
}

func refreshPawnMovesGrid(position [2]int8, side uint8) [8][8]bool {
	var freshGrid [8][8]bool
	var movement int8
	switch side {
	case 0:
		movement = 1
	case 1:
		movement = -1
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

	return freshGrid
}
