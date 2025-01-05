package idphotosdk

import (
	"context"
	"io"
	"net/http"
)

type IdphotoCropRequest struct {
	InputImage io.Reader `json:"-"`

	InputImageBase64 string `json:"-"`

	Height int `json:"height"`

	Width int `json:"width"`

	FaceDetectModel string `json:"face_detect_model"`

	Hd bool `json:"hd"`

	Dpi int `json:"dpi"`

	HeadMeasureRatio float64 `json:"head_measure_ratio"`

	HeadHeightRatio float64 `json:"head_height_ratio"`

	TopDistanceMax float64 `json:"top_distance_max"`

	TopDistanceMin float64 `json:"top_distance_min"`
}

func NewIdphotoCropRequest() *IdphotoCropRequest {
	return &IdphotoCropRequest{}
}

func (r *IdphotoCropRequest) Request(ctx context.Context) (*http.Request, error) {
	if r == nil {
		return nil, ErrRequestIsNil
	}
	method := http.MethodPost
	url := "idphoto_crop"
	return request(ctx, method, url, r.InputImage, r.InputImageBase64, r)
}

type IdphotoCropResponse struct {
	baseResponse

	ImageBase64Standard string `json:"image_base64_standard"`

	ImageBase64Hd string `json:"image_base64_hd"`
}

func (c *Client) IdphotoCrop(ctx context.Context, req *IdphotoCropRequest) (*IdphotoCropResponse, error) {
	return c.idphotoCrop(ctx, req)
}

func (c *Client) idphotoCrop(ctx context.Context, req *IdphotoCropRequest) (*IdphotoCropResponse, error) {
	var rsp IdphotoCropResponse
	err := c.Do(ctx, req, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
