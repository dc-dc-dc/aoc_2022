package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	draw = map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
	}
	win = map[rune]int{
		'A': 2,
		'B': 3,
		'C': 1,
	}
	lose = map[rune]int{
		'A': 3,
		'B': 1,
		'C': 2,
	}
)

// opp you type    val
// A   X   Rock    1
// B   Y   Paper   2
// C   Z   Scissor 3
func calcScore(me, op rune) int {
	if me == 'X' {
		return lose[op]
	}
	if me == 'Y' {
		return draw[op]
	}
	return win[op]
}

func convPlay(play rune) int {
	if play == 'Y' {
		return 3
	}
	if play == 'Z' {
		return 6
	}
	return 0
}

func main() {

	file, err := os.Open("day_2/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var score int
	for scanner.Scan() {
		temp := scanner.Bytes()
		score += calcScore(rune(temp[2]), rune(temp[0]))
		score += convPlay(rune(temp[2]))
	}
	fmt.Println(score)
}
