package tests

import (
	"redsunsetbackend/cryptography"
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestKeyManagement(t *testing.T) {
	privateKey1 := cryptography.GetPrivateKey()
	privateKey2 := cryptography.GetPrivateKey()
	assert(t, privateKey1, privateKey2)
	privateKey3 := cryptography.GetPrivateKey()
	assert(t, privateKey1, privateKey3)
}
