package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type line struct {
	pairs []*pair
}

type pair struct {
	row    int
	column int
}

type grid struct {
	minRow    int
	maxRow    int
	minColumn int
	maxColumn int
	contents  *[][]rune
}

func (g *grid) set(r, c int, ch rune) {
	column := c - g.minColumn

	(*g.contents)[r][column] = ch
}

func (g *grid) get(r, c int) rune {
	column := c - g.minColumn
	return (*g.contents)[r][column]
}

func (g *grid) drop(r, c int) bool {
	next := r
	for {
		if next > g.maxRow {
			return false
		}
		if c <= g.minColumn {
			g.grow(left)
		} else if c > g.maxColumn {
			g.grow(right)
		}

		switch g.get(next, c) {
		case 0:
		default:
			if g.get(next, c-1) == 0 {
				c--
			} else if g.get(next, c+1) == 0 {
				c++
			} else {
				g.set(next-1, c, 'o')
				return true
			}
		}

		next++
	}
}

type side bool

const (
	left  = false
	right = true
)

func (g *grid) grow(s side) {
	for r, row := range *g.contents {
		if s == right {
			(*g.contents)[r] = append(row, 0)
		} else {
			(*g.contents)[r] = append([]rune{0}, row...)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func makeGrid(lines []*line, g *grid) *grid {
	height := g.maxRow + 1
	width := g.maxColumn - g.minColumn
	grid := make([][]rune, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]rune, width+1)
	}

	g.contents = &grid
	g.set(0, 500, '+')

	for _, line := range lines {
		last := line.pairs[0]
		for i := 1; i < len(line.pairs); i++ {
			next := line.pairs[i]
			if last.row == next.row {
				start := min(last.column, next.column)
				end := max(last.column, next.column)

				for c := start; c <= end; c++ {
					g.set(last.row, c, '#')
				}
			} else {
				start := min(last.row, next.row)
				end := max(last.row, next.row)
				for r := start; r <= end; r++ {
					g.set(r, last.column, '#')
				}
			}
			last = next
		}
	}

	return g
}

func main() {
	lines := []*line{}
	minRow, minColumn := math.MaxInt, math.MaxInt
	maxRow, maxColumn := 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()

		pairs := []*pair{}
		num := 0
		j := 0
		for j < len(str) {
			ch := str[j]
			switch ch {
			case ' ':
				minRow = min(minRow, num)
				maxRow = max(maxRow, num)
				pairs[len(pairs)-1].row = num
				num = 0
				j += 3
			case ',':
				minColumn = min(minColumn, num)
				maxColumn = max(maxColumn, num)
				pairs = append(pairs, &pair{
					column: num,
				})
				num = 0
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num *= 10
				num += int(ch) - 0x30
			}
			j++

			if j == len(str) {
				pairs[len(pairs)-1].row = num
			}
		}

		lines = append(lines, &line{
			pairs: pairs,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	g := makeGrid(lines, &grid{
		minRow:    minRow,
		maxRow:    maxRow,
		minColumn: minColumn,
		maxColumn: maxColumn,
	})

	count := 0
	for {
		if !g.drop(0, 500) {
			fmt.Println("done")
			break
		}
		count++
	}

	for r, row := range *g.contents {
		fmt.Printf("%3d ", r)
		for _, column := range row {
			if column == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(column))
			}
		}
		fmt.Println()
	}

	fmt.Println("Drops: ", count)
}
