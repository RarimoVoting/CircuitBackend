package cryptography

import (
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
)

func EddsaSignature(hash *big.Int) *babyjub.Signature {
	privKey := GetPrivateKey()
	signature := privKey.SignPoseidon(hash)

	return signature
}
