package cryptography

import (
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

func PoseidonHash(data []byte) *big.Int {
	inputBigInt := big.NewInt(0)
	inputBigInt.SetBytes(data)
	input := []*big.Int{inputBigInt}

	hash, _ := poseidon.Hash(input)

	return hash
}
