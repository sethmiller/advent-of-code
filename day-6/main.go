package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasDupes(str string) bool {
	fmt.Println(str)
	set := map[rune]interface{}{}
	for _, ch := range str {
		if _, found := set[ch]; found {
			return true
		}

		set[ch] = nil
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	line := ""

	for scanner.Scan() {
		line = scanner.Text()
		break
	}

	length := len(line)

	for index := range line[:length-4] {
		if !hasDupes(line[index : index+4]) {
			fmt.Printf("Offset: %d\n", index+4)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
}
