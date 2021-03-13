package engines

import (
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

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

// positionScore is a tuple for tracking the value of a position and the move that should
// be made to achieve that value.
type positionScore struct {
	// value is a heuristic value of the board. For more, see evaluatePosition.
	value float64
	// move is the optimal next move that should be made for the current player to achieve
	// the value `positionScore.value` of the position.
	move *chess.Move
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

// randomBool generates a pseudorandom true or false value -- a 'coin flip'.
// Useful for determining which player is assigned White.
func randomBool() bool {
	rand.Seed(time.Now().UnixNano())
	outcome := rand.Intn(1)
	return outcome == 1
}
