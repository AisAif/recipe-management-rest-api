package storage

import (
	"mime/multipart"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type s3Storage struct {
	uploader *s3manager.Uploader
	s3Client *s3.S3
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

func (s s3Storage) GetURL(path string) (url string, err error) {
	if viper.GetString("AWS_URL") != "" {
		return viper.GetString("AWS_URL") + "/" + strings.ReplaceAll(path, " ", "%20"), nil
	} else {
		req, _ := s.s3Client.GetObjectAclRequest(&s3.GetObjectAclInput{
			Bucket: aws.String(viper.GetString("AWS_BUCKET")),
			Key:    aws.String(path),
		})

		return req.Presign(15 * time.Minute)
	}
}

func (s s3Storage) Delete(path string) error {
	_, err := s.s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(viper.GetString("AWS_BUCKET")),
		Key:    aws.String(path),
	})

	if err != nil {
		return err
	}

	return nil
}
