package chessboard

func (c *Chessboard) pieceExists(position [2]uint8) bool {
	return c.matrix[position[0]][position[1]].exists
}

func (c *Chessboard) canMove(fromPosition [2]uint8, toPosition [2]uint8) bool {
	if fromPosition[0] == toPosition[0] && fromPosition[1] == toPosition[1] {
		return false
	}
	p := c.matrix[fromPosition[0]][fromPosition[1]]
	if p.side != c.playingSide {
		return false
	}
	switch p.kind {
	case 'p':
		return canPawnMove(p.position, toPosition)
	case 'b':
		return canBishopMove(p.position, toPosition)
	case 'n':
		return canKnightMove(p.position, toPosition)
	case 'r':
		return canRookMove(p.position, toPosition)
	case 'q':
		return canQueenMove(p.position, toPosition)
	case 'k':
		return canKingMove(p.position, toPosition)
	default:
		return c.matrix[toPosition[0]][toPosition[1]].exists
	}
}

// validation for each piece kind:

func canPawnMove(fromPosition [2]uint8, toPosition [2]uint8) bool {
	return true
}
func canBishopMove(fromPosition [2]uint8, toPosition [2]uint8) bool {
	return true
}
func canKnightMove(fromPosition [2]uint8, toPosition [2]uint8) bool {
	return true
}
func canRookMove(fromPosition [2]uint8, toPosition [2]uint8) bool {
	return true
}
func canQueenMove(fromPosition [2]uint8, toPosition [2]uint8) bool {
	return true
}
func canKingMove(fromPosition [2]uint8, toPosition [2]uint8) bool {
	return true
}
