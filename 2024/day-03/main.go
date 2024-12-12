package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var muls = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		matches := muls.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])

			sum += r * l
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)

}
