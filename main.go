package main

import (
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

	service := driveService.New()
	service.UploadFile(fileToUpload)
}
