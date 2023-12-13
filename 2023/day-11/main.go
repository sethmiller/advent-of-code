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
	rows := []bool{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)

		if len(cols) == 0 {
			cols = make([]bool, len(line))
		}

		occupied := false
		for col, ch := range line {
			if ch == '#' {
				cols[col] = true
				occupied = true
			}
		}

		rows = append(rows, occupied)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	expansion := 999_999
	stars := map[Coords]interface{}{}
	extraCols := 0
	for col, occupied := range cols {
		if !occupied {
			extraCols++
			continue
		}

		extraRows := 0
		for row, str := range grid {
			if !rows[row] {
				extraRows++
				continue
			}

			if str[col] == '#' {
				co := Coords{r: row + (extraRows * expansion), c: col + (extraCols * expansion)}
				stars[co] = nil
			}
		}
	}

	dists := map[pair]int{}
	for from := range stars {
		for to := range stars {
			if _, ok := dists[pair{from: from, to: to}]; !ok && from != to {
				dists[pair{from: from, to: to}] = from.dist(to)
				dists[pair{from: to, to: from}] = 0
			}
		}
	}

	sum := 0
	for _, d := range dists {
		sum += d
	}

	fmt.Println(sum)
}
