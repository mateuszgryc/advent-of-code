package common

import (
	"fmt"
	"testing"
)

func Test_convertRuneToPriority(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want int
	}{
		{"a to 1", 'a', 1},
		{"z to 26", 'z', 26},
		{"A to 27", 'A', 27},
		{"Z to 52", 'Z', 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertRuneToPriority(tt.r); got != tt.want {
				t.Errorf("ConvertRuneToPriority() = %v, want %v", got, tt.want)
			}
		})
	}

	for i := 0; i < 128; i++ {
		fmt.Printf("Char %c Unicode: %U pos: %d \n", i, i, i)
	}
}
