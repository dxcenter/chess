package models

//go:generate reform

import (
	"database/sql"
	"fmt"
	"github.com/notnil/chess"
)

var (
	ErrGameNotStarted = fmt.Errorf("game is not started")
)

//reform:games
type game struct {
	engine *chess.Game `reform:"-"`

	Id              int    `reform:"id,pk"`
	PlayersPairId   int    `reform:"players_pair_id"`
	Status          string `reform:"status"`
	InvitedPlayerId *int   `reform:"invited_player_id"`
}

var (
	games        map[int]*game
	updatedGames map[int]*game
)

func init() {
	games = map[int]*game{}
	updatedGames = map[int]*game{}
}

func NewGame(initiatorPlayerId, invitedPlayerId int) *game {
	result := &game{
		engine: chess.NewGame(),
	}

	result.InvitedPlayerId = &invitedPlayerId
	result.setPlayersPair(initiatorPlayerId, invitedPlayerId)
	err := result.Save()
	if err != nil {
		panic(err)
	}

	return result
}

func prefetchGame(gameId int) bool {
	if games[gameId] != nil {
		return true
	}

	game, err := Game.First(gameId)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		panic(err)
	}

	games[gameId] = &game
	return true
}

func GetGame(gameId int) *game {
	prefetchGame(gameId)
	return games[gameId]
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

func (g game) wrapperError(f func() error) error {
	g.check()
	return f()
}

func (g game) GetStatus() Status {
	return g.wrapperStatus(func() Status {
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
	return g.wrapperError(func() error {
		return g.engine.MoveStr(move)
	})
}
