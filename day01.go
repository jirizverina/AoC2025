package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day01_input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	calculate(50, reader)
}

func calculate(start_pos int32, reader *bufio.Reader) {
	var curr_pos, stopped_at_zero, passed_zero int32 = start_pos, 0, 0

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

		is_positive := line[0] == 'R'

		rotate_by, err := strconv.ParseInt(line[1:(len(line)-1)], 10, 32)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Position: %d, IsPositive: %t, RotateBy: %d, PassedZero: %d, StoppedAtZero: %d\n", curr_pos, is_positive, rotate_by, passed_zero, stopped_at_zero)

		passed_zero += int32(rotate_by) / 100

		if is_positive {
			curr_pos += int32(rotate_by) % 100
			if curr_pos > 99 {
				curr_pos = curr_pos - 100
				passed_zero++
			}

		} else {
			prev_pos := curr_pos
			curr_pos -= int32(rotate_by) % 100
			if((curr_pos < 0 && prev_pos != 0) || curr_pos == 0) {
				passed_zero++
			}

			if curr_pos < 0 {
				curr_pos = curr_pos + 100
			}
		}

		if curr_pos == 0 {
			stopped_at_zero++
		}

		if eof {
			break
		}
	}

	fmt.Println()
	fmt.Println(stopped_at_zero)
	fmt.Println()
	fmt.Println(passed_zero)

}
