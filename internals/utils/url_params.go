package utils

import (
	"fmt"
	"net/url"
)

func UrlGenerator(baseURL string, parameters map[string]string) (string, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", err
	}

	queryParams := url.Query()

	for key, param := range parameters {
		queryParams.Set(key, param)
	}

	url.RawQuery = queryParams.Encode()

	return url.String(), nil
}
