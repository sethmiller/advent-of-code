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

type Sensor struct {
	row      int
	column   int
	beacon   *Beacon
	distance int
}

func (s *Sensor) setBeacon(b *Beacon) {
	s.beacon = b
	rows := max(b.row, s.row) - min(b.row, s.row)
	columns := max(b.column, s.column) - min(b.column, s.column)

	s.distance = rows + columns
}

type Beacon struct {
	row     int
	column  int
	sensors []*Sensor
}

type Pair struct {
	start int
	end   int
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

const lower = 0
const upper = 4_000_000

func check(s *Sensor, target int) *Pair {
	d := max(s.row, target) - min(s.row, target)
	if d <= s.distance {
		width := ((s.distance - d) * 2) + 1
		// fmt.Printf("  sr(%d) tr(%d) sd(%d) d(%d) w(%d)\n", s.row, target, s.distance, d, width)
		half := (width - 1) / 2
		start := s.column - half
		end := s.column + half
		// Bound the result by range within which we are searching
		pair := &Pair{start: max(lower, start), end: min(end, upper)}
		if start > upper || end < lower {
			// We intersected with the line outside of the range within which we are searching
			return nil
		}

		return pair
	}

	return nil
}

func overlap(l1 *Pair, l2 *Pair) *Pair {
	// l1 start is inside l2
	if l1.start >= l2.start && l1.start <= l2.end {
		// is fully contained
		if l1.end <= l2.end {
			return l2
		}

		return &Pair{start: l2.start, end: l1.end}
	}

	// l2 start is inside l1
	if l2.start >= l1.start && l2.start <= l1.end {
		// is fully contained
		if l2.end <= l1.end {
			return l1
		}

		return &Pair{start: l1.start, end: l2.end}
	}

	// Adjacent
	if l1.end+1 == l2.start {
		return &Pair{start: l1.start, end: l2.end}
	}

	if l2.end+1 == l1.start {
		return &Pair{start: l2.start, end: l1.end}
	}

	// No overlap
	return nil
}

var coords = regexp.MustCompile(`^[^x]*x=([-\d]*), y=([-\d]*)$`)

func main() {
	minRow, minColumn := math.MaxInt, math.MaxInt
	maxRow, maxColumn := 0, 0
	beacons := map[string]*Beacon{}
	sensors := map[string]*Sensor{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		halves := strings.Split(line, ":")

		m1 := coords.FindStringSubmatch(halves[0])
		r, _ := strconv.Atoi(m1[2])
		c, _ := strconv.Atoi(m1[1])
		minRow = min(minRow, r)
		maxRow = max(maxRow, r)
		minColumn = min(minColumn, c)
		maxColumn = max(maxColumn, c)
		key := fmt.Sprintf("%d,%d", r, c)

		sensor := Sensor{
			row:    r,
			column: c,
		}

		sensors[key] = &sensor

		m2 := coords.FindStringSubmatch(halves[1])
		r, _ = strconv.Atoi(m2[2])
		c, _ = strconv.Atoi(m2[1])
		minRow = min(minRow, r)
		maxRow = max(maxRow, r)
		minColumn = min(minColumn, c)
		maxColumn = max(maxColumn, c)

		key = fmt.Sprintf("%d,%d", r, c)

		if b, found := beacons[key]; found {
			(&sensor).setBeacon(b)

			b.sensors = append(b.sensors, &sensor)
		} else {
			beacon := Beacon{
				row:     r,
				column:  c,
				sensors: []*Sensor{&sensor},
			}
			beacons[key] = &beacon
			(&sensor).setBeacon(&beacon)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	lines := []*Pair{}
outer:
	for r := lower; r <= upper; r++ {
		for _, s := range sensors {
			if line := check(s, r); line != nil {
				l := 0
				for l < len(lines) {
					existing := lines[l]
					if joined := overlap(line, existing); joined != nil {
						// This line overlaps another. Remove the existing line.
						lines = append(lines[:l], lines[l+1:]...)
						// and update the line with the joined lines
						line = joined
					} else {
						l++
					}
				}
				lines = append(lines, line)

				if len(lines) == 1 && lines[0].start == lower && lines[0].end == upper {
					// fmt.Printf("Row (%d) is full\n", r)
					break
				}
			}
		}

		switch len(lines) {
		case 0:
			panic(fmt.Sprintf("Oh no (%d)\n", r))
		case 1:
			// if lines[0].start == lower && lines[0].end == upper {
			// 	fmt.Printf("Row (%d) is full\n", r)
			// }
		case 2:
			fmt.Printf("It's probably (%d)\n", r)
			for _, line := range lines {
				fmt.Printf("  (%d, %d)\n", line.start, line.end)
			}
			fmt.Printf("Answer: %d\n", r+((lines[0].end+1)*4_000_000))
			break outer
		default:
			panic(fmt.Sprintf("Weird (%d) - (%d)\n", r, len(lines)))
		}

		lines = make([]*Pair, 0)
	}
}
