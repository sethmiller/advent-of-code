package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func diff(nums []int) int {
	diffs := make([]int, len(nums)-1)
	nonzero := false
	for i := 1; i < len(nums); i++ {
		diffs[i-1] = nums[i] - nums[i-1]
		if diffs[i-1] != 0 {
			nonzero = true
		}
	}

	add := 0
	if nonzero {
		add = diff(diffs)
	}

	return add + diffs[len(diffs)-1]
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		data := split(scanner.Text())
		fmt.Println("> ", data)

		i := data[len(data)-1] + diff(data)
		sum += i
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)
}
