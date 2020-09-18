package driveService

import (
	"google.golang.org/api/drive/v3"
)

type DriveService struct {
	serviceInstance *drive.Service
}
