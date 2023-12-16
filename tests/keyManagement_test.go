package tests

import (
	"math/big"
	"redsunsetbackend/cryptography"
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestKeyManagementPrivateKeyStorage(t *testing.T) {
	privateKey1 := cryptography.GetPrivateKey()
	privateKey2 := cryptography.GetPrivateKey()
	assert(t, privateKey1, privateKey2)
	privateKey3 := cryptography.GetPrivateKey()
	assert(t, privateKey1, privateKey3)
}

func TestPoseidon(t *testing.T) {
	input := big.NewInt(1)
	hashLeftRight := cryptography.PoseidonHashLeftRight(input, input)

	result := big.NewInt(0)
	result.SetString("217234377348884654691879377518794323857294947151490278790710809376325639809", 10)
	assert(t, hashLeftRight, result)
}
