package kDrive

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strings"
)

type FileService interface {
	GetDirectoryAndFiles(ctx context.Context, queryParams map[string]string, fileId FileId) (*List, error)
	// Upload(context.Context, *FileRequest) (*object.File, error)
	Download(ctx context.Context, queryParams map[string]string, fileId FileId) (*FileStream, error)
	Thumbnail(ctx context.Context, queryParams map[string]string, fileId FileId) (*List, error)
}

type FileClient struct {
	apiClient *Client
}

func (fc FileClient) GetDirectoryAndFiles(ctx context.Context, queryParams map[string]string, fileId FileId) (*List, error) {
	urlPath := "files"
	if fileId != "" {
		urlPath += fmt.Sprintf("/%s/files", fileId.String())
	}
	res, err := fc.apiClient.request(ctx, http.MethodGet, urlPath, queryParams, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response List

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

//func (fc FileClient) Upload(ctx context.Context, request *FileRequest) (*DirectoryList, error) {
//	urlPath := "files"
//	if fileId != "" {
//		urlPath = fmt.Sprintf("%s/files", fileId.String())
//	}
//	res, err := fc.apiClient.request(ctx, http.MethodGet, urlPath, nil, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	defer func() {
//		if errClose := res.Body.Close(); errClose != nil {
//			log.Println("failed to close body, should never happen")
//		}
//	}()
//
//	var response DirectoryList
//
//	err = json.NewDecoder(res.Body).Decode(&response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &response, nil
//}

func (fc FileClient) Download(ctx context.Context, queryParams map[string]string, fileId FileId) (*FileStream, error) {
	urlPath := fmt.Sprintf("files/%s/download", fileId.String())
	res, err := fc.apiClient.request(ctx, http.MethodGet, urlPath, queryParams, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	fileName, err := getFileName(*res)
	fileType, err := getFileType(*res)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}
	file := buf.Bytes()

	fileStream := FileStream{
		Name: fileName,
		Type: fileType,
		File: file,
	}
	return &fileStream, nil
}

func (fc FileClient) Thumbnail(ctx context.Context, queryParams map[string]string, fileId FileId) (*List, error) {
	urlPath := "files"
	if fileId != "" {
		urlPath = fmt.Sprintf("%s/files", fileId.String())
	}
	res, err := fc.apiClient.request(ctx, http.MethodGet, urlPath, queryParams, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response List

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func getFileName(file http.Response) (string, error) {
	contentDisposition := file.Header.Get("Content-Disposition")
	if contentDisposition == "" {
		return "", errors.New("Content-Disposition header not found")
	}

	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error parsing Content-Disposition header: %s", err))
	}

	filename, ok := params["filename"]
	if !ok {
		return "", errors.New("filename not found in content-disposition header")
	}

	// Unerlaubte Zeichen aus dem Dateinamen entfernen
	filename = strings.Map(func(r rune) rune {
		if r == '/' || r == '\\' {
			return -1
		}
		return r
	}, filename) // Bild von der HTTP-Antwort in die Datei streamen
	return filename, nil
}

func getFileType(file http.Response) (string, error) {
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		return "", errors.New("content-type header not found")
	}
	return contentType, nil
}
