package main

import (
	"bufio"
	"fmt"
	"os"
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

	rights := map[int]int{}

	for _, v := range right {
		rights[v]++
	}

	sum := 0
	for _, v := range left {
		sum += v * rights[v]
	}

	fmt.Println(sum)
}
