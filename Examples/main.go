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

	// Get a list of directories and files in the root or a subdirectory
	directorsFilesList, err := client.File.GetDirectoryAndFiles(context.Background(), kDrive.FileId(""))
	if err != nil {
		fmt.Println(err)
	}

	// Call download method of the client
	fileStream, err := client.File.Download(context.Background(), kDrive.FileId("38840"))
	if err != nil {
		fmt.Println(err)
	}

	// Create file for writing the image -> remove the path before the variable name to store the file in the current directory
	file, err := os.Create(fmt.Sprintf("/Users/username/Desktop/%s", fileStream.Name))
	if err != nil {
		fmt.Printf("Error creating file: %s", err)
		return
	}
	defer file.Close()

	// Load byte array into the created file.
	_, err = io.Copy(file, bytes.NewReader(fileStream.File))
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = tempToken
	_ = directorsFilesList
}

func init() {
	log.Println("[INIT MAIN]: Status -> Main-init is loaded")
	godotenv.Load()
}
