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

	defer fileToUpload.Close()

	service := driveService.New()
	service.UploadService(fileToUpload)
	//driveService.UploadFile(fileToUpload)
	//driveService.UploadService2(service,fileToUpload)
}
