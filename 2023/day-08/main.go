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
	return n.name
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

	curr := nodes["AAA"]
	step := 0

	visited := 0
	for {
		if curr.name == "ZZZ" {
			break
		}

		dir := steps[step]
		var next *Node
		if dir == 'L' {
			next = curr.left
		} else {
			next = curr.right
		}

		visited++
		step++
		if step >= len(steps) {
			step = 0
		}

		curr = next
	}

	fmt.Println(visited)

}
