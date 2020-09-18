package main

import (
	"googleDriveClient/driveService"
	"log"
	"os"
	"path"
)

func main() {
	fileToUpload, err := os.Open(path.Join("/home/gaspar", "Dropbox", "Documentos", "Gastos.xlsx"))

	if err != nil {
		log.Fatalln()
	}

	defer fileToUpload.Close()

	service := driveService.New()
	preparedFile, err := driveService.PrepareFile(fileToUpload)

	if err != nil {
		log.Fatalln(err.Error())
	}
	service.Upload(preparedFile)
	//driveService.UploadFile(fileToUpload)
	//driveService.UploadService2(service,fileToUpload)
}
