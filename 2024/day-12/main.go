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

type Dir = int

const (
	North Dir = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

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

	// left
	if n.column > 0 {
		neighbors = n.checkAndAppend(neighbors, grid[n.row][n.column-1])
	}

	// right
	if n.column < columns-1 {
		neighbors = n.checkAndAppend(neighbors, grid[n.row][n.column+1])
	}

	// down
	if n.row < rows-1 {
		neighbors = n.checkAndAppend(neighbors, grid[n.row+1][n.column])
	}

	return neighbors
}

func (n *Node) allNeighbors(grid [][]*Node) map[Dir]*Node {
	neighbors := map[Dir]*Node{
		North:     nil,
		NorthEast: nil,
		East:      nil,
		SouthEast: nil,
		South:     nil,
		SouthWest: nil,
		West:      nil,
		NorthWest: nil,
	}

	rows := len(grid)
	columns := len(grid[0])

	// up
	if n.row > 0 {
		neighbors[North] = grid[n.row-1][n.column]
	}

	// left
	if n.column > 0 {
		neighbors[West] = grid[n.row][n.column-1]
	}

	// right
	if n.column < columns-1 {
		neighbors[East] = grid[n.row][n.column+1]
	}

	// down
	if n.row < rows-1 {
		neighbors[South] = grid[n.row+1][n.column]
	}

	// up -> right
	if n.row > 0 && n.column < columns-1 {
		neighbors[NorthEast] = grid[n.row-1][n.column+1]
	}

	// up -> left
	if n.row > 0 && n.column > 0 {
		neighbors[NorthWest] = grid[n.row-1][n.column-1]
	}

	// down -> right
	if n.row < rows-1 && n.column < columns-1 {
		neighbors[SouthEast] = grid[n.row+1][n.column+1]
	}

	// down -> left
	if n.row < rows-1 && n.column > 0 {
		neighbors[SouthWest] = grid[n.row+1][n.column-1]
	}

	return neighbors
}

func external(a *Node, b *Node, diag *Node, source *Node) bool {
	return a != nil && b != nil && a.value == source.value && a.value == b.value && a.value != diag.value
}

func internal(a *Node, b *Node, source *Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a != nil && b == nil && a.value != source.value {
		return true
	}

	if a == nil && b != nil && b.value != source.value {
		return true
	}

	if a != nil && b != nil && a.value != source.value && b.value != source.value {
		return true
	}

	return false
}

func (n *Node) corners(grid [][]*Node) int {
	corners := 0
	a := n.allNeighbors(grid)

	// internal corners
	if internal(a[North], a[East], n) {
		corners++
	}
	if internal(a[East], a[South], n) {
		corners++
	}
	if internal(a[South], a[West], n) {
		corners++
	}
	if internal(a[West], a[North], n) {
		corners++
	}

	// external corners
	if external(a[North], a[East], a[NorthEast], n) {
		corners++
	}
	if external(a[East], a[South], a[SouthEast], n) {
		corners++
	}
	if external(a[South], a[West], a[SouthWest], n) {
		corners++
	}
	if external(a[West], a[North], a[NorthWest], n) {
		corners++
	}
	return corners
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
		corners := 0
		var value rune
		for node := range region {
			corners += node.corners(grid)
			value = node.value
		}

		fmt.Printf("%c %d * %d\n", value, len(region), corners)

		total += len(region) * corners
	}

	fmt.Println(len(regions), total)
}
