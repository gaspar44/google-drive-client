package driveService

import (
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

func (srv *DriveService) Upload(fileToUpload *PreparatedFileToUpload) (*drive.File, error) {
	file := &drive.File{
		MimeType: fileToUpload.MimeType,
		Name:     fileToUpload.Name,
		Parents:  []string{FOLDER_ID},
	}

	uploadedFile, err := srv.ServiceInstance.Files.Create(file).Media(fileToUpload.file).Do()

	if err != nil {
		return nil, err
	}

	return uploadedFile, err
}

func (srv *DriveService) getFolderContents() ([]*drive.File, error) {
	fields := []googleapi.Field{"nextPageToken,files(id,fileExtension, name)"}
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
