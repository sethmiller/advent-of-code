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

const target = 2_000_000

func check(s *Sensor) *[2]int {
	d := max(s.row, target) - min(s.row, target)
	if d <= s.distance {
		width := ((s.distance - d) * 2) + 1
		fmt.Printf("  sr(%d) tr(%d) sd(%d) d(%d) w(%d)\n", s.row, target, s.distance, d, width)
		half := (width - 1) / 2
		start := s.column - half
		end := s.column + half
		return &[2]int{start, end}
	}

	return nil
}

func overlap(l1 *[2]int, l2 *[2]int) *[2]int {
	// l1 start is inside l2
	if l1[0] >= l2[0] && l1[0] <= l2[1] {
		// is fully contained
		if l1[1] <= l2[1] {
			return l2
		}

		return &[2]int{l2[0], l1[1]}
	}

	// l2 start inside l1
	if l2[0] >= l1[0] && l2[0] <= l1[1] {
		// is fully contained
		if l2[1] <= l1[1] {
			return l1
		}

		return &[2]int{l1[0], l2[1]}
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

	lines := []*[2]int{}
	for _, s := range sensors {
		if line := check(s); line != nil {
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
		}
	}

	covered := 0
	for _, line := range lines {
		fmt.Println(line)
		covered += line[1] - line[0]
	}

	fmt.Printf("Covered: %d", covered)

}
