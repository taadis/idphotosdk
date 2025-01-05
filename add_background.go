package idphotosdk

import (
	"context"
	"io"
	"net/http"
)

type AddBackgroundRequest struct {
	InputImage       io.Reader `json:"-"`
	InputImageBase64 string    `json:"-"`

	Color *string `json:"color,omitempty"`

	Kb *int `json:"kb,omitempty"`

	Render *int `json:"render,omitempty"`

	Dpi *int `json:"dpi,omitempty"`
}

func (r *AddBackgroundRequest) Request(ctx context.Context) (*http.Request, error) {
	if r == nil {
		return nil, ErrRequestIsNil
	}

	method := http.MethodPost
	url := "add_background"
	return request(ctx, method, url, r.InputImage, r.InputImageBase64, r)
}

func NewAddBackgroudRequest() *AddBackgroundRequest {
	return &AddBackgroundRequest{}
}

type AddBackgroundResponse struct {
	baseResponse
	ImageBase64 string `json:"image_base64"`
}

func (c *Client) AddBackgroud(ctx context.Context, req *AddBackgroundRequest) (*AddBackgroundResponse, error) {
	return c.addBackgroud(ctx, req)
}

func (c *Client) addBackgroud(ctx context.Context, req *AddBackgroundRequest) (*AddBackgroundResponse, error) {
	var rsp AddBackgroundResponse
	err := c.Do(ctx, req, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}
