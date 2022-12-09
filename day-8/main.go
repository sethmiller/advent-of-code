package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func clear(height rune, trees string) bool {
	for _, h := range trees {
		if h >= height {
			return false
		}
	}

	return true
}

func main() {
	rowOriented := []string{}
	columnOriented := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		trees := strings.Split(line, "")

		for c, tree := range trees {
			if len(columnOriented) <= c {
				// Should just be hit on the first row
				columnOriented = append(columnOriented, "")
			}
			columnOriented[c] += tree
		}

		rowOriented = append(rowOriented, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	visible := 0
	rowCount := len(rowOriented)
	colCount := len(columnOriented)
	for r, row := range rowOriented {
		for c, height := range row {
			if c == 0 || r == 0 || r == rowCount-1 || c == colCount-1 {
				// Edges
				visible++
			} else {
				// left
				if clear(height, row[0:c]) {
					visible++
					continue
				}
				// right
				if clear(height, row[c+1:]) {
					visible++
					continue
				}
				// up
				if clear(height, columnOriented[c][0:r]) {
					visible++
					continue
				}
				// down
				if clear(height, columnOriented[c][r+1:]) {
					visible++
					continue
				}

			}
		}
	}

	fmt.Println(visible)
}
