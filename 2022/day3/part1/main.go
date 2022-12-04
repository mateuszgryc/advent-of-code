package main

import (
	"advent-of-code/2022/day3/common"
	"bufio"
	"os"
)

type Rucksack struct {
	items       []string
	compartment map[uint8]int
	priority    int
}

func (r *Rucksack) GetPriority() int {
	return r.priority
}

func NewRucksack(items string) int {
	r := Rucksack{items: []string{items}}
	r.compartment = make(map[uint8]int, len(r.items)/2)
	for i := 0; i < len(items)/2; i++ {
		r.compartment[items[i]] = 1
	}
	for i := len(items) / 2; i < len(items); i++ {
		item := items[i]
		if v, ok := r.compartment[item]; ok && v == 1 {
			r.priority += common.ConvertRuneToPriority(rune(item))
			r.compartment[item] = 0
		}
	}
	return r.priority
}

func main() {
	file, err := os.Open("2022/day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		r := NewRucksack(line)
		sum += r
	}
	println(sum)
}
