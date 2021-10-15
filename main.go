package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var url string = "s3.amazonaws.com/"
var proto string = "https://"
var win string = "\033[32;1m[✓] "
var lose string = "\033[31;1m[!] "
var info string = "\033[36;7m[*] "
var end string = "\033[0m"

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

func s3query(target string) {
	fmt.Println(info + "Starting s3.amazonaws.com query!" + end)

	common := []string{"test", "dev", "bucket", "s3", "aws", "prd", "prod", "pub", "public", "production", "development", "testing", "archive", "backup", "web", "devops", "sec", "secure", "hidden", "secret", "staging", "download"}
	connectors := []string{"-", "_", ""}

	for _, com := range common {
		for _, conns := range connectors {
			ur := proto + url + target + conns + com
			_, stat, err := getCode(ur)
			if err != nil {

				// fmt.Println(lose + "Failed to get! " + err.Error() + end)
			}
			if stat == 404 || stat == 401 || stat == 500 || stat == 400 || stat == 0 {
				fmt.Printf(lose+"%s\t|\t%d"+end, ur, stat)
			} else if stat == 403 {
				fmt.Printf(info+"%s\t|\t%d"+end, ur, stat)
			} else {
				fmt.Printf(win+"%s\t|\t%d"+end, ur, stat)
			}
			fmt.Println("\033[0m")
		}
	}

	for _, com := range common {
		for _, conns := range connectors {
			ur := proto + target + conns + com + "." + url
			_, stat, err := getCode(ur)
			if err != nil {

				// fmt.Println(lose + "Failed to get! " + err.Error() + end)
			}
			if stat == 404 || stat == 401 || stat == 500 || stat == 400 || stat == 0 {
				fmt.Printf(lose+"%s\t|\t%d"+end, ur, stat)
			} else if stat == 403 {
				fmt.Printf(info+"%s\t|\t%d"+end, ur, stat)
			} else {
				fmt.Printf(win+"%s\t|\t%d"+end, ur, stat)
			}
			fmt.Println("\033[0m")
		}
	}
}

func main() {
	fmt.Println(info + "Welcome to s3num v1.0!" + end)

	args := os.Args[1:]
	target := args[0]

	fmt.Println(info + "Target: " + target + end)

	s3query(target)
}
