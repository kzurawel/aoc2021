package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var h, d, aim int
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		amt, err := strconv.Atoi(parts[1])
		check(err)

		switch parts[0] {
		case "forward":
			h += amt
			d += (amt * aim)
		case "up":
			aim -= amt
		case "down":
			aim += amt
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("h: %i, d: %i", h, d)
	fmt.Println(h * d)
}
