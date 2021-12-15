package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	inputbytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	input := strings.TrimSpace(string(inputbytes))
	inputs := strings.Split(input, ",")

	var fish [9]int64
	for i := 0; i < len(inputs); i++ {
		num, err := strconv.ParseInt(inputs[i], 10, 64)
		if err != nil {
			panic(err)
		}
		fish[num]++
	}

	fmt.Println("starting fish are", fish)

	for i := 0; i < 256; i++ {
		var newfish [9]int64
		newfish[8] = fish[0]
		for i := 0; i < 8; i++ {
			newfish[i] = fish[i+1]
		}
		newfish[6] += fish[0]

		fish = newfish
		fmt.Println("fish is now", fish)

		var fishcount int64
		for i := 0; i < 9; i++ {
			fishcount += fish[i]
		}
		fmt.Println("day", i+1, "has", fishcount, "fish")
	}

	fmt.Println(len(fish))
}
