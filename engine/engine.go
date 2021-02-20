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

// minimax estimates the value of a position pos using the minimax algorithm,
// which minimizes loss for the given player, assuming the other player always
// makes an optimal move.
func minimax(pos *chess.Position, depth int, maxPlayer bool) (posValue float64) {
	if depth == 0 || len(pos.ValidMoves()) == 0 {
		return evaluatePosition(pos)
	}

	if maxPlayer {
		posValue = -9999
		for _, m := range pos.ValidMoves() {
			posValue = math.Max(posValue, minimax(pos.Update(m), depth - 1, false))
		}
	} else {
		posValue = 9999
		for _, m := range pos.ValidMoves() {
			posValue = math.Min(posValue, minimax(pos.Update(m), depth - 1, true))
		}
	}

	return posValue
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
