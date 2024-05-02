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
	switch p.kind {
	case 'p':
		return canPawnMove(c, p, toPosition)
	case 'b':
		return canBishopMove(c, p, toPosition)
	case 'n':
		return canKnightMove(p, toPosition)
	case 'r':
		return canRookMove(c, p, toPosition)
	case 'q':
		return canQueenMove(c, p, toPosition)
	case 'k':
		return canKingMove(p, toPosition)
	default:
		return c.matrix[toPosition[0]][toPosition[1]].exists
	}
}

// validation for each piece kind:

func canPawnMove(c *Chessboard, p Piece, to [2]uint8) bool {
	if p.movesGrid[to[0]][to[1]] {
		if p.position[1] == to[1] && !c.matrix[to[0]][to[1]].exists {
			return true
		}
		if (p.position[1]+1 == to[1] || p.position[1]-1 == to[1]) && c.matrix[to[0]][to[1]].exists {
			return true
		}
		return false
	} else {
		return false
	}
}
func canBishopMove(c *Chessboard, p Piece, to [2]uint8) bool {
	return true
}
func canKnightMove(p Piece, to [2]uint8) bool {
	if p.movesGrid[to[0]][to[1]] {
		return true
	}
	return false
}
func canRookMove(c *Chessboard, p Piece, to [2]uint8) bool {
	return true
}
func canQueenMove(c *Chessboard, p Piece, to [2]uint8) bool {
	return true
}
func canKingMove(p Piece, to [2]uint8) bool {
	if p.movesGrid[to[0]][to[1]] {
		return true
	}
	return false
}
