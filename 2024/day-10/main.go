package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Node struct {
	row    int
	column int
	value  int
}

func (n *Node) eq(o *Node) bool {
	return n.row == o.row && n.column == o.column
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

func buildPath(end *Node, steps map[*Node]*Node) []*Node {
	path := []*Node{}
	for step, _ := end, true; step != nil; step = steps[step] {
		path = append(path, step)
	}

	return path
}

func distance(n *Node, end *Node) float64 {
	a := math.Abs(float64(n.row - end.row))
	b := math.Abs(float64(n.column - end.column))

	return math.Sqrt((a*a + b*b))
}

func smallest(toVisit map[*Node]interface{}, m map[*Node]float64) *Node {
	var min *Node
	minScore := math.MaxFloat64
	for node, score := range m {
		_, found := toVisit[node]
		if found && (min == nil || score < minScore) {
			min = node
			minScore = score
		}
	}

	return min
}

// Based on the pseudo code in https://en.wikipedia.org/wiki/A*_search_algorithm
func aStar(start *Node, end *Node, grid [][]*Node, h func(p *Node, o *Node) float64) []*Node {
	toVisit := map[*Node]interface{}{start: nil}
	steps := map[*Node]*Node{}
	scores := map[*Node]int{start: 0}
	weightedScores := map[*Node]float64{start: h(start, end)}

	for len(toVisit) > 0 {
		current := smallest(toVisit, weightedScores)
		if end.eq(current) {
			return buildPath(current, steps)
		}

		delete(toVisit, current)
		for _, neighbor := range current.neighbors(grid) {
			score := scores[current] + 1
			if currentScore, found := scores[neighbor]; !found || score < currentScore {
				steps[neighbor] = current
				scores[neighbor] = score
				weightedScores[neighbor] = float64(score) + h(neighbor, end)
				if _, found := toVisit[neighbor]; !found {
					toVisit[neighbor] = nil
				}
			}
		}
	}

	return nil
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
			path := aStar(start, end, grid, distance)

			if path != nil {
				sum++
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
