package requests

import (
	"encoding/json"
	"math/big"
	"net/http"
	"redsunsetbackend/cryptography"
	"redsunsetbackend/merkletree"
	"redsunsetbackend/verification"

	"github.com/labstack/echo/v4"
)

func HandleVerifyPhoto(context echo.Context) error {
	var imageVerificationRequest ImageVerificationRequest
	if err := json.NewDecoder(context.Request().Body).Decode(&imageVerificationRequest); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": "Unable to parse a request",
		})
	}
	if !verification.Verify(imageVerificationRequest.PhotoReal.ImageBytes, imageVerificationRequest.PhotoPassport.ImageBytes) {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": "Verification of the provided photo failed",
		})
	}
	hashRealPhoto := cryptography.PoseidonHash(imageVerificationRequest.PhotoReal.ImageBytes)
	hashPassportPhoto := cryptography.PoseidonHash(imageVerificationRequest.PhotoPassport.ImageBytes)

	verificationHash := cryptography.PoseidonHashLeftRight(hashRealPhoto, hashPassportPhoto)

	signature := cryptography.EddsaSignature(verificationHash)

	return context.JSON(http.StatusOK, map[string]any{
		"photoHash": verificationHash.String(),
		"signature": signature,
	})
}

var MerkleTree *merkletree.MerkleTree

func HandleProviderInclusionProof(context echo.Context) error {
	if MerkleTree == nil {
		return context.JSON(http.StatusInternalServerError, map[string]any{
			"msg": "Merkle Tree is NOT defined on the server",
		})
	}
	leaveHash := big.NewInt(0)
	leaveHashString := context.Param("leaveHash")

	if leaveHashString == "" {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": "Unable to retrieve leaveHash param",
		})
	}
	leaveHash.SetString(leaveHashString, 10)
	branch, order, isOk := MerkleTree.GetMerkleBranch(leaveHash)

	if !isOk {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": "Unable to find requested leave",
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"branch": branch,
		"order":  order,
	})
}

func HandleProviderMerkleRoot(context echo.Context) error {
	return nil
}

func HandleProviderList(context echo.Context) error {
	return nil
}

func HandleUpdateMerkleRoot(context echo.Context) error {
	return nil
}
