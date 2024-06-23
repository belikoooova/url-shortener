package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"strings"
)

type Sha256WithBase58Shortener struct {
}

func NewSha256WithBase58Shortener() *Sha256WithBase58Shortener {
	return &Sha256WithBase58Shortener{}
}

func (sh Sha256WithBase58Shortener) Shorten(url string) (string, error) {
	url = strings.TrimSpace(url)

	algo := sha256.New()
	algo.Write([]byte(url))

	generatedNumber := new(big.Int).SetBytes(algo.Sum(nil)).Uint64()

	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}
