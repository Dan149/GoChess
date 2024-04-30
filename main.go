package main

import (
	"fmt"
	"gochess/chessboard"
)

func main() {
	fmt.Print("Welcome to the golang-based CLI chess game.\n\n")
	board := chessboard.NewBoard()
	board.Print()
	board.GetPiece(6, 7).PrintMovesGrid()
}
