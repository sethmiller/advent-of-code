package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoitoa(str string) string {
	i, _ := strconv.Atoi(str)

	return fmt.Sprintf("%d", i)
}

func mul(str string, times int) string {
	i, _ := strconv.Atoi(str)

	return fmt.Sprintf("%d", i*times)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stones := strings.Split(scanner.Text(), " ")
	fmt.Println(stones)

	for i := 0; i < 25; i++ {
		updated := []string{}
		for _, stone := range stones {
			if stone == "0" {
				updated = append(updated, "1")
			} else if len(stone)%2 == 0 {
				updated = append(updated, stone[0:len(stone)/2])
				updated = append(updated, atoitoa((stone[len(stone)/2:])))
			} else {
				updated = append(updated, mul(stone, 2024))
			}
		}

		stones = updated
	}

	fmt.Println(len(stones))
}
