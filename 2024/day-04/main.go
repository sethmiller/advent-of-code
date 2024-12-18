package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(a rune, b rune) bool {
	return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := [][]rune{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		grid = append(grid, make([]rune, len(line)))

		for i, ch := range line {
			grid[row][i] = ch
		}

		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	count := 0
	height := len(grid)
	for i, line := range grid {
		width := len(line)
		for j, ch := range line {
			if ch == 'A' {
				if i == 0 || i == width-1 {
					continue
				}
				if j == 0 || j == height-1 {
					continue
				}

				ul := grid[i-1][j-1]
				ur := grid[i-1][j+1]
				ll := grid[i+1][j-1]
				lr := grid[i+1][j+1]

				if check(ul, lr) && check(ur, ll) {
					fmt.Printf("(%d, %d)\n", i, j)
					fmt.Printf("(%c, %c), (%c, %c)\n", ul, ur, ll, lr)
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
