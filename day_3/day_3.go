package main

import (
	"bufio"
	"fmt"
	"os"
)

// opp you type    val
// A   X   Rock    1
// B   Y   Paper   2
// C   Z   Scissor 3

func main() {
	file, err := os.Open("day_3/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) / 2
		sack1, sack2 := line[:mid], line[mid:]
		fmt.Printf("%s : %s\n", sack1, sack2)
		s1 := map[byte]any{}
		s2 := map[byte]any{}
		var found byte
		for i := range sack1 {
			t1 := sack1[i]
			t2 := sack2[i]
			if t1 == t2 {
				found = t1
				break
			}
			if _, ok := s2[t1]; ok {
				found = t1
				break
			}
			if _, ok := s1[t2]; ok {
				found = t2
				break
			}
			s1[t1] = nil
			s2[t2] = nil
		}
		res := rune(found) - 'a'
		if res > 0 {
			res += 1
		} else {
			res = rune(found) - 'A' + 27
		}
		fmt.Printf("char %c val: %d \n", rune(found), res)
		result += int(res)
	}
	fmt.Println(result)
}
