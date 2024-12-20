package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Page struct {
	val   string
	after []*Page
}

func returnOrCreate(node *Page, val string) *Page {
	if node == nil {
		return &Page{
			val:   val,
			after: make([]*Page, 0),
		}
	}

	return node
}

func containsAll(page *Page, after []string) bool {
	for _, val := range after {
		// fmt.Printf("Checking %s for %s\n", page.val, val)
		found := false
		for _, child := range page.after {
			if child.val == val {
				found = true
				break
			}
		}

		if !found {
			// fmt.Printf("%s did not containsAll\n", page.val)
			return false
		}
	}

	return true
}

func containsNone(page *Page, after []string) bool {
	for _, val := range after {
		// fmt.Printf("Checking %s -> %s\n", page.val, val)
		for _, child := range page.after {
			if child.val == val {
				// fmt.Printf("%s did not containNone\n", page.val)
				return false
			}
		}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	nodes := map[string]*Page{}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if len(line) == 0 {
			break
		}

		tokens := strings.Split(line, "|")

		left := returnOrCreate(nodes[tokens[0]], tokens[0])
		right := returnOrCreate(nodes[tokens[1]], tokens[1])

		fmt.Println(left, right)
		left.after = append(left.after, right)

		nodes[tokens[0]] = left
		nodes[tokens[1]] = right
	}

	sum := 0
books:
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		pages := strings.Split(line, ",")
		for i, page := range pages {
			if i < len(pages)-1 && !containsAll(nodes[page], pages[i+1:]) {
				fmt.Println("nope")
				continue books
			}

			// This isn't doing anything, apparently
			if i > 1 && !containsNone(nodes[page], pages[:i]) {
				fmt.Println("noper")
				continue books
			}
		}

		fmt.Println("yup")
		middle, _ := strconv.Atoi(pages[(len(pages) / 2)])
		sum += middle
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)
}
