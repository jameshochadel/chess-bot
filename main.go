package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

func main() {
	eng, err := uci.New("stockfish")
	if err != nil {
		panic(err)
	}
	defer eng.Close()

	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}

	game := chess.NewGame()
	for game.Outcome() == chess.NoOutcome {
		cmdPos := uci.CmdPosition{Position: game.Position()}
		cmdGo := uci.CmdGo{MoveTime: time.Second / 100}
		if err = eng.Run(cmdPos, cmdGo); err != nil {
			panic(err)
		}
		move := eng.SearchResults().BestMove
		if err := game.Move(move); err != nil {
			panic(err)
		}
	}

	fmt.Println(game.String())
}

func minimax(pos chess.Position, depth int, maxPlayer bool) (bestMove chess.Move, posValue float32) {
	/*
	if depth == 0 or len(game.ValidMoves) == 0,
		return evaluatePosition(pos)
	if maxPlayer {

	} else {

	}
	*/
	return chess.Move{}, 0
}

var pieceVals = map[string]float32 {
	"p": 1,
	"n": 3,
	"b": 3,
	"r": 6,
	"q": 9,
}

func evaluatePosition(pos chess.Position) float32 {
	str := pos.String()

	var acc float32

	for piece, value := range pieceVals {
		acc += float32(strings.Count(str, piece)) * value
	}
	strings.Count(str, "q")
	// count number of each piece
	return 0.0
}
