package idphotosdk

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestSetKb(t *testing.T) {

	t.Run("invalid_source", func(t *testing.T) {
		req := NewSetKbRequest()
		_, err := client.SetKb(ctx, req)
		if err == nil {
			t.Fatalf("want non-nil but got %v", err)
		}
		if errors.Is(err, ErrInvalidInputImage) {
			t.Fatalf("want error msg but got:%v", err)
		}
	})

	t.Run("with-input-image", func(t *testing.T) {
		req := NewSetKbRequest()
		req.InputImage = file
		req.Kb = 0
		req.Dpi = 300
		rsp, err := client.SetKb(ctx, req)
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

		req := NewSetKbRequest()
		req.InputImageBase64 = base64Str
		req.Kb = 0
		req.Dpi = 300
		rsp, err := client.SetKb(ctx, req)
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

}
