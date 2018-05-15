package models

//go:generate reform

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"hash"
	"strconv"
	"strings"
)

//reform:players
type player struct {
	Id           int     `reform:"id,pk"`
	Nickname     *string `reform:"nickname,unique"`
	PasswordHash *string `reform:"password_hash"`
	Source       *string `reform:"source,index"`
}

func NewPlayer() *player {
	return &player{}
}

type PlayerI interface {
	GetPlayerId() int
	NewGame(invitiedPlayerId int) *game
	MyGamesScope() *gameScope
	VisibleGamesScope() *gameScope
}

const (
	SALT_SIZE   int = 8
	HASH_ROUNDS int = 79
)

var defaultHashFunct func() hash.Hash

func init() {
	defaultHashFunct = sha512.New
}

func (p player) GetPlayerId() int {
	return p.Id
}

func (p player) NewGame(invitedPlayerId int) *game {
	return NewGame(p.Id, invitedPlayerId)
}

func HashPassword(password string) string {
	return CryptoHash([]byte(password), nil, nil)
}

func genSalt() (salt []byte) {
	salt = make([]byte, SALT_SIZE)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	return
}

func CryptoHash(input []byte, salt []byte, hashfn func() hash.Hash) string {
	if hashfn == nil {
		hashfn = defaultHashFunct
	}

	if len(salt) == 0 {
		salt = genSalt()
	}

	return "sha512," + strconv.Itoa(HASH_ROUNDS) + "," + hex.EncodeToString(salt) + "," + hex.EncodeToString(pbkdf2.Key(input, salt, HASH_ROUNDS, 64, hashfn))
}

func ParseCryptoHash(fullhash string) (hashfn func() hash.Hash, salt []byte, hashself []byte) {
	var err error
	words := strings.Split(fullhash, ",")

	switch words[0] {
	case "sha512":
		hashfn = sha512.New
	default:
		panic(fmt.Errorf("Unknown hash algorithm: \"%s\"", words[0]))
	}

	salt, err = hex.DecodeString(words[2])
	if err != nil {
		panic(err)
	}

	hashself, err = hex.DecodeString(words[3])
	if err != nil {
		panic(err)
	}

	return hashfn, salt, hashself
}

func (player player) CheckPassword(password []byte) bool {
	if player.PasswordHash == nil {
		return false
	}

	hashfunct, salt, _ := ParseCryptoHash(*player.PasswordHash)
	passwordHash := CryptoHash(password, salt, hashfunct)

	return (passwordHash == *player.PasswordHash)
}

func (p player) MyGamesScope() *gameScope {
	return Game.Where(`players_pair_id IN (SELECT id FROM players_pairs WHERE player_id_0 = ? OR player_id_1 = ?)`, p.Id, p.Id)
}
func (p player) VisibleGamesScope() *gameScope {
	return Game.Where(`is_public = 1 OR id IN (SELECT game_id FROM watchers WHERE player_id = ?) OR players_pair_id IN (SELECT id FROM players_pairs WHERE player_id_0 = ? OR player_id_1 = ?)`, p.Id, p.Id, p.Id)
}
