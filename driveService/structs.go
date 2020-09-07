package driveService

import "google.golang.org/api/drive/v3"

type clientUser struct {
	Email      string `json:"client_email"`
	PrivateKey string `json:"private_key"`
}

type DriveService struct {
	serviceInstance *drive.Service
}
