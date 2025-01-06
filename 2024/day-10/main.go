package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	row    int
	column int
	value  int
}

func (from *Node) checkAndAppend(neighbors []*Node, to *Node) []*Node {
	if from.value == to.value-1 {
		return append(neighbors, to)
	}

	return neighbors
}

func (n *Node) neighbors(grid [][]*Node) []*Node {
	neighbors := []*Node{}
	rows := len(grid)
	columns := len(grid[0])

	// up
	if n.row > 0 {
		neighbors = n.checkAndAppend(neighbors, grid[n.row-1][n.column])
	}

	// down
	if n.column > 0 {
		neighbors = n.checkAndAppend(neighbors, grid[n.row][n.column-1])
	}

	// right
	if n.column < columns-1 {
		neighbors = n.checkAndAppend(neighbors, grid[n.row][n.column+1])
	}

	// left
	if n.row < rows-1 {
		neighbors = n.checkAndAppend(neighbors, grid[n.row+1][n.column])
	}

	return neighbors
}

func bfs(start *Node, end *Node, grid [][]*Node) int {
	queue := []*Node{start}
	paths := 0

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next == end {
			paths++
			continue
		}

		queue = append(queue, next.neighbors(grid)...)
	}

	return paths
}

func main() {
	grid := [][]*Node{}
	row := 0
	trailheads := map[*Node]interface{}{}
	ends := map[*Node]interface{}{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		grid = append(grid, []*Node{})
		for i, r := range line {
			value, _ := strconv.Atoi(string(r))
			grid[row] = append(grid[row], &Node{
				row:    row,
				column: i,
				value:  value,
			})

			if r == '0' {
				trailheads[grid[row][i]] = nil
			} else if r == '9' {
				ends[grid[row][i]] = nil
			}
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	sum := 0
	for start := range trailheads {
		for end := range ends {
			paths := bfs(start, end, grid)

			sum += paths
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
