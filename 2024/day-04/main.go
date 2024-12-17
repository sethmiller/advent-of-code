package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

var xmas = regexp.MustCompile("XMAS")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rows := []string{}
	columns := []string{}
	var rotateLeft []string
	var rotateRight []string
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		rows = append(rows, line)

		height := (2 * len(line)) - 1
		middle := len(line) - 1
		row := len(rows) - 1

		if row == 0 {
			rotateLeft = make([]string, height)
			rotateRight = make([]string, height)
		}
		for i, ch := range line {
			if len(columns) <= i {
				columns = append(columns, "")
			}

			rotateLeftRow := (middle + row) - i
			rotateRightRow := row + i

			rotateLeft[rotateLeftRow] = rotateLeft[rotateLeftRow] + string(ch)
			rotateRight[rotateRightRow] = string(ch) + rotateRight[rotateRightRow]

			columns[i] = columns[i] + string(ch)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(rotateLeft)
	fmt.Println(rotateRight)

	count := 0
	for _, dir := range [][]string{rows, columns, rotateLeft, rotateRight} {
		for _, str := range dir {
			// fmt.Printf("Checking %s...\n", str)
			matches := xmas.FindAllStringSubmatch(str, -1)
			count += len(matches)

			// fmt.Printf("Checking %s...\n", reverse(str))
			matches = xmas.FindAllStringSubmatch(reverse(str), -1)
			count += len(matches)
		}
	}

	fmt.Println(count)

}
