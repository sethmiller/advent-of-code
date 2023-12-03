package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		current := 0
		for i := 0; i < len(line); i++ {
			ch := int(line[i])
			if ch >= 48 && ch <= 57 {
				current += ch - 48
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			ch := int(line[i])
			if ch >= 48 && ch <= 57 {
				current *= 10
				current += ch - 48
				break
			}
		}
		sum += current
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)
}
