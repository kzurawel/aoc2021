package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkSlice(slice []string, pos int, o2 bool) []string {
	if len(slice) == 1 {
		return slice
	}

	// find most common bit
	var zeroes, ones int
	for _, data := range slice {
		char := strings.Split(data, "")[pos]
		switch char {
		case "0":
			zeroes++
		case "1":
			ones++
		}
	}

	var criteria string
	if o2 == true {
		if zeroes > ones {
			criteria = "0"
		}
		if ones > zeroes || ones == zeroes {
			criteria = "1"
		}
	}
	if o2 == false {
		if zeroes < ones || ones == zeroes {
			criteria = "0"
		}
		if ones < zeroes {
			criteria = "1"
		}
	}

	output := make([]string, 0, 600)
	for _, data := range slice {
		char := strings.Split(data, "")[pos]
		if char == criteria {
			output = append(output, data)
		}
	}

	return output
}

func main() {
	var zeroes, ones [12]int
	var gamma, epsilon strings.Builder
	var o2s, co2s []string
	var o2, co2 string

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		o2s = append(o2s, s.Text())
		co2s = append(co2s, s.Text())
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

	fmt.Println("checking o2s", len(o2s))

	output := checkSlice(o2s, 0, true)
	for i := 1; i < 12; i++ {
		output = checkSlice(output, i, true)
		if len(output) == 1 {
			fmt.Println("on pass", i+1, "o2 value is", output[0])
			o2 = output[0]
			break
		}
	}

	fmt.Println("checking co2s", len(co2s))

	output2 := checkSlice(co2s, 0, false)
	for j := 1; j < 12; j++ {
		output2 = checkSlice(output2, j, false)
		if len(output2) == 1 {
			fmt.Println("on pass", j+1, "co2 value is", output2[0])
			co2 = output2[0]
			break
		}
	}

	o2num, err := strconv.ParseInt(o2, 2, 64)
	if err != nil {
		panic(err)
	}
	co2num, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("o2 is", o2num, "co2 is", co2num)
	fmt.Println("final value", o2num*co2num)
}
