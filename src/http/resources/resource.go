package resources

type Resource[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}
