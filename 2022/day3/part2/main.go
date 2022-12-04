package main

import (
	"advent-of-code/2022/day3/common"
	"bufio"
	"os"
)

func NewGroupRucksack(groupItems []string) int {
	compartment := make(map[uint8]int)

	for i := 0; i < len(groupItems); i++ {
		items := groupItems[i]
		for j := 0; j < len(items); j++ {
			c := items[j]
			if v, _ := compartment[c]; v == i {
				compartment[c] = v + 1
				if compartment[c] == 3 {
					return common.ConvertRuneToPriority(rune(c))
				}
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("2022/day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	q := make([]string, 3)
	index := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		q[index] = line
		index++
		if index == 3 {
			sum += NewGroupRucksack(q)
			index = 0
		}
	}
	println(sum)
}
