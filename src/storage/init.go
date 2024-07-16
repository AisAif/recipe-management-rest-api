package storage

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

type storageInterface interface {
	Store(path string, f *multipart.FileHeader) (filePath string, err error)
	GetURL() (url string, err error)
	Delete(path string) error
}

func InitStorage() {
	if viper.GetString("STORAGE_TYPE") == "S3" {
		creds := credentials.NewStaticCredentials(
			viper.GetString("AWS_ACCESS_KEY_ID"),
			viper.GetString("AWS_SECRET_ACCESS_KEY"),
			"",
		)

		sess := session.Must(session.NewSession(&aws.Config{
			Endpoint:         aws.String(viper.GetString("AWS_ENDPOINT")),
			Region:           aws.String(viper.GetString("AWS_DEFAULT_REGION")),
			S3ForcePathStyle: aws.Bool(true),
			Credentials:      creds,
		}))

		uploader := s3manager.NewUploader(sess)
		s3Client := s3.New(sess)

		Storage = &s3Storage{
			uploader: uploader,
			s3Client: s3Client,
		}
	} else {
		Storage = &localStorage{}
	}
}

var Storage storageInterface
