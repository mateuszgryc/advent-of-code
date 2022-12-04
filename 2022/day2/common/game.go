package common

import "strings"

type PaperRockScissorsGame struct {
	points   int
	GameRule func(opponent, picked int8) GameState
}

type GameState int8

const (
	LOSE GameState = iota
	DRAW
	WIN
)

func (g *PaperRockScissorsGame) Play(opponent, picked int8) {
	if g.GameRule == nil {
		g.GameRule = g.DefaultRule
	}
	switch g.GameRule(opponent, picked) {
	case WIN:
		g.points += 6
	case DRAW:
		g.points += 3
	}
}

func (g *PaperRockScissorsGame) DefaultRule(opponent, picked int8) GameState {
	g.points += int(picked) + 1
	if opponent == picked {
		return DRAW
	}
	if (opponent+1)%3 == picked {
		return WIN
	}
	return LOSE
}

func (g *PaperRockScissorsGame) SneakyRule(opponent, picked int8) GameState {
	switch picked {
	case 0:
		g.points += int(((opponent-1)%3+3)%3 + 1) // previous figure and add one point because modulo starts from 0
		return LOSE
	case 1:
		g.points += int((opponent)%3) + 1 // the same figure
		return DRAW
	case 2:
		g.points += int((opponent+1)%3) + 1 // next figure
		return WIN
	default:
		panic("this should never happen")
	}
}

func (g *PaperRockScissorsGame) ParseInput(sc Scanner) (bool, int8, int8) {
	if sc.Scan() {
		line := sc.Text()
		strategy := strings.Split(line, " ")
		opponent := strategy[0][0] % 65
		picked := strategy[1][0] % 88
		return true, int8(opponent), int8(picked)
	}
	return false, 0, 0
}

func (g *PaperRockScissorsGame) GetPoints() int {
	return g.points
}

type Scanner interface {
	Scan() bool
	Text() string
}
