package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inBounds(ll, lr, rl, rr int) bool {
	return (ll <= rl && rl <= lr) && (ll <= rr && rr <= lr)
}

func main() {
	file, err := os.Open("day_4/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		splts := strings.Split(scanner.Text(), ",")
		left := strings.Split(splts[0], "-")
		right := strings.Split(splts[1], "-")
		leftMin, _ := strconv.Atoi(left[0])
		leftMax, _ := strconv.Atoi(left[1])
		rightMin, _ := strconv.Atoi(right[0])
		rightMax, _ := strconv.Atoi(right[1])
		if inBounds(leftMin, leftMax, rightMin, rightMax) || inBounds(rightMin, rightMax, leftMin, leftMax) {
			fmt.Printf("left: %v, right: %v \n", left, right)
			result += 1
		}
	}
	fmt.Println(result)
}
