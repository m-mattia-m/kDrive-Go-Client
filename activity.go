package kDrive

import (
	"context"
)

type FileId string

func (fId FileId) String() string {
	return string(fId)
}

type ActivityService interface {
	Get(context.Context, FileId) (*File, error)
	Upload(context.Context, *FileRequest) (*File, error)
}

type ActivityClient struct {
	apiClient *Client
}

func (a ActivityClient) Get(ctx context.Context, id FileId) (*File, error) {
	//TODO implement me
	panic("implement me")
}

func (a ActivityClient) Upload(ctx context.Context, request *FileRequest) (*File, error) {
	//TODO implement me
	panic("implement me")
}

type File struct {
	Name string `json:"name"`
	file []byte `json:"file"`
}

type FileRequest struct {
	file []byte `json:"file"`
}
