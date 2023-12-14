package main

import (
	"redsunsetbackend/cryptography"
	"redsunsetbackend/requests"

	"github.com/labstack/echo/v4"
)

func main() {
	cryptography.GetPrivateKey()
	e := echo.New()

	e.GET("/verifyPhoto", requests.HandleVerifyPhoto)
	e.GET("/signerInclusionProof", requests.HandleSignerInclusionProof)
	e.GET("/providerMerkleRoot", requests.HandleSignerInclusionProof)
	e.GET("/providerList", requests.HandleSignerInclusionProof)
	e.POST("/updateMerkleRoot", requests.HandleUpdateMerkleRoot)

}
