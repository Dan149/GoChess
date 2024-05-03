package chessboard

import "fmt"

func (p *Piece) PrintMovesGrid() {
	if p.exists {
		fmt.Println("Displaying Piece @ [row col] ->", p.position, "theorically possible moves:")
		for x := 7; x >= 0; x-- {
			for y := range p.movesGrid[x] {
				if p.movesGrid[x][y] {
					fmt.Print("▣ ")
				} else {
					fmt.Print("▧ ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Piece at given location does not exist.")
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
func (c *Chessboard) PrintIsOnCheck() {
	if c.IsOnCheck() {
		if c.sideOnCheck == 0 {
			fmt.Println("White's king is on check.")
		} else {
			fmt.Println("Black's king is on check.")
		}
	}
}
