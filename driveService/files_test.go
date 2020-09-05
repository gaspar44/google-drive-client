package driveService

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const EXPECTED_NAME string = "holis.txt"

func TestParseNameFullPath(t *testing.T) {
	var assert = assert.New(t)
	var inputName = "/hola/como/estas/holis.txt"

	result := parseName(inputName)
	assert.NotNil(result)
	assert.Equal(EXPECTED_NAME, result)
}

func TestParseNameFileWithoutPath(t *testing.T) {
	var assert = assert.New(t)
	var inputName = EXPECTED_NAME

	result := parseName(inputName)
	assert.NotNil(result)
	assert.Equal(EXPECTED_NAME, result)
}

func TestParseNameFileWithDotPath(t *testing.T) {
	var assert = assert.New(t)
	var inputName = "./" + EXPECTED_NAME

	result := parseName(inputName)
	assert.NotNil(result)
	assert.Equal(EXPECTED_NAME, result)
}
