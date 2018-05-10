package game

import (
	"fmt"
	"github.com/notnil/chess"
)

var (
	ErrGameNotStarted = fmt.Errorf("game is not started")
)

var g *chess.Game

func NewGame() {
	g = chess.NewGame()
	chess.UseNotation(&chess.LongAlgebraicNotation{})(g)
}

type Status struct {
	SquareMap map[chess.Square]chess.Piece
	History []*chess.Position
}

func GetStatus() Status {
	if g == nil {
		return Status{}
	}
	return Status{
		History:   g.Positions(),
		SquareMap: g.Position().Board().SquareMap(),
	}
}

func MoveStr(move string) error {
	if g == nil {
		return ErrGameNotStarted
	}
	return g.MoveStr(move)
}
