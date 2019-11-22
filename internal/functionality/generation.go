package functionality

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/ikcilrep/twister/internal/parser"
)

var notPositiveLengthError error = errors.New("Length of data which exists has to be positive.")

func randomBytes(length int) ([]byte, error) {
	if length < 0 {
		return nil, notPositiveLengthError
	}
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func GenerateKey(arguments *parser.Arguments) error {
	keyBytes, err := randomBytes(arguments.KeySize)

	if err != nil {
		return err
	}

	if arguments.KeyOutput.IsBinary {
		arguments.KeyOutput.Writer.Write(keyBytes)
	} else {
		arguments.KeyOutput.Writer.Write([]byte(fmt.Sprintf("%v\n", base64.StdEncoding.EncodeToString(keyBytes))))
	}
	return nil
}