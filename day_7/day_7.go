package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NodeFile struct {
	name string
	size int
}

type Node struct {
	name      string
	parent    *Node
	totalsize int
	children  []*Node
	files     []*NodeFile
}

func NewNode(name string, parent *Node) *Node {
	return &Node{
		name:     name,
		parent:   parent,
		children: make([]*Node, 0),
		files:    make([]*NodeFile, 0),
	}
}

func dfs(nodes *Node) int {
	if nodes.totalsize != 0 {
		return nodes.totalsize
	}
	if len(nodes.children) == 0 && len(nodes.files) == 0 {
		return 0
	}
	totalsize := 0
	for _, node := range nodes.children {
		totalsize += dfs(node)
	}
	for _, file := range nodes.files {
		totalsize += file.size
	}
	nodes.totalsize = totalsize
	return totalsize
}

func PrintTree(nodes []*Node) {
	next := []*Node{}
	for _, s := range nodes {
		fmt.Printf("%s %d\n", s.name, s.totalsize)
		next = append(next, s.children...)
	}
	if len(next) > 0 {
		PrintTree(next)
	}
}

func Calc(nodes []*Node) int {
	next := []*Node{}
	var res int
	for _, s := range nodes {
		if s.totalsize < 100_000 {
			res += s.totalsize
		}
		fmt.Printf("%s %d\n", s.name, s.totalsize)
		next = append(next, s.children...)
	}
	if len(next) > 0 {
		res += Calc(next)
	}
	return res
}

func FilterSizes(nodes []*Node, size int) []int {
	next := []*Node{}
	res := []int{}
	for _, s := range nodes {
		if s.totalsize >= size {
			res = append(res, s.totalsize)
		}
		fmt.Printf("%s %d\n", s.name, s.totalsize)
		next = append(next, s.children...)
	}
	if len(next) > 0 {
		res = append(res, FilterSizes(next, size)...)
	}
	return res
}

func main() {
	file, err := os.Open("day_7/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	root := NewNode("root", nil)
	for scanner.Scan() {
		line := scanner.Text()
		splts := strings.Split(line, " ")
		if strings.HasPrefix(line, "$") {
			// command
			if splts[1] == "cd" {
				if splts[2] == ".." {
					root = root.parent
				} else {
					temp := NewNode(splts[2], root)
					root.children = append(root.children, temp)
					root = temp
				}
			} else {
				// fmt.Println(splts[1])
			}
		} else if strings.HasPrefix(line, "dir") {
			// dir
		} else {
			// file with size
			// fmt.Println(splts[0], splts[1])
			size, _ := strconv.Atoi(splts[0])
			root.files = append(root.files, &NodeFile{
				name: splts[1],
				size: size,
			})
		}
	}
	for root.parent.name != "root" {
		root = root.parent
	}
	dfs(root)
	// temp := []int{}
	numToFind := root.totalsize - (maxSize - updateSize)
	temp := FilterSizes([]*Node{root}, numToFind)
	sort.Ints(temp)

	fmt.Printf("%v\n", temp[0])
}

var (
	maxSize    = 70_000_000
	updateSize = 30_000_000
)
