package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/m-mattia-m/kdrive-go-client"
	"io"
	"log"
	"os"
)

func main() {
	client := kDrive.NewClient(
		kDrive.DriveId(os.Getenv("KDRIVE_ID")),
		kDrive.Token(os.Getenv("KDRIVE_TOKEN")),
	)

	tempToken := client.Token.String()

	// TODO: type does not work -> , "type": kDrive.TypeDir
	queryParams := map[string]string{"per_page": "25", "order": kDrive.OrderAsc}

	// Get a list of directories and files in the root or a subdirectory
	directorsFilesList, err := client.File.GetDirectoryAndFiles(context.Background(), queryParams, kDrive.FileId("38837"))
	if err != nil {
		fmt.Println(err)
	}

	// Call download method of the client
	fileStream, err := client.File.Download(context.Background(), queryParams, kDrive.FileId("39387"))
	if err != nil {
		fmt.Println(err)
	}

	// Create file for writing the image -> remove the path before the variable name to store the file in the current directory
	file, err := os.Create(fmt.Sprintf("/Users/mattiamueggler/Desktop/%s", fileStream.Name))
	if err != nil {
		fmt.Printf("Error creating file: %s", err)
	}
	defer file.Close()

	// Load byte array into the created file.
	_, err = io.Copy(file, bytes.NewReader(fileStream.File))
	if err != nil {
		fmt.Println(err)
	}

	queryParams = map[string]string{}

	// Call GetPrivateLink-Function to generate a public-access-link
	privateLink, err := client.Link.GetPrivateLink(context.Background(), queryParams, kDrive.FileId("39387"))
	if err != nil {
		fmt.Println(err)
	}

	// Call GetSharedLink-Function to generate a private-access-link
	sharedLinkBody := kDrive.SharedLinkBody{
		CanComment:  true,
		CanDownload: true,
		CanEdit:     true,
		CanSeeInfo:  true,
		CanSeeStats: true,
		Right:       kDrive.RightPublic,
		ValidUntil:  0,
	}
	publicLink, err := client.Link.GetSharedLink(context.Background(), queryParams, kDrive.FileId("39409"), sharedLinkBody)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("")

	_ = tempToken
	_ = directorsFilesList
	_ = privateLink
	_ = publicLink
}

func init() {
	log.Println("[INIT MAIN]: Status -> Main-init is loaded")
	godotenv.Load()
}
