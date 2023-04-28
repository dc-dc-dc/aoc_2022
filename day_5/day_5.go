package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prependInt(x []rune, y rune) []rune {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
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
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		count, _ := strconv.Atoi(line[1])
		src, _ := strconv.Atoi(line[3])
		dst, _ := strconv.Atoi(line[5])
		for count > 0 {
			val := crates[src-1][0]
			crates[src-1] = crates[src-1][1:]
			crates[dst-1] = prependInt(crates[dst-1], val)
			count -= 1
		}
	}
	for _, c := range crates {
		res = append(res, c[0])
	}
	fmt.Println(string(res))
}
