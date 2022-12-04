package main

import (
	"advent-of-code/2022/day2/common"
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2022/day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	game := common.PaperRockScissorsGame{}
	game.GameRule = game.SneakyRule

	sc := bufio.NewScanner(file)

	for ok, opponent, picked := game.ParseInput(sc); ok; ok, opponent, picked = game.ParseInput(sc) {
		game.Play(opponent, picked)
	}

	log.Println(game.GetPoints())
}
