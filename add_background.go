package idphotosdk

import (
	"context"
	"io"
	"net/http"
)

type AddBackgroundRequest struct {
	InputImage       io.Reader `json:"-"`
	InputImageBase64 string    `json:"-"`

	Color string `json:"color"`

	Kb int `json:"kb"`

	Render int `json:"render"`

	Dpi int `json:"dpi"`
}

func (r *AddBackgroundRequest) Request(ctx context.Context) (*http.Request, error) {
	if r == nil {
		return nil, ErrRequestIsNil
	}

	method := http.MethodPost
	url := "add_background"
	return request(ctx, method, url, r.InputImage, r.InputImageBase64, r)
}

func NewAddBackgroundRequest() *AddBackgroundRequest {
	return &AddBackgroundRequest{}
}

type AddBackgroundResponse struct {
	baseResponse
	ImageBase64 string `json:"image_base64"`
}

func (c *Client) AddBackground(ctx context.Context, req *AddBackgroundRequest) (*AddBackgroundResponse, error) {
	return c.addBackground(ctx, req)
}

func (c *Client) addBackground(ctx context.Context, req *AddBackgroundRequest) (*AddBackgroundResponse, error) {
	var rsp AddBackgroundResponse
	err := c.Do(ctx, req, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}
