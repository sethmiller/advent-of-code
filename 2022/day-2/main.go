package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
	loss     = 0
	tie      = 3
	win      = 6
)

var scoresPartA = map[string]int{
	"A X": rock + tie,
	"A Y": paper + win,
	"A Z": scissors + loss,

	"B X": rock + loss,
	"B Y": paper + tie,
	"B Z": scissors + win,

	"C X": rock + win,
	"C Y": paper + loss,
	"C Z": scissors + tie,
}

// X, Loss
// Y, Tie
// Z, Win
var scoresPartB = map[string]int{
	"A X": scissors + loss,
	"A Y": rock + tie,
	"A Z": paper + win,

	"B X": rock + loss,
	"B Y": paper + tie,
	"B Z": scissors + win,

	"C X": paper + loss,
	"C Y": scissors + tie,
	"C Z": rock + win,
}

func main() {
	totalPartA := 0
	totalPartB := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		score, found := scoresPartA[line]
		if found {
			totalPartA += score
		} else {
			fmt.Printf("No score found for %s in `A`\n", line)
		}

		score, found = scoresPartB[line]
		if found {
			totalPartB += score
		} else {
			fmt.Printf("No score found for %s in `B`\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Printf("Total Part A: %d\n", totalPartA)
	fmt.Printf("Total Part B: %d\n", totalPartB)
}
