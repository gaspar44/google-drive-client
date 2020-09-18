package main

import (
	"google.golang.org/api/drive/v3"
	"googleDriveClient/driveService"
	"log"
	"os"
	"path"
)

func main() {
	fileToUpload, err := os.Open(path.Join("/home/gaspar", "Dropbox", "Documentos", "hola.txt"))

	if err != nil {
		log.Fatalln()
	}

	defer fileToUpload.Close()

	service := driveService.New()
	file := &drive.File{
		MimeType: "text/plain",
		Name:     fileToUpload.Name(),
		Parents:  []string{driveService.FOLDER_ID},

	}
	service.Files.Create(file).Media(fileToUpload).Do()
	//service.UploadFile(fileToUpload)
}
