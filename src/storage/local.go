package storage

import (
	"mime/multipart"
)

type localStorage struct{}

func (s localStorage) Store(path string, f *multipart.FileHeader) (filePath string, err error) {
	return "", nil
}

func (s localStorage) GetURL() (url string, err error) {
	return "", nil
}

func (s localStorage) Delete() error {
	return nil
}
