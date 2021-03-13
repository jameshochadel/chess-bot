package stockfishvs

import (
	"fmt"
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
