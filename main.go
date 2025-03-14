package main

import (
	// "errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	urlsStr string
)

func init() {
	flag.StringVar(&urlsStr, "s", "", "list of urls separeted by commas")
	flag.Usage = func() {
		fmt.Printf("Ussage:%v -s <comma sepperated urls>\n", os.Args[0])
	}
	flag.Parse()
}

func main() {
	var wg sync.WaitGroup
	urls := strings.Split(urlsStr, ",")
	results := make(chan string, len(urls))
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
			results <- fmt.Sprintf("%v - %v", url, stat)
		}(url)
	}
	wg.Wait()
	close(results)
	for res := range results {
		fmt.Println(res)
	}
}

func statusCheck(site string, client *http.Client) (string, error) {
	resp, err := client.Get(site)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Status, nil
}
