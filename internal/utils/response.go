package utils

type ErrorResponse struct {
	Code    int
	Message string
}

type SuccessResponse struct {
	Code    int
	Message string
	Data    interface{}
}
