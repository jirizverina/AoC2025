package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"math"
)

func Day02() {
	fmt.Println("Day 02")

	f, err := os.Open("inputs/day02_input.txt")
	if(err != nil) {
		panic(err)
	}
	
	reader := bufio.NewReader(f)
	_ = reader

	var firstResult, secondResult int64 = 0, 0

	for {
		eof := false
		firstIdString, err := reader.ReadString('-')
		if(err != nil) {
			if(err != io.EOF) {
				panic(err)
			}
			eof = true
		}

		if(len(firstIdString) == 0) {
			break;
		}

		firstId, err := strconv.Atoi(firstIdString[:len(firstIdString) - 1])
		if(err != nil) {
			panic(err)
		}

		secondIdString, err := reader.ReadString(',')
		if(err != nil) {
			if(err != io.EOF) {
				panic(err)
			}
			eof = true
		}

		if(len(secondIdString) == 0) {
			break;
		}

		secondId, err := strconv.Atoi(secondIdString[:len(secondIdString) - 1])
		if(err != nil) {
			panic(err)
		}

		for i:=firstId; i<=secondId; i++ {
			digits := int(math.Log10(float64(i))) + 1

			//first part
			if  digits & 1 == 0 {
				delimiter := int(math.Pow10(digits / 2))

				firstPart := i / delimiter
				secondPart := i - (firstPart * delimiter)

				if secondPart == firstPart {
					firstResult += int64(i)
				}
			}

			//second part
			if(digits < 2) {
				continue;
			}

			for j:= 1; j <= digits / 2; j++ {
				if digits % j != 0 {
					continue;
				}

				delimiter := int(math.Pow10(digits - j))

				needle := i / delimiter
				
				// if i == 446446 {
				// 	fmt.Printf("Delimimter: %d, Needle: %d, J: %d\n", delimiter, needle, j)
				// }

				isValid := true
				for k:= 1; k <= (digits - j) / j; k++ {
					next := (i % (delimiter / int(math.Pow10(j * (k - 1))))) / (delimiter / int(math.Pow10(j * k)))

					// if i == 446446 {
					// 	fmt.Printf("Needle: %d, Next: %d, J: %d, K: %d\n", needle, next, j, k)
					// }

					if needle != next {
						isValid = false;
						break;
					}
				}

				if isValid {
					secondResult += int64(i);
					// fmt.Printf("Found number: %d\n", i)
					break;
				}
			}
		}



		if(eof) {
			break;
		}
	}

	fmt.Printf("Part 1: %d\n", firstResult)
	fmt.Printf("Part 2: %d\n", secondResult)
}
