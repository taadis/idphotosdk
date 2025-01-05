package idphotosdk

import (
	"context"
	"io"
	"net/http"
)

type SetKbRequest struct {
	InputImage io.Reader `json:"-"`

	InputImageBase64 string `json:"-"`

	Kb int `json:"kb"`

	Dpi int `json:"dpi"`
}

func NewSetKbRequest() *SetKbRequest {
	return &SetKbRequest{}
}

func (r *SetKbRequest) Request(ctx context.Context) (*http.Request, error) {
	if r == nil {
		return nil, ErrRequestIsNil
	}
	method := http.MethodPost
	url := "set_kb"
	return request(ctx, method, url, r.InputImage, r.InputImageBase64, r)
}

type SetKbResponse struct {
	baseResponse

	ImageBase64 string `json:"image_base64"`
}

func (c *Client) SetKb(ctx context.Context, req *SetKbRequest) (*SetKbResponse, error) {
	return c.setKb(ctx, req)
}

func (c *Client) setKb(ctx context.Context, req *SetKbRequest) (*SetKbResponse, error) {
	var rsp SetKbResponse
	err := c.Do(ctx, req, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
