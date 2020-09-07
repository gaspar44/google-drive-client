package driveService

import (
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"strings"

	"log"
	"net/http"
	"os"
)

func (srv *DriveService) UploadFile(fileToUpload *os.File) error {
	contentType, err := getMimeTypeFile(fileToUpload)

	if err != nil {
		log.Fatalln(err.Error())
	} else if contentType == "" {
		log.Fatalln("unable to determinate mime type")
	}

	fileWithoutAbsolutePath := parseName(fileToUpload.Name())

	file := &drive.File{
		MimeType: contentType,
		Name:     fileWithoutAbsolutePath,
		Parents:  []string{FOLDER_ID},
	}

	//_, err = srv.serviceInstance.Files.Create(file).Media(fileToUpload).Do()

	folderContent, err := getFolderContents(srv)
	fmt.Println(file.Name)
	fmt.Println(folderContent[0].Name)

	if err != nil {
		return err
	}

	return nil
}

func getFolderContents(srv *DriveService) ([]*drive.File, error) {
	fields := []googleapi.Field{"nextPageToken,files(id,fileExtension, name)"}
	query := "'" + FOLDER_ID + "' in parents"

	driverFileList, err := srv.serviceInstance.Files.List().Spaces("drive").Q(query).Fields(fields...).Do()
	//response, err := srv.serviceInstance.Files.Get(FOLDER_ID).Download()

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

	return http.DetectContentType(buffer), nil

}

func (srv *DriveService) checkIfFileExistsOrIsNew(fileToCheck *os.File) bool {
	return true
}

func parseName(absolutePathOfFile string) string {
	lastIndex := strings.LastIndex(absolutePathOfFile, "/")

	if lastIndex == -1 {
		return absolutePathOfFile
	}

	return absolutePathOfFile[lastIndex+1 : len(absolutePathOfFile)]
}
