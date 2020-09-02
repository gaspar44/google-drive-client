package main

import (
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"log"
	"net/http"
	"os"
)

func uploadFile(service *drive.Service, fileToUpload *os.File) {
	contentType, err := getMimeTypeFile(fileToUpload)

	if err != nil {
		log.Fatalln(err.Error())
	} else if contentType == "" {
		log.Fatalln("unable to determinate mime type")
	}

	file := &drive.File{
		MimeType: contentType,
		Name:     fileToUpload.Name(),
		Parents:  []string{FOLDER_ID},
	}

	//_, err = service.Files.Create(file).Media(fileToUpload).Do()
	fields := []googleapi.Field{"nextPageToken,files(id,fileExtension, name)"}

	list,err := service.Files.List().Spaces("drive").Q("'" + FOLDER_ID + "' in parents").Fields(fields...).Do()
	prueba := list.Files

	fmt.Println(prueba[0].Name)
	fmt.Println(file.Name)

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getMimeTypeFile(fileToDetect *os.File) (string, error) {

	buffer := make([]byte, 512)

	_, err := fileToDetect.Read(buffer)

	if err != nil {
		return "", err
	}

	contenType := http.DetectContentType(buffer)

	return contenType, nil

}
