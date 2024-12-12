package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		parts := strings.Split(line, "   ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])

		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for row, _ := range left {
		sum += int(math.Abs(float64(left[row] - right[row])))
	}

	fmt.Println(sum)
}
