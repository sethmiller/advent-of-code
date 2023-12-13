package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coords struct {
	r int
	c int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func (co Coords) dist(o Coords) int {
	r := co.r - o.r
	c := co.c - o.c

	return abs(r) + abs(c)
}

type pair struct {
	from Coords
	to   Coords
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := []string{}
	var cols []bool
	for scanner.Scan() {
		line := scanner.Text()

		if len(cols) == 0 {
			cols = make([]bool, len(line))
		}

		grid = append(grid, line)
		empty := true
		for col, ch := range line {
			if ch == '#' {
				cols[col] = true
				empty = false
			}
		}

		if empty {
			grid = append(grid, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	stars := map[Coords]interface{}{}
	extra := 0
	for col, occupied := range cols {
		if !occupied {
			extra++
		}

		for row, str := range grid {
			if str[col] == '#' {
				co := Coords{r: row, c: col + extra}
				stars[co] = nil
			}
		}
	}

	dists := map[pair]int{}
	for from := range stars {
		for to := range stars {
			if from != to {
				dists[pair{from: from, to: to}] = from.dist(to)
			}
		}
	}

	sum := 0

	for _, d := range dists {
		sum += d
	}

	fmt.Println(sum / 2)
}
