package main

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


var userHome string
var googleDriveConfigHome string

const GOOGLE_DRIVE_CONFIG_HOME string = ".google_drive"
const FOLDER_ID string = "1YiH4GVqveyq_xdKBwdbfK7ylaXjXStnN"

func getDriverClient() *http.Client{
	secrets, err := ioutil.ReadFile(path.Join(googleDriveConfigHome, "secrets.json"))

	if err != nil {
		log.Fatalln(err.Error())
	}

	var user clientUser
	json.Unmarshal(secrets,&user)

	config := &jwt.Config{
		Email:         user.Email,
		PrivateKey:    []byte(user.PrivateKey),
		Scopes:        []string {
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}

	 httpClient := config.Client(context.Background())

	return httpClient

}
func checkAndSetGoogleDriveHome(pathToHome string) {
	_, err := os.Stat(pathToHome)
	if os.IsNotExist(err) {
		log.Println("not found. Creating in: " + pathToHome)
		err = os.Mkdir(pathToHome, 755)

		if err != nil {
			log.Fatalln(err.Error())
		}

		// SECURITY OF CLIENT FOR FUTURE
		return
	}

	log.Println("Already exists")
}

func init() {
	var exists bool
	userHome := os.Getenv("HOME")
	googleDriveConfigHome, exists = os.LookupEnv("GOOGLE_DRIVE_CONFIG_HOME")

	if !exists {
		googleDriveConfigHome = path.Join(userHome, GOOGLE_DRIVE_CONFIG_HOME)
		checkAndSetGoogleDriveHome(googleDriveConfigHome)
	}
}

func uploadFile(service *drive.Service, fileToUpload  *os.File) {
	file := &drive.File{
		MimeType: "application/vnd.ms-excel",
		Name: fileToUpload.Name(),
		Parents: []string{FOLDER_ID},
	}

	_,err := service.Files.Create(file).Media(fileToUpload).Do()

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	driverClient := getDriverClient()
	fileToUpload, err := os.Open(path.Join("/home/gaspar","Dropbox","Documentos","Gastos.xlsx"))

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer fileToUpload.Close()

	service, err := drive.New(driverClient)

	if err != nil {
		log.Fatalln(err.Error())
	}

	uploadFile(service,fileToUpload)



}

func createFile(service *drive.Service) {

}
