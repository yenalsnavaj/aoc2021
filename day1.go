package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day1(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var p, c int
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err.Error())
		}
		if p != 0 && v > p {
			c++
		}
		p = v
	}

	fmt.Printf("day1:\t\t%d values increased\n", c)
}

func day1_pt2(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var values []int
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err.Error())
		}
		values = append(values, v)
	}

	var c int
	for i := 3; i < len(values); i++ {
		// window one is i-3 -> i-1
		// window two i-2 -> i
		sum1 := values[i-3] + values[i-2] + values[i-1]
		sum2 := values[i-2] + values[i-1] + values[i]
		if sum2 > sum1 {
			c++
		}
	}

	fmt.Printf("day1_pt2:\t%d values increased\n", c)
}
