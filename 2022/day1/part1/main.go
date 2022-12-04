package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	bestElfNumber := 0
	bestElfCalories := 0

	elfNumber := 0
	elfCalories := 0
	for sc.Scan() {
		x, err := strconv.Atoi(sc.Text())
		if err != nil {
			if bestElfCalories < elfCalories {
				bestElfNumber = elfNumber
				bestElfCalories = elfCalories
			}

			elfNumber++
			elfCalories = 0
		} else {
			elfCalories += x
		}
	}
	log.Println(bestElfCalories)
	log.Println(bestElfNumber)
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}
