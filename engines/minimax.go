package engines

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/notnil/chess"
)

type MinimaxEngine struct {}

// SuggestedMove calculates the most advantageous move for the player. It currently works
// synchronously, but can probably be made concurrent with goroutines.
func (e *MinimaxEngine) SuggestedMove(pos *chess.Position, maxPlayer bool) *chess.Move {
	startingDepth := 3
	return minimax(pos, startingDepth, maxPlayer).move
}

// positionScore is a tuple for tracking the value of a position and the move that should
// be made to achieve that value.
type positionScore struct {
	// value is a heuristic value of the board. For more, see evaluatePosition.
	value float64
	// move is the optimal next move that should be made for the current player to achieve
	// the value `positionScore.value` of the position.
	move *chess.Move
}

// minimax estimates the value of a position pos using the minimax algorithm,
// which minimizes loss for the given player, assuming the other player always
// makes an optimal move.
func minimax(pos *chess.Position, depth int, maxPlayer bool) *positionScore {
	if depth == 0 || len(pos.ValidMoves()) == 0 {
		return &positionScore{
			value: evaluatePosition(pos),
		}
	}
	pos.Hash()

	bestScore := &positionScore{}
	var comparator func(float64, float64) float64

	if maxPlayer {
		bestScore.value = -9999
		comparator = math.Max
	} else {
		bestScore.value = 9999
		comparator = math.Min
	}

	for _, m := range pos.ValidMoves() {
		fmt.Printf("Evaluating move %v\n", m)
		candidate := &positionScore{
			move:  m,
			value: minimax(pos.Update(m), depth-1, !maxPlayer).value,
		}

		bestScore = best(comparator, bestScore, candidate)
	}

	return bestScore
}

// best compares two positions and returns the 'better' position, as determined by the
// comparator function. In the case that the positions are of equal value, position `a`
// is returned.
func best(comparator func(float64, float64) float64, a, b *positionScore) *positionScore {
	if a.value == b.value {
		return a
	} else {
		best := comparator(a.value, b.value)
		if best == a.value {
			return a
		} else {
			return b
		}
	}
}

// pieceVals uses Forsyth-Edwards Notation (FEN) conventions to represent pieces.
var pieceVals = map[string]float64{
	"P": 1,
	"N": 3,
	"B": 3,
	"R": 6,
	"Q": 9,
	"K": 9999,
	"p": -1,
	"n": -3,
	"b": -3,
	"r": -6,
	"q": -9,
	"k": -9999,
}

// evaluatePosition returns a heuristic value of the board, where positive values
// are advantageous for the white (max) player and negative values are advantageous
// for the black (min) player.
func evaluatePosition(pos *chess.Position) float64 {
	str := pos.String()
	var acc float64

	for piece, value := range pieceVals {
		acc += float64(strings.Count(str, piece)) * value
	}

	return acc
}

// randomBool generates a pseudorandom true or false value -- a 'coin flip'.
// Useful for determining which player is assigned White.
func randomBool() bool {
	rand.Seed(time.Now().UnixNano())
	outcome := rand.Intn(1)
	return outcome == 1
}
