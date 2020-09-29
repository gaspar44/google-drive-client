package driveService

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"os"
	"strings"
)

type PreparedFileToUpload struct {
	file        *bytes.Reader
	MimeType    string
	Name        string
	Md5CheckSum string
}

func PrepareFile(fileToUpload *os.File) (*PreparedFileToUpload, error) {
	fileWithoutAbsolutePath := parseName(fileToUpload.Name())
	data, err := ioutil.ReadAll(fileToUpload)
	checkSum := calculateCheckSum(data)

	if err != nil {
		return nil, err
	}

	mimeType := mimetype.Detect(data)
	fileToUse := bytes.NewReader(data)

	return &PreparedFileToUpload{
		file:        fileToUse,
		MimeType:    mimeType.String(),
		Name:        fileWithoutAbsolutePath,
		Md5CheckSum: checkSum,
	}, nil
}

func calculateCheckSum(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func parseName(absolutePathOfFile string) string {
	lastIndex := strings.LastIndex(absolutePathOfFile, "/")

	if lastIndex == -1 {
		return absolutePathOfFile
	}

	return absolutePathOfFile[lastIndex+1 : len(absolutePathOfFile)]
}