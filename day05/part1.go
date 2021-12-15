package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func (c Coord) New(s string) Coord {
	vals := strings.Split(s, ",")
	xNum, _ := strconv.Atoi(vals[0])
	yNum, _ := strconv.Atoi(vals[1])
	return Coord{
		x: xNum,
		y: yNum,
	}
}

func main() {
	var grid [1000][1000]int
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var pt1, pt2 Coord
	for s.Scan() {
		line := strings.Split(s.Text(), " -> ")
		pt1 = pt1.New(line[0])
		pt2 = pt2.New(line[1])

		// vertical / horizontal only
		if pt1.x != pt2.x && pt1.y != pt2.y {
			continue
		}

		// draw line to grid
		if pt1.x == pt2.x {
			// vertical line, change y
			if pt1.y > pt2.y {
				for i := pt2.y; i <= pt1.y; i++ {
					grid[pt1.x][i]++
				}
			} else {
				for i := pt1.y; i <= pt2.y; i++ {
					grid[pt1.x][i]++
				}
			}
		} else {
			// horizontal line, change x
			if pt1.x > pt2.x {
				for i := pt2.x; i <= pt1.x; i++ {
					grid[i][pt1.y]++
				}
			} else {
				for i := pt1.x; i <= pt2.x; i++ {
					grid[i][pt1.y]++
				}
			}
		}
	}

	// calculate results
	var points int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				points++
			}
		}
	}

	fmt.Println("points with 2+ lines:", points)
}
