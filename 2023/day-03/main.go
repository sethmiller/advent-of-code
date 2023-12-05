package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func checkRow(line string, start, end int) bool {
	if len(line) == 0 {
		return true
	}

	for i := max(start-1, 0); i < min(end+1, len(line)-1); i++ {
		if line[i] != '.' {
			return true
		}
	}

	return false
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

		start := -1
		for i := 0; i < len(line); i++ {
			ch := int(line[i])
			if start < 0 && ch >= 48 && ch <= 57 {
				start = i
			} else if start >= 0 && (ch < 48 || ch > 57 || i == len(line)-1) {
				end := i
				if ch >= 48 && ch <= 57 {
					end = i + 1
				}
				segment := line[start:end]
				// previous row
				if row > 0 && checkRow(rows[row-1], start, end) {
					sum += atoi(segment)
				} else if (start > 0 && line[start-1] != '.') || (end < len(line)-1 && line[end] != '.') {
					// current row
					sum += atoi(segment)
				} else if !last && checkRow(rows[row+1], start, end) {
					// next row
					sum += atoi(segment)
				}

				start = -1
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
