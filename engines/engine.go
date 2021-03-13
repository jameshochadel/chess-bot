package engines

import (
	"github.com/notnil/chess"
)

type Engine interface {
	SuggestedMove(pos *chess.Position, maxPlayer bool) *chess.Move
}
