package driveService

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
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
var service *drive.Service

func init() {
	var exists bool
	userHome := os.Getenv("HOME")
	googleDriveConfigHome, exists = os.LookupEnv("GOOGLE_DRIVE_CONFIG_HOME")

	if !exists {
		googleDriveConfigHome = path.Join(userHome, GOOGLE_DRIVE_CONFIG_HOME)
		checkAndSetGoogleDriveHome(googleDriveConfigHome)
	}

}

func New() *DriveService {
	httpClient := getDriverClient()
	newService, err := drive.New(httpClient)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return &DriveService{
		serviceInstance: newService,
	}
}

func getDriverClient() *http.Client {
	secrets, err := ioutil.ReadFile(path.Join(googleDriveConfigHome, "secrets.json"))

	if err != nil {
		log.Fatalln(err.Error())
	}

	var user clientUser
	json.Unmarshal(secrets, &user)

	config := &jwt.Config{
		Email:      user.Email,
		PrivateKey: []byte(user.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}

	return config.Client(context.Background())
}
