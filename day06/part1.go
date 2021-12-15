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

	var fish []int
	for i := 0; i < len(inputs); i++ {
		num, err := strconv.ParseInt(inputs[i], 10, 64)
		if err != nil {
			panic(err)
		}
		fish = append(fish, num)
	}

	for i := 0; i < 80; i++ {
		var newfish []int
		for j := 0; j < len(fish); j++ {
			cur := fish[j]
			if cur == 0 {
				cur = 6
				newfish = append(newfish, cur)
				newfish = append(newfish, 8)
			} else {
				cur--
				newfish = append(newfish, cur)
			}
		}
		fish = newfish
	}

	fmt.Println(len(fish))
}
