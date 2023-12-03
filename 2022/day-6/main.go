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
	offset := 14

	for scanner.Scan() {
		line = scanner.Text()
		break
	}

	length := len(line)

	for index := range line[:length-offset] {
		if !hasDupes(line[index : index+offset]) {
			fmt.Printf("Offset: %d\n", index+offset)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
}
