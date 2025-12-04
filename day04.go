package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Day04() {
	fmt.Println("\nDay 04")
	f, err := os.Open("inputs/day04_input.txt")
	if err != nil {
		fmt.Println("Could not read input.")
		return
	}
	defer f.Close()

	// scanner := bufio.NewScanner(f)
	// day04Part1(scanner)
	day04Part2(f)
}

func day04Part1(scanner *bufio.Scanner) {

	var lineFirst, lineSecond, lineThird string
	var sum int

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		panic("Boom!")
	}

	lineFirst = scanner.Text()
	_break := false

	for {
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			_break = true
		}
		lineSecond = scanner.Text()
		sum += giveResult(&lineFirst, &lineThird, &lineSecond)
		if _break {
			break
		}

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			_break = true
		}

		lineThird = scanner.Text()
		sum += giveResult(&lineSecond, &lineFirst, &lineThird)
		if _break {
			break
		}

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			_break = true
		}

		lineFirst = scanner.Text()
		sum += giveResult(&lineThird, &lineSecond, &lineFirst)

		if _break {
			break
		}
	}

	//TODO solve again

	fmt.Println(sum)

}

func giveResult(currentLine *string, prevLine *string, nextLine *string) int {
	if currentLine == nil || *currentLine == "" {
		panic("Boom!")
	}

	var result int
	var lineSize int = len(*currentLine)

	for i, r := range *currentLine {
		if r != '@' {
			continue
		}

		var paper int

		if prevLine != nil && *prevLine != "" {
			if i > 0 && (*prevLine)[i-1] == '@' {
				paper++
			}
			if (*prevLine)[i] == '@' {
				paper++
			}
			if i < lineSize-1 && (*prevLine)[i+1] == '@' {
				paper++
			}
		}

		if i > 0 && (*currentLine)[i-1] == '@' {
			paper++
		}
		if i < lineSize-1 && (*currentLine)[i+1] == '@' {
			paper++
		}

		if nextLine != nil && *nextLine != "" {
			if i > 0 && (*nextLine)[i-1] == '@' {
				paper++
			}
			if (*nextLine)[i] == '@' {
				paper++
			}
			if i < lineSize-1 && (*nextLine)[i+1] == '@' {
				paper++
			}
		}

		if paper < 4 {
			result++
		}
	}

	return result
}

func day04Part2(file *os.File) {

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var columnCount int
	for ; bytes[columnCount] != '\n'; columnCount++ {
	}
	columnCount++

	var result int

	for {
		keepRunning := false

		for i, r := range bytes {
			if r != '@' {
				continue
			}

			var papersAround int
			column := i % columnCount

			if i > 0 && bytes[i-1] == '@' {
				papersAround++
			}

			if column < columnCount-1 && bytes[i+1] == '@' {
				papersAround++
			}

			if i > columnCount {
				if column > 0 && bytes[i-columnCount-1] == '@' {
					papersAround++
				}

				if bytes[i-columnCount] == '@' {
					papersAround++
				}

				if column < columnCount-1 && bytes[i-columnCount+1] == '@' {
					papersAround++
				}
			}

			//NOTE: nejspis off by 1 error
			if i < len(bytes)-columnCount {
				if column > 0 && bytes[i+columnCount-1] == '@' {
					papersAround++
				}

				if bytes[i+columnCount] == '@' {
					papersAround++
				}

				if column < columnCount-1 && bytes[i+columnCount+1] == '@' {
					papersAround++
				}

			}

			if papersAround < 4 {
				bytes[i] = '.'
				result++
				keepRunning = true
			}

		}
		if !keepRunning {
			break
		}
	}

	fmt.Printf("Part 2: %d", result)
}
