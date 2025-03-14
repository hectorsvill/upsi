package main

import (
	// "errors"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("ping")
	urls := []string{"http://www.google.com", "https://x.com/", "https://time.is/"}
	stat, err := statusCheck(urls[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stat)

}

func statusCheck(site string) (string, error) {
	resp, err := http.Get(site)
	if err != nil {
		return "", err
	}
	return resp.Status, nil
}
