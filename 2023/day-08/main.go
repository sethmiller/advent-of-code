package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	name  string
	left  *Node
	right *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%s: L->%s, R->%s", n.name, n.left.name, n.right.name)
}

// too dumb and lazy: https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
func prime_factors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	steps := scanner.Text()
	scanner.Scan()

	nodes := map[string]*Node{}
	for scanner.Scan() {
		line := scanner.Text()
		name := line[0:3]
		left := line[7:10]
		right := line[12:15]

		var node *Node
		var ok bool

		if node, ok = nodes[name]; !ok {
			node = &Node{
				name:  name,
				left:  nil,
				right: nil,
			}

			nodes[name] = node
		}

		if _, ok = nodes[left]; !ok {
			nodes[left] = &Node{
				name: left,
			}
		}

		if _, ok = nodes[right]; !ok {
			nodes[right] = &Node{
				name: right,
			}
		}

		node.left = nodes[left]
		node.right = nodes[right]
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	currs := []*Node{}
	for _, n := range nodes {
		if n.name[2] == 'A' {
			currs = append(currs, n)
		}
	}

	// look for loops
	loops := []int{}
	for _, c := range currs {
		fmt.Println(c)
		step := 0
		count := 0
		visited := map[string]int{}
		curr := c
		for {
			if curr.name[2] == 'Z' {
				key := fmt.Sprintf("%s%d", curr.name, step)
				if i, ok := visited[key]; ok {
					loops = append(loops, i)
					break
				}
				fmt.Printf("%s -> %d\r", curr.name, step)
				visited[key] = count
			}

			dir := steps[step]
			if dir == 'L' {
				curr = curr.left
			} else {
				curr = curr.right
			}

			count++
			step++
			if step >= len(steps) {
				step = 0
			}
		}
	}

	p := map[int]interface{}{}
	for _, i := range loops {
		for _, f := range prime_factors(i) {
			p[f] = nil
		}
	}

	product := 1
	for k, _ := range p {
		product *= k
	}

	fmt.Println(product)
}
