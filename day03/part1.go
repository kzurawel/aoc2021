package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var zeroes, ones [12]int
	var gamma, epsilon strings.Builder

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		c := strings.Split(s.Text(), "")
		for i, bit := range c {
			switch bit {
			case "0":
				zeroes[i]++
			case "1":
				ones[i]++
			}
		}
	}

	for i := 0; i < 12; i++ {
		if zeroes[i] > ones[i] {
			gamma.WriteRune('0')
			epsilon.WriteRune('1')
		} else {
			gamma.WriteRune('1')
			epsilon.WriteRune('0')
		}
	}

	fmt.Println(gamma.String())
	fmt.Println(epsilon.String())

	gammaNum, err := strconv.ParseInt(gamma.String(), 2, 0)
	if err != nil {
		panic(err)
	}
	epsilonNum, err := strconv.ParseInt(epsilon.String(), 2, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(gammaNum * epsilonNum)
}
