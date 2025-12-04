package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Day01() {
	fmt.Println("Day 01")

	f, err := os.Open("inputs/day01_input.txt")
	if err != nil {
		fmt.Println("Could not load input for day 01.")
		return;
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	calculateDay01(50, reader)
}

func calculateDay01(startPos int32, reader *bufio.Reader) {
	var currPos, stoppedAtZero, passedZero int32 = startPos, 0, 0

	for {
		eof := false

		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			eof = true
		}

		if len(line) == 0 {
			break
		}

		isPositive := line[0] == 'R'

		rotateBy, err := strconv.ParseInt(line[1:(len(line)-1)], 10, 32)
		if err != nil {
			panic(err)
		}

		passedZero += int32(rotateBy) / 100

		if isPositive {
			currPos += int32(rotateBy) % 100
			if currPos > 99 {
				currPos = currPos - 100
				passedZero++
			}

		} else {
			prevPos := currPos
			currPos -= int32(rotateBy) % 100
			if((currPos < 0 && prevPos != 0) || currPos == 0) {
				passedZero++
			}

			if currPos < 0 {
				currPos = currPos + 100
			}
		}

		if currPos == 0 {
			stoppedAtZero++
		}

		if eof {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", stoppedAtZero)
	fmt.Printf("Part 2: %d \n", passedZero)
	fmt.Println()
}
