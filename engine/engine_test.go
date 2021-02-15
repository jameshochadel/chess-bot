package engine

import (
	"fmt"
	"testing"

	"github.com/notnil/chess"
)

func TestEvaluatePosition(t *testing.T) {
	cases := []struct{
		Pos *chess.Position
		Expected float64
	}{
		{ &chess.Position{}, 0.0 },
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v + %v", tc.Pos.Board().String(), tc.Expected), func(t *testing.T) {
			actual := evaluatePosition(tc.Pos)
			if actual != tc.Expected {
				t.Fatal("failed")
			}
		})		
	}
}
