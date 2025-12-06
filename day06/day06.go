package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type mathOperation uint8

const (
	Add mathOperation = iota
	Multiply
)

type Expression struct {
	numbers   []int64
	operation mathOperation
}



func Solve() {
	fmt.Println("\nDay 06")
	part1()
	part2()
}

func part1() {
	f, err := os.Open("inputs/day06_input.txt")
	if err != nil {
		fmt.Println("Could not load input for day 06.")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var result int64
	expressions := make([]*Expression, 0, 8)

	for ok := scanner.Scan(); ok; {
		var currentColumn int
		line := scanner.Text()

		ok = scanner.Scan()

		for i := 0; i < len(line); i++ {
			r := line[i]

			if r == ' ' {
				continue
			}


			if ok {
				startOfNumber := i

				for ; i < len(line); i++ {
					if line[i] == ' ' {
						break
					}
				}

				num, err := strconv.ParseInt(line[startOfNumber:i], 10, 64)
				if err != nil {
					panic(err)
				}

				if currentColumn >= len(expressions) {
					newExpression := new(Expression)
					newExpression.numbers = []int64{num}
					expressions = append(expressions, newExpression)
				} else {
					numbers := &expressions[currentColumn].numbers
					*numbers = append(*numbers, num)
				}
			} else {
				switch r {
				case '+':
					expressions[currentColumn].operation = Add
				case '*':
					expressions[currentColumn].operation = Multiply
				default:
					panic("Unknown operation")
				}
			}

			currentColumn++
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, r := range expressions {
		var expressionResult int64
		if r.operation == Multiply {
			expressionResult = 1
		}

		for _, n := range r.numbers {
			switch r.operation {
			case Add:
				expressionResult += n
			case Multiply:
				expressionResult *= n
			default:
				panic("Unknown operation")
			}
		}

		result += expressionResult
	}


	fmt.Printf("Part 1: %d\n", result)
}

type Number struct {
	number  int64
	isValid bool
}

func part2() {
	f, err := os.Open("inputs/day06_input.txt")
	if err != nil {
		fmt.Println("Could not load input for day 06.")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := make([]*Number, 0, 8)
	var result int64

	var lastAddedIdx int

	for ok := scanner.Scan(); ok; {
		line := scanner.Text()

		ok = scanner.Scan()

		for i := 0; i < len(line); i++ {
			r := line[i]

			if ok {
				numbers = append(numbers, new(Number))
				if r == ' ' {
					continue
				}

				num := int64(r - '0')

				value := numbers[i]
				value.number = value.number*10 + num
				value.isValid = true
			} else {
				var expressionResult int64
				switch r {
				case '+':
					for j := lastAddedIdx; j < len(numbers); j++ {
						value := numbers[j]
						if !value.isValid {
							lastAddedIdx++
							break
						}

						expressionResult += value.number
						lastAddedIdx++
					}
				case '*':
					expressionResult = 1
					for j := lastAddedIdx; j < len(numbers); j++ {
						value := numbers[j]
						if !value.isValid {
							lastAddedIdx++
							break
						}

						expressionResult *= value.number
						lastAddedIdx++
					}
				case ' ': continue
				default:
					panic("Unknown operation")
				}
				result += expressionResult
			}
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part 2: %d\n", result)

}
