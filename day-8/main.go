package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func clear(height rune, trees string) int {
	for i, h := range trees {
		if h >= height {
			return i + 1
		}
	}

	return len(trees)
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

	scenic := 0
	for r, row := range rowOriented {
		for c, height := range row {
			score := 1
			// left
			score *= clear(height, reverse(row[0:c]))
			// right
			score *= clear(height, row[c+1:])
			// up
			score *= clear(height, reverse(columnOriented[c][0:r]))
			// down
			score *= clear(height, columnOriented[c][r+1:])

			if score > scenic {
				scenic = score
			}
		}
	}

	fmt.Println(scenic)
}
