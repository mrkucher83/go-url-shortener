package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

// sha256f gets hash from user input
func sha256f(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// base58Encoded encodes hash using base58
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateShortLink(initialLink, userId string) string {
	urlHashBytes := sha256f(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
