package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var muls = regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	on := true
	for scanner.Scan() {
		line := scanner.Text()

		matches := muls.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				on = true
				fmt.Println("on")
				continue
			}
			if match[0] == "don't()" {
				on = false
				fmt.Println("off")
				continue
			}

			if !on {
				fmt.Printf("skipping %s\n", match[0])
				continue
			}

			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])

			fmt.Printf("%s, %d*%d = %d\n", match[0], l, r, l*r)

			sum += r * l
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)

}
