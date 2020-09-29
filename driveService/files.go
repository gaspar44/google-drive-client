package driveService

import (
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"log"
)

func (srv *DriveService) Upload(fileToUpload *PreparedFileToUpload) (*drive.File, error) {
	file := &drive.File{
		MimeType: fileToUpload.MimeType,
		Name:     fileToUpload.Name,
		Parents:  []string{FOLDER_ID},
		/*Md5Checksum: fileToUpload.Md5CheckSum,*/
	}

	uploadedFile, err := srv.ServiceInstance.Files.Create(file).Media(fileToUpload.file).Do()

	if err != nil {
		return nil, err
	}

	return uploadedFile, err
}

func (srv *DriveService) getFolderContents() ([]*drive.File, error) {
	fields := []googleapi.Field{"nextPageToken,files(id,fileExtension, name,md5Checksum)"}
	query := "'" + FOLDER_ID + "' in parents"

	driverFileList, err := srv.ServiceInstance.Files.List().Spaces("drive").Q(query).Fields(fields...).Do()
	//response, err := srv.serviceInstance.Files.Get(FOLDER_ID).Download()

	if err != nil {
		return nil, err
	}

	return driverFileList.Files, nil
}

/*func (srv *DriveService) checkIfFileExistsOrIsNew(fileToCheck *os.File) bool {
	return true
}
*/

func (srv *DriveService) WatchRemoteFolderContents() {
	token, err := srv.getFolderContents()

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(token[0].Md5Checksum)

}
