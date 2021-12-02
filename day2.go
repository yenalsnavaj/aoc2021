package main

import (
	"bufio"
	"fmt"
	"os"
)

func day2(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var x, d int
	for scanner.Scan() {
		line := scanner.Text()

		// scan command and param
		var cmd string
		var param int
		fmt.Sscanf(line, "%s %d", &cmd, &param)
		switch cmd {
		case "forward":
			x += param
		case "up":
			d -= param
		case "down":
			d += param
		}
	}

	fmt.Printf("day2:\t\tposition = %d, depth = %d, result = %d\n", x, d, x*d)
}

func day2_pt2(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var x, d, a int
	for scanner.Scan() {
		line := scanner.Text()

		// scan command and param
		var cmd string
		var param int
		fmt.Sscanf(line, "%s %d", &cmd, &param)
		switch cmd {
		case "forward":
			x += param
			d += a * param
		case "up":
			a -= param
		case "down":
			a += param
		}
	}

	fmt.Printf("day2_pt2:\tposition = %d, depth = %d, aim = %d, result = %d\n", x, d, a, x*d)
}
