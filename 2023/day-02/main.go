package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors = map[string]int{
	"green": 13,
	"red":   12,
	"blue":  14,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

		fail := false
		rounds := strings.Split(parts[1], ";")
	round:
		for _, round := range rounds {
			values := strings.Split(round, ",")
			for _, value := range values {
				yup := strings.Split(strings.TrimSpace(value), " ")
				count, _ := strconv.Atoi(yup[0])
				color := yup[1]
				if count > colors[color] {
					fail = true
					break round
				}
			}
		}

		if !fail {
			sum += id
		}

	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

}
