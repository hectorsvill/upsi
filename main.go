package main

import (
	// "errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	urls := []string{
		// "http://www.google.com",
		"https://time.is/",
		"http://www.youtube.com",
		"http://www.google.co.kr",
		// "http://www.alipay.com",
	}

	for _, url := range urls {
		wg.Add(1)
		client := &http.Client{
			Timeout: 3 * time.Second,
		}
		go func(url string) {
			defer wg.Done()
			stat, err := statusCheck(url, client)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%v - %v\n", url, stat)
		}(url)
	}
	wg.Wait()
}

func statusCheck(site string, client *http.Client) (string, error) {
	resp, err := client.Get(site)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Status, nil
}
