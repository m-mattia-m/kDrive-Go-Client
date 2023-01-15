package kDrive

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type LinkService interface {
	GetSharedLink(ctx context.Context, queryParams map[string]string, fileId FileId, body SharedLinkBody) (*SharedLink, error)
	GetPublicLink(ctx context.Context, queryParams map[string]string, fileId FileId, body SharedLinkBody) (*SharedLink, error)
	GetPrivateLink(ctx context.Context, queryParams map[string]string, fileId FileId) (*PrivateLink, error)
}

type LinkClient struct {
	apiClient *Client
}

func (lc LinkClient) GetSharedLink(ctx context.Context, queryParams map[string]string, fileId FileId, body SharedLinkBody) (*SharedLink, error) {
	res, err := lc.apiClient.request(ctx, http.MethodPost, fmt.Sprintf("files/%s/link", fileId.String()), queryParams, body)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response SharedLink

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPublicLink this is an alias for GetSharedLink -> please use the official function
func (lc LinkClient) GetPublicLink(ctx context.Context, queryParams map[string]string, fileId FileId, body SharedLinkBody) (*SharedLink, error) {
	return lc.GetSharedLink(ctx, queryParams, fileId, body)
}

func (lc LinkClient) GetPrivateLink(ctx context.Context, queryParams map[string]string, fileId FileId) (*PrivateLink, error) {
	if fileId.String() == "" {
		return nil, errors.New("has no fileId")
	}
	url := fmt.Sprintf("%s/app/drive/%s/redirect/%s", lc.apiClient.appUrl, lc.apiClient.DriveId, fileId.String())
	return &PrivateLink{
		FileId:  FileId(fileId),
		DriveId: DriveId(lc.apiClient.DriveId),
		Url:     url,
	}, nil
}
