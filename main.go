package main

import (
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
	fileLoc string
)

func init() {
	flag.StringVar(&urlsStr, "s", "", "list of urls separated by commas")
	flag.StringVar(&fileLoc, "f", "", "file with list of urls separated by new line")
	flag.Usage = func() {
		printUsage()
	}
	flag.Parse()
}

func main() {
	if len(os.Args) == 3 {
		if os.Args[1] == "-s" {
			urls := strings.Split(urlsStr, ",")
			upsi(urls)
		} else if os.Args[1] == "-f" {
			urls, err := getUrlsFromFile()
			if err != nil {
				fmt.Println(err)
				return
			}
			upsi(urls)
		}
	} else {
		printUsage()
	}
}

func getUrlsFromFile() ([]string, error) {
	data, err := os.ReadFile(fileLoc)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	urlsStr := string(data)
	var urls []string
	for _, url := range strings.Split(urlsStr, "\n") {
		if url != "" {
			urls = append(urls, url)
		}
	}
	return urls, nil
}

func upsi(urls []string) {
	var wg sync.WaitGroup
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
				print:  fmt.Sprintf("URL: %-42v | Status: %v", url, stat),
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
	s := fmt.Sprintf("%v -s \"list of urls separated by commas\"\n", os.Args[0])
	f := fmt.Sprintf("%v -f \"file location:list of urls separated by new line\"\n", os.Args[0])
	fmt.Printf("Usage: \n%v%v", s, f)
}

func statusCheck(site string, client *http.Client) (string, error) {
	resp, err := client.Get(site)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Status, nil
}
