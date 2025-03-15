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

type Result struct {
	url    string
	status string
	print  string
}

var (
	urlsStr string
)

func init() {
	flag.StringVar(&urlsStr, "s", "", "list of urls separeted by commas")
	flag.Usage = func() {
		printUsage()
	}
	flag.Parse()
}

func main() {
	if len(os.Args) == 3 {
		upsi()
	} else {
		printUsage()
	}

}

func upsi() {
	var wg sync.WaitGroup
	urls := strings.Split(urlsStr, ",")
	results := make(chan Result, len(urls))
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
			res := Result{
				url:    url,
				status: stat,
				print:  fmt.Sprintf("%v - %v", url, stat),
			}

			results <- res
		}(url)
	}
	wg.Wait()
	close(results)
	for res := range results {
		fmt.Println(res.print)
	}

}

func printUsage() {
	fmt.Printf("Ussage:\n%v -s \"list of urls separeted by commas\"\n", os.Args[0])
}

func statusCheck(site string, client *http.Client) (string, error) {
	resp, err := client.Get(site)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Status, nil
}
