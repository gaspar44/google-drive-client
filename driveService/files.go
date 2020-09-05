package driveService

import (
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"strings"

	/*"googleDriveClient"*/
	"log"
	"net/http"
	"os"
)

func UploadFile(fileToUpload *os.File) error {
	contentType, err := getMimeTypeFile(fileToUpload)

	if err != nil {
		log.Fatalln(err.Error())
	} else if contentType == "" {
		log.Fatalln("unable to determinate mime type")
	}

	fileWithoutAbsolutePath := parseName(fileToUpload.Name())

	if fileWithoutAbsolutePath == "" {
		return &errorParsingFileName{Message: "Can not parse name"}
	}

	file := &drive.File{
		MimeType: contentType,
		Name:     fileWithoutAbsolutePath,
		Parents:  []string{FOLDER_ID},
	}

	_, err = service.Files.Create(file).Media(fileToUpload).Do()

	folderContent, err := getFolderContents()
	fmt.Println(file.Name)
	fmt.Println(folderContent[0].Name)

	if err != nil {
		return err
	}

	return nil
}

func getFolderContents() ([]*drive.File, error) {
	fields := []googleapi.Field{"nextPageToken,files(id,fileExtension, name)"}
	query := "'" + FOLDER_ID + "' in parents"

	driverFileList, err := service.Files.List().Spaces("drive").Q(query).Fields(fields...).Do()
	service.Files.Get(FOLDER_ID).Download()

	if err != nil {
		return nil, err
	}

	return driverFileList.Files, nil
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

func checkIfFileExistsOrIsNew(fileToCheck *os.File) bool {
	return true
}

func parseName(absolutePathOfFile string) string {
	lastIndex := strings.LastIndex(absolutePathOfFile, "/")

	if lastIndex == -1 {
		return ""
	}
	/*	beforeFileName := lastIndex + len(absolutePathOfFile)

		if beforeFileName >= len(absolutePathOfFile) {
			return ""
		}*/

	return absolutePathOfFile[lastIndex+1 : len(absolutePathOfFile)]
}
