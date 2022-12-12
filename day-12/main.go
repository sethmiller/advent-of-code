package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Node struct {
	row    int
	column int
	value  rune
}

func (n *Node) eq(o *Node) bool {
	return n.row == o.row && n.column == o.column
}

func (from *Node) checkAndAppend(neighbors []*Node, to *Node) []*Node {
	// I'm stupid
	dest := to.value
	if dest == 'E' {
		dest = 'z'
	}

	if from.value == 'S' || dest <= from.value+1 {
		return append(neighbors, to)
	}

	return neighbors
}

func (n *Node) neighbors(grid [][]*Node) []*Node {
	neighbors := []*Node{}
	rows := len(grid)
	columns := len(grid[0])

	if n.row > 0 {
		// up
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

	if n.row < rows-1 {
		// left
		neighbors = n.checkAndAppend(neighbors, grid[n.row+1][n.column])
	}

	return neighbors
}

func (n Node) String() string {
	return fmt.Sprintf("%d,%d", n.row, n.column)
}

func buildPath(end *Node, steps map[*Node]*Node) []*Node {
	path := []*Node{}
	for step, _ := end, true; step != nil; step, _ = steps[step] {
		path = append(path, step)
	}

	return path
}

func distance(n *Node, end *Node) float64 {
	// y = mx + b
	// a^2 + b^2 = c^2
	// sqrt(a^2+b^2)

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
			fmt.Println("Done")
			return buildPath(current, steps)
		}

		delete(toVisit, current)
		for _, neighbor := range current.neighbors(grid) {
			score := scores[current] + 1 // Always one away
			// `!found` is the same as infinity
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

	panic("oh no! We're lost")
}

func main() {
	grid := [][]*Node{}
	row := 0
	var start, end *Node

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		grid = append(grid, []*Node{})
		line := scanner.Text()
		for i, r := range line {
			grid[row] = append(grid[row], &Node{
				row:    row,
				column: i,
				value:  r,
			})

			if r == 'S' {
				start = grid[row][i]
			} else if r == 'E' {
				end = grid[row][i]
			}
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	path := aStar(start, end, grid, distance)

	visualize := make([][]rune, len(grid))
	for i := range visualize {
		visualize[i] = make([]rune, len(grid[0]))
	}

	steps := len(path)
	for i, step := range path {
		if i < steps-1 {
			prev := path[i+1]
			r := 'V'
			if prev.row > step.row {
				r = '^'
			} else if prev.column > step.column {
				r = '<'
			} else if prev.column < step.column {
				r = '>'
			}

			visualize[prev.row][prev.column] = r
		}
	}

	for r, row := range visualize {
		for c, ch := range row {
			if ch == 0 {
				fmt.Print(string(grid[r][c].value))
			} else {
				fmt.Print(string(ch))
			}
		}
		fmt.Println()
	}

	// Subtract 1 for since the start square doesn't count
	fmt.Printf("Steps: %d\n", len(path)-1)
}
