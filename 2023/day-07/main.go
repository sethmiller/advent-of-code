package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	High = iota
	Pair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

var values = map[rune]int{
	'A': 12, 'K': 11, 'Q': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1, 'J': 0,
}

type Hand struct {
	cards string
	rank  int
	bid   int
}

func rank(str string) int {
	m := map[rune]int{}
	for _, ch := range str {
		i := m[ch]
		m[ch] = i + 1
	}

	jokers := m['J']
	delete(m, 'J')

	if jokers > 0 {
		most := ' '
		for k, v := range m {
			if v > m[most] {
				most = k
			}
		}

		m[most] = m[most] + jokers
	}

	items := len(m)
	if items == 1 {
		return FiveKind
	}

	if items == 4 {
		return Pair
	}

	if items == 2 {
		for _, v := range m {
			if v == 3 || v == 2 {
				return FullHouse
			}

			if v == 4 || v == 1 {
				return FourKind
			}

			break
		}
	}

	if items == 3 {
		pairs := 0
		for _, v := range m {
			if v == 2 {
				pairs++
			}

			if v == 3 {
				return ThreeKind
			}
		}

		if pairs == 2 {
			return TwoPair
		}

		return Pair
	}

	return High
}

func sorter(a, b Hand) bool {
	if a.rank == b.rank {
		for i, ch := range a.cards {
			if ch != rune(b.cards[i]) {
				return values[ch] < values[rune(b.cards[i])]
			}
		}
	}

	return a.rank < b.rank
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hands := []Hand{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		cards := parts[0]
		bid, _ := strconv.Atoi(strings.Trim(parts[1], " "))

		hands = append(hands, Hand{
			cards: cards,
			rank:  rank(cards),
			bid:   bid,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	sort.Slice(hands, func(a, b int) bool {
		return sorter(hands[a], hands[b])
	})

	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}

	fmt.Println(sum)
}
