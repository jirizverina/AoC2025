package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func Day03() {
	fmt.Println("\nDay 03")

	f, err := os.Open("inputs/day03_input.txt")
	if err != nil {
		fmt.Println("Could not load input for day 03.")
		return;
	}
	defer f.Close()
	
	reader := bufio.NewReader(f)
	_ = reader

	var sumFirstPart, sumSecondPart int64 = 0, 0

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
			break;
		}

		line = line[:len(line)-1]
		
		//First part
		{
			var first, second byte = 0, 0
			var firstIdx int = -1

			for i := len(line) - 2; i >= 0; i-- {
				b := line[i]

				if(b >= first) {
					first = b
					firstIdx = i
				}
			}

			for i:= len(line) - 1; i > firstIdx; i-- {
				b := line[i]
				if b > second {
					second = b
				}
			}

			sumFirstPart += (int64(first - '0') * 10) + int64(second - '0')
		}

		//Second part
		{
			var minIdx int = -1

			var numberToAdd int64 = 0
			for currIdx := range 12 {
				var value byte = 0

				minIdx2 := minIdx
				for i := len(line) - 12 + currIdx; i > minIdx2; i-- {
					b := line[i]
					
					// fmt.Printf("Current idx: %d, Value: %d, MinIdx: %d, New value: %d, I: %d\n", currIdx, value - '0', minIdx, b - '0', i)
					if b >= value {
						value = b
						minIdx = i
					}
				}

				numberToAdd += int64(value - '0') * int64(math.Pow10(11 - currIdx))
				value = 0
			}

			// fmt.Printf("Line: %s, NumberToAdd: %d\n\n", line, numberToAdd)
			sumSecondPart += numberToAdd;
		}


		if eof {
			break;
		}
	}

	fmt.Printf("Part 1: %d\n", sumFirstPart)
	fmt.Printf("Part 2: %d\n", sumSecondPart)
}
