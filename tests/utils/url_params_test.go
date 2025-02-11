package utils_test

import (
	"fintual-cli/internals/utils"
	"testing"
)

func TestUrlParamsMaker(t *testing.T) {
	url, err := utils.UrlGenerator("http://localhost:8000", map[string]string{
		"queryParam1": "query1",
		"queryParam2": "query2",
	})

	if err != nil {
		t.Fatalf("Error making the url for %v", err)
	}

	matchedUrl := "http://localhost:8000?queryParam1=query1&queryParam2=query2"

	if matchedUrl != url {
		t.Fatalf("Left does not match with right value\n%s != %s ", matchedUrl, url)
	}
}
