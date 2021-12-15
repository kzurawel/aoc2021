package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSum(s [3]int) int {
	var r int
	for _, v := range s {
		r += v
	}
	return r
}

func main() {
	var incs int
	var readings [3]int
	f, err := os.Open("./day1input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i := 0; i < 3; i++ {
		scanner.Scan()
		r, err := strconv.Atoi(scanner.Text())
		check(err)
		readings[i] = r
	}

	fmt.Println(readings)
	fmt.Println(getSum(readings))

	for scanner.Scan() {
		sum := getSum(readings)
		readings[0] = readings[1]
		readings[1] = readings[2]
		current, err := strconv.Atoi(scanner.Text())
		check(err)
		readings[2] = current

		newSum := getSum(readings)

		if newSum > sum {
			incs++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(incs)
}
