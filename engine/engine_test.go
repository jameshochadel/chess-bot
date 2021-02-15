package engine

import (
	"fmt"
	"testing"

	"github.com/notnil/chess"
)

func TestEvaluatePosition(t *testing.T) {
	cases := []struct{
		Name string
		FEN string
		Expected float64
	}{
		{ "starting position", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", 0.0 },
		{ "starting position", "rnbqkbnr/ppppppp1/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", 1.0 },
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v + %v", tc.FEN, tc.Expected), func(t *testing.T) {
			updateFunc, err := chess.FEN(tc.FEN)
			if err != nil {
				t.Fatal("applying FEN to game", err)
			}
			game := chess.NewGame(updateFunc)

			actual := evaluatePosition(game.Position())
			if actual != tc.Expected {
				t.Fatalf("expected %v, got %v", tc.Expected, actual)
			}
		})		
	}
}
