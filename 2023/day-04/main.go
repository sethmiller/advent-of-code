package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func atoi(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func toSet(arr []string) map[int]interface{} {
	m := make(map[int]interface{})
	for _, s := range arr {
		t := strings.Trim(s, " ")
		if len(t) > 0 {
			m[atoi(t)] = nil
		}
	}

	return m
}

func union(a, b map[int]interface{}) []int {
	found := []int{}

	for k := range b {
		if _, ok := a[k]; ok {
			found = append(found, k)
		}
	}

	return found
}

func split(str string) ([]string, int) {
	chunks := []string{}
	n := 0
	for _, str := range strings.Split(str, " ") {
		if str == "|" {
			n = len(chunks)
			continue
		}
		if len(str) > 0 {
			chunks = append(chunks, str)
		}
	}

	return chunks, n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts, n := split(line[strings.Index(line, ":"):])

		winners := toSet(parts[0:n])
		tries := toSet(parts[n:])

		found := union(winners, tries)
		fmt.Printf("%d -> (%d) (%d)\n", found, len(winners), len(tries))
		if len(found) > 0 {
			sum += int(math.Pow(2, float64(len(found)-1)))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)
}
