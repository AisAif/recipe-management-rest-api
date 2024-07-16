package resources

import "github.com/rosberry/go-pagination"

type Resource[T any] struct {
	Data     T                   `json:"data"`
	Message  string              `json:"message"`
	Errors   any                 `json:"errors"`
	PageInfo pagination.PageInfo `json:"page_info"`
}
