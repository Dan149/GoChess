package chessboard

func (c *Chessboard) movePiece(fromPosition [2]uint8, toPosition [2]uint8) {
	if c.pieceExists(fromPosition) {
		p := c.matrix[fromPosition[0]][fromPosition[1]]
		p.position = toPosition
		c.matrix[toPosition[0]][toPosition[1]] = p
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
