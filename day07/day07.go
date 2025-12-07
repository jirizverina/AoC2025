package day07

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	fmt.Println("\nDay 07")
	part1()
}

func part1() {
	f, err := os.Open("inputs/day07_input.txt")
	if err != nil {
		fmt.Println("Could not load input for day 07.")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var result uint32
	lasers := make(map[int]any)
	
	for scanner.Scan() {
		line := scanner.Text()
		
		for i, r := range line {
			if len(lasers) == 0 {
				if r != 'S' {
					continue
				}

				lasers[i] = nil
				break;
			} else {
				if r != '^' {
					continue
				}

				if _, ok := lasers[i]; ok {
					delete(lasers, i)
					result++

					if i > 0 {
						lasers[i - 1] = nil
					}
					if i < len(line) - 1 {
						lasers[i + 1] = nil
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d", result)
}

