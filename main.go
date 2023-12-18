package main

import (
	"redsunsetbackend/merkletree"
	"redsunsetbackend/requests"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	requests.MerkleTree = merkletree.NewMerkleTree(1)
	requests.MerkleTree.BuildMerkleTreeMock()

	e.GET("/verifyPhoto", requests.HandleVerifyPhoto)
	e.GET("/providerInclusionProof/:leafHash", requests.HandleProviderInclusionProof)
	e.GET("/providerMerkleRoot", requests.HandleProviderMerkleRoot)
	e.GET("/providerList", requests.HandleProviderList)
	// e.GET("/circuitTestInputs", requests.HandleCircuitTestInput)
	e.POST("/updateMerkleRoot", requests.HandleUpdateMerkleRoot)

	e.Start(":3000")
}
