package main

import (
	"fmt"
	"log"

	"github.com/notnil/chess"

	"github.com/jameshochadel/chess-bot/engine"
)

func main() {
	game := chess.NewGame()

	for game.Outcome() == chess.NoOutcome {
		// how to do context with timeout? basically want to go as far as possible until timeout.
		// would need to return some results eagerly though?
		// maybe benchmark to get a sense of how far it can get
		bestMove := engine.BestMove(game.Position(), game.Position().Turn() == chess.White)
		if err := game.Move(bestMove); err != nil {
			wrapped := fmt.Errorf("making move %v: %w", bestMove, err)
			log.Fatalf("game state %v, err %v", game.String(), wrapped)
		}
	}

	fmt.Println(game.String())
}
