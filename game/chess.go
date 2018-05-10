package game

import (
	"github.com/notnil/chess"
)

var g *chess.Game

func NewGame() {
	g = chess.NewGame()
}
func Get() chess.Game {
	return *g
}

type Status struct {
	SquareMap map[chess.Square]chess.Piece
	History []*chess.Position
}

func GetStatus() Status {
	return Status{
		History:   g.Positions(),
		SquareMap: g.Position().Board().SquareMap(),
	}
}

func MoveStr(move string) error {
	return g.MoveStr(move)
}
