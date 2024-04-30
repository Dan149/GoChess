package main

import (
	"fmt"
	"gochess/chessboard"
)

func main() {
	fmt.Print("Welcome to the golang-based CLI chess game.\n\n")
	board := chessboard.NewBoard()
	board.TryMovingPiece([2]uint8{0, 1}, [2]uint8{3, 3})
	board.Print()
	board.GetPiece(3, 3).PrintMovesGrid()
}
