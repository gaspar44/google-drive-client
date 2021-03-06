package driveService

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

const GOOGLE_DRIVE_CONFIG_HOME string = ".google_drive"
const FOLDER_ID string = "1YiH4GVqveyq_xdKBwdbfK7ylaXjXStnN"

var userHome string
var googleDriveConfigHome string

func init() {
	var exists bool
	userHome := os.Getenv("HOME")
	googleDriveConfigHome, exists = os.LookupEnv("GOOGLE_DRIVE_CONFIG_HOME")

	if !exists {
		googleDriveConfigHome = path.Join(userHome, GOOGLE_DRIVE_CONFIG_HOME)
		checkAndSetGoogleDriveHome(googleDriveConfigHome)
	}

}

type DriveService struct {
	ServiceInstance *drive.Service
}

func New() *DriveService {
	httpClient := getDriverClient()
	newService, err := drive.New(httpClient)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return &DriveService{
		ServiceInstance: newService,
	}
}

func getDriverClient() *http.Client {
	secrets, err := ioutil.ReadFile(path.Join(googleDriveConfigHome, "secrets.json"))

	if err != nil {
		log.Fatalln(err.Error())
	}

	config, err := google.JWTConfigFromJSON(secrets, drive.DriveScope, drive.DriveReadonlyScope)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return config.Client(context.Background())
}
