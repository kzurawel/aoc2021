package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	data   [5][5]BoardSquare
	winner bool
}

func (b Board) Mark(num int) Board {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if b.data[r][c].value == num {
				b.data[r][c].marked = true
			}
		}
	}
	return b
}

func (b Board) CheckForWin() bool {
	var result bool
	// check each row
	for r := 0; r < 5; r++ {
		if b.data[r][0].marked == true && b.data[r][1].marked == true &&
			b.data[r][2].marked == true && b.data[r][3].marked == true &&
			b.data[r][4].marked == true {
			result = true
		}
	}

	// check each col
	for c := 0; c < 5; c++ {
		if b.data[0][c].marked == true && b.data[1][c].marked == true &&
			b.data[2][c].marked == true && b.data[3][c].marked == true &&
			b.data[4][c].marked == true {
			result = true
		}
	}

	return result
}

func (b Board) CalculateWinScore(multiplier int) int {
	var result int
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if b.data[r][c].marked == false {
				result += b.data[r][c].value
			}
		}
	}
	return result * multiplier
}

type BoardSquare struct {
	value  int
	marked bool
}

func countNonWinners(boards []Board) int {
	boardsLeft := len(boards)
	for i := 0; i < len(boards); i++ {
		if boards[i].winner {
			boardsLeft = boardsLeft - 1
		}
	}
	return boardsLeft
}

func main() {
	var boards []Board
	boards = append(boards, Board{})

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan()
	inputs := strings.Split(s.Text(), ",")
	s.Scan()

	row := 0
	for s.Scan() {
		if row == 0 {
			boards = append(boards, Board{})
		}
		if row != 5 {
			curRow := strings.Fields(s.Text())
			for c, val := range curRow {
				boards[len(boards)-1].data[row][c].value, _ = strconv.Atoi(val)
			}
			row++
		} else {
			row = 0
		}
	}

	// Boards structure built out, now time to check inputs
	var winningBoard Board
	var winningInput int

InputTest:
	for _, input := range inputs {
		in, _ := strconv.Atoi(input)
		for i := 0; i < len(boards); i++ {
			boards[i] = boards[i].Mark(in)
			winner := boards[i].CheckForWin()
			if winner {
				boards[i].winner = true
				count := countNonWinners(boards)
				if count == 0 {
					winningBoard = boards[i]
					winningInput = in
					break InputTest
				}
			}
		}
	}

	// found winning board
	fmt.Println(winningInput)
	fmt.Println(winningBoard)
	fmt.Println()
	fmt.Println("winning score:", winningBoard.CalculateWinScore(winningInput))
}
