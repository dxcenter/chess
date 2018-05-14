package models

//go:generate reform

import (
	"fmt"
	"github.com/notnil/chess"
)

var (
	ErrGameNotStarted = fmt.Errorf("game is not started")
)

//reform:games
type game struct {
	engine *chess.Game `reform:"-"`

	Id            int    `reform:"id,pk"`
	PlayersPairId int    `reform:"players_pair_id"`
	Status        string `reform:"status"`
}

func NewGame(playerId0, playerId1 int) (result *game) {
	result = &game{
		engine: chess.NewGame(),
	}

	result.setPlayersPair(player0, player1)
	err := result.Save()
	if err != nil {
		checkErr(err)
		return nil
	}
	return result
}

type Status struct {
	SquareMap map[chess.Square]chess.Piece
	History   []*chess.Position
}

func (g game) check() {
	if g.engine == nil {
		panic("The game is not initialized")
	}
}

func (g game) wrapperStatus(f func() Status) Status {
	g.check()
	return f()
}

func (g game) wrapperError(f func() error) Error {
	g.check()
	return f()
}

func (g game) GetStatus() Status {
	return g.wrapperStatus(func() {
		if g == nil {
			return Status{}
		}
		return Status{
			History:   g.engine.Positions(),
			SquareMap: g.engine.Position().Board().SquareMap(),
		}
	})
}

func (g *game) setPlayersPair(playerId0, playerId1 int) {
	g.PlayersPairId = getPlayersPair(playerId0, playerId1).Id
}

func (g *game) MoveStr(move string) error {
	return g.wrapperError(func() {
		if g == nil {
			return ErrGameNotStarted
		}
		return g.engine.MoveStr(move)
	})
}
