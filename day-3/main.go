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
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()

		runes := []map[rune]interface{}{
			make(map[rune]interface{}),
			make(map[rune]interface{}),
			make(map[rune]interface{}),
		}

		for pos, bag := range []string{first, second, third} {
			for _, letter := range bag {
				runes[pos][letter] = nil
			}
		}

		for letter := range runes[0] {
			_, inBag2 := runes[1][letter]
			_, inBag3 := runes[2][letter]

			if inBag2 && inBag3 {
				total += runeToValue(letter)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Printf("Value: %d\n", total)
}
