package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type filetype int

const (
	file filetype = iota
	dir
)

type node struct {
	name     string
	children map[string]*node
	parent   *node
	ft       filetype
	size     int
	total    int
}

func (n *node) findOrCreateChild(name string, ft filetype, size int) *node {
	child, found := n.children[name]
	if !found {
		child = &node{
			name:     name,
			size:     size,
			ft:       ft,
			parent:   n,
			children: make(map[string]*node),
		}
		n.children[name] = child
	}

	return child
}

func (n node) path() string {
	path := n.name
	for dir := n.parent; dir != nil; dir = dir.parent {
		path = dir.name + "/" + path
	}

	return path
}

func (n *node) updateSizes() int {
	total := n.size
	for _, v := range n.children {
		total += v.updateSizes()
	}

	n.total = total

	return total
}

func (n node) walk(f func(n *node)) {
	f(&n)
	for _, v := range n.children {
		v.walk(f)
	}
}

var size = regexp.MustCompile(`\d*`)

func main() {
	root := &node{
		name:     "",
		children: make(map[string]*node),
	}

	pwd := root

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch tok := parts[0]; {
		case tok == "$":
			switch parts[1] {
			case "cd":
				if parts[2] == "/" {
					pwd = root
				} else if parts[2] == ".." {
					if pwd != root {
						pwd = pwd.parent
					}
				} else {
					pwd = pwd.findOrCreateChild(parts[2], dir, 0)
				}
			case "ls":
			}
		case tok == "dir":
			pwd.findOrCreateChild(parts[1], dir, 0)
		case size.MatchString(tok):
			fileSize, _ := strconv.Atoi(parts[0])
			pwd.findOrCreateChild(parts[1], file, fileSize)
		default:
			panic("Uh oh!")
		}
	}

	sum := 0
	root.updateSizes()
	root.walk(func(n *node) {
		fmt.Printf("%s -> %d\n", n.path(), n.size)
		if n.ft == dir && n.total <= 100_000 {
			sum += n.total
		}
	})

	fmt.Printf("Sum: %d", sum)
}
