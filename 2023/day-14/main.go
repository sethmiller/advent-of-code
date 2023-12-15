package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, line)
		for offset, ch := range line {
			if ch == 'O' {
				// fmt.Println("Found O at ", len(grid)-1, offset)
				for i := len(grid) - 1; i >= 0; i-- {
					// fmt.Println("Looking at ", string(grid[i-1][offset]), i, offset)
					if i == 0 || grid[i-1][offset] != '.' {
						// fmt.Println("moving", len(grid)-1, offset, "to", i, offset)
						source := []byte(grid[len(grid)-1])
						source[offset] = '.'
						grid[len(grid)-1] = string(source)

						target := []byte(grid[i])
						target[offset] = 'O'
						grid[i] = string(target)
						break
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println()

	sum := 0
	height := len(grid)
	for offset, l := range grid {
		found := 0
		for _, ch := range l {
			if ch == 'O' {
				found++
			}
		}

		sum += (height - offset) * found
	}

	fmt.Println(sum)

}
