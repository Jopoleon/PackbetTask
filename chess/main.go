package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("you should specify knight's position on desk in format [a-h][1-8] \n")
	}
	position := os.Args[1]
	rightPlaceFormat, err := regexp.MatchString("^[a-h][1-8]$", position)
	if err != nil {
		log.Fatal("cant match knight position on desk to string, error: ", err)
	}
	if !rightPlaceFormat {
		log.Fatalln("Wrong position format. You should specify position in format [a-h][1-8] (for example b4, e3, etc...)")
	}
	moves := knightMoves(position)
	for _, move := range moves {
		fmt.Println(move)
	}

}

//knightMoves returns slice of string with what turns can knight make
func knightMoves(position string) []string {
	col, row := int(position[0]-96), 9-int(position[1]-48)
	var positions []string
	if col-2 > 0 && row+1 < 9 {
		positions = append(positions, intToPosition(col-2, row+1))
	}
	if col-2 > 0 && row-1 > 0 {
		positions = append(positions, intToPosition(col-2, row-1))
	}
	if col-1 > 0 && row+2 < 9 {
		positions = append(positions, intToPosition(col-1, row+2))
	}
	if col-1 > 0 && row-2 > 0 {
		positions = append(positions, intToPosition(col-1, row-2))
	}

	if col+2 < 9 && row-1 > 0 {
		positions = append(positions, intToPosition(col+2, row-1))
	}
	if col+2 < 9 && row+1 < 9 {
		positions = append(positions, intToPosition(col+2, row+1))
	}
	if col+1 < 9 && row+2 < 9 {
		positions = append(positions, intToPosition(col+1, row+2))
	}
	if col+1 < 9 && row-2 > 0 {
		positions = append(positions, intToPosition(col+1, row-2))
	}
	return positions
}

func intToPosition(col int, row int) string {
	return string([]byte{byte(col + 96), byte(9 - row + 48)}[:])
}
