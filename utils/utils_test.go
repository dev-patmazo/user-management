package utils

import (
	"encoding/base64"
	"testing"
)

func TestBasicAuthGenerator(t *testing.T) {
	username := "testuser"
	password := "testpassword"
	expectedAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	auth := BasicAuthGenerator(username, password)

	if auth != expectedAuth {
		t.Errorf("Expected %s, but got %s", expectedAuth, auth)
	}
}
