package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	label string
	value int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	parts := strings.Split(line, ",")

	m := map[int][]entry{}

	for _, part := range parts {
		hash := 0
		for _, ch := range part {
			if ch == '=' || ch == '-' {
				break
			}
			hash = (hash + int(ch)) * 17
			hash = hash % 256
		}

		fmt.Println("Hash", hash, part)

		var label string
		var value int
		deleting := false
		if part[len(part)-1] == '-' {
			deleting = true
			label = part[0 : len(part)-1]
		} else {
			s := strings.Split(part, "=")
			label = s[0]
			value, _ = strconv.Atoi(s[1])
		}

		if list, ok := m[hash]; ok {
			found := false
			for i, item := range list {
				if item.label == label {
					fmt.Println("Found existing", label)
					found = true
					if deleting {
						fmt.Println("deleting existing", label)
						fmt.Println("before", m[hash])
						m[hash] = append(m[hash][0:i], m[hash][i+1:]...)
						fmt.Println("after", m[hash])
						break
					} else {
						fmt.Println("replacing value")
						item.value = value
						m[hash][i] = item
						break
					}
				}
			}

			if !found && !deleting {
				fmt.Println("Adding to existing", label)
				m[hash] = append(m[hash], entry{
					label: label,
					value: value,
				})
			} else {
				fmt.Println("Skipping found delete", label)
			}
		} else {
			if !deleting {
				fmt.Println("Creating new", label)
				m[hash] = []entry{
					{
						label: label,
						value: value,
					},
				}
			} else {
				fmt.Println("Skipping delete", label)
			}
		}
	}

	sum := 0
	for i := 0; i < 256; i++ {
		if items, ok := m[i]; ok {
			for slot, item := range items {
				fmt.Println(i, slot, item.label, (i+1)*(slot+1)*item.value)
				sum += (i + 1) * (slot + 1) * item.value
			}
		}
	}

	fmt.Println(sum)
}
