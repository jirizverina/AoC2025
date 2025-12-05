package day05

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Range struct {
	from uint64
	to uint64
	isValid bool
}

type readingState uint8

const (
	StateRanges readingState = iota
	StateIngredients
)

func Solve() {
	fmt.Println("\nDay 05")

	f, err := os.Open("inputs/day05_input.txt")
	if err != nil {
		fmt.Println("Could not load input for day 05.")
		return;
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ranges []*Range
	var freshIngredientsCount, allFreshIngredientsCount uint64

	state := StateRanges

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line == "\n" {
			state = StateIngredients
			continue
		}

		switch state {
		case StateRanges:
			r, err := parseRange(&line)
			if err != nil {
				fmt.Printf("Line is %q\n", line)
				panic(err)
			}

			ranges = append(ranges, r)
		case StateIngredients:
			ingredient, err := parseIngredients(&line)
			if err != nil {
				panic(err)
			}

			if isIngredientFresh(ranges, ingredient) {
				freshIngredientsCount++
			}
			
		default: panic("Unknown state.")
		}
		
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	mergeRanges(ranges)

	for _, r := range ranges {
		if r.isValid {
			allFreshIngredientsCount += uint64(r.to - r.from) + 1
		}
	}

	fmt.Printf("Part 1: %d\n", freshIngredientsCount)
	fmt.Printf("Part 2: %d\n", allFreshIngredientsCount)
}

func mergeRanges(ranges []*Range) {
	for _, r1 := range ranges {
		if !r1.isValid {
			continue
		}

		for _, r2 := range ranges {
			if r1 == r2 || !r2.isValid {
				continue
			}

			if (r1.from <= r2.from && r1.to >= r2.from) || (r2.from <= r1.from && r2.to >= r1.from) {
				r2.from = min(r1.from, r2.from)
				r2.to = max(r1.to, r2.to)
				r1.isValid = false
			}
		}
	}
}

func isIngredientFresh(freshIngredients []*Range, ingredient uint64) bool {
	for _, r := range freshIngredients {
		if(r.from <= ingredient && r.to >= ingredient) {
			return true
		}
	}

	return false
}

func parseIngredients(line *string) (uint64, error) {
	i, err := strconv.ParseInt(*line, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(i), nil
}

func parseRange(line *string) (*Range, error) {
	for i, r := range *line {
		if r != '-' {
			continue
		}

		from, err := strconv.ParseInt((*line)[0:i], 10, 64)
		if err != nil {
			return nil, err
		}

		to, err := strconv.ParseInt((*line)[i+1:len(*line)], 10, 64)
		if err != nil {
			return nil, err
		}

		result := Range {from: uint64(from), to: uint64(to), isValid: true }
		return &result, nil
	}

	return nil, errors.New("Could not parse range: '-' not found.")
}

