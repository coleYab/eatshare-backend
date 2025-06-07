package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func ComparePassword(password, hashed string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return err
	}

	return nil
}

func Base64Encode(data string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) (string, error) {
	val, err := base64.RawStdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	return string(val), nil
}
