package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

func checkProtocol(domain string) {
	domain = strings.TrimSpace(domain)
	if domain == "" {
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	httpSupported := false
	httpsSupported := false
	http404 := false
	https404 := false

	// Check HTTP
	httpURL := "http://" + domain
	resp, err := client.Get(httpURL)
	if err == nil {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			httpSupported = true
		} else if resp.StatusCode == 404 {
			http404 = true
		}
		resp.Body.Close()
	}

	// Check HTTPS
	httpsURL := "https://" + domain
	resp, err = client.Get(httpsURL)
	if err == nil {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			httpsSupported = true
		} else if resp.StatusCode == 404 {
			https404 = true
		}
		resp.Body.Close()
	}

	// Output to console
	if httpSupported && httpsSupported {
		fmt.Printf("%s: Supports both HTTP and HTTPS\n", domain)
	} else if httpSupported {
		fmt.Printf("%s: Supports HTTP only\n", domain)
	} else if httpsSupported {
		fmt.Printf("%s: Supports HTTPS only\n", domain)
	} else {
		fmt.Printf("%s: No valid response from HTTP or HTTPS\n", domain)
	}

	// Save to files
	mu.Lock()
	defer mu.Unlock()

	// Save to Result-200.txt if any protocol succeeded
	if httpSupported || httpsSupported {
		file200, err := os.OpenFile("Result-200.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			fmt.Fprintln(file200, domain)
			file200.Close()
		}
	}

	// Save to Result-404.txt if any protocol returned 404
	if http404 || https404 {
		file404, err := os.OpenFile("Result-404.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			fmt.Fprintln(file404, domain)
			file404.Close()
		}
	}
}

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		go checkProtocol(domain) // Run in goroutine for parallelism
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	// Wait a bit for goroutines to finish (simple way, not ideal for production)
	time.Sleep(15 * time.Second)
}
