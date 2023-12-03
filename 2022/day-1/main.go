package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	tallies := []int{0}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			tallies = append(tallies, 0)
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			(tallies[len(tallies)-1]) += calories
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	count := len(tallies)
	sort.Ints(tallies)
	total := tallies[count-3] + tallies[count-2] + tallies[count-1]

	fmt.Printf("Most calories: %d\n", tallies[count-1])
	fmt.Printf("Sum of three most calories: %d\n", total)
}
