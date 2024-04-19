package utils

import "encoding/base64"

// BasicAuthGenerator generates a basic authentication string using the provided username and password.
// It encodes the username and password using base64 encoding and returns the encoded string.
func BasicAuthGenerator(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
