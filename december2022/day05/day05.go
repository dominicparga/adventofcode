package day05

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseStack(stackContent string) []stack[string] {
	// Input might look like
	//     [C]
	//     [V] [R]
	// [Z] [Q] [F]
	//  1   2   3

	stackContent = strings.ReplaceAll(stackContent, "    ", "[.] ")
	stackContent = strings.ReplaceAll(stackContent, " ", "")
	stackContent = strings.ReplaceAll(stackContent, "[", "")
	stackContent = strings.ReplaceAll(stackContent, "]", "")
	stackRowList := strings.Split(stackContent, "\n")
	stackRowList = stackRowList[:len(stackRowList)-1]

	// .C.
	// .VR
	// ZQF

	stackList := []stack[string]{}
	for i := 0; i < len(stackRowList[0]); i++ {
		stackList = append(stackList, stack[string]{top: nil, size: 0})
	}
	for i := len(stackRowList) - 1; i >= 0; i-- {
		row := stackRowList[i]
		for col, char := range row {
			str := string(char)
			if str != "." {
				stackList[col].Push(str)
			}
		}
	}

	return stackList
}

func parseMoves(moveStrList string) []move {
	moveList := []move{}
	for _, moveStr := range strings.Split(moveStrList, "\n") {
		moveStrList := strings.Split(moveStr, " ")
		// move 1 from 5 to 6
		count, _ := strconv.Atoi(moveStrList[1])
		fromIdx, _ := strconv.Atoi(moveStrList[3])
		toIdx, _ := strconv.Atoi(moveStrList[5])
		// input indices begin with 1
		moveList = append(moveList, move{
			fromIdx: fromIdx - 1,
			toIdx:   toIdx - 1,
			count:   count,
		})
	}
	return moveList
}

func Run(args []string) error {
	config := parseCmdline(args)
	var err error
	stackAndMovesByteList, err := os.ReadFile(config.stackAndMovesFilepath)
	if err != nil {
		return err
	}
	stackAndMovesContentList := strings.Split(string(stackAndMovesByteList), "\n\n")
	stackList := parseStack(stackAndMovesContentList[0])
	moveList := parseMoves(stackAndMovesContentList[1])

	for _, move := range moveList {
		stackList[move.fromIdx].MoveTo(&stackList[move.toIdx], move.count)
	}

	finalString := ""
	for _, stack := range stackList {
		finalString += stack.top.value
	}
	fmt.Println("Tops of all stacks:", finalString)

	return err
}
