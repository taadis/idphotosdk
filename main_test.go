package idphotosdk

import (
	"context"
	"os"
	"testing"
)

const HIVISION_IDPHOTO_BASE_URL = "HIVISION_IDPHOTO_BASE_URL"
const HIVISION_IDPHOTO_INPUT_IMAGE = "HIVISION_IDPHOTO_INPUT_IMAGE"
const HIVISION_IDPHOTO_INPUT_IMAGE_BASE64 = "HIVISION_IDPHOTO_INPUT_IMAGE_BASE64"

func getHivisionIdphotoBaseUrl() string {
	return os.Getenv(HIVISION_IDPHOTO_BASE_URL)
}

func getHivisionIdphotoInputImage() string {
	return os.Getenv(HIVISION_IDPHOTO_INPUT_IMAGE)
}

func getHivisionIdphotoInputImageBase64() string {
	return os.Getenv(HIVISION_IDPHOTO_INPUT_IMAGE_BASE64)
}

var inputImagePath string
var file *os.File
var client *Client
var ctx context.Context

func TestMain(m *testing.M) {
	var err error
	ctx = context.Background()
	inputImagePath = getHivisionIdphotoInputImage()
	file, err = os.Open(inputImagePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	client = NewClient(
		WithBaseUrl(getHivisionIdphotoBaseUrl()),
		WithLogEnabled(false),
	)
	os.Exit(m.Run())
}

func TestTempFile(t *testing.T) {
	// create a temporary file for testing
	tmpfile, err := os.CreateTemp(os.TempDir(), "test-image-*.png")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// t.Logf("remove temp file name=%s", tmpfile.Name())
		// name=/var/folders/kz/w7qjphf50w5121d_fb6r2fl40000gn/T/test-image-3681701983.png
		err := os.Remove(tmpfile.Name())
		if err != nil {
			t.Fatal(err)
		}
	}()
}
