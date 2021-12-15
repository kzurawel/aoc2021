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

func main() {
	var incs, lastRead int

	f, err := os.Open("./day1input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		check(err)

		if lastRead == 0 {
			lastRead = current
			continue
		}

		if current > lastRead {
			incs++
		}

		lastRead = current
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(incs)
}
