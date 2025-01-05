package idphotosdk

import (
	"context"
	"io"
	"net/http"
)

type HumanMattingRequest struct {
	InputImage io.Reader `json:"-"`

	InputImageBase64 string `json:"-"`

	HumanMattingModel string `json:"human_matting_model"`

	Dpi int `json:"dpi,omitempty"`
}

func NewHumanMattingRequest() *HumanMattingRequest {
	return &HumanMattingRequest{}
}

func (r *HumanMattingRequest) Request(ctx context.Context) (*http.Request, error) {
	if r == nil {
		return nil, ErrRequestIsNil
	}

	method := http.MethodPost
	url := "human_matting"
	return request(ctx, method, url, r.InputImage, r.InputImageBase64, r)
}

type HumanMattingResponse struct {
	baseResponse

	ImageBase64 string `json:"image_base64"`
}

func (c *Client) HumanMatting(ctx context.Context, req *HumanMattingRequest) (*HumanMattingResponse, error) {
	return c.humanMatting(ctx, req)
}

func (c *Client) humanMatting(ctx context.Context, req *HumanMattingRequest) (*HumanMattingResponse, error) {
	var rsp HumanMattingResponse
	err := c.Do(ctx, req, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
