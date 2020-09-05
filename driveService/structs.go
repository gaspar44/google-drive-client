package driveService

type clientUser struct {
	Email      string `json:"client_email"`
	PrivateKey string `json:"private_key"`
}

type errorParsingFileName struct {
	Message string
}

func (errorParsing *errorParsingFileName) Error() string {
	return errorParsing.Message
}
