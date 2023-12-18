package merkletree

import (
	"fmt"
	"math/big"
	"redsunsetbackend/cryptography"
)

type MerkleTree struct {
	Tree            []*big.Int
	Leaves          []*big.Int
	TREE_DEPTH      int
	LEAVES_SIZE     int
	TREE_ARRAY_SIZE int
}

func NewMerkleTree(treeDepth int) *MerkleTree {
	mt := &MerkleTree{}
	mt.TREE_DEPTH = treeDepth
	mt.LEAVES_SIZE = 1 << mt.TREE_DEPTH
	mt.TREE_ARRAY_SIZE = mt.LEAVES_SIZE * 2
	return mt
}

func (mt *MerkleTree) BuildMerkleTree(array []*big.Int) {
	mt.Leaves = array
	size := len(array)
	mt.Tree = make([]*big.Int, size*2)

	mt.buildRecursive(0)
}

func (mt *MerkleTree) buildRecursive(current int) {
	if current >= len(mt.Tree)-len(mt.Leaves)-1 {
		mt.Tree[current] = mt.Leaves[current-mt.LEAVES_SIZE+1]
		return
	}
	mt.buildRecursive(current*2 + 1)
	mt.buildRecursive(current*2 + 2)
	mt.Tree[current] = cryptography.PoseidonHashLeftRight(mt.Tree[current*2+1], mt.Tree[current*2+2])
}

func (mt *MerkleTree) GetMerkleBranch(leafValue *big.Int) ([]*big.Int, int, bool) {
	for i := 0; i < mt.LEAVES_SIZE; i++ {
		if mt.Tree[mt.LEAVES_SIZE+i].String() == leafValue.String() {
			answer := []*big.Int{}
			index := mt.LEAVES_SIZE + i
			for index > 1 {
				pair := index - 1
				if index%2 == 1 {
					pair = index + 1
				}
				answer = append(answer, mt.Tree[pair])
				index = index / 2
			}
			return answer, mt.LEAVES_SIZE + 1, true
		}
	}
	return []*big.Int{}, 0, false
}

func (mt *MerkleTree) PrintMerkleTree() {
	for i := 0; i < len(mt.Tree); i++ {
		fmt.Println(i, mt.Tree[i])
	}
}

// 3 >= 8 / 2 - 1
//      0
//   1     2
//  3 4   5 6
