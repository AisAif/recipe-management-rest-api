package storage

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type s3Storage struct {
	uploader *s3manager.Uploader
}

func (s s3Storage) Store(path string, f *multipart.FileHeader) (filePath string, err error) {

	file, err := f.Open()
	if err != nil {
		return "", err
	}

	u := uuid.New()

	filePath = path + "/" + u.String() + "." + f.Filename

	_, err = s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(viper.GetString("AWS_BUCKET")),
		Key:    aws.String(filePath),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func (s s3Storage) GetURL() (url string, err error) {
	return "", nil
}

func (s s3Storage) Delete() error {
	return nil
}
