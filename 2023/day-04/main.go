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

func inc(vals []int, pos, n, mul int) []int {
	for i := pos; i < pos+n; i++ {
		if i >= len(vals) {
			vals = append(vals, 1)
		}

		vals[i] = vals[i] + mul
	}

	return vals
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	num := 0
	cards := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(cards) <= num {
			cards = append(cards, 1)
		}

		parts, n := split(line[strings.Index(line, ":"):])

		winners := toSet(parts[0:n])
		tries := toSet(parts[n:])

		found := union(winners, tries)
		cards = inc(cards, num+1, len(found), cards[num])

		num++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
	sum := 0

	for _, i := range cards {
		sum += i
	}
	fmt.Println(sum)
}
