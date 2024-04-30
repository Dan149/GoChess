package chessboard

import "fmt"

type Piece struct {
	side     uint8    // either 0: white | 1: black
	kind     byte     // (p)awn | (b)ishop | k(n)ight | (r)ook | (q)ueen | (k)ing
	position [2]uint8 // x & y coordinates
	exists   bool
	movesGrid [8][8]bool // places where the piece can theorically go
}

func (p *Piece) PrintMovesGrid() {
	fmt.Println("Displaying Piece @ [row col] ->", p.position, "theorically possible moves:")
	for x:= 7; x >= 0; x--{
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

func NewPiece(side uint8, kind byte, position [2]uint8) Piece {
	piece := Piece{side: side, kind: kind, position: position, exists: true}
	piece.refreshMovesGrid()
	return piece
}

type Chessboard struct {
	matrix      [8][8]Piece
	turn        uint
	playingSide uint8 // same logic as for the Piece.side value
	sideOnCheck uint8 // 2 if none
}

func (c *Chessboard) IsOnCheck() bool {
	return c.sideOnCheck != 2
}

func (c *Chessboard) GetPiece(row uint8, col uint8) *Piece {
	return &c.matrix[row][col]
}

func (c *Chessboard) Print() {
	for i := 7; i >= 0; i-- {
		for j := range c.matrix[i] {
			if c.matrix[i][j].kind == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print(string(c.matrix[i][j].kind) + " ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

//

func NewBoard() *Chessboard {
	board := new(Chessboard)
	var i uint8
	for i = 0; i < 8; i++ {
		var j uint8
		for j = 0; j < 8; j++ {
			var side uint8 // 0 white, 1 black, 2 if no piece
			switch i {
			case 0, 1:
				side = 0
			case 6, 7:
				side = 1
			default:
				side = 2
			}
			if side != 2 {
				var kind byte
				switch i {
				case 1, 6:
					kind = 'p'
				case 0, 7:
					switch j {
					case 0, 7:
						kind = 'r'
					case 1, 6:
						kind = 'n'
					case 2, 5:
						kind = 'b'
					case 3:
						kind = 'q'
					case 4:
						kind = 'k'
					}
				}
				board.matrix[i][j] = NewPiece(side, kind, [2]uint8{i, j})
				board.turn = 1

			}
		}
	}
	return board
}
