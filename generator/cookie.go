package generator

import (
	"fmt"
	"encoding/base64"
)

func CookieValue(userID string) (string, error) {
	randomString, err := StringBase64(32)

	if err != nil {
		return "", err
	}

	cookieValue := fmt.Sprintf("%s:%s", randomString, userID)
	return base64.StdEncoding.EncodeToString([]byte(cookieValue)), nil
}
