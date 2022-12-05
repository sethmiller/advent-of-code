package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node[V any] struct {
	next  *node[V]
	value *V
}

type stack[V any] struct {
	head *node[V]
}

func (s *stack[V]) push(value *V) {
	head := node[V]{value: value, next: s.head}
	s.head = &head
}

func (s *stack[V]) pop() *V {
	head := s.head
	if head == nil {
		return nil
	}

	s.head = s.head.next
	return head.value
}

func (s stack[V]) peek() *V {
	if s.head == nil {
		return nil
	}

	return s.head.value
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Read the initial state
	header := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		header = append(header, line)
	}

	// Build the stacks
	stackCount := len(strings.Fields(header[len(header)-1]))
	stacks := make([]stack[rune], stackCount)
	for level := len(header) - 2; level >= 0; level-- {
		line := header[level]
		for i := 0; i < len(line); i += 4 {
			chunk := line[i : i+3]
			if chunk != "   " {
				value := rune(chunk[1])
				stacks[i/4].push(&value)
			}
		}
	}

	// Follow the steps
	for scanner.Scan() {
		// move 6 from 2 to 1
		tokens := strings.Split(scanner.Text(), " ")
		count, _ := strconv.Atoi(tokens[1])
		source, _ := strconv.Atoi(tokens[3])
		dest, _ := strconv.Atoi(tokens[5])

		for i := 0; i < count; i++ {
			stacks[dest-1].push(stacks[source-1].pop())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	for i := 0; i < len(stacks); i++ {
		fmt.Printf("%s", string(*stacks[i].peek()))
	}
}
