package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day3(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	// count bits
	var zeros [12]int
	var ones [12]int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		for i, c := range line {
			switch c {
			case '0':
				zeros[i]++
			case '1':
				ones[i]++
			}
		}

	}

	// build result
	var gamma, epsilon int
	for i := 0; i < 12; i++ {

		if zeros[i] > ones[i] {
			gamma += 0 << (11 - i)
			epsilon += 1 << (11 - i)
		} else {
			gamma += 1 << (11 - i)
			epsilon += 0 << (11 - i)
		}

	}

	fmt.Printf("day3:\t\tgamma = %d, epsilon = %d, result = %d\n", gamma, epsilon, gamma*epsilon)
}

func day3_pt2(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	// count bits and capture lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// oxygen generator
	ogr := calculate(lines, true)
	csr := calculate(lines, false)

	fmt.Printf("day3_pt2:\toxygen generator rating = %d, CO2 scrubber rating = %d, result = %d\n", ogr, csr, ogr*csr)
}

// return slice containing matches
func filter(lines []string, pos int, match byte) (result []string) {

	for _, line := range lines {
		if line[pos] == match {
			result = append(result, line)
		}
	}
	return
}

func countBits(lines []string) (zeros, ones [12]int) {
	for _, line := range lines {
		for i, c := range line {
			switch c {
			case '0':
				zeros[i]++
			case '1':
				ones[i]++
			}
		}

	}
	return
}

func calculate(lines []string, isogr bool) int64 {
	var value int64
	var err error
	r := lines
	for i := 0; ; {
		zeros, ones := countBits(r)
		var c byte
		if ones[i] == zeros[i] {
			if isogr {
				c = '1'
			} else {
				c = '0'
			}
		} else if ones[i] > zeros[i] {
			if isogr {
				c = '1'
			} else {
				c = '0'
			}
		} else {
			if isogr {
				c = '0'
			} else {
				c = '1'
			}
		}
		r = filter(r, i, c)
		if len(r) == 1 {
			value, err = strconv.ParseInt(r[0], 2, 32)
			if err != nil {
				panic(err.Error())
			}
			break
		}
		i++
	}
	return value
}
