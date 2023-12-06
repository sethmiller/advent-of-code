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

func dist(hold, duration int) int {
	// dist = hold * (hold - duration)
	return hold * (duration - hold)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	times := split(scanner.Text()[10:])
	scanner.Scan()
	distances := split(scanner.Text()[10:])

	product := 1
	for index, time := range times {
		middle := int(math.Ceil(float64(time) / float64(2)))
		count := 0
		for t := middle; t < time; t++ {
			d := dist(t, time)
			if d > distances[index] {
				count++
			} else {
				break
			}
		}

		for t := middle - 1; t > 1; t-- {
			d := dist(t, time)
			if d > distances[index] {
				count++
			} else {
				break
			}
		}

		product *= count
	}

	fmt.Println(product)
}
