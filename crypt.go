package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

func signPath(path string) (string, error) {
	mac := hmac.New(sha256.New, conf.Key)
	mac.Write(conf.Salt)
	mac.Write([]byte(path))
	token := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return token, nil
}

func validatePath(token, path string) error {
	messageMAC, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return errors.New("Invalid token encoding")
	}

	mac := hmac.New(sha256.New, conf.Key)
	mac.Write(conf.Salt)
	mac.Write([]byte(path))
	expectedMAC := mac.Sum(nil)

	if !hmac.Equal(messageMAC, expectedMAC) {
		return errors.New("Invalid token")
	}

	return nil
}
