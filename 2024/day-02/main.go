package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const min = 1
const max = 3

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		parts := strings.Split(line, " ")

		levels := []int{}
		for _, v := range parts {
			i, _ := strconv.Atoi(v)
			levels = append(levels, i)
		}

		asc := levels[0]-levels[1] < 0
		for i := range levels[:len(levels)-1] {
			delta := levels[i] - levels[i+1]

			if (delta < 0 && !asc) || (delta > 0 && asc) {
				break
			}

			absDelta := int(math.Abs(float64(delta)))
			if absDelta < min || absDelta > max {
				break
			}

			if i == len(levels)-2 {
				safe++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(safe)

}
