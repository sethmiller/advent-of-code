package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type target = int
type operation = func(i int) int

type Monkey struct {
	items       []int
	op          operation
	test        int
	yes         target
	no          target
	inspections int
}

func (m Monkey) HasStuff() bool {
	return len(m.items) > 0
}

func (m *Monkey) Inspect(ms *Monkeys) (int, target) {
	m.inspections += 1
	count := len(m.items)
	worry := m.items[count-1]
	m.items = m.items[0 : count-1]

	worry = m.op(worry) % ms.max

	if worry%m.test == 0 {
		// return ms.Find(worry), m.yes
		return worry, m.yes
	}

	// return Ms.Find(worry), m.no
	return worry, m.no
}

func (m *Monkey) ReceiveItem(i int) {
	m.items = append(m.items, i)
}

type Monkeys struct {
	monkeys []*Monkey
	tests   []int
	max     int
}

func NewMonkeys(monkeys []*Monkey) *Monkeys {
	tests := make([]int, len(monkeys))
	max := 1
	for i, m := range monkeys {
		tests[i] = m.test
		max *= m.test
	}
	return &Monkeys{
		monkeys: monkeys,
		tests:   tests,
		max:     max,
	}
}

// This was dumb but I'm leaving it. Big whoop. Wanna fight about it?
func (m *Monkeys) Find(n int) int {
	memo := make([]int, len(m.tests))
	for i, t := range m.tests {
		memo[i] = n % t
	}

	needed := len(memo)

	for curr := 0; curr < m.max; curr++ {
		found := 0
		for i, t := range m.tests {
			if n%t != memo[i] {
				found++
			}
		}

		if found == needed {
			return curr
		}
	}

	return n
}

func parseOp(line string) operation {
	tokens := strings.Split(line, " ")
	left := tokens[2]
	op := tokens[3]
	right := tokens[4]
	return func(i int) int {
		a, b := i, i
		if l, err := strconv.Atoi(left); err == nil {
			a = l
		}

		if r, err := strconv.Atoi(right); err == nil {
			b = r
		}

		if op == "*" {
			return a * b
		}

		return a + b
	}
}

func parseItems(line string) []int {
	parts := strings.Split(line, ", ")
	items := make([]int, len(parts))
	for i, str := range parts {
		val, _ := strconv.Atoi(str)
		items[i] = val
	}

	return items
}

func main() {
	monkeys := []*Monkey{}
	scanner := bufio.NewScanner(os.Stdin)
	currentLine := 0
	monkey := new(Monkey)
	for scanner.Scan() {
		line := scanner.Text()

		switch currentLine {
		case 0:
			currentLine++
		case 1:
			monkey.items = parseItems(line[18:])
			currentLine++
		case 2:
			monkey.op = parseOp(line[13:])
			currentLine++
		case 3:
			test, _ := strconv.Atoi(line[21:])
			monkey.test = test
			currentLine++
		case 4:
			yes, _ := strconv.Atoi(line[29:])
			monkey.yes = yes
			currentLine++
		case 5:
			no, _ := strconv.Atoi(line[30:])
			monkey.no = no
			currentLine++
		case 6:
			currentLine = 0
			monkeys = append(monkeys, monkey)
			monkey = new(Monkey)
		}
	}

	monkeys = append(monkeys, monkey)

	ms := NewMonkeys(monkeys)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	for _, m := range monkeys {
		monk := *m
		fmt.Println(monk)
	}

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {
			for monkey.HasStuff() {
				item, target := monkey.Inspect(ms)
				//fmt.Printf("Monkey %d looked at %d and will give it to %d\n", i, item, target)
				monkeys[target].ReceiveItem(item)
			}
		}
	}

	inspections := make([]int, len(monkeys))

	for i, m := range monkeys {
		inspections = append(inspections, m.inspections)
		fmt.Printf("Monkey %d => %d\n", i, m.inspections)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Printf("Business of the monkeys: %d\n", inspections[0]*inspections[1])
}
