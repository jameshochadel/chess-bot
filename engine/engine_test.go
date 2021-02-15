package engine

import (
	"fmt"
	"testing"

	"github.com/notnil/chess"
)

func TestEvaluatePosition(t *testing.T) {
	cases := []struct{
		Name string
		PositionFEN string
		Expected float64
	}{
		{ "starting position", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", 0.0 },
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v + %v", tc.PositionFEN, tc.Expected), func(t *testing.T) {
			pos := chess.Position{}
			pos.UnmarshalText([]byte(tc.PositionFEN))
			actual := evaluatePosition(&pos)
			if actual != tc.Expected {
				t.Fatal("failed")
			}
		})		
	}
}
