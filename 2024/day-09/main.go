package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const EMPTY = -1

func atoi(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func atoiAll(strs []string) []int {
	ints := make([]int, len(strs))
	for i, val := range strs {
		ints[i] = atoi(val)
	}

	return ints
}

func checksum(vals []int) int {
	sum := 0
	for i, val := range vals {
		if val == EMPTY {
			continue
		}
		sum += i * val
	}

	return sum
}

func expand(layout []int) []int {
	expanded := []int{}
	for i := 0; i < len(layout); {
		index := i / 2
		size := layout[i]
		i++
		blank := layout[i]
		i++

		for j := 0; j < size; j++ {
			expanded = append(expanded, index)
		}

		for j := 0; j < blank; j++ {
			expanded = append(expanded, EMPTY)
		}
	}

	return expanded
}

func findGap(disk []int, size int, stop int) int {
	count := 0
	for i := 0; i <= stop; i++ {
		val := disk[i]
		if val != EMPTY && count >= size {
			return i - count
		}

		if val != EMPTY {
			count = 0
		}

		if val == EMPTY {
			count++
		}
	}

	return -1
}

func print(disk []int) string {
	str := ""
	for _, val := range disk {
		ch := fmt.Sprintf("%d", val)

		if val == EMPTY {
			ch = "."
		}
		str += ch
	}

	return str
}

func defrag(disk []int, layout []int) []int {
	defraged := make([]int, len(disk))
	copy(defraged, disk)

	fromEnd := len(defraged)
	for i := len(layout) - 1; i >= 0; {
		sourceBlank := layout[i]
		i--
		sourceSize := layout[i]
		i--

		fromEnd -= sourceBlank + sourceSize

		fromStart := findGap(defraged, sourceSize, fromEnd)
		if fromStart == EMPTY {
			continue
		}

		for m := 0; m < sourceSize; m++ {
			// fmt.Printf("moving %d from %d to %d\n", defraged[fromEnd+m], fromStart+m, fromEnd+m)
			defraged[fromStart+m] = defraged[fromEnd+m]
			defraged[fromEnd+m] = -1
		}
		//fmt.Println(print(defraged))
	}

	return defraged
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	line += "0"

	layout := atoiAll(strings.Split(line, ""))

	disk := expand(layout)
	fmt.Println(print(disk))
	defraged := defrag(disk, layout)
	//fmt.Println(print(defraged))
	sum := checksum(defraged)

	fmt.Println(sum)

}
