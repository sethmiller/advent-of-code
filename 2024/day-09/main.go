package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const EMPTY = -1

func atoi(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func atoiAll(strs []string) []int {
	ints := make([]int, len(strs))
	for i, val := range strs {
		ints[i] = atoi(val)
	}

	return ints
}

func checksum(vals []int) int {
	sum := 0
	for i, val := range vals {
		if val == EMPTY {
			break
		}
		sum += i * val
	}

	return sum
}

func expand(disk []int) []int {
	expanded := []int{}
	for i := 0; i < len(disk); {
		index := i / 2
		size := disk[i]
		i++
		blank := disk[i]
		i++

		for j := 0; j < size; j++ {
			expanded = append(expanded, index)
		}

		for j := 0; j < blank; j++ {
			expanded = append(expanded, EMPTY)
		}
	}

	return expanded
}

func defrag(disk []int) []int {
	last := 0
	defraged := make([]int, len(disk))
	for i := range defraged {
		defraged[i] = -1
	}
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == EMPTY {
			continue
		}
		for j := last; j <= i; j++ {
			if disk[j] == EMPTY {
				defraged[j] = disk[i]
				last = j + 1
				break
			}
			defraged[j] = disk[j]
		}
	}

	return defraged
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	line += "0"

	disk := atoiAll(strings.Split(line, ""))

	expanded := expand(disk)
	fmt.Println(expanded)
	defraged := defrag(expanded)
	fmt.Println(defraged)
	sum := checksum(defraged)

	fmt.Println(sum)

}
