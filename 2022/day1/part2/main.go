package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type elfTree struct {
	calories []int
}

func (c elfTree) addNode(calories int) {
	if c.calories[0] < calories {
		c.calories[0] = calories
	}
	for i := 0; i < len(c.calories)-1; i++ {
		if c.calories[i] < c.calories[i+1] {
			return
		}
		t := c.calories[i+1]
		c.calories[i+1] = c.calories[i]
		c.calories[i] = t
	}
}

func (c elfTree) getCalories() int {
	sum := 0
	for _, calory := range c.calories {
		sum += calory
	}
	return sum
}

func main() {
	file, err := os.Open("/Users/mateuszgryc/personal/advent-of-code/1-1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	t := elfTree{calories: make([]int, 3)}
	elfCalories := 0
	for sc.Scan() {
		x, err := strconv.Atoi(sc.Text())
		if err != nil {
			t.addNode(elfCalories)
			elfCalories = 0
		} else {
			elfCalories += x
		}
	}
	t.addNode(elfCalories)
	log.Println(t.getCalories())

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}
