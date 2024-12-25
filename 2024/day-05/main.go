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

func inOrder(page *Page, after []string) bool {
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

func swap(arr []string, i int, j int) {
	memo := arr[i]
	arr[i] = arr[j]
	arr[j] = memo
}

func changed(before, after []string) bool {
	for i, val := range before {
		if val != after[i] {
			return true
		}
	}

	return false
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

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		pages := strings.Split(line, ",")
		fixed := true

		before := make([]string, len(pages))
		copy(before, pages)
		for fixed {
			fixed = false
			for i := 0; i < len(pages); i++ {
				page := pages[i]
				if i < len(pages)-1 && !inOrder(nodes[page], pages[i+1:]) {
					fmt.Println("nope")
					swap(pages, i, i+1)
					fixed = true
				}
			}
		}

		if changed(before, pages) {
			fmt.Println("yup")
			fmt.Println(pages)
			middle, _ := strconv.Atoi(pages[(len(pages) / 2)])
			sum += middle
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(sum)
}
