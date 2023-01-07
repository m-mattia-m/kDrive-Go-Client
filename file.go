package kDrive

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FileService interface {
	GetDirectoryAndFiles(ctx context.Context, fileId FileId) (*DirectoryList, error)
	//Upload(context.Context, *FileRequest) (*File, error)
}

type FileClient struct {
	apiClient *Client
}

func (fc FileClient) GetDirectoryAndFiles(ctx context.Context, fileId FileId) (*DirectoryList, error) {
	urlPath := "files"
	if fileId != "" {
		urlPath = fmt.Sprintf("%s/files", fileId.String())
	}
	res, err := fc.apiClient.request(ctx, http.MethodGet, urlPath, nil, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response DirectoryList

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

//func (fc FileClient) Upload(ctx context.Context, request *FileRequest) (*File, error) {
//	//TODO implement me
//	panic("implement me")
//}
