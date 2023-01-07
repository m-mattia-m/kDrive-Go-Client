package kDrive_test

import (
	"context"
	"encoding/json"
	"fmt"
	"kDrive"
	"net/http"
	"reflect"
	"testing"
)

func TestFileClient(t *testing.T) {

	t.Run("GetDirectoryAndFiles-root", func(t *testing.T) {
		var tests = []struct {
			name       string
			filePath   string
			statusCode int
			id         kDrive.DriveId
			fileId     kDrive.FileId // for the root the id is an empty string
			want       *kDrive.List
			wantErr    bool
			err        error
		}{
			{
				name:       "return the directories and files from root",
				id:         "some_id",
				filePath:   "testdata/fileDirectory_root.json",
				statusCode: http.StatusOK,
				want: &kDrive.List{
					Result: "success",
					Data: []kDrive.FileDirectoryList{
						{
							Id:             17,
							Name:           "Files",
							Type:           "dir",
							Status:         "",
							Visibility:     "",
							DriveId:        550642,
							Depth:          1,
							CreatedBy:      927543,
							CreatedAt:      0,
							AddedAt:        1663532308,
							LastModifiedAt: 1673084771,
							ParentId:       1,
							Color:          "",
						},
						{
							Id:             34621,
							Name:           "test.txt",
							Type:           "file",
							Status:         "",
							Visibility:     "",
							DriveId:        550642,
							Depth:          1,
							CreatedBy:      927543,
							CreatedAt:      1673044838,
							AddedAt:        1673044838,
							LastModifiedAt: 1673044844,
							ParentId:       1,
							Size:           6,
							HasThumbnail:   true,
							HasOnlyOffice:  false,
							MimeType:       "text/plain",
							ExtensionType:  "code",
						},
					},
					Page:         1,
					ItemsPerPage: 10,
					ResponseAt:   1673092847,
				},
				wantErr: false,
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				c := newMockedClient(t, test.filePath, test.statusCode)
				client := kDrive.NewClient("some_kDrive_Id", "some_token", kDrive.WithHTTPClient(c))
				got, err := client.File.GetDirectoryAndFiles(context.Background(), test.fileId)

				if (err != nil) != test.wantErr {
					t.Errorf("Get() error = %v, wantErr %v", err, test.wantErr)
					return
				}
				if !reflect.DeepEqual(got, test.want) {
					fmt.Println("---------------------")
					fmt.Println("Want")
					json2, err := json.Marshal(test.want)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(string(json2))
					fmt.Println("---------------------")
					fmt.Println("Got")
					json, err := json.Marshal(got)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(string(json))
					fmt.Println("---------------------")
					t.Errorf("Get() got = %v, want %v", got, test.want)
				}
			})
		}
	})

	t.Run("GetDirectoryAndFiles-subFolder", func(t *testing.T) {
		var tests = []struct {
			name       string
			filePath   string
			statusCode int
			id         kDrive.DriveId
			fileId     kDrive.FileId
			want       *kDrive.List
			wantErr    bool
			err        error
		}{
			{
				name:       "return the directories and files from a subfolder",
				id:         "some_id",
				filePath:   "testdata/fileDirectory_folder.json",
				statusCode: http.StatusOK,
				want: &kDrive.List{
					Result: "success",
					Data: []kDrive.FileDirectoryList{
						{
							Id:             23,
							Name:           "Private",
							Type:           "dir",
							Status:         "",
							Visibility:     "",
							DriveId:        550642,
							Depth:          2,
							CreatedBy:      927543,
							CreatedAt:      1663532975,
							AddedAt:        1663532977,
							LastModifiedAt: 1673084771,
							ParentId:       17,
							Color:          "",
						},
						{
							Id:             24,
							Name:           "Work",
							Type:           "dir",
							Status:         "",
							Visibility:     "",
							DriveId:        550642,
							Depth:          2,
							CreatedBy:      927543,
							CreatedAt:      1671783867,
							AddedAt:        1671814379,
							LastModifiedAt: 1673037804,
							ParentId:       17,
							Color:          "",
						},
						{
							Id:             25,
							Name:           "Documents",
							Type:           "dir",
							Status:         "",
							Visibility:     "",
							DriveId:        550642,
							Depth:          2,
							CreatedBy:      927543,
							CreatedAt:      1631468505,
							AddedAt:        1663533274,
							LastModifiedAt: 1672985209,
							ParentId:       17,
							Color:          "",
						},
					},
					Total:        0,
					Page:         1,
					Pages:        0,
					ItemsPerPage: 10,
					ResponseAt:   1673100753,
				},
				wantErr: false,
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				c := newMockedClient(t, test.filePath, test.statusCode)
				client := kDrive.NewClient("some_kDrive_Id", "some_token", kDrive.WithHTTPClient(c))
				got, err := client.File.GetDirectoryAndFiles(context.Background(), test.fileId)

				if (err != nil) != test.wantErr {
					t.Errorf("Get() error = %v, wantErr %v", err, test.wantErr)
					return
				}
				if !reflect.DeepEqual(got, test.want) {

					t.Errorf("Get() got = %v, want %v", got, test.want)
				}
			})
		}
	})

	t.Run("GetDirectoryAndFiles-error", func(t *testing.T) {
		var tests = []struct {
			name       string
			filePath   string
			statusCode int
			id         kDrive.DriveId
			fileId     kDrive.FileId
			want       *kDrive.List
			wantErr    bool
			err        error
		}{
			{
				name:       "return the directories and files from a subfolder",
				id:         "some_id",
				filePath:   "testdata/fileDirectory_folder.json",
				statusCode: http.StatusOK,
				want: &kDrive.List{
					Result: "success",
					Data: []kDrive.FileDirectoryList{
						{
							Id:             23,
							Name:           "Private",
							Type:           "dir",
							Status:         "",
							Visibility:     "",
							DriveId:        550642,
							Depth:          2,
							CreatedBy:      927543,
							CreatedAt:      1663532975,
							AddedAt:        1663532977,
							LastModifiedAt: 1673084771,
							ParentId:       17,
							Color:          "",
						},
						{
							Id:             24,
							Name:           "Work",
							Type:           "dir",
							Status:         "",
							Visibility:     "",
							DriveId:        550642,
							Depth:          2,
							CreatedBy:      927543,
							CreatedAt:      1671783867,
							AddedAt:        1671814379,
							LastModifiedAt: 1673037804,
							ParentId:       17,
							Color:          "",
						},
						{
							Id:             25,
							Name:           "Documents",
							Type:           "dir",
							Status:         "",
							Visibility:     "",
							DriveId:        550642,
							Depth:          2,
							CreatedBy:      927543,
							CreatedAt:      1631468505,
							AddedAt:        1663533274,
							LastModifiedAt: 1672985209,
							ParentId:       17,
							Color:          "",
						},
					},
					Total:        0,
					Page:         1,
					Pages:        0,
					ItemsPerPage: 10,
					ResponseAt:   1673100753,
				},
				wantErr: false,
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				c := newMockedClient(t, test.filePath, test.statusCode)
				client := kDrive.NewClient("some_kDrive_Id", "some_token", kDrive.WithHTTPClient(c))
				got, err := client.File.GetDirectoryAndFiles(context.Background(), test.fileId)

				if (err != nil) != test.wantErr {
					t.Errorf("Get() error = %v, wantErr %v", err, test.wantErr)
					return
				}
				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("Get() got = %v, want %v", got, test.want)
				}
			})
		}
	})

}
