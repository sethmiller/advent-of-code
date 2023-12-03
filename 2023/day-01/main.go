package main

import (
	"bufio"
	"fmt"
	"os"
)

var numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		current := 0
	outer:
		for i := 0; i < len(line); i++ {
			ch := int(line[i])
			if ch >= 48 && ch <= 57 {
				current += ch - 48
				break
			} else {
				for n := 1; n < len(numbers); n++ {
					if len(line) > i+len(numbers[n]) && line[i:i+len(numbers[n])] == numbers[n] {
						i += len(numbers[n])
						current += n
						break outer
					}
				}
			}
		}

	outer2:
		for i := len(line) - 1; i >= 0; i-- {
			ch := int(line[i])
			if ch >= 48 && ch <= 57 {
				current *= 10
				current += ch - 48
				break
			} else {
				for n := 1; n < len(numbers); n++ {
					start := i - len(numbers[n]) + 1
					end := i + 1
					if start >= 0 && line[start:end] == numbers[n] {
						i -= len(numbers[n])
						current *= 10
						current += n
						break outer2
					}
				}
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
