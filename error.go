package kDrive

import "fmt"

type ErrorCode string

//type Error struct {
//	Object  ObjectType `json:"object"`
//	Status  int        `json:"status"`
//	Code    ErrorCode  `json:"code"`
//	Message string     `json:"message"`
//}

type Error struct {
	Result      string         `json:"result"`
	ErrorResult ApiErrorResult `json:"error"`
}

type ApiErrorResult struct {
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Context     ApiErrorContext `json:"context"`
}

type ApiErrorContext struct {
	Model string `json:"model"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.ErrorResult)
}

type RateLimitedError struct {
	Message string
}

func (e *RateLimitedError) Error() string {
	return e.Message
}
