package idphotosdk

import "errors"

var (
	ErrInvalidInputImage = errors.New("invalid input image or base64 is empty")

	ErrRequestIsNil = errors.New("request is nil")
)
