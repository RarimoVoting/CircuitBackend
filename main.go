package main

import (
	"redsunsetbackend/requests"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/verifyPhoto", requests.HandleVerifyPhoto)
	e.GET("/signerInclusionProof", requests.HandleSignerInclusionProof)
	e.GET("/providerMerkleRoot", requests.HandleSignerInclusionProof)
	e.GET("/providerList", requests.HandleSignerInclusionProof)
	e.POST("/updateMerkleRoot", requests.HandleUpdateMerkleRoot)

	e.Start(":3000")
}
