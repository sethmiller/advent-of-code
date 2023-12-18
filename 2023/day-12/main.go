package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func split(str string) []int {
	chunks := []int{}
	for _, str := range strings.Split(str, ",") {
		if len(str) > 0 {
			chunks = append(chunks, atoi(str))
		}
	}

	return chunks
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func chunker(str string) []string {
	chunks := []string{}
	for _, str := range strings.Split(str, ".") {
		if len(str) > 0 {
			chunks = append(chunks, str)
		}
	}

	return chunks
}

func score(wanted int, str string) int {
	fmt.Println(" ", str, wanted)
	if len(str) == wanted {
		fmt.Println(" Even")
		return 1
	}

	has := 0
	for _, ch := range str {
		if ch == '#' {
			has++
		}
	}

	fmt.Println(" Missing", wanted, has, wanted-has)
	fmt.Println(" Options", len(str)-has)

	// product := 1
	// for i := ; i > wanted-has; i-- {
	// 	product *= i
	// }

	return (len(str) - has) * wanted
}

func satisfies(groups []int, str string) int {
	chunks := chunker(str)
	i := 0
	for len(groups) != len(chunks) {
		if len(chunks[i]) >= groups[i] {
			length := min(groups[i]+1, len(chunks[i]))
			tail := append([]string{}, chunks[min(i+1, len(chunks)):]...)
			chunks = append(chunks[0:i], chunks[i][0:length-1], chunks[i][length:])
			chunks = append(chunks, tail...)
		}
		i++
	}

	fmt.Println(chunks)

	product := 1
	for i, chunk := range chunks {
		product *= score(groups[i], chunk)
	}

	return product
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		parts := strings.Split(line, " ")
		groups := split(parts[1])
		curr := satisfies(groups, parts[0])
		fmt.Println(">>", curr)
		sum += curr
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)
}
