package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	parts := strings.Split(line, ",")

	sum := 0
	for _, part := range parts {
		cv := 0
		for _, ch := range part {
			cv = (cv + int(ch)) * 17
			cv = cv % 256
		}
		sum += cv
	}

	fmt.Println(sum)
}
