package tests

import (
	"fmt"
	"math/big"
	"redsunsetbackend/cryptography"
	"redsunsetbackend/merkletree"
	"testing"
)

func TestMerkleTree(t *testing.T) {
	mt := merkletree.NewMerkleTree(8)
	keyArray := make([]*big.Int, mt.LEAVES_SIZE)
	for i := 0; i < mt.LEAVES_SIZE; i++ {
		keyArray[i] = cryptography.PoseidonHashPoint(cryptography.GetPublicKey())
	}
	mt.BuildMerkleTree(keyArray)
	branch, _, isOk := mt.GetMerkleBranch(keyArray[0])

	assert(t, isOk, true)
	for i := 0; i < len(branch)-1; i++ {
		assert(t, cryptography.PoseidonHashLeftRight(branch[i], branch[i]), branch[i+1])
	}
	fmt.Println(branch)
	mt.PrintMerkleTree()
}
