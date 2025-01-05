package idphotosdk

import (
	"context"
	"io"
	"net/http"
)

type WatermarkRequest struct {
	InputImage io.Reader `json:"-"`

	InputImageBase64 string `json:"-"`

	Text string `json:"text"`

	Size int `json:"size"`

	Opacity float64 `json:"opacity"`

	Angle int `json:"angle"`

	Color string `json:"color"`

	Space int `json:"space"`

	Dpi int `json:"dpi"`
}

func NewWatermarkRequest() *WatermarkRequest {
	return &WatermarkRequest{}
}

func (r *WatermarkRequest) Request(ctx context.Context) (*http.Request, error) {
	if r == nil {
		return nil, ErrRequestIsNil
	}

	method := http.MethodPost
	url := "watermark"
	return request(ctx, method, url, r.InputImage, r.InputImageBase64, r)
}

type WatermarkResponse struct {
	baseResponse

	ImageBase64 string `json:"image_base64"`
}

func (c *Client) Watermark(ctx context.Context, req *WatermarkRequest) (*WatermarkResponse, error) {
	return c.watermark(ctx, req)
}

func (c *Client) watermark(ctx context.Context, req *WatermarkRequest) (*WatermarkResponse, error) {
	var rsp WatermarkResponse
	err := c.Do(ctx, req, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
