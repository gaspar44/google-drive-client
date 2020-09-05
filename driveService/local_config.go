package driveService

import (
	"log"
	"os"
)

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
