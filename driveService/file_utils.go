package driveService

import (
	"bytes"
	"github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"os"
	"strings"
)

type PreparatedFileToUpload struct {
	file     *bytes.Reader
	MimeType string
	Name     string
}

func PrepareFile(fileToUpload *os.File) (*PreparatedFileToUpload, error) {
	fileWithoutAbsolutePath := parseName(fileToUpload.Name())
	data, err := ioutil.ReadAll(fileToUpload)

	if err != nil {
		return nil, err
	}

	mimeType := mimetype.Detect(data)
	fileToUse := bytes.NewReader(data)

	return &PreparatedFileToUpload{
		file:     fileToUse,
		MimeType: mimeType.String(),
		Name:     fileWithoutAbsolutePath,
	}, nil
}

func parseName(absolutePathOfFile string) string {
	lastIndex := strings.LastIndex(absolutePathOfFile, "/")

	if lastIndex == -1 {
		return absolutePathOfFile
	}

	return absolutePathOfFile[lastIndex+1 : len(absolutePathOfFile)]
}
