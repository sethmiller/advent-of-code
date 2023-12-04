package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		rounds := strings.Split(parts[1], ";")
		maxes := map[string]int{}
		for _, round := range rounds {
			values := strings.Split(round, ",")
			for _, value := range values {
				yup := strings.Split(strings.TrimSpace(value), " ")
				count, _ := strconv.Atoi(yup[0])
				color := yup[1]
				maxes[color] = max(maxes[color], count)
			}
		}

		product := 1
		for _, v := range maxes {
			product *= v
		}

		sum += product
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

}
