package idphotosdk

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestIdphoto(t *testing.T) {
	ctx := context.Background()
	client := NewClient(
		WithBaseUrl(getHivisionIdphotoBaseUrl()),
		WithLogEnabled(true),
	)

	t.Run("invalid_source", func(t *testing.T) {
		req := NewIdphotoRequest()
		_, err := client.Idphoto(ctx, req)
		if err == nil {
			t.Fatalf("want non-nil but got %v", err)
		}
		if err.Error() != "invalid input image or base64 is empty" {
			t.Fatalf("want error msg but got:%v", err)
		}
	})

	t.Run("with-input-image-standard", func(t *testing.T) {
		req := NewIdphotoRequest()
		req.InputImage = file
		req.Height = 413
		req.Width = 295
		req.Hd = false
		req.HumanMattingModel = HumanMattingModel_ModnetPhotographicPortraitMatting
		req.FaceDetectModel = FaceDetectModel_Mtcnn
		rsp, err := client.Idphoto(ctx, req)
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
		req := NewIdphotoRequest()
		req.InputImageBase64 = base64Str
		rsp, err := client.Idphoto(ctx, req)
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

}
