package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const debug = true

type node struct {
	val      *int
	children *[]*node
}

type packets []*node

var _ sort.Interface = packets{}

func (p packets) Len() int                  { return len(p) }
func (p packets) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }
func (p packets) Less(left, right int) bool { return treeDiff(p[left], p[right]) == Failure }

func (n *node) isChild() bool {
	return n.val != nil
}

func (n *node) hasChildren() bool {
	return !n.isChild() && len(*n.children) > 0
}

func label(b status, s string) status {
	if debug {
		fmt.Printf(" %s -> %d\n", s, b)
	}

	return b
}

func convert(child *node) *node {
	return &node{children: &[]*node{child}}
}

type status = int

const (
	Failure  status = -1
	Continue status = 0
	Success  status = 1
)

func treeDiff(left, right *node) status {
	result := Continue

	if left.isChild() && right.isChild() {
		if *(left.val) < *(right.val) {
			return label(Success, fmt.Sprintf("A left lower (%d, %d)", *(left.val), *(right.val)))
		} else if *(left.val) > *(right.val) {
			return label(Failure, fmt.Sprintf("B left higher (%d, %d)", *(left.val), *(right.val)))
		}
		return label(Continue, fmt.Sprintf("C (%d, %d)", *(left.val), *(right.val)))
	} else if left.isChild() {
		return label(treeDiff(convert(left), right), "D (converting left, see ^)")
	} else if right.isChild() {
		return label(treeDiff(left, convert(right)), "E (converting right, see ^)")
	} else if !left.hasChildren() && right.hasChildren() {
		return label(Success, "F (left was empty, right was not)")
	} else if left.hasChildren() && !right.hasChildren() {
		return label(Failure, "G (left was not empty, right was)")
	} else {
		for c, lChild := range *(left.children) {
			if c >= len(*right.children) {
				return label(Failure, "I (no more right - Failure)")
			}

			rChild := (*right.children)[c]

			// Compare the next child in `left` to the next child in `right`
			if result = treeDiff(lChild, rChild); result != Continue {
				return label(result, "J (got terminal result)")
			}

			if c+1 == len(*left.children) && c+1 < len(*right.children) {
				return label(Success, "K (left was empty, right was not)")
			}
		}
	}

	return label(result, "L (fell through. see ^)")
}

func parseTree(str *string) ([]*node, *string) {
	found := []*node{}
	num := intp(0)
	hasNum := false
	for len(*str) > 0 {
		s := (*str)[0]
		*str = (*str)[1:]

		switch s {
		case '[':
			children, _ := parseTree(str)
			found = append(found, &node{children: &children})
		case ']':
			if hasNum {
				found = append(found, &node{val: num})
			}
			return found, str
		case ',':
			if hasNum {
				found = append(found, &node{val: num})
				hasNum = false
				num = intp(0)
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			*num *= 10
			*num += int(s) - 0x30
			hasNum = true
		}
	}

	return found, str
}

func intp(i int) *int {
	return &i
}

func printTree(tree *node, level int) {
	if tree == nil || tree.children == nil {
		fmt.Println("{nil}")
		return
	}
	fmt.Println(strings.Repeat("  ", level-1), "[")
	for _, child := range *(tree.children) {
		if child.children != nil {
			printTree(child, level+1)
		} else {
			val := *(child.val)
			fmt.Printf("%s {%d}\n", strings.Repeat("  ", level), val)
		}
	}

	fmt.Println(strings.Repeat("  ", level-1), "]")
}

func printOriginalInput(tree *node) string {
	str := "["
	for i, child := range *(tree.children) {

		if child.children != nil {
			str += printOriginalInput(child)
		} else {
			val := *(child.val)
			str = fmt.Sprintf("%s%d", str, val)
		}

		if i < len(*tree.children)-1 {
			str += ","
		}
	}

	return str + "]"
}

func main() {
	p := packets{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		nodes, _ := parseTree(&line)
		tree := nodes[0]
		p = append(p, tree)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	sort.Sort(sort.Reverse(p))

	for i, tree := range p {
		fmt.Printf("%d => %s\n", i+1, printOriginalInput(tree))
	}

	// fmt.Printf("Sum: %d\n", sum)
}
