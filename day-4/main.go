package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		a := strings.Split(parts[0], "-")
		b := strings.Split(parts[1], "-")

		aLower, _ := strconv.Atoi(a[0])
		aUpper, _ := strconv.Atoi(a[1])
		bLower, _ := strconv.Atoi(b[0])
		bUpper, _ := strconv.Atoi(b[1])

		if (aLower >= bLower && aUpper <= bUpper) || (bLower >= aLower && bUpper <= aUpper) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Printf("Count %d\n", count)

}
