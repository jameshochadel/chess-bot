package engines

import (
	"fmt"
	"math"
	"strings"

	"github.com/notnil/chess"
)

type AlphaBetaEngine struct{}

// SuggestedMove calculates the most advantageous move for the player. It currently works
// synchronously, but can probably be made concurrent with goroutines.
func (e *AlphaBetaEngine) SuggestedMove(pos *chess.Position, maxPlayer bool) *chess.Move {
	startingDepth := 3
	return e.alphabeta(pos, startingDepth, maxPlayer).move
}

// minimax estimates the value of a position pos using the minimax algorithm,
// which minimizes loss for the given player, assuming the other player always
// makes an optimal move.
func (e *AlphaBetaEngine) alphabeta(pos *chess.Position, depth int, maxPlayer bool) *positionScore {
	if depth == 0 || len(pos.ValidMoves()) == 0 {
		return &positionScore{
			value: e.evaluatePosition(pos),
		}
	}

	bestScore := &positionScore{}
	var comparator func(float64, float64) float64

	if maxPlayer {
		bestScore.value = math.Inf(-1)
		comparator = math.Max
	} else {
		bestScore.value = math.Inf(1)
		comparator = math.Min
	}

	for _, m := range pos.ValidMoves() {
		fmt.Printf("Evaluating move %v\n", m)
		candidate := &positionScore{
			move:  m,
			value: e.alphabeta(pos.Update(m), depth-1, !maxPlayer).value,
		}

		bestScore = best(comparator, bestScore, candidate)
	}

	return bestScore
}

// evaluatePosition returns a heuristic value of the board, where positive values
// are advantageous for the white (max) player and negative values are advantageous
// for the black (min) player.
func (e *AlphaBetaEngine) evaluatePosition(pos *chess.Position) float64 {
	str := pos.String()
	var acc float64

	for piece, value := range pieceVals {
		acc += float64(strings.Count(str, piece)) * value
	}

	return acc
}
