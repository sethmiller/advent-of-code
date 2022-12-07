package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	total_space    = 70_000_000
	required_space = 30_000_000
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

	root.updateSizes()
	used := root.total
	free := total_space - used
	needed := required_space - free

	var found *node
	root.walk(func(n *node) {
		fmt.Printf("%s -> %d\n", n.path(), n.total)
		if n.ft == dir && n.total >= needed && (found == nil || n.total < found.total) {
			found = n
		}
	})

	fmt.Println()
	fmt.Printf("Used: %d\n", used)
	fmt.Printf("Free: %d\n", free)
	fmt.Printf("Needed: %d\n", needed)
	fmt.Println()
	fmt.Printf("Found: (%s, %d)", found.path(), found.total)
}
