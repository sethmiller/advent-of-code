package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value  rune
	row    int
	column int
	used   bool
}

func (from *Node) checkAndAppend(neighbors []*Node, to *Node) []*Node {
	if from.value == to.value {
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

func (n *Node) walls(grid [][]*Node) int {
	return 4 - len(n.neighbors(grid))
}

func walk(n *Node, grid [][]*Node) map[*Node]interface{} {
	queue := []*Node{n}
	group := map[*Node]interface{}{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if _, exists := group[node]; exists {
			continue
		}

		queue = append(queue, node.neighbors((grid))...)
		node.used = true
		group[node] = nil
	}

	return group
}

func main() {
	grid := [][]*Node{}
	row := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		grid = append(grid, make([]*Node, len(line)))
		for col, ch := range line {
			grid[row][col] = &Node{
				value:  ch,
				row:    row,
				column: col,
			}
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	regions := []map[*Node]interface{}{}
	for _, nodes := range grid {
		for _, node := range nodes {
			if node.used {
				continue
			}

			region := walk(node, grid)
			regions = append(regions, region)
		}
	}

	total := 0
	for _, region := range regions {
		walls := 0
		for node := range region {
			walls += node.walls(grid)
		}

		total += len(region) * walls
	}

	fmt.Println(len(regions), total)
}
