package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operand = int

const (
	Multiply Operand = iota
	Add
)

func atoi(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)

	return i
}

func atoiAll(strs []string) []int64 {
	ints := make([]int64, len(strs))
	for i, val := range strs {
		ints[i] = atoi(val)
	}

	return ints
}

func resolves(answer int64, parts []int64, total int64) bool {
	product := total * parts[0]
	sum := total + parts[0]

	if len(parts) == 1 {
		if product == answer || sum == answer {
			return true
		}
	} else {
		if resolves(answer, parts[1:], product) {
			return true
		}

		if resolves(answer, parts[1:], sum) {
			return true
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		halves := strings.Split(line, ": ")
		answer := atoi(halves[0])
		parts := atoiAll(strings.Split(halves[1], " "))

		if resolves(answer, parts[1:], parts[0]) {
			fmt.Println(answer, parts)
			sum += int(answer)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)
}
