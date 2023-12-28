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

	e.POST("/verifyPhoto", requests.HandleVerifyPhoto)
	e.GET("/providerInclusionProof/:leafHash", requests.HandleProviderInclusionProof)
	e.GET("/providerMerkleRoot", requests.HandleProviderMerkleRoot)
	e.GET("/providerList", requests.HandleProviderList)
	e.POST("/updateMerkleRoot", requests.HandleUpdateMerkleRoot)
	e.POST("/storeBytes", requests.HandleStoreBytes)

	e.Start(":80")
}
