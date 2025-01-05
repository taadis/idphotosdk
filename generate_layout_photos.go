package idphotosdk

import (
	"context"
	"io"
	"net/http"
)

type GenerateLayoutPhotosRequest struct {
	InputImage io.Reader `json:"-"`

	InputImageBase64 string `json:"-"`

	Height *int `json:"height,omitempty"`

	Width *int `json:"width,omitempty"`

	Kb *int `json:"kb,omitempty"`

	Dpi *int `json:"dpi,omitempty"`
}

func (r *GenerateLayoutPhotosRequest) Request(ctx context.Context) (*http.Request, error) {
	if r == nil {
		return nil, ErrRequestIsNil
	}

	method := http.MethodPost
	url := "generate_layout_photos"
	return request(ctx, method, url, r.InputImage, r.InputImageBase64, r)
}

func NewGenerateLayoutPhotosRequest() *GenerateLayoutPhotosRequest {
	return &GenerateLayoutPhotosRequest{}
}

type GenerateLayoutPhotosResponse struct {
	baseResponse

	ImageBase64 string `json:"image_base64"`
}

func (c *Client) GenerateLayoutPhotos(ctx context.Context, req *GenerateLayoutPhotosRequest) (*GenerateLayoutPhotosResponse, error) {
	return c.generateLayoutPhotos(ctx, req)
}

func (c *Client) generateLayoutPhotos(ctx context.Context, req *GenerateLayoutPhotosRequest) (*GenerateLayoutPhotosResponse, error) {
	var rsp GenerateLayoutPhotosResponse
	err := c.Do(ctx, req, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
