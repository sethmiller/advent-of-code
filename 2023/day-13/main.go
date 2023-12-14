package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func rotate(blob []string) []string {
	out := make([]string, len(blob[0]))
	for i := 0; i < len(blob[0]); i++ {
		str := ""
		for j := 0; j < len(blob); j++ {
			str = str + string(blob[j][i])
		}

		out[i] = str
	}

	return out
}

func check(blob []string) int {
	for i := 1; i < len(blob); i++ {
		if blob[i-1] == blob[i] {
			mirror := true
			for j := 1; j <= min(len(blob)-i, i)-1; j++ {
				if blob[i-j-1] != blob[i+j] {
					mirror = false
					break
				}
			}

			if mirror {
				return i
			}
		}
	}

	return 0
}

func reflection(blob []string) int {
	return (check(blob) * 100) + check(rotate(blob))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	blob := []string{}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			sum += reflection(blob)
			blob = []string{}
		} else {
			blob = append(blob, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	sum += reflection(blob)

	fmt.Println(sum)
}
