package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	row       int
	column    int
	beenThere map[string]interface{}
}

func (p position) adjacent(other position) bool {
	deltaRow := p.row - other.row
	deltaColumn := p.column - other.column

	return math.Abs(float64(deltaRow)) <= 1 && math.Abs(float64(deltaColumn)) <= 1
}

func (p *position) move(direction string) {
	switch direction {
	case "U":
		p.row++
	case "D":
		p.row--
	case "R":
		p.column++
	case "L":
		p.column--
	}
}

func (p *position) moveToward(target position) {
	deltaRow := target.row - p.row
	deltaColumn := target.column - p.column
	if math.Abs(float64(deltaRow)) > 1 || math.Abs(float64(deltaColumn)) > 1 {
		if deltaRow < 0 {
			p.row--
		}
		if deltaRow > 0 {
			p.row++
		}
		if deltaColumn > 0 {
			p.column++
		}
		if deltaColumn < 0 {
			p.column--
		}

		p.beenThere[fmt.Sprintf("%d,%d", p.row, p.column)] = nil
	}
}

func (p position) visited() int {
	return len(p.beenThere)
}

func main() {
	head := position{}
	tail := position{
		beenThere: map[string]interface{}{"0,0": nil},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		move := strings.Split(line, " ")
		direction := move[0]
		distance, _ := strconv.Atoi(move[1])

		for i := 0; i < distance; i++ {
			head.move(direction)
			tail.moveToward(head)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Printf("Visited: %d", tail.visited())
}
