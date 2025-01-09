package idphotosdk

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestAddBackground(t *testing.T) {
	ctx := context.Background()
	client := NewClient(
		WithBaseUrl(getHivisionIdphotoBaseUrl()),
		WithLogEnabled(true),
	)

	t.Run("invalid_source", func(t *testing.T) {
		req := NewAddBackgroundRequest()
		_, err := client.AddBackground(ctx, req)
		if err == nil {
			t.Fatalf("want non-nil but got %v", err)
		}
		if err.Error() != "invalid input image or base64 is empty" {
			t.Fatalf("want error msg but got:%v", err)
		}
	})

	t.Run("with-input-image", func(t *testing.T) {
		req := NewAddBackgroundRequest()
		req.InputImage = file
		rsp, err := client.AddBackground(ctx, req)
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

		addBackgroundReq := NewAddBackgroundRequest()
		addBackgroundReq.InputImageBase64 = idphotoRsp.ImageBase64Standard
		addBackgroundRsp, err := client.AddBackground(ctx, addBackgroundReq)
		if err != nil {
			t.Fatalf("request failed, error:%v", err)
		}
		if addBackgroundRsp == nil {
			t.Fatalf("want non-nil but got %v", addBackgroundRsp)
		}
		if addBackgroundRsp.Status != true {
			t.Fatalf("want true but got %v", addBackgroundRsp.Status)
		}
		if addBackgroundRsp.ImageBase64 == "" {
			t.Fatal("want image base64 standard but got empty")
		}
	})

}
