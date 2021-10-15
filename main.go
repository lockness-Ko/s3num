package main

import (
	"io"
	"net/http"
	"os"
)

var url string = "https://s3.amazonaws.com/"

func getCode(url string) (string, int, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", 0, err
	}
	return (string)(body), response.StatusCode, nil
}

func main() {
	// common := []string{"test", "dev", "bucket", "s3", "aws", "prd", "prod", "pub", "public", "production", "development", "testing", "archive", "backup", "web", "devops", "sec", "secure", "hidden", "secret", "staging", "download"}
	// connectors := []string{"-", "_", ""}

	args := os.Args[1:]
	target := args[0]

	url = url + target

	, stat, err := getCode(url)
}
