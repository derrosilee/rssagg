package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the headers of a Http Request
// Example:
// Authorization: ApiKey { insert Apikey here }
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authentication Credentials Provided")
	}

	vals := strings.Split(val, "")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "Apikey" {
		return "", errors.New("malformed first part of the header")
	}
	return "", nil
}
