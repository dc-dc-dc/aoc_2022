package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type MaxIntHeatp []int

func (h MaxIntHeatp) Len() int           { return len(h) }
func (h MaxIntHeatp) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeatp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeatp) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeatp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	file, err := os.Open("day_1_data")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	h := &MaxIntHeatp{}
	heap.Init(h)

	var current int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			heap.Push(h, current)
			current = 0
			continue
		}
		// convert to int and add to current
		res, _ := strconv.Atoi(scanner.Text())
		current += res
	}
	var total, count int
	for count < 3 {
		res := heap.Pop(h).(int)
		total += res
		count += 1
	}

	fmt.Printf("%d\n", total)
}
