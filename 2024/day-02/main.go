package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const min = 1
const max = 3

func remove(slice []int, s int) []int {
	n := make([]int, len(slice)-1)

	index := 0
	for i, v := range slice {
		if i != s {
			n[index] = v
			index++
		}
	}

	return n
}

func test(levels []int) bool {
	asc := levels[0]-levels[1] < 0
	for i := range levels[:len(levels)-1] {
		delta := levels[i] - levels[i+1]

		if (delta < 0 && !asc) || (delta > 0 && asc) {
			return false
		}

		absDelta := int(math.Abs(float64(delta)))
		if absDelta < min || absDelta > max {
			return false
		}

		if i == len(levels)-2 {
			return true
		}
	}

	fmt.Println("unsafe")
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		parts := strings.Split(line, " ")

		levels := []int{}
		for _, v := range parts {
			i, _ := strconv.Atoi(v)
			levels = append(levels, i)
		}

		if test(levels) {
			safe++
			continue
		}

		for i := range levels {
			next := remove(levels, i)
			if test(next) {
				safe++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(safe)

}
