package chessboard

func (c *Chessboard) pieceExists(position [2]uint8) bool {
	return c.matrix[position[0]][position[1]].exists
}

func (c *Chessboard) canMove(fromPosition [2]uint8, toPosition [2]uint8) bool {
	if fromPosition == toPosition {
		return false
	}
	if c.matrix[toPosition[0]][toPosition[1]].kind == 'k' {
		return false
	}
	p := c.matrix[fromPosition[0]][fromPosition[1]]
	if c.matrix[toPosition[0]][toPosition[1]].exists && c.matrix[toPosition[0]][toPosition[1]].side == p.side {
		return false
	}
	if p.side != c.playingSide {
		return false
	}
	if !p.movesGrid[toPosition[0]][toPosition[1]] {
		return false
	}
	switch p.kind {
	case 'p':
		return canPawnMove(c, p, toPosition)
	case 'b':
		return canBishopMove(c, p, toPosition)
	case 'n':
		return true
	case 'r':
		return canRookMove(c, p, toPosition)
	case 'q':
		return canQueenMove(c, p, toPosition)
	case 'k':
		return true
	default:
		return c.matrix[toPosition[0]][toPosition[1]].exists
	}
}

//

func absoluteDiffUint8(x, y uint8) uint8 {
	if x < y {
		return y - x
	}
	return x - y
}

// validation for each piece kind:

func canPawnMove(c *Chessboard, p Piece, to [2]uint8) bool {
	if p.position[1] == to[1] && !c.matrix[to[0]][to[1]].exists {
		return true
	}
	if (p.position[1]+1 == to[1] || p.position[1]-1 == to[1]) && c.matrix[to[0]][to[1]].exists {
		return true
	}
	return false
}
func canBishopMove(c *Chessboard, p Piece, to [2]uint8) bool {
	rowDirection := int(1)
	colDirection := int(1) // up-right by default
	if p.position[0] > to[0] {
		rowDirection = -1
		if p.position[1] > to[1] { // down-left
			colDirection = -1
		} // else remains down-right
	} else { // up
		if p.position[1] > to[1] { // up-left
			colDirection = -1
		} // else remains up-right
	}
	var x uint8
	for x = 1; x <= absoluteDiffUint8(p.position[1], to[1]); x++ {
		xPos0 := int(p.position[0]) + (int(x) * rowDirection)
		xPos1 := int(p.position[1]) + (int(x) * colDirection)
		xCell := c.matrix[xPos0][xPos1]
		if xCell.exists && uint8(xPos1) != to[1] {
			return false
		} else if uint8(xPos1) == to[1] {
			return true
		}
	}
	return false
}
func canRookMove(c *Chessboard, p Piece, to [2]uint8) bool {
	return true
}
func canQueenMove(c *Chessboard, p Piece, to [2]uint8) bool {
	return true
}
