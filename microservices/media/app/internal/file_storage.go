package internal

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/alexandr-io/backend/media/data"

	"github.com/gofiber/fiber/v2"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob" // Used to open a Bucket for file storage
)

var mediaPath = os.Getenv("MEDIA_PATH")
var mediaURI = os.Getenv("MEDIA_URI")

// UploadFile upload a file on the storage server
func UploadFile(ctx context.Context, file []byte, path string) error {

	// Open a connection to the bucket.
	bucket, err := blob.OpenBucket(ctx, mediaURI+"://"+mediaPath)
	if err != nil {
		return err
	}
	defer bucket.Close()

	w, err := bucket.NewWriter(ctx, path, nil)
	if err != nil {
		return err
	}
	_, err = w.Write(file)
	if err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

// DownloadFile download a file from the storage  server
func DownloadFile(ctx context.Context, path string) (*data.File, error) {

	var fileObject data.File

	// Open a connection to the bucket.
	bucket, err := blob.OpenBucket(ctx, mediaURI+"://"+mediaPath)
	if err != nil {
		return nil, err
	}
	defer bucket.Close()

	file, err := bucket.NewReader(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileObject.Data, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fileObject.ContentType = file.ContentType()

	return &fileObject, err
}

// DeleteFile delete a file from the storage server
func DeleteFile(ctx context.Context, path string) error {

	// Open a connection to the bucket.
	bucket, err := blob.OpenBucket(ctx, mediaURI+"://"+mediaPath)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	defer bucket.Close()

	err = bucket.Delete(ctx, path)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
