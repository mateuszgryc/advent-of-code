package main

import "testing"

func TestPaperRockScissorsGame_Play(t *testing.T) {
	tests := []struct {
		name    string
		points  int
		oponent int8
		picked  int8
	}{
		{name: "rock vs rock", points: 4, oponent: 0, picked: 0},
		{name: "rock vs paper", points: 8, oponent: 0, picked: 1},
		{name: "rock vs scissors", points: 3, oponent: 0, picked: 2},
		{name: "paper vs rock", points: 1, oponent: 1, picked: 0},
		{name: "paper vs paper", points: 5, oponent: 1, picked: 1},
		{name: "paper vs scissors", points: 9, oponent: 1, picked: 2},
		{name: "scissors vs rock", points: 6 + 1, oponent: 2, picked: 0},
		{name: "scissors vs paper", points: 0 + 2, oponent: 2, picked: 1},
		{name: "scissors vs scissors", points: 3 + 3, oponent: 2, picked: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &PaperRockScissorsGame{}
			g.Play(tt.oponent, tt.picked)
			if g.GetPoints() != tt.points {
				t.Fatalf("%s: want %d, got %d", tt.name, tt.points, g.GetPoints())
			}
		})
	}
}

func TestPaperRockScissorsGame_sneakRule(t *testing.T) {
	tests := []struct {
		name       string
		opponent   int8
		picked     int8
		wantPoints int
	}{
		{name: "rock vs LOSE", wantPoints: 3, opponent: 0, picked: 0},
		{name: "rock vs DRAW", wantPoints: 1, opponent: 0, picked: 1},
		{name: "rock vs WIN", wantPoints: 2, opponent: 0, picked: 2},
		{name: "paper vs LOSE", wantPoints: 1, opponent: 1, picked: 0},
		{name: "paper vs DRAW", wantPoints: 2, opponent: 1, picked: 1},
		{name: "paper vs WIN", wantPoints: 3, opponent: 1, picked: 2},
		{name: "scissors vs LOSE", wantPoints: 2, opponent: 2, picked: 0},
		{name: "scissors vs DRAW", wantPoints: 3, opponent: 2, picked: 1},
		{name: "scissors vs WIN", wantPoints: 1, opponent: 2, picked: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &PaperRockScissorsGame{}
			if _ = g.sneakRule(tt.opponent, tt.picked); g.points != tt.wantPoints {
				t.Errorf("sneakRule() = %v, want %v", g.points, tt.wantPoints)
			}
		})
	}
}
