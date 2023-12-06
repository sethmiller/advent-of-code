package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func atoi(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func split(str string) []int {
	chunks := []int{}
	for _, str := range strings.Split(str, " ") {
		if len(str) > 0 {
			chunks = append(chunks, atoi(str))
		}
	}

	return chunks
}

type Chapter struct {
	name   string
	ranges []Range
}

func (c Chapter) Dest(src int) int {
	for _, r := range c.ranges {
		if r.In(src) {
			return r.Dest(src)
		}
	}

	return src
}

type Range struct {
	start int
	end   int
	dest  int
}

func (r Range) In(i int) bool {
	return i >= r.start && i <= r.end
}

func (r Range) Dest(src int) int {
	return r.dest + (src - r.start)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	seeds := split(scanner.Text()[6:])
	scanner.Scan()

	order := 0
	chapters := map[int]Chapter{}
	for scanner.Scan() {
		header := scanner.Text()
		names := strings.Split(header[:len(header)-5], "-to-")
		ch := Chapter{
			name:   names[1],
			ranges: []Range{},
		}

		for scanner.Scan() {
			values := split(scanner.Text())
			if len(values) == 0 {
				break
			}

			ch.ranges = append(ch.ranges, Range{
				start: values[1],
				end:   values[1] + values[2] - 1,
				dest:  values[0],
			})
		}

		chapters[order] = ch
		order++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	min := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println("Next...")
		start := seeds[i]
		length := seeds[i+1]
		// Too dumb to math it; do brute force
		// Probably want check whole blocks at a time instead of just one. Next year.
		for j := start; j < start+length; j++ {
			next := j
			for c := 0; c < len(chapters); c++ {
				ch := chapters[c]
				next = ch.Dest(next)
			}

			if next < min {
				min = next
			}
		}

	}

	fmt.Println(min)
}
