package requests

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"os"
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

	providerSignatureView := []string{signature.R8.X.String(), signature.R8.Y.String(), cryptography.GetPublicKey().X.String(), cryptography.GetPublicKey().Y.String(), signature.S.String()}
	providerMerkleBranchView := []string{}

	for i := 0; i < len(branch); i++ {
		providerMerkleBranchView = append(providerMerkleBranchView, branch[i].String())
	}

	return context.JSON(http.StatusOK, map[string]any{
		"realPhotoHash":        hashRealPhoto.String(),
		"passPhotoHash":        hashPassportPhoto.String(),
		"providerSignature":    providerSignatureView,
		"providerMerkleRoot":   MerkleTree.GetMerkleRoot().String(),
		"providerMerkleBranch": providerMerkleBranchView,
		"providerMerkleOrder":  fmt.Sprint(order),
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

func HandleStoreBytes(context echo.Context) error {
	fmt.Println("Handing Store Bytes request..")
	var serializedPassport SerializedPassport
	if err := json.NewDecoder(context.Request().Body).Decode(&serializedPassport); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": messages.UNABLE_TO_PARSE_REQUEST,
		})
	}
	// fmt.Println(serializedPassport.PassportBytes)
	passportName := getRandomPassportName()
	WriteFile(serializedPassport.PassportBytes, passportName)

	fmt.Println("OK")

	return context.JSON(http.StatusOK, map[string]any{
		"msg": "Data has been stored",
		"id":  passportName,
	})
}

func getRandomPassportName() string {
	return fmt.Sprint(rand.Int63()) + "passport"
}

func WriteFile(data []byte, filename string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		file.Write(data)
		fmt.Println("Done")
	}
	file.Close()
	os.Rename(filename, "./storedPassports/"+filename)
}

func HandleGetBytes(context echo.Context) error {
	storageId := context.Param("id")

	fmt.Println(storageId, "<---")

	if storageId == "" {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": messages.UNABLE_TO_PARSE_REQUEST,
		})
	}

	filename := "./storedPassports/" + storageId

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to parse the file:", err)
		return context.JSON(http.StatusBadRequest, map[string]any{
			"msg": messages.UNABLE_TO_FIND_ID,
		})
	}

	return context.JSON(http.StatusOK, map[string]any{
		"data": fmt.Sprint(data),
	})

}
