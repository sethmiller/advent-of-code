package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	name      string
	rate      int
	neighbors []*Node
}

func (n *Node) eq(o *Node) bool {
	return n.name == o.name
}

func buildPath(end *Node, steps map[*Node]*Node) []*Node {
	path := []*Node{}
	for step, _ := end, true; step != nil; step, _ = steps[step] {
		path = append(path, step)
	}

	return path
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
func aStar(start *Node, end *Node, seconds int, h func(p *Node, s int) float64) []*Node {
	toVisit := map[*Node]interface{}{start: nil}
	steps := map[*Node]*Node{}
	scores := map[*Node]int{start: 0}
	weightedScores := map[*Node]float64{start: h(start, seconds)}

	for len(toVisit) > 0 {
		current := smallest(toVisit, weightedScores)
		if end.eq(current) {
			return buildPath(current, steps)
		}

		delete(toVisit, current)
		for _, neighbor := range current.neighbors {
			score := scores[current] + 1 // Always one away
			// `!found` is the same as infinity
			if currentScore, found := scores[neighbor]; !found || score < currentScore {
				steps[neighbor] = current
				scores[neighbor] = score
				weightedScores[neighbor] = float64(score) + h(neighbor)
				if _, found := toVisit[neighbor]; !found {
					toVisit[neighbor] = nil
				}
			}
		}
	}

	return nil
}

func value(p *Node, stepsRemaining int) float64 {
	return float64(p.rate * stepsRemaining)
}

var coords = regexp.MustCompile(`^Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)

func main() {
	nodes := map[string]*Node{}
	var start *Node

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matches := coords.FindStringSubmatch(line)
		name := matches[1]
		rate, _ := strconv.Atoi(matches[2])
		neighbors := strings.Split(matches[3], ", ")

		node, found := nodes[name]

		if !found {
			node = &Node{
				name:      name,
				rate:      rate,
				neighbors: make([]*Node, 0),
			}

			nodes[name] = node
			if start == nil {
				start = node
			}
		}

		for _, neighbor := range neighbors {
			n, found := nodes[neighbor]

			if !found {
				n = &Node{
					name:      neighbor,
					neighbors: make([]*Node, 0),
				}

				nodes[neighbor] = n
			}

			node.neighbors = append(node.neighbors, n)
		}

	}
	steps := aStar(start, nil, value)
	fmt.Println(nodes)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

}
