package requests

import (
	"encoding/json"
	"math/big"
	"net/http"
	"redsunsetbackend/cryptography"
	"redsunsetbackend/merkletree"
	"redsunsetbackend/messages"
	"redsunsetbackend/verification"

	"github.com/labstack/echo/v4"
)

func HandleVerifyPhoto(context echo.Context) error {
	var imageVerificationRequest ImageVerificationRequest
	if err := json.NewDecoder(context.Request().Body).Decode(&imageVerificationRequest); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": messages.UNABLE_TO_PARSE_REQUEST,
		})
	}
	if !verification.Verify(imageVerificationRequest.PhotoReal.ImageBytes, imageVerificationRequest.PhotoPassport.ImageBytes) {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": messages.VERIFICATION_PHOTO_FAILED,
		})
	}
	hashRealPhoto := cryptography.PoseidonHash(imageVerificationRequest.PhotoReal.ImageBytes)
	hashPassportPhoto := cryptography.PoseidonHash(imageVerificationRequest.PhotoPassport.ImageBytes)

	verificationHash := cryptography.PoseidonHashLeftRight(hashRealPhoto, hashPassportPhoto)

	signature := cryptography.EddsaSignature(verificationHash)

	pubKeyHash := cryptography.PoseidonHashLeftRight(cryptography.GetPublicKey().X, cryptography.GetPublicKey().Y)

	branch, order, isOk := MerkleTree.GetMerkleBranch(pubKeyHash)

	if !isOk {
		return context.JSON(http.StatusInternalServerError, map[string]any{
			"msg": messages.UNABLE_TO_VERIFY_PROVIDER,
		})
	}

	return context.JSON(http.StatusOK, map[string]any{
		"hashRealPhoto":     hashRealPhoto.String(),
		"hashPassportPhoto": hashPassportPhoto.String(),
		"photoHash":         verificationHash.String(),
		"signature":         signature,
		"root":              MerkleTree.GetMerkleRoot(),
		"branch":            branch,
		"order":             order,
		"pubKeyHash":        pubKeyHash,
		"pubKey":            cryptography.GetPublicKey(),
	})
}

var MerkleTree *merkletree.MerkleTree

func HandleProviderInclusionProof(context echo.Context) error {
	if MerkleTree == nil {
		return context.JSON(http.StatusInternalServerError, map[string]any{
			"msg": messages.MT_NOT_DEFINED,
		})
	}
	leafHash := big.NewInt(0)
	leafHashString := context.Param("leafHash")

	if leafHashString == "" {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": messages.UNABLE_TO_FIND_LEAFHASH_PARAM,
		})
	}
	leafHash.SetString(leafHashString, 10)
	branch, order, isOk := MerkleTree.GetMerkleBranch(leafHash)

	if !isOk {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": messages.UNABLE_TO_FIND_LEAF,
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"branch": branch,
		"order":  order,
	})
}

func HandleProviderMerkleRoot(context echo.Context) error {
	if MerkleTree == nil {
		return context.JSON(http.StatusInternalServerError, map[string]any{
			"msg": messages.MT_NOT_DEFINED,
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"root": MerkleTree.GetMerkleRoot(),
	})
}

func HandleProviderList(context echo.Context) error {
	if MerkleTree == nil {
		return context.JSON(http.StatusInternalServerError, map[string]any{
			"msg": messages.MT_NOT_DEFINED,
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"providers": MerkleTree.GetProviderList(),
	})
}

func HandleUpdateMerkleRoot(context echo.Context) error {
	return nil
}
