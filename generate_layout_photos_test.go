package idphotosdk

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestGenerateLayoutPhotos(t *testing.T) {
	t.Run("invalid_source", func(t *testing.T) {
		req := NewGenerateLayoutPhotosRequest()
		_, err := client.GenerateLayoutPhotos(ctx, req)
		if err == nil {
			t.Fatalf("want non-nil but got %v", err)
		}
		if err.Error() != "invalid input image or base64 is empty" {
			t.Fatalf("want error msg but got:%v", err)
		}
	})

	t.Run("with-input-image", func(t *testing.T) {
		req := NewGenerateLayoutPhotosRequest()
		req.InputImage = file
		rsp, err := client.GenerateLayoutPhotos(ctx, req)
		if err != nil {
			t.Fatalf("request failed, error:%v", err)
		}
		if rsp == nil {
			t.Fatalf("want non-nil but got %v", rsp)
		}
		if rsp.Status != true {
			t.Fatalf("want true but got %v", rsp.Status)
		}
		if rsp.ImageBase64 == "" {
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

		idphotoReq := NewGenerateLayoutPhotosRequest()
		idphotoReq.InputImageBase64 = base64Str
		idphotoRsp, err := client.GenerateLayoutPhotos(ctx, idphotoReq)
		if err != nil {
			t.Fatalf("request failed, error:%v", err)
		}
		if idphotoRsp == nil {
			t.Fatalf("want non-nil but got %v", idphotoRsp)
		}
		if idphotoRsp.Status != true {
			t.Fatalf("want true but got %v", idphotoRsp.Status)
		}
		if idphotoRsp.ImageBase64 == "" {
			t.Fatal("want image base64 standard but got empty")
		}
	})

}
