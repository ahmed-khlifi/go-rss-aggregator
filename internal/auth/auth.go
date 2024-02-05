package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
	Get the apiKey from the headers of a HTTP Request
	Authorization: ApiKey {KEY}
*/
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing Authorization header")
	}

	vals := strings.Split(val, " ")
	if(len(vals) != 2 || strings.ToLower(vals[0]) != "apikey") {
		return "", errors.New("invalid Authorization header")
	}

	return vals[1], nil
}