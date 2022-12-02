package main

import (
	"bufio"
	"fmt"
	"os"
)

// Loss: 0
// Draw: 3
// Win: 6

// A: Rock
// B: Paper
// C: Scissors

// X: Rock (1)
// Y: Paper (2)
// Z: Scissors (3)

const inputFilePath = "2022/day02/input"

var rock, paper, scissors = 1, 2, 3
var lose, draw, win = 0, 3, 6

var scoreMapping1 = map[string]int{
	"A X": rock + draw,     // rock vs ROCK
	"A Y": paper + win,     // rock vs PAPER
	"A Z": scissors + lose, // rock vs SCISSORS
	"B X": rock + lose,     // paper vs ROCK
	"B Y": paper + draw,    // paper vs PAPER
	"B Z": scissors + win,  // paper vs SCISSORS
	"C X": rock + win,      // scissors vs ROCK
	"C Y": paper + lose,    // scissors vs PAPER
	"C Z": scissors + draw, // scissors vs SCISSORS
}

var scoreMapping2 = map[string]int{
	"A X": lose + scissors, // throw this to lose to elf's rock
	"A Y": draw + rock,     // throw this to draw to elf's rock
	"A Z": win + paper,     // throw this to win to elf's rock
	"B X": lose + rock,     // throw this to lose to elf's paper
	"B Y": draw + paper,    // throw this to draw to elf's paper
	"B Z": win + scissors,  // throw this to win to elf's paper
	"C X": lose + paper,    // throw this to lose to elf's scissors
	"C Y": draw + scissors, // throw this to draw to elf's scissors
	"C Z": win + rock,      // throw this to win to elf's scissors
}

func main() {
	fmt.Printf("Final Score Not Knowing Strategy: %d\n", GetScore(scoreMapping1))
	fmt.Printf("Final Score Knowing Strategy: %d\n", GetScore(scoreMapping2))
}

func GetScore(scoreMap map[string]int) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println(err)
	} else {
		if file != nil {
			defer file.Close()
			return calculateScore(scoreMap, file)
		}
	}
	return 0
}

func calculateScore(scoreMap map[string]int, file *os.File) int {
	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		score += scoreMap[line]
	}
	return score
}
