package main

import (
	"fmt"
	"gochess/chessboard"
	"strconv"
	"strings"

	"github.com/inancgumus/screen"
)

const HELP string = `
GoChess commands:
----------------
move <row>;<col> <row>;<col> : Moves piece from position 1 to position 2.
info <row>;<col>             : Display movement help for piece at given location.

q|quit|exit                  : Quits the program.
h|?|help                     : Prints this message.

`

func main() {
	board := chessboard.NewBoard()
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Print("Welcome to the golang-based CLI chess game.\n\n")
	for {
		board.PrintBoard()
		fmt.Println()
		board.PrintTurnNumber()
		fmt.Print(" | ")
		board.PrintPlayingSide()
		fmt.Println("\nType '?' to display the help message.")
		board.PrintIsOnCheck()
		fmt.Println()
		//
		input := make([]string, 3)
		fmt.Print("$ ")
		fmt.Scanln(&input[0], &input[1], &input[2])
		if input[0] == "q" || input[0] == "quit" || input[0] == "exit" {
			break
		} else if input[0] == "h" || input[0] == "?" || input[0] == "help" {
			screen.Clear()
			screen.MoveTopLeft()
			fmt.Print(HELP)
		} else {
			screen.Clear()
			screen.MoveTopLeft()
			if input[0] == "move" && len(input) == 3 {
				fmt.Print(input, "\n\n")
				board.TryMovingPiece(parseMovementString(input))
			} else if input[0] == "info" {
				position, uintErr := parsePiecePosition(input)
				if uintErr == 1 {
					fmt.Println("Error: bad position input.")
				} else {
					board.GetPiece(position[0], position[1]).PrintMovesGrid()
				}
			}
		}
	}
}

func parsePiecePosition(inputSlice []string) ([2]uint8, uint8) {
	posSlice := strings.Split(inputSlice[1], ";")
	parsedPos0, err0 := strconv.ParseUint(posSlice[0], 10, 8)
	parsedPos1, err1 := strconv.ParseUint(posSlice[1], 10, 8)
	if err0 != nil || err1 != nil {
		return [2]uint8{}, 1
	}
	return [2]uint8{uint8(parsedPos0), uint8(parsedPos1)}, 0
}

func parseMovementString(movementSlice []string) ([2]uint8, [2]uint8) {
	fromPos := strings.Split(movementSlice[1], ";")
	toPos := strings.Split(movementSlice[2], ";")
	fromPos0Uint, err0 := strconv.ParseUint(fromPos[0], 10, 8)
	fromPos1Uint, err1 := strconv.ParseUint(fromPos[1], 10, 8)
	toPos0Uint, err2 := strconv.ParseUint(toPos[0], 10, 8)
	toPos1Uint, err3 := strconv.ParseUint(toPos[1], 10, 8)
	if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
		return [2]uint8{}, [2]uint8{}
	}
	return [2]uint8{uint8(fromPos0Uint), uint8(fromPos1Uint)}, [2]uint8{uint8(toPos0Uint), uint8(toPos1Uint)}

}
