package engine

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/notnil/chess"
)

// bestMove calculates the most advantageous move for the player. It currently works
// synchronously, but can probably be made concurrent with goroutines.
func BestMove(pos *chess.Position, maxPlayer bool) *chess.Move {
	var scoredMoves map[chess.Move]float64
	startingDepth := 10
	for _, move := range pos.ValidMoves() {
		scoredMoves[*move] = minimax(pos.Update(move), startingDepth, maxPlayer)
	}

	var (
		bestMove  *chess.Move
		bestValue float64
	)

	// compareFunc := func(a float64, b float64, )

	for m, v := range scoredMoves {
		if bestMove == nil {
			bestMove = &m
			bestValue = v
		} else {
			if maxPlayer && bestValue < v {
				bestMove = &m
				bestValue = v
			} else if !maxPlayer && v < bestValue {
				bestMove = &m
				bestValue = v
			}
		}
	}
	return bestMove
}

func minimax(pos *chess.Position, depth int, maxPlayer bool) (posValue float64) {
	if depth == 0 || len(pos.ValidMoves()) == 0 {
		return evaluatePosition(pos)
	}
	if maxPlayer {
		var maxMove float64
		for i, m := range pos.ValidMoves() {
			p := pos.Update(m)
			if i == 0 {
				maxMove = minimax(p, depth - 1, false)
			} else {
				maxMove = math.Max(maxMove, minimax(p, depth - 1, false))
			}
		}
	} else {
		var minMove float64
		for i, m := range pos.ValidMoves() {
			p := pos.Update(m)
			if i == 0 {
				minMove = minimax(p, depth - 1, true)
			} else {
				minMove = math.Min(minMove, minimax(p, depth - 1, true))
			}
		}
	}

	return 0
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
	strings.Count(str, "q")
	// count number of each piece
	return 0.0
}

// randomBool generates a pseudorandom true or false value -- a 'coin flip'.
// Useful for determining which player is assigned White.
func randomBool() bool {
	rand.Seed(time.Now().UnixNano())
	outcome := rand.Intn(1)
	return outcome == 1
}
