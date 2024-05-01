package chessboard

import "fmt"

func (p *Piece) PrintMovesGrid() {
	fmt.Println("Displaying Piece @ [row col] ->", p.position, "theorically possible moves:")
	for x := 7; x >= 0; x-- {
		for y := range p.movesGrid[x] {
			if p.movesGrid[x][y] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func (c *Chessboard) PrintTurnNumber() {
	fmt.Print("Turn N°", c.turn)
}

func (c *Chessboard) PrintPlayingSide() {
	if c.playingSide == 0 {
		fmt.Print("White's turn.")
	} else if c.playingSide == 1 {
		fmt.Print("Black's turn.")
	}
}

func (c *Chessboard) PrintBoard() {
	fmt.Print("  -----------------\n")
	for i := 7; i >= 0; i-- {
		fmt.Print(i, "| ")
		for j := range c.matrix[i] {
			if c.matrix[i][j].kind == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print(string(c.matrix[i][j].kind) + " ")
			}
		}
		fmt.Print("|")
		fmt.Println("")
	}
	fmt.Print("  -----------------\n   ")
	for i := 0; i <= 7; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("")
}
