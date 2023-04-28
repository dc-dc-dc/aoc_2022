package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prependInt(x []rune, y []rune) []rune {
	x = append(x, y...)
	copy(x[len(y):], x)
	copy(x[:len(y)], y)
	return x
}

func main() {
	file, err := os.Open("day_5/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []rune{}
	// construct the crates
	var crates [][]rune = nil
	for scanner.Scan() {
		if crates == nil {
			cols := (len(scanner.Text()) + 1) / 4
			crates = make([][]rune, cols)
		}
		if len(scanner.Text()) == 0 {
			break
		}
		for i, c := range scanner.Text() {
			col := (i + 1) / 4
			if (i+1)%2 != 0 || c == ' ' || c-'A' < 0 {
				continue
			}
			// fmt.Printf("%c %d\n", c, c-'A')
			crates[col] = append(crates[col], c)
		}
	}
	// do some
	// for _, c := range crates {
	// 	fmt.Println(string(c))
	// }

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		count, _ := strconv.Atoi(line[1])
		src, _ := strconv.Atoi(line[3])
		dst, _ := strconv.Atoi(line[5])
		temp := make([]rune, count)
		copy(temp, crates[src-1][:count])

		crates[dst-1] = append(temp, crates[dst-1]...) //prependInt(crates[dst-1], toAdd)
		crates[src-1] = crates[src-1][count:]
		for _, c := range crates {
			fmt.Println(string(c))
		}
	}

	for _, c := range crates {
		res = append(res, c[0])
	}
	fmt.Println(string(res))
}
