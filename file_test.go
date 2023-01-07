package kDrive_test

import (
	"context"
	"kDrive"
	"net/http"
	"reflect"
	"testing"
)

func TestDatabaseClient(t *testing.T) {
	//timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	//if err != nil {
	//	t.Fatal(err)
	//}

	t.Run("GetDirectoryAndFiles", func(t *testing.T) {
		var tests = []struct {
			name       string
			filePath   string
			statusCode int
			id         kDrive.DriveId
			fileId     kDrive.FileId
			want       *kDrive.DirectoryList
			wantErr    bool
			err        error
		}{
			{
				name:       "returns a list of files and directories",
				id:         "some_id",
				filePath:   "testdata/fileDirectory_root.json",
				statusCode: http.StatusOK,
				want: &kDrive.DirectoryList{
					Result: "success",
					Data: []kDrive.FileDirectory{
						{
							Directory: kDrive.Directory{
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
						},
						{
							Directory: kDrive.Directory{
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
							},
							Size:          6,
							HasThumbnail:  true,
							HasOnlyOffice: false,
							MimeType:      "text/plain",
							ExtensionType: "code",
						},
					},
					Total:        0,
					Page:         1,
					Pages:        0,
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
					t.Errorf("Get() got = %v, want %v", got, test.want)
				}
			})
		}
	})

}
