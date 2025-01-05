package idphotosdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"reflect"
	"time"
)

type HumanMattingModel string

const (
	HumanMattingModel_ModnetPhotographicPortraitMatting HumanMattingModel = "modnet_photographic_portrait_matting"
	HumanMattingModel_HivisionModnet                    HumanMattingModel = "hivision_modnet"
	HumanMattingModel_Rmbg1_4                           HumanMattingModel = "rmbg-1.4"
	HumanMattingModel_BirefnetV1Lite                    HumanMattingModel = "birefnet-v1-lite"
)

type IdphotoRequest struct {
	InputImage io.Reader `json:"-"`

	InputImageBase64 string `json:"-"`

	Height *int `json:"height"`

	Width *int `json:"width"`

	HumanMattingModel HumanMattingModel `json:"human_matting_model"`

	FaceDetectModel string `json:"face_detect_model"`

	Hd bool `json:"hd"`

	Dpi *int `json:"dpi"`

	FaceAlignment bool `json:"face_alignment"`

	HeadHeightRatio *float64 `json:"head_height_ratio"`

	HeadMeasureRatio *float64 `json:"head_measure_ratio"`

	TopDistanceMin *float64 `json:"top_distance_min"`

	TopDistanceMax *float64 `json:"top_distance_max"`

	BrightnessStrength *float64 `json:"brightness_strength"`

	ContrastStrength *float64 `json:"contrast_strength"`

	SharpenStrength *float64 `json:"sharpen_strength"`

	SaturationStrength *float64 `json:"saturation_strength"`
}

func (r *IdphotoRequest) Request(ctx context.Context) (*http.Request, error) {
	if r == nil {
		return nil, ErrRequestIsNil
	}

	method := http.MethodPost
	url := "idphoto"

	return request(ctx, method, url, r.InputImage, r.InputImageBase64, r)
}

func request(ctx context.Context, method string, url string, inputImage io.Reader, inputImageBase64 string, params interface{}) (*http.Request, error) {
	if inputImage == nil && inputImageBase64 == "" {
		return nil, errors.New("invalid input image or base64 is empty")
	}

	paramsBuf, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var fields map[string]interface{}
	err = json.Unmarshal(paramsBuf, &fields)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// that's not work
	// for k, v := range fields {
	// 	err = writer.WriteField(k, fmt.Sprintf("%v", v))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// using reflect to write struct fields
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("json")
		if fieldName == "" || fieldName == "-" {
			continue
		}
		fieldValue := v.Field(i).Interface()
		if reflect.ValueOf(fieldValue).Kind() == reflect.Ptr && reflect.ValueOf(fieldValue).IsNil() {
			continue
		}
		err = writer.WriteField(fieldName, fmt.Sprintf("%v", fieldValue))
		if err != nil {
			return nil, err
		}
		slog.Info("write field", "fieldName", fieldName, "fieldValue", fieldValue)
	}

	if inputImage != nil {
		filename := fmt.Sprintf("%d.jpg", time.Now().UnixNano())
		part, err := writer.CreateFormFile("input_image", filename)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, inputImage)
		if err != nil {
			return nil, err
		}
	} else {
		err = writer.WriteField("input_image_base64", inputImageBase64)
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

func NewIdphotoRequest() *IdphotoRequest {
	return &IdphotoRequest{}
}

type IdphotoResponse struct {
	Status bool `json:"status"`

	ImageBase64Standard string `json:"image_base64_standard,omitempty"`

	ImageBase64Hd string `json:"image_base64_hd,omitempty"`
}

func (c *Client) Idphoto(ctx context.Context, req *IdphotoRequest) (*IdphotoResponse, error) {
	return c.idphoto(ctx, req)
}

func (c *Client) idphoto(ctx context.Context, req *IdphotoRequest) (*IdphotoResponse, error) {
	var rsp IdphotoResponse
	err := c.Do(ctx, req, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}
