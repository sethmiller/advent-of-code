package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type record struct {
	stone string
	depth int
}

var memo = map[record]int{}

func trimOrZero(str string) string {
	trimmed := strings.TrimLeft(str, "0")

	if len(trimmed) == 0 {
		return "0"
	}
	return trimmed
}

func mul(str string, times int) string {
	i, _ := strconv.Atoi(str)

	return fmt.Sprintf("%d", i*times)
}

func dfs(stone string, depth int) int {
	if depth == 0 {
		return 1
	}
	if val, exists := memo[record{stone: stone, depth: depth}]; exists {
		return val
	}

	var val int

	if stone == "0" {
		val = dfs("1", depth-1)
	} else if len(stone)%2 == 0 {
		val = dfs(trimOrZero((stone[len(stone)/2:])), depth-1) + dfs(stone[0:len(stone)/2], depth-1)
	} else {
		val = dfs(mul(stone, 2024), depth-1)
	}

	memo[record{stone: stone, depth: depth}] = val

	return val
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fullset := strings.Split(scanner.Text(), " ")
	fmt.Println(fullset)

	length := 0
	for _, next := range fullset {
		fmt.Println(next)
		length += dfs(next, 75)
	}

	fmt.Println()
	fmt.Println(length)
}
