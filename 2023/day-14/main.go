package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func rotate(blob []string) []string {
	height := len(blob[0])
	width := len(blob)
	out := make([]string, height)
	for i := 0; i < height; i++ {
		str := make([]byte, width)
		for j := 0; j < width; j++ {
			str[j] = blob[width-j-1][i]
		}

		out[i] = string(str)
	}

	return out
}

func print(board []string, n int) {
	for i := 4 - n; i > 0; i-- {
		board = rotate(board)
	}
	for _, l := range board {
		fmt.Println(l)
	}

	fmt.Println()
}

const (
	North = iota
	West
	South
	East
)

type Record struct {
	dir   int
	board []string
}

func (r Record) String() string {
	str := fmt.Sprintf("%d", r.dir)
	str += strings.Join(r.board, ",")

	return str
}

func score(board []string, n int) int {
	for i := 4 - n; i > 0; i-- {
		board = rotate(board)
	}
	sum := 0
	height := len(board)
	for offset, l := range board {
		found := 0
		for _, ch := range l {
			if ch == 'O' {
				found++
			}
		}

		sum += (height - offset) * found
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	board := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		board = append(board, line)
	}

	visited := map[string]interface{}{}
	scores := map[int]int{}
	steps := 0
	dir := North

	for {
		fmt.Println("Next", dir, steps, score(board, dir))
		for height, line := range board {
			for offset, ch := range line {
				if ch == 'O' {
					for i := height; i >= 0; i-- {
						if i == 0 || board[i-1][offset] != '.' {
							source := []byte(board[height])
							source[offset] = '.'
							board[height] = string(source)

							target := []byte(board[i])
							target[offset] = 'O'
							board[i] = string(target)
							break
						}
					}
				}
			}
		}

		r := Record{dir: dir, board: board}
		if _, ok := visited[r.String()]; dir == North && ok {
			fmt.Println("looped", dir, steps, score(board, dir))
			break
		}

		visited[r.String()] = nil
		scores[steps] = score(board, dir)
		dir = (dir + 1) % 4
		board = rotate(board)
		steps++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	i := 4_000_000_000 % 41

	fmt.Println(scores[i])

}
