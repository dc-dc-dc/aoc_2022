package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day_6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	scanner.Scan()
	line := scanner.Bytes()
	bmap := map[byte]any{}
	left := 0
	for right, s := range line {
		fmt.Println(left, right, s)
		if (right - left) == 14 {
			result = right
			break
		}
		_, ok := bmap[s]
		for left < right && ok {
			delete(bmap, line[left])
			_, ok = bmap[s]
			left++
			continue
		}
		bmap[s] = nil
	}

	fmt.Println(result)
}
