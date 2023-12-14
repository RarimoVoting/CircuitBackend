package cryptography

import (
	"io"
	"log"
	"math/big"
	"os"

	"github.com/iden3/go-iden3-crypto/babyjub"
)

const PRIVATE_KEY_FILE_NAME = "private.txt"

func GetPrivateKey() babyjub.PrivKeyScalar {
	privateKeyString, exist := getPrivateKeyIfExist(PRIVATE_KEY_FILE_NAME)

	if !exist {
		return initNewPrivateKey()
	}
	privateKeyInt := big.NewInt(0)
	privateKeyInt.SetString(privateKeyString, 10)
	return *babyjub.NewPrivKeyScalar(privateKeyInt)
}

func initNewPrivateKey() babyjub.PrivKeyScalar {
	privKey := babyjub.NewRandPrivKey()
	privKeyString := privKey.Scalar().BigInt().String()
	createAndWriteToFile(PRIVATE_KEY_FILE_NAME, privKeyString)
	return *privKey.Scalar()
}

func createAndWriteToFile(fileName string, data string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}

func getPrivateKeyIfExist(fileName string) (string, bool) {
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		return "", false
	}

	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Open(fileName)
	// Read data from file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Print the data
	return string(data), true
}
