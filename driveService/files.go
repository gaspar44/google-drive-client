package driveService

import (
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"strings"

	"os"
)

func UploadService(service *drive.Service, fileToUpload *os.File) (*drive.File, error) {
	fileWithoutAbsolutePath := parseName(fileToUpload.Name())

	file := &drive.File{
		//MimeType: "text/plain",
		Name:     fileWithoutAbsolutePath,
		Parents:  []string{FOLDER_ID},
	}

	uploadedFile, err := service.Files.Create(file).Media(fileToUpload).Do()

	if err != nil {
		return nil, err
	}

	return uploadedFile, err
}

func getFolderContents(srv *drive.Service) ([]*drive.File, error) {
	fields := []googleapi.Field{"nextPageToken,files(id,fileExtension, name)"}
	query := "'" + FOLDER_ID + "' in parents"

	driverFileList, err := srv.Files.List().Spaces("drive").Q(query).Fields(fields...).Do()
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

	//return mimetype.Detect(buffer).String(), nil
	return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", nil
	//return http.DetectContentType(buffer), nil
	//return http.DetectContentType(buffer), nil

}

/*func (srv *DriveService) checkIfFileExistsOrIsNew(fileToCheck *os.File) bool {
	return true
}
*/
func parseName(absolutePathOfFile string) string {
	lastIndex := strings.LastIndex(absolutePathOfFile, "/")

	if lastIndex == -1 {
		return absolutePathOfFile
	}

	return absolutePathOfFile[lastIndex+1 : len(absolutePathOfFile)]
}
