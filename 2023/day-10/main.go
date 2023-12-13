package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dir = int

const (
	North Dir = iota
	South
	West
	East
)

type Coords struct {
	r int
	c int
}

var dirs = map[byte][]int{
	'|': {North, South},
	'-': {West, East},
	'L': {North, East},
	'J': {North, West},
	'7': {South, West},
	'F': {South, East},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := []string{}
	start := Coords{}
	rowCount := 0
	for scanner.Scan() {
		str := scanner.Text()
		data = append(data, str)
		for c, b := range str {
			if b == 'S' {
				start.r = rowCount
				start.c = c
			}
		}

		rowCount++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	starters := []Coords{}
	// North
	if start.r > 0 {
		r := start.r - 1
		c := start.c
		b := data[r][c]
		if b == '|' || b == '7' || b == 'F' {
			starters = append(starters, Coords{r: r, c: c})
		}
	}
	// West
	if start.c > 0 {
		r := start.r
		c := start.c - 1
		b := data[r][c]
		if b == '-' || b == 'L' || b == 'F' {
			starters = append(starters, Coords{r: r, c: c})
		}
	}
	// South
	if start.r < rowCount-1 {
		r := start.r + 1
		c := start.c
		b := data[r][c]
		if b == '|' || b == 'L' || b == 'J' {
			starters = append(starters, Coords{r: r, c: c})
		}
	}
	// East
	if start.c < len(data[0])-1 {
		r := start.r
		c := start.c + 1
		b := data[r][c]
		if b == '-' || b == '7' || b == 'J' {
			starters = append(starters, Coords{r: r, c: c})
		}
	}

	if len(starters) != 2 {
		panic("Oops")
	}

	for _, s := range starters {
		stack := []Coords{s}
		visited := map[Coords]interface{}{start: nil}
		// out:
		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if _, ok := visited[curr]; ok {
				continue
			}

			visited[curr] = nil

			value := data[curr.r][curr.c]
			// fmt.Printf("%+v -> %c\n", curr, value)
			if d, ok := dirs[value]; ok {
				for _, dir := range d {
					r := curr.r
					c := curr.c
					switch dir {
					case North:
						r--
					case South:
						r++
					case East:
						c++
					case West:
						c--
					default:
						panic("oops")
					}

					co := Coords{r: r, c: c}
					// fmt.Printf("Looking at %+v\n", co)
					if r >= 0 && c >= 0 && r < len(data) && c < len(data[0]) {
						if _, ok := visited[co]; !ok {
							// fmt.Println("Adding")
							stack = append(stack, co)
						}
					}
				}
			}
		}

		for r, row := range data {
			for c, v := range row {
				if _, ok := visited[Coords{r: r, c: c}]; ok {
					fmt.Print("*")
				} else {
					fmt.Printf("%c", v)
				}
			}
			fmt.Println()
		}

		fmt.Println(len(visited) / 2)
	}
}
