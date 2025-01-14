package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dir = int

const (
	North Dir = iota
	East
	South
	West
)

type Node struct {
	x   int
	y   int
	dir Dir
}

func delta(dir Dir) []int {
	switch dir {
	case North:
		return []int{0, -1}
	case East:
		return []int{1, 0}
	case South:
		return []int{0, 1}
	case West:
		return []int{-1, 0}
	}

	panic("oops")
}

func print(grid [][]byte) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

func copyGrid(_grid [][]byte) [][]byte {
	grid := make([][]byte, len(_grid))
	for i, line := range _grid {
		grid[i] = make([]byte, len(line))
		copy(grid[i], line)
	}

	return grid
}

func loops(grid [][]byte, x, y int) bool {
	visited := map[Node]interface{}{}
	dir := North
	width := len(grid[0])
	height := len(grid)
	for {
		deltas := delta(dir)
		nextX := x + deltas[0]
		nextY := y + deltas[1]
		node := Node{x: nextX, y: nextY, dir: dir}
		if _, exists := visited[node]; exists {
			return true
		}

		visited[node] = nil
		if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
			return false
		}

		next := grid[nextY][nextX]
		if next == '#' || next == 'O' {
			dir = (dir + 1) % 4
			continue
		}

		x = nextX
		y = nextY
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := [][]byte{}
	x, y := 0, 0
	dir := North
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		grid = append(grid, []byte(line))
		for i, val := range line {
			if val == '^' {
				x = i
				y = len(grid)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Printf("Found start at (%d, %d)\n", x, y)

	orig := copyGrid(grid)
	startX := x
	startY := y
	width := len(grid[0])
	height := len(grid)
	count := 0
	for {
		deltas := delta(dir)
		nextX := x + deltas[0]
		nextY := y + deltas[1]
		if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
			grid[y][x] = 'X'
			blocked := copyGrid(orig)
			blocked[y][x] = 'O'
			if loops(blocked, startX, startY) {
				count++
				// print(blocked)
				// fmt.Println()
			}
			break
		}

		next := grid[nextY][nextX]
		if next == '#' {
			dir = (dir + 1) % 4
			continue
		}

		current := grid[y][x]
		if current != 'X' {
			grid[y][x] = 'X'
			blocked := copyGrid(orig)
			blocked[y][x] = 'O'
			if loops(blocked, startX, startY) {
				count++
				// print(blocked)
				// fmt.Println()
			}
		}
		x = nextX
		y = nextY
	}

	fmt.Println(count)
}
