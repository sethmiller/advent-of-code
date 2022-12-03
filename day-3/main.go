package main

import (
	"bufio"
	"fmt"
	"os"
)

func runeToValue(r rune) int {
	if r > 96 {
		return int(r) - 96
	}

	return int(r) - 38
}

func main() {
	total := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		found := map[rune]interface{}{}
		line := scanner.Text()
		first := line[0 : len(line)/2]
		second := line[len(line)/2:]

		for _, letter := range first {
			found[letter] = nil
		}

		for _, letter := range second {
			if _, exists := found[letter]; exists {
				total += runeToValue(letter)
				break
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Printf("Value: %d\n", total)
}
