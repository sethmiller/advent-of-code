package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var total int = 0
var cycles = []int{20, 60, 100, 140, 180, 220}

func inc(i *int, pc *int) int {
	// for _, count := range cycles {
	// 	if *pc == count {
	// 		return *i * count
	// 	}
	// }

	pixel := *pc % 40
	if pixel == 0 {
		fmt.Println()
	}
	if *i-1 == pixel || *i == pixel || *i+1 == pixel {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	(*pc)++

	return 0
}

func main() {
	pc := new(int)
	val := new(int)
	*val = 1
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		switch parts[0] {
		case "noop":
			sum += inc(val, pc)
		case "addx":
			i, _ := strconv.Atoi(parts[1])
			sum += inc(val, pc)
			sum += inc(val, pc)
			*val += i
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
}
