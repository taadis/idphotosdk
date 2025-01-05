package idphotosdk

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestIdphotoCrop(t *testing.T) {

	t.Run("invalid_source", func(t *testing.T) {
		req := NewIdphotoCropRequest()
		_, err := client.IdphotoCrop(ctx, req)
		if err == nil {
			t.Fatalf("want non-nil but got %v", err)
		}
		if err.Error() != "invalid input image or base64 is empty" {
			t.Fatalf("want error msg but got:%v", err)
		}
	})

	t.Run("with-input-image", func(t *testing.T) {
		req := NewIdphotoCropRequest()
		req.InputImage = file
		req.Height = 413
		req.Width = 295
		req.FaceDetectModel = "mtcnn"
		req.Hd = true
		req.Dpi = 300
		req.HeadMeasureRatio = 0.2
		req.HeadHeightRatio = 0.45
		req.TopDistanceMax = 0.12
		req.TopDistanceMin = 0.1
		rsp, err := client.IdphotoCrop(ctx, req)
		if err != nil {
			t.Fatalf("request failed, error:%v", err)
		}
		if rsp == nil {
			t.Fatalf("want non-nil but got %v", rsp)
		}
		if rsp.Status != true {
			t.Fatalf("want true but got %v", rsp.Status)
		}
		if rsp.ImageBase64Standard == "" {
			t.Fatal("want image base64 standard but got empty")
		}
	})

	t.Run("with-input-image-base64", func(t *testing.T) {
		fileBytes, err := os.ReadFile(inputImagePath)
		if err != nil {
			t.Fatal(err)
		}
		base64Str := base64.StdEncoding.EncodeToString(fileBytes)
		base64Str = fmt.Sprintf("data:image/jpeg;base64,%s", base64Str)

		idphotoReq := NewIdphotoRequest()
		idphotoReq.InputImageBase64 = base64Str
		idphotoReq.Hd = false
		idphotoRsp, err := client.Idphoto(ctx, idphotoReq)
		if err != nil {
			t.Fatalf("request failed, error:%v", err)
		}
		if idphotoRsp == nil {
			t.Fatalf("want non-nil but got %v", idphotoRsp)
		}
		if idphotoRsp.Status != true {
			t.Fatalf("want true but got %v", idphotoRsp.Status)
		}
		if idphotoRsp.ImageBase64Standard == "" {
			t.Fatal("want image base64 standard but got empty")
		}

		idphotoCropReq := NewIdphotoCropRequest()
		idphotoCropReq.InputImageBase64 = idphotoRsp.ImageBase64Standard
		idphotoCropReq.Height = 413
		idphotoCropReq.Width = 295
		idphotoCropReq.FaceDetectModel = "mtcnn"
		idphotoCropReq.Hd = true
		idphotoCropReq.Dpi = 300
		idphotoCropReq.HeadMeasureRatio = 0.2
		idphotoCropReq.HeadHeightRatio = 0.45
		idphotoCropReq.TopDistanceMax = 0.12
		idphotoCropReq.TopDistanceMin = 0.1
		idphotoCropRsp, err := client.IdphotoCrop(ctx, idphotoCropReq)
		if err != nil {
			t.Fatalf("request failed, error:%v", err)
		}
		if idphotoCropRsp == nil {
			t.Fatalf("want non-nil but got %v", idphotoCropRsp)
		}
		if idphotoCropRsp.Status != true {
			t.Fatalf("want true but got %v", idphotoCropRsp.Status)
		}
		if idphotoCropRsp.ImageBase64Standard == "" {
			t.Fatal("want image base64 standard but got empty")
		}
		if idphotoCropRsp.ImageBase64Hd == "" {
			t.Fatal("want image base64 hd but got empty")
		}
	})

}
