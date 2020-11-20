package internal

import (
	"context"
	"github.com/alexandr-io/backend/media/data"
	"github.com/gofiber/fiber/v2"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	"io/ioutil"
	"os"
)

var mediaPath = os.Getenv("MEDIA_PATH")
var mediaURI = os.Getenv("MEDIA_URI")

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

func DownloadFile(ctx context.Context, path string) (*data.File, error) {

	fileObject := new(data.File)

	// Open a connection to the bucket.
	bucket, err := blob.OpenBucket(ctx, mediaURI+"://"+mediaPath)
	if err != nil {
		return fileObject, err
	}
	defer bucket.Close()

	file, err := bucket.NewReader(ctx, path, nil)
	if err != nil {
		return fileObject, err
	}
	defer file.Close()

	fileObject.Data, err = ioutil.ReadAll(file)
	if err != nil {
		return fileObject, err
	}
	fileObject.ContentType = file.ContentType()

	return fileObject, err
}

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
