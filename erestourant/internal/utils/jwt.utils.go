package utils

import (
	"encoding/base64"
	"fmt"
)

// CreateBasicToken generates a basic token by encoding "email:userID" in base64
func CreateBasicToken(email string, userID string) (string, error) {
	plain := fmt.Sprintf("%s:%s", email, userID)
	encoded := base64.StdEncoding.EncodeToString([]byte(plain))
	return encoded, nil
}
