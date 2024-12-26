package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func getOrCreate(points map[rune][]Point, r rune) []Point {
	p, exists := points[r]
	if exists {
		return p
	}

	return make([]Point, 0)
}

func add(points map[Point]interface{}, p Point, w int, h int) {
	if p.x < 0 || p.x >= w || p.y < 0 || p.y >= h {
		return
	}

	points[p] = nil
}

func addAll(points map[Point]interface{}, p Point, dx int, dy int, w int, h int) {
	x := p.x
	y := p.y

	for x >= 0 && x < w && y >= 0 && y < h {
		add(points, Point{x: x, y: y}, w, h)
		x -= dx
		y -= dy
	}

	x = p.x
	y = p.y
	for x >= 0 && x < w && y >= 0 && y < h {
		add(points, Point{x: x, y: y}, w, h)
		x += dx
		y += dy
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := []string{}
	row := 0
	antennae := map[rune][]Point{}
	antinodes := map[Point]interface{}{}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		grid = append(grid, line)
		for i, ch := range line {
			if ch != '.' {
				points := getOrCreate(antennae, ch)
				antennae[ch] = append(points, Point{i, row})
			}
		}

		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	width := len(grid[0])
	height := row

	for _, points := range antennae {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				dx := points[j].x - points[i].x
				dy := points[j].y - points[i].y

				addAll(antinodes, points[j], dx, dy, width, height)
			}
		}
	}

	fmt.Println(len(antinodes))
}
