package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isNum(ch byte) bool {
	return ch >= 48 && ch <= 57
}

func checkRow(line string, pos int) []int {
	center := line[pos]
	left := []byte{}
	right := []byte{}
	for i := pos + 1; i < len(line); i++ {
		if isNum(line[i]) {
			right = append(right, line[i])
		} else {
			break
		}
	}
	for i := pos - 1; i >= 0; i-- {
		if isNum(line[i]) {
			left = append([]byte{line[i]}, left...)
		} else {
			break
		}
	}

	if isNum(center) {
		return []int{atoi(fmt.Sprintf("%s%c%s", left, center, right))}
	}

	vals := []int{}
	if len(left) > 0 {
		vals = append(vals, atoi(string(left)))
	}
	if len(right) > 0 {
		vals = append(vals, atoi(string(right)))
	}

	return vals
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rows := []string{
		scanner.Text(),
	}
	sum := 0
	row := 0
	last := false
	for !last {
		if scanner.Scan() {
			rows = append(rows, scanner.Text())
		} else {
			last = true
		}

		line := rows[row]

		for i := 0; i < len(line); i++ {
			if line[i] == '*' {
				vals := []int{}
				if row > 0 {
					vals = append(vals, checkRow(rows[row-1], i)...)
				}

				vals = append(vals, checkRow(line, i)...)

				if !last {
					vals = append(vals, checkRow(rows[row+1], i)...)
				}

				if len(vals) == 2 {
					sum += (vals[0] * vals[1])
				}
			}
		}

		row++
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

}
