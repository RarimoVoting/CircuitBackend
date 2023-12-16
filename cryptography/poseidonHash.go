package cryptography

import (
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

func PoseidonHash(data []byte) *big.Int {
	inputBigInt := big.NewInt(0)
	inputBigInt.SetBytes(data)
	input := []*big.Int{inputBigInt}

	hash, _ := poseidon.Hash(input)

	return hash
}

func PoseidonHashLeftRight(left *big.Int, right *big.Int) *big.Int {
	input := []*big.Int{left, right}

	hash, _ := poseidon.Hash(input)
	return hash
}

func PoseidonHashPoint(point *babyjub.Point) *big.Int {
	return PoseidonHashLeftRight(point.X, point.Y)
}
