package requests

import (
	"encoding/json"
	"net/http"
	"redsunsetbackend/cryptography"
	"redsunsetbackend/verification"

	"github.com/labstack/echo/v4"
)

func HandleVerifyPhoto(context echo.Context) error {
	var imageVerificationRequest ImageVerificationRequest
	if err := json.NewDecoder(context.Request().Body).Decode(&imageVerificationRequest); err != nil {
		return err
	}
	if !verification.Verify(imageVerificationRequest.PhotoReal.ImageBytes, imageVerificationRequest.PhotoPassport.ImageBytes) {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": "Verification of the provided photo failed",
		})
	}
	hashRealPhoto := cryptography.PoseidonHash(imageVerificationRequest.PhotoReal.ImageBytes)
	hashPassportPhoto := cryptography.PoseidonHash(imageVerificationRequest.PhotoPassport.ImageBytes)

	verificationHash := cryptography.PoseidonHashLeftRight(hashRealPhoto, hashPassportPhoto)

	signature := cryptography.EddsaSignature(*verificationHash)

	return context.JSON(http.StatusOK, map[string]any{
		"photoHash": verificationHash.String(),
		"signature": signature.String(),
	})
}

func HandleSignerInclusionProof(context echo.Context) error {
	return nil
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
