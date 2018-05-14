package models

//go:generate reform

import (
	"fmt"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
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

type PlayerI interface {
	GetPlayerId() int
	NewGame(invitiedPlayerId int) *game
}

const (
	SALT_SIZE int = 8
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

	return "sha512,"+strconv.Itoa(HASH_ROUNDS)+"," + hex.EncodeToString(salt) + "," + hex.EncodeToString(pbkdf2.Key(input, salt, HASH_ROUNDS, 64, hashfn))
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
