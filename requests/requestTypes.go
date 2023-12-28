package requests

type ImageVerificationRequest struct {
	PhotoReal     Image
	PhotoPassport Image
}

type Image struct {
	ImageBytes []byte
}

type SerializedPassport struct {
	PassportBytes []byte
}
